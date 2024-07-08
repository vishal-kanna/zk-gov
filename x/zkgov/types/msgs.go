package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	_ sdk.Msg = &RegisterCommitmentRequest{}
	_ sdk.Msg = &RegisterUserRequest{}
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

func NewRegisterUserRequest() *RegisterUserRequest {
	return &RegisterUserRequest{}
}
func (msg RegisterUserRequest) ValidateBasic() error {
	return nil
}
