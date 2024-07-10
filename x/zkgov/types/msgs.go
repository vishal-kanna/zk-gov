package types

import (
	"errors"
	"fmt"

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
	if len([]byte(msg.Commitment)) == COMMITMENT_SIZE {
		return errors.New(fmt.Sprintf("commitment should of size %d bytes", COMMITMENT_SIZE))
	}

	if len([]byte(msg.Sender)) == COMMITMENT_SIZE {
		return errors.New(fmt.Sprintf("Send should be a valid chain address of size %d bytes", USER_SIZE))
	}
	// TODO: sender should be valid secp256k1 address

	return nil
}

func NewMsgVoteProposal() *MsgVoteProposal {
	return &MsgVoteProposal{}
}
func (msg MsgVoteProposal) ValidateBasic() error {
	if len([]byte(msg.Nullifier)) == NULLIFIER_SIZE {
		return errors.New(fmt.Sprintf("nullifier should of size %d bytes", NULLIFIER_SIZE))
	}

	if len([]byte(msg.ProposalStateRoot)) == MERKLE_ROOT_SIZE {
		return errors.New(fmt.Sprintf("merkle root should be of size %d bytes", USER_SIZE))
	}
	return nil
}
