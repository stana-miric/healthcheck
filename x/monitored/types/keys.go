package types

const (
	// ModuleName defines the module name
	ModuleName = "monitored"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_monitored"
)

var (
	// PortKey defines the key to store the port ID in store
	PortKey = KeyPrefix("monitored-port-")
	// HealthcheckChannelKey defines the key to store the channel for communication with the healthcheck chain
	HealthcheckChannelKey = KeyPrefix("health-check-channel")
	// LAstCheckinKey defines the key to store the last checkin infromation in the store
	LastCheckinKey = KeyPrefix("last-checkin")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
