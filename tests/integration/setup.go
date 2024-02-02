package integration

import (
	"encoding/json"
	monitoredApp "healthcheck/app/monitored"
	registryApp "healthcheck/app/registry"
	registryTypes "healthcheck/x/healthcheck/types"
	commonTypes "healthcheck/x/types"
	"testing"

	ibctesting "github.com/cosmos/ibc-go/v6/testing"
	"github.com/stretchr/testify/suite"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"
)

const (
	RegistryChainID  = "registry"
	MonitoredChainID = "monitored"
)

type SetupRegistryCallback func(t *testing.T, coordinator *ibctesting.Coordinator) (*ibctesting.TestChain, *registryApp.App)
type SetupMonitoredCallback func(t *testing.T, coordinator *ibctesting.Coordinator) (*ibctesting.TestChain, *monitoredApp.App)

type HealthcheckTestSuite struct {
	suite.Suite

	setupRegistryCallback  SetupRegistryCallback
	setupMonitoredCallback SetupMonitoredCallback

	coordinator    *ibctesting.Coordinator
	registryChain  *ibctesting.TestChain
	monitoredChain *ibctesting.TestChain
	registryApp    *registryApp.App
	monitoredApp   *monitoredApp.App

	path *ibctesting.Path
}

func NewHealthcheckTestSuite() *HealthcheckTestSuite {
	healthcheckSuite := new(HealthcheckTestSuite)

	healthcheckSuite.setupRegistryCallback = func(t *testing.T, coordinator *ibctesting.Coordinator) (
		*ibctesting.TestChain,
		*registryApp.App,
	) {
		t.Helper()
		ibctesting.DefaultTestingAppInit = SetupRegistryApp
		registry := ibctesting.NewTestChain(t, coordinator, RegistryChainID)
		coordinator.Chains[RegistryChainID] = registry

		return registry, registry.App.(*registryApp.App)
	}

	healthcheckSuite.setupMonitoredCallback = func(t *testing.T, coordinator *ibctesting.Coordinator) (
		*ibctesting.TestChain,
		*monitoredApp.App,
	) {
		t.Helper()
		ibctesting.DefaultTestingAppInit = SetupMonitoredApp
		monitored := ibctesting.NewTestChain(t, coordinator, MonitoredChainID)
		coordinator.Chains[MonitoredChainID] = monitored

		return monitored, monitored.App.(*monitoredApp.App)
	}

	return healthcheckSuite
}

// SetupTest sets up in-mem state before every test
func (suite *HealthcheckTestSuite) SetupTest() {
	suite.coordinator = ibctesting.NewCoordinator(suite.T(), 0)
	suite.registryChain, suite.registryApp = suite.setupRegistryCallback(suite.T(), suite.coordinator)
	suite.monitoredChain, suite.monitoredApp = suite.setupMonitoredCallback(suite.T(), suite.coordinator)

	// create clients and connection
	suite.path = ibctesting.NewPath(suite.monitoredChain, suite.registryChain) // clientID, connectionID, channelID empty
	suite.coordinator.SetupConnections(suite.path)                             // clientID, connectionID
	suite.Require().Equal("07-tendermint-0", suite.path.EndpointA.ClientID)
	suite.Require().Equal("connection-0", suite.path.EndpointA.ConnectionID)
	suite.Require().Equal("07-tendermint-0", suite.path.EndpointB.ClientID)
	suite.Require().Equal("connection-0", suite.path.EndpointB.ConnectionID)

	//register monitored chain
	var newMonitoredChain = registryTypes.MonitoredChain{
		ChainId:      MonitoredChainID,
		ConnectionId: suite.path.EndpointA.ConnectionID,
	}
	suite.registryApp.HealthcheckKeeper.SetMonitoredChain(suite.registryChain.GetContext(), newMonitoredChain)

	// - channel config
	suite.path.EndpointA.ChannelConfig.PortID = commonTypes.MonitoredPortID
	suite.path.EndpointB.ChannelConfig.PortID = commonTypes.HealthcheckPortID
	suite.path.EndpointA.ChannelConfig.Version = commonTypes.Version
	suite.path.EndpointB.ChannelConfig.Version = commonTypes.Version
	suite.coordinator.CreateChannels(suite.path) // setup channel
	suite.Require().Equal("channel-0", suite.path.EndpointA.ChannelID)
	suite.Require().Equal("channel-0", suite.path.EndpointB.ChannelID)
}

func SetupRegistryApp() (ibctesting.TestingApp, map[string]json.RawMessage) {
	db := dbm.NewMemDB()
	encCdc := registryApp.MakeEncodingConfig()
	app := registryApp.New(
		log.NewNopLogger(),
		db,
		nil,
		true,
		map[int64]bool{},
		registryApp.DefaultNodeHome,
		5,
		encCdc,
		EmptyAppOptions{},
	)
	return app, registryApp.NewDefaultGenesisState(encCdc.Marshaler)
}

func SetupMonitoredApp() (ibctesting.TestingApp, map[string]json.RawMessage) {
	db := dbm.NewMemDB()
	encCdc := monitoredApp.MakeEncodingConfig()
	app := monitoredApp.New(
		log.NewNopLogger(),
		db,
		nil,
		true,
		map[int64]bool{},
		monitoredApp.DefaultNodeHome,
		5,
		encCdc,
		EmptyAppOptions{},
	)
	return app, monitoredApp.NewDefaultGenesisState(encCdc.Marshaler)
}

// EmptyAppOptions is a stub implementing AppOptions
type EmptyAppOptions struct{}

// Get implements AppOptions
func (ao EmptyAppOptions) Get(o string) interface{} {
	return nil
}
