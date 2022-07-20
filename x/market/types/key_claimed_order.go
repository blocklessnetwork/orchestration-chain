package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ClaimedOrderKeyPrefix is the prefix to retrieve all ClaimedOrder
	ClaimedOrderKeyPrefix = "ClaimedOrder/value/"
)

// ClaimedOrderKey returns the store key to retrieve a ClaimedOrder from the index fields
func ClaimedOrderKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
