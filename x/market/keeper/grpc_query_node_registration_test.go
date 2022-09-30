package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/blocklessnetwork/orchestration-chain/testutil/keeper"
	"github.com/blocklessnetwork/orchestration-chain/testutil/nullify"
	"github.com/blocklessnetwork/orchestration-chain/x/market/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestNodeRegistrationQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.MarketKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNNodeRegistration(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetNodeRegistrationRequest
		response *types.QueryGetNodeRegistrationResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetNodeRegistrationRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetNodeRegistrationResponse{NodeRegistration: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetNodeRegistrationRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetNodeRegistrationResponse{NodeRegistration: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetNodeRegistrationRequest{
				Index: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.NodeRegistration(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestNodeRegistrationQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.MarketKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNNodeRegistration(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllNodeRegistrationRequest {
		return &types.QueryAllNodeRegistrationRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.NodeRegistrationAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.NodeRegistration), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.NodeRegistration),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.NodeRegistrationAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.NodeRegistration), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.NodeRegistration),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.NodeRegistrationAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.NodeRegistration),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.NodeRegistrationAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
