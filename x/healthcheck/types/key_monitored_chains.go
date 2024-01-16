package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// MonitoredChainKeyPrefix is the prefix to retrieve all MonitoredChain
	MonitoredChainKeyPrefix = "MonitoredChain/value/"
)

// MonitoredChainKey returns the store key to retrieve a MonitoredChain from the index fields
func MonitoredChainKey(
	chainId string,
) []byte {
	var key []byte

	chainIdBytes := []byte(chainId)
	key = append(key, chainIdBytes...)
	key = append(key, []byte("/")...)

	return key
}
