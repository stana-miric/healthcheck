package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// MonitoredChainKeyPrefix is the prefix to retrieve all MonitoredChain
	MonitoredChainKeyPrefix = "MonitoredChain/value/"
	// ChainToChannelKeyPrefix is the prefix to retrieve channel by monitored chain id
	ChainToChannelKeyPrefix = "ChainToChannel/value/"
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

// ChainToChannelKey returns the store key to retrieve a channel of the MonitoredChain from the index fields
func ChainToChannelKey(
	chainId string,
) []byte {
	var key []byte

	chainIdBytes := []byte(chainId)
	key = append(key, chainIdBytes...)
	key = append(key, []byte("/")...)

	return key
}
