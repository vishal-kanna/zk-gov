package keeper

import (
	"encoding/binary"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	CommitmentKey       = []byte{0x01}
	CommitmentSeqPrefix = []byte{0x02}
	UserKey             = []byte{0x03}
	UserSeqPrefix       = []byte{0x04}
)

func CommitmentStoreKey(commitmentID uint64) []byte {
	commitmentIDBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(commitmentIDBytes, commitmentID)
	key := append(CommitmentKey, commitmentIDBytes...)
	return key
}
func ParseCommitmentStoreKey(key []byte) uint64 {
	return sdk.BigEndianToUint64(key[1:])
}

func UserStoreKey(userId uint64) []byte {
	userIdBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(userIdBytes, userId)
	key := append(UserKey, userIdBytes...)
	return key
}
