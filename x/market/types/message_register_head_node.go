package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRegisterHeadNode = "register_head_node"

var _ sdk.Msg = &MsgRegisterHeadNode{}

func NewMsgRegisterHeadNode(creator string, nodeId string, nodePort string, nodeIp string) *MsgRegisterHeadNode {
	return &MsgRegisterHeadNode{
		Creator:  creator,
		NodeId:   nodeId,
		NodePort: nodePort,
		NodeIp:   nodeIp,
	}
}

func (msg *MsgRegisterHeadNode) Route() string {
	return RouterKey
}

func (msg *MsgRegisterHeadNode) Type() string {
	return TypeMsgRegisterHeadNode
}

func (msg *MsgRegisterHeadNode) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRegisterHeadNode) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRegisterHeadNode) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
