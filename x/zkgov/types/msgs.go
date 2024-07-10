package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	_ sdk.Msg = &MsgRegisterUser{}
	_ sdk.Msg = &MsgVoteProposal{}
)

func NewMsgRegisterUser(commitment string, sender string, proposalID uint64) *MsgRegisterUser {
	return &MsgRegisterUser{Commitment: commitment, Sender: sender, ProposalId: proposalID}
}
func (msg MsgRegisterUser) ValidateBasic() error {
	// TODO: commitment size
	// TODO: sender should be valid secp256k1 address
	// TODO: proposalID not nill
	return nil
}

func NewMsgVoteProposal() *MsgVoteProposal {
	return &MsgVoteProposal{}
}
func (msg MsgVoteProposal) ValidateBasic() error {
	// TODO: null size
	// TODO: state root size
	return nil
}
