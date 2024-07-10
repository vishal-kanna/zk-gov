package types

import (
	"encoding/binary"
)

const (
	// ModuleName defines the name of the nft module
	ModuleName = "zk-gov"

	// StoreKey is the default store key for nft
	StoreKey = ModuleName

	// RouterKey is the message route for nft
	RouterKey = ModuleName
)

var (
	CommitmentsKey = []byte{0x01}
	UsersKey       = []byte{0x02}
	MerkleRootKey  = []byte{0x03}
	NullifiersKey  = []byte{0x04}
)

func CommitmentsStoreKey(proposalID uint64) []byte {
	proposalIDBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(proposalIDBytes, proposalID)

	key := append(CommitmentsKey, proposalIDBytes...)

	return key
}

func NullifiersStoreKey(proposalID uint64) []byte {
	proposalIDBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(proposalIDBytes, proposalID)

	key := append(NullifiersKey, proposalIDBytes...)

	return key
}

func UsersStoreKey(proposalID uint64) []byte {

	proposalIDBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(proposalIDBytes, proposalID)

	key := append(UsersKey, proposalIDBytes...)

	return key
}

func MerkleRootStoreKey(proposalID uint64) []byte {
	proposalIDBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(proposalIDBytes, proposalID)
	key := append(MerkleRootKey, proposalIDBytes...)
	return key

}
