package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"github.com/txlabs/blockless-chain/testutil/sample"
)

func TestMsgSubmitCompletedOrder_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgSubmitCompletedOrder
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgSubmitCompletedOrder{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgSubmitCompletedOrder{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
