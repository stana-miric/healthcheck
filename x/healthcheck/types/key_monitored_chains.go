package types

import (
	"encoding/binary"
)

var _ binary.ByteOrder

const (
	// MonitoredChainKeyPrefix is the prefix to retrieve all MonitoredChain
	MonitoredChainKeyPrefix byte = iota
	// ChainToChannelKeyPrefix is the prefix to retrieve channel by monitored chain id
	ChainToChannelKeyPrefix
)

// MonitoredChainKey returns the store key to retrieve a MonitoredChain from the index fields
func MonitoredChainKey(chainId string) []byte {
	return append([]byte{MonitoredChainKeyPrefix}, []byte(chainId)...)
}

// ChainToChannelKey returns the store key to retrieve a channel of the MonitoredChain from the index fields
func ChainToChannelKey(chainId string) []byte {
	return append([]byte{ChainToChannelKeyPrefix}, []byte(chainId)...)
}
