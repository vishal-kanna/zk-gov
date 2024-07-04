package keeper

import (
	"encoding/binary"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	CommitmentKey       = []byte{0x01}
	CommitmentSeqPrefix = []byte{0x02}
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
