package types

import "encoding/hex"

type (
	Hash [32]byte
)

var (
	EmptyHash = Hash{}
)

func (hash *Hash) String() string {
	return string(hash[:])
}

func (hash *Hash) MarshalJSON() ([]byte, error) {
	return []byte(`"` + hex.EncodeToString(hash[:]) + `"`), nil
}
