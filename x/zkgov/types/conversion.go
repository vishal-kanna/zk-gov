package types

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
)

func HexStringToBytes(data string) ([]byte, error) {
	bytes, err := hex.DecodeString(data)
	if err != nil {
		return nil, fmt.Errorf("invalid hex string: %v", err)
	}
	return bytes, nil
}

func BytesToHexString(data []byte) string {
	return hex.EncodeToString(data)
}

func MarshalVoteOption(voteOption VoteOption) []byte {
	vote := uint64(0)
	if voteOption == VoteOption_VOTE_OPTION_YES {
		vote = uint64(1)
	}

	voteOptionBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(voteOptionBytes, vote)

	return voteOptionBytes
}

func UnMarshalVoteOption(voteOptionBytes []byte) VoteOption {
	vote := binary.BigEndian.Uint64(voteOptionBytes)
	voteOption := VoteOption_VOTE_OPTION_NO
	if vote == 1 {
		voteOption = VoteOption_VOTE_OPTION_YES
	}
	return voteOption
}
