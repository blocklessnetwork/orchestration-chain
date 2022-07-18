package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSubmitCompletedOrder = "submit_completed_order"

var _ sdk.Msg = &MsgSubmitCompletedOrder{}

func NewMsgSubmitCompletedOrder(creator string, orderIndex string) *MsgSubmitCompletedOrder {
	return &MsgSubmitCompletedOrder{
		Creator:    creator,
		OrderIndex: orderIndex,
	}
}

func (msg *MsgSubmitCompletedOrder) Route() string {
	return RouterKey
}

func (msg *MsgSubmitCompletedOrder) Type() string {
	return TypeMsgSubmitCompletedOrder
}

func (msg *MsgSubmitCompletedOrder) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSubmitCompletedOrder) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSubmitCompletedOrder) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
