package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// AliasesKeyPrefix is the prefix to retrieve all Aliases
	AliasesKeyPrefix = "Aliases/value/"
)

// AliasesKey returns the store key to retrieve a Aliases from the index fields
func AliasesKey(
	address string,
) []byte {
	var key []byte

	addressBytes := []byte(address)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)

	return key
}
