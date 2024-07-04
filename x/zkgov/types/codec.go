package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/legacy"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

// RegisterLegacyAminoCodec registers the necessary x/bank interfaces and concrete types
// on the provided LegacyAmino codec. These types are used for Amino JSON serialization.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	legacy.RegisterAminoMsg(cdc, &RegisterCommitmentRequest{}, "zkgov/RegisterCommitment")
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&RegisterCommitmentRequest{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
