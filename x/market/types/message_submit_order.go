package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSubmitOrder = "submit_order"

var _ sdk.Msg = &MsgSubmitOrder{}

func NewMsgSubmitOrder(creator string, functionId string, fuel string) *MsgSubmitOrder {
	return &MsgSubmitOrder{
		Creator:    creator,
		FunctionId: functionId,
		Fuel:       fuel,
	}
}

func (msg *MsgSubmitOrder) Route() string {
	return RouterKey
}

func (msg *MsgSubmitOrder) Type() string {
	return TypeMsgSubmitOrder
}

func (msg *MsgSubmitOrder) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSubmitOrder) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSubmitOrder) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
