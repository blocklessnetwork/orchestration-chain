package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgClaimOrderWork = "claim_order_work"

var _ sdk.Msg = &MsgClaimOrderWork{}

func NewMsgClaimOrderWork(creator string, orderIndex string) *MsgClaimOrderWork {
	return &MsgClaimOrderWork{
		Creator:    creator,
		OrderIndex: orderIndex,
	}
}

func (msg *MsgClaimOrderWork) Route() string {
	return RouterKey
}

func (msg *MsgClaimOrderWork) Type() string {
	return TypeMsgClaimOrderWork
}

func (msg *MsgClaimOrderWork) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgClaimOrderWork) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgClaimOrderWork) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
