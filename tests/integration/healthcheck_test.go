package integration

import (
	registryTypes "healthcheck/x/healthcheck/types"
	"healthcheck/x/monitored/types"
	commonTypes "healthcheck/x/types"
	"testing"
	"time"

	clienttypes "github.com/cosmos/ibc-go/v6/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v6/modules/core/04-channel/types"
	ibctesting "github.com/cosmos/ibc-go/v6/testing"
	"github.com/stretchr/testify/suite"
)

func TestHealthcheckTestSuite(t *testing.T) {
	// Run tests
	suite.Run(t, NewHealthcheckTestSuite())
}

func (s *HealthcheckTestSuite) TestHealthcheck() {

	//Send few healthcheck packets
	for i := 0; i < 5; i++ {
		ctx := s.monitoredChain.GetContext()
		monitoredBlockTime := ctx.BlockTime()
		monitoredBlockHeight := uint64(ctx.BlockHeight())
		packet := commonTypes.HealthcheckUpdateData{
			Block:     monitoredBlockHeight,
			Timestamp: uint64(monitoredBlockTime.UnixNano()),
		}
		packetData, err := types.ModuleCdc.MarshalJSON(&packet)
		s.Require().Nil(err)

		timeout := uint64(monitoredBlockTime.Add(time.Hour).UnixNano()) //packet timeout
		sendHealthcheckPacket(s, s.path, timeout, packetData)

		monitoredChainEntry, found := s.registryApp.HealthcheckKeeper.GetMonitoredChain(s.registryChain.GetContext(), MonitoredChainID)
		s.Require().True(found)
		s.Require().True(monitoredBlockHeight == monitoredChainEntry.Status.Block)
		s.Require().True(uint64(monitoredBlockTime.UnixNano()) == monitoredChainEntry.Status.Timestamp)
	}
}

func (s *HealthcheckTestSuite) TestHealthcheckTimeout() {

	//Send one healthcheck packet
	ctx := s.monitoredChain.GetContext()
	monitoredBlockTime := ctx.BlockTime()
	monitoredBlockHeight := uint64(ctx.BlockHeight())
	packet := commonTypes.HealthcheckUpdateData{
		Block:     monitoredBlockHeight,
		Timestamp: uint64(monitoredBlockTime.UnixNano()),
	}
	packetData, err := types.ModuleCdc.MarshalJSON(&packet)
	s.Require().Nil(err)

	timeout := uint64(monitoredBlockTime.Add(time.Hour).UnixNano()) //packet timeout
	sendHealthcheckPacket(s, s.path, timeout, packetData)

	monitoredChainEntry, found := s.registryApp.HealthcheckKeeper.GetMonitoredChain(s.registryChain.GetContext(), MonitoredChainID)
	s.Require().True(found)
	s.Require().True(monitoredBlockHeight == monitoredChainEntry.Status.Block)
	s.Require().True(uint64(monitoredBlockTime.UnixNano()) == monitoredChainEntry.Status.Timestamp)
	s.Require().True(monitoredChainEntry.Status.Status == string(registryTypes.Active))
	channelID, found := s.registryApp.HealthcheckKeeper.GetChainToChannelMap(s.registryChain.GetContext(), MonitoredChainID)
	s.Require().True(found)
	channel, found := s.registryApp.GetIBCKeeper().ChannelKeeper.GetChannel(s.registryChain.GetContext(), commonTypes.HealthcheckPortID, channelID)
	s.Require().True(found)
	s.Require().True(channel.State == channeltypes.OPEN)
	updateInterval := int(monitoredChainEntry.UpdateInterval)
	timeoutInterval := int(monitoredChainEntry.TimeoutInterval)

	//Move registry chain to reach update interval which will set monitored chain to inactive
	for i := 0; i <= updateInterval; i++ {
		s.coordinator.CommitBlock(s.registryChain)
	}

	monitoredChainEntry, found = s.registryApp.HealthcheckKeeper.GetMonitoredChain(s.registryChain.GetContext(), MonitoredChainID)
	s.Require().True(found)
	s.Require().True(monitoredChainEntry.Status.Status == string(registryTypes.Inactive))
	channelID, found = s.registryApp.HealthcheckKeeper.GetChainToChannelMap(s.registryChain.GetContext(), MonitoredChainID)
	s.Require().True(found)
	channel, found = s.registryApp.GetIBCKeeper().ChannelKeeper.GetChannel(s.registryChain.GetContext(), commonTypes.HealthcheckPortID, channelID)
	s.Require().True(found)
	s.Require().True(channel.State == channeltypes.OPEN)

	//Move registry chain to reach timeout interval which will close the channel and remove it from chainToChannel map
	for i := 0; i <= timeoutInterval; i++ {
		s.coordinator.CommitBlock(s.registryChain)
	}
	channel, found = s.registryApp.GetIBCKeeper().ChannelKeeper.GetChannel(s.registryChain.GetContext(), commonTypes.HealthcheckPortID, channelID)
	s.Require().True(found)
	s.Require().True(channel.State == channeltypes.CLOSED)
	_, found = s.registryApp.HealthcheckKeeper.GetChainToChannelMap(s.registryChain.GetContext(), MonitoredChainID)
	s.Require().False(found)

}

func sendHealthcheckPacket(s *HealthcheckTestSuite, path *ibctesting.Path, timeoutTimestamp uint64, data []byte) channeltypes.Packet {
	sequence, err := path.EndpointA.SendPacket(clienttypes.Height{}, timeoutTimestamp, data)
	s.Require().NoError(err)

	packet := s.newHealthcheckPacket(data, sequence, path, clienttypes.Height{}, timeoutTimestamp)

	err = path.EndpointB.RecvPacket(packet)
	s.Require().NoError(err)
	return packet
}

func (s *HealthcheckTestSuite) newHealthcheckPacket(data []byte, sequence uint64, path *ibctesting.Path, timeoutHeight clienttypes.Height, timeoutTimestamp uint64) channeltypes.Packet {
	return channeltypes.NewPacket(data, sequence,
		path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID,
		path.EndpointB.ChannelConfig.PortID, path.EndpointB.ChannelID,
		timeoutHeight, timeoutTimestamp)
}
