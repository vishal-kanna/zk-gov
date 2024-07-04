package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	_ sdk.Msg = &RegisterCommitmentRequest{}
)

func NewRegisterCommitmentRequest(commitment string) *RegisterCommitmentRequest {
	return &RegisterCommitmentRequest{Commitment: commitment}
}
func (msg RegisterCommitmentRequest) ValidateBasic() error {
	if msg.Commitment == "" {
		return EmptyCommitment
	}
	return nil
}
