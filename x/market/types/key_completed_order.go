package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// CompletedOrderKeyPrefix is the prefix to retrieve all CompletedOrder
	CompletedOrderKeyPrefix = "CompletedOrder/value/"
)

// CompletedOrderKey returns the store key to retrieve a CompletedOrder from the index fields
func CompletedOrderKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
