package abci_test

import (
	"testing"

	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sei-protocol/sei-chain/x/dex/keeper"
	"github.com/sei-protocol/sei-chain/x/dex/keeper/abci"
	"github.com/sei-protocol/sei-chain/x/dex/types"
)

const (
	SupportedFeatures = "iterator,staking,stargate"
	TestContract      = "sei14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9sh9m79m"
)

func TestHandleBBNewBlock(t *testing.T) {
	// this test only ensures that HandleBBNewBlock doesn't crash. The actual logic
	// is tested in module_test.go where an actual wasm file is deployed and invoked.
	wasmkeeper.TestingStakeParams.MinCommissionRate = sdk.NewDecWithPrec(5, 2)
	ctx, wasmkeepers := wasmkeeper.CreateTestInput(t, false, SupportedFeatures)
	ctx = ctx.WithGasMeter(sdk.NewInfiniteGasMeter())
	dexKeeper := keeper.Keeper{WasmKeeper: *wasmkeepers.WasmKeeper}
	dexKeeper.SetContract(ctx, &types.ContractInfoV2{
		ContractAddr: TestContract,
		RentBalance:  100000000,
	})
	wrapper := abci.KeeperWrapper{Keeper: &dexKeeper}
	wrapper.HandleBBNewBlock(ctx, TestContract, 1)
}
