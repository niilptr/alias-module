package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "units/testutil/keeper"
	"units/testutil/nullify"
	"units/x/alias/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestAliasesQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.AliasKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNAliases(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetAliasesRequest
		response *types.QueryGetAliasesResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetAliasesRequest{
				Address: msgs[0].Address,
			},
			response: &types.QueryGetAliasesResponse{Aliases: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetAliasesRequest{
				Address: msgs[1].Address,
			},
			response: &types.QueryGetAliasesResponse{Aliases: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetAliasesRequest{
				Address: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Aliases(wctx, tc.request)
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

func TestAliasesQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.AliasKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNAliases(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllAliasesRequest {
		return &types.QueryAllAliasesRequest{
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
			resp, err := keeper.AliasesAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Aliases), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Aliases),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.AliasesAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Aliases), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Aliases),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.AliasesAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Aliases),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.AliasesAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
