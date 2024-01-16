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
	PortKey               = KeyPrefix("monitored-port-")
	HealthcheckChannelKey = KeyPrefix("health-check-channel")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
