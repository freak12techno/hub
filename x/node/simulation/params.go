package simulation

import (
	"fmt"
	"math/rand"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	simulationtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/sentinel-official/hub/x/node/types"
)

const (
	MaxDepositAmount    = 1 << 18
	MaxInactiveDuration = 1 << 18
)

func ParamChanges(_ *rand.Rand) []simulationtypes.ParamChange {
	return []simulationtypes.ParamChange{
		simulation.NewSimParamChange(
			types.ModuleName,
			string(types.KeyDeposit),
			func(r *rand.Rand) string {
				return sdk.NewInt64Coin(
					sdk.DefaultBondDenom,
					r.Int63n(MaxDepositAmount),
				).String()
			},
		),
		simulation.NewSimParamChange(
			types.ModuleName,
			string(types.KeyInactiveDuration),
			func(r *rand.Rand) string {
				return fmt.Sprintf(
					"%s",
					time.Duration(r.Int63n(MaxInactiveDuration))*time.Millisecond,
				)
			},
		),
	}
}
