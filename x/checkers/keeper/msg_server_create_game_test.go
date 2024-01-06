package keeper_test

import (
	//"context"
	"testing"

	// keepertest "github.com/alice/checkers/testutil/keeper"
	// "github.com/alice/checkers/x/checkers"
	// "github.com/alice/checkers/x/checkers/keeper"
	// "github.com/alice/checkers/x/checkers/testutil"
	"github.com/alice/checkers/x/checkers/types"
	// sdk "github.com/cosmos/cosmos-sdk/types"
	// "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestCreateGame(t *testing.T) {
    msgServer, context := setupMsgServer(t)
    createResponse, err := msgServer.CreateGame(context, &types.MsgCreateGame{
        Creator: alice,
        Black:   bob,
        Red:     carol,
    })
    require.Nil(t, err)
    require.EqualValues(t, types.MsgCreateGameResponse{
        GameIndex: "", // TODO: update with a proper value when updated
    }, *createResponse)
}
