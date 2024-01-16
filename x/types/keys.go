package types

const (
	// Version defines the current version the IBC healtcheck and monitored
	// module supports
	Version = "1"

	// ProviderPortID is the default port id the healthcheck CCV module binds to
	HealthcheckPortID = "healthcheck"

	// ConsumerPortID is the default port id the monitored CCV module binds to
	MonitoredPortID = "monitored"
)
