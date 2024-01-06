package keeper

import (
	"context"
	"strconv"

	rules "github.com/alice/checkers/x/checkers/rules"
	"github.com/alice/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateGame(goCtx context.Context, msg *types.MsgCreateGame) (*types.MsgCreateGameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	systemInfo, found := k.Keeper.GetSystemInfo(ctx)
	if !found {
		panic("System Info not found")
	}
	newIndex := strconv.FormatInt(int64(systemInfo.NextId), 10)

	newGame := rules.New()
	storedGame := types.StoredGame{
		Index: newIndex,
		Board: newGame.String(),
		Turn:  rules.PieceStrings[newGame.Turn],
		Black: msg.Black,
		Red:   msg.Red,
	}

	_ = ctx

	return &types.MsgCreateGameResponse{}, nil
}
