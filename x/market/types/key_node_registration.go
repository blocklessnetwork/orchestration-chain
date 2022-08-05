package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// NodeRegistrationKeyPrefix is the prefix to retrieve all NodeRegistration
	NodeRegistrationKeyPrefix = "NodeRegistration/value/"
)

// NodeRegistrationKey returns the store key to retrieve a NodeRegistration from the index fields
func NodeRegistrationKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
