package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// MonitoredChainsKeyPrefix is the prefix to retrieve all MonitoredChains
	MonitoredChainsKeyPrefix = "MonitoredChains/value/"
)

// MonitoredChainsKey returns the store key to retrieve a MonitoredChains from the index fields
func MonitoredChainsKey(
	chainId string,
) []byte {
	var key []byte

	chainIdBytes := []byte(chainId)
	key = append(key, chainIdBytes...)
	key = append(key, []byte("/")...)

	return key
}
