// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	nexus "github.com/axelarnetwork/axelar-core/x/nexus/exported"
	"github.com/axelarnetwork/axelar-core/x/reward/exported"
	rewardtypes "github.com/axelarnetwork/axelar-core/x/reward/types"
	cosmossdktypes "github.com/cosmos/cosmos-sdk/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/tendermint/tendermint/libs/log"
	"sync"
)

// Ensure, that RewarderMock does implement rewardtypes.Rewarder.
// If this is not the case, regenerate this file with moq.
var _ rewardtypes.Rewarder = &RewarderMock{}

// RewarderMock is a mock implementation of rewardtypes.Rewarder.
//
// 	func TestSomethingThatUsesRewarder(t *testing.T) {
//
// 		// make and configure a mocked rewardtypes.Rewarder
// 		mockedRewarder := &RewarderMock{
// 			GetParamsFunc: func(ctx cosmossdktypes.Context) rewardtypes.Params {
// 				panic("mock out the GetParams method")
// 			},
// 			GetPoolFunc: func(ctx cosmossdktypes.Context, name string) exported.RewardPool {
// 				panic("mock out the GetPool method")
// 			},
// 			LoggerFunc: func(ctx cosmossdktypes.Context) log.Logger {
// 				panic("mock out the Logger method")
// 			},
// 		}
//
// 		// use mockedRewarder in code that requires rewardtypes.Rewarder
// 		// and then make assertions.
//
// 	}
type RewarderMock struct {
	// GetParamsFunc mocks the GetParams method.
	GetParamsFunc func(ctx cosmossdktypes.Context) rewardtypes.Params

	// GetPoolFunc mocks the GetPool method.
	GetPoolFunc func(ctx cosmossdktypes.Context, name string) exported.RewardPool

	// LoggerFunc mocks the Logger method.
	LoggerFunc func(ctx cosmossdktypes.Context) log.Logger

	// calls tracks calls to the methods.
	calls struct {
		// GetParams holds details about calls to the GetParams method.
		GetParams []struct {
			// Ctx is the ctx argument value.
			Ctx cosmossdktypes.Context
		}
		// GetPool holds details about calls to the GetPool method.
		GetPool []struct {
			// Ctx is the ctx argument value.
			Ctx cosmossdktypes.Context
			// Name is the name argument value.
			Name string
		}
		// Logger holds details about calls to the Logger method.
		Logger []struct {
			// Ctx is the ctx argument value.
			Ctx cosmossdktypes.Context
		}
	}
	lockGetParams sync.RWMutex
	lockGetPool   sync.RWMutex
	lockLogger    sync.RWMutex
}

// GetParams calls GetParamsFunc.
func (mock *RewarderMock) GetParams(ctx cosmossdktypes.Context) rewardtypes.Params {
	if mock.GetParamsFunc == nil {
		panic("RewarderMock.GetParamsFunc: method is nil but Rewarder.GetParams was just called")
	}
	callInfo := struct {
		Ctx cosmossdktypes.Context
	}{
		Ctx: ctx,
	}
	mock.lockGetParams.Lock()
	mock.calls.GetParams = append(mock.calls.GetParams, callInfo)
	mock.lockGetParams.Unlock()
	return mock.GetParamsFunc(ctx)
}

// GetParamsCalls gets all the calls that were made to GetParams.
// Check the length with:
//     len(mockedRewarder.GetParamsCalls())
func (mock *RewarderMock) GetParamsCalls() []struct {
	Ctx cosmossdktypes.Context
} {
	var calls []struct {
		Ctx cosmossdktypes.Context
	}
	mock.lockGetParams.RLock()
	calls = mock.calls.GetParams
	mock.lockGetParams.RUnlock()
	return calls
}

// GetPool calls GetPoolFunc.
func (mock *RewarderMock) GetPool(ctx cosmossdktypes.Context, name string) exported.RewardPool {
	if mock.GetPoolFunc == nil {
		panic("RewarderMock.GetPoolFunc: method is nil but Rewarder.GetPool was just called")
	}
	callInfo := struct {
		Ctx  cosmossdktypes.Context
		Name string
	}{
		Ctx:  ctx,
		Name: name,
	}
	mock.lockGetPool.Lock()
	mock.calls.GetPool = append(mock.calls.GetPool, callInfo)
	mock.lockGetPool.Unlock()
	return mock.GetPoolFunc(ctx, name)
}

// GetPoolCalls gets all the calls that were made to GetPool.
// Check the length with:
//     len(mockedRewarder.GetPoolCalls())
func (mock *RewarderMock) GetPoolCalls() []struct {
	Ctx  cosmossdktypes.Context
	Name string
} {
	var calls []struct {
		Ctx  cosmossdktypes.Context
		Name string
	}
	mock.lockGetPool.RLock()
	calls = mock.calls.GetPool
	mock.lockGetPool.RUnlock()
	return calls
}

// Logger calls LoggerFunc.
func (mock *RewarderMock) Logger(ctx cosmossdktypes.Context) log.Logger {
	if mock.LoggerFunc == nil {
		panic("RewarderMock.LoggerFunc: method is nil but Rewarder.Logger was just called")
	}
	callInfo := struct {
		Ctx cosmossdktypes.Context
	}{
		Ctx: ctx,
	}
	mock.lockLogger.Lock()
	mock.calls.Logger = append(mock.calls.Logger, callInfo)
	mock.lockLogger.Unlock()
	return mock.LoggerFunc(ctx)
}

// LoggerCalls gets all the calls that were made to Logger.
// Check the length with:
//     len(mockedRewarder.LoggerCalls())
func (mock *RewarderMock) LoggerCalls() []struct {
	Ctx cosmossdktypes.Context
} {
	var calls []struct {
		Ctx cosmossdktypes.Context
	}
	mock.lockLogger.RLock()
	calls = mock.calls.Logger
	mock.lockLogger.RUnlock()
	return calls
}

// Ensure, that RefunderMock does implement rewardtypes.Refunder.
// If this is not the case, regenerate this file with moq.
var _ rewardtypes.Refunder = &RefunderMock{}

// RefunderMock is a mock implementation of rewardtypes.Refunder.
//
// 	func TestSomethingThatUsesRefunder(t *testing.T) {
//
// 		// make and configure a mocked rewardtypes.Refunder
// 		mockedRefunder := &RefunderMock{
// 			DeletePendingRefundFunc: func(ctx cosmossdktypes.Context, req rewardtypes.RefundMsgRequest)  {
// 				panic("mock out the DeletePendingRefund method")
// 			},
// 			GetPendingRefundFunc: func(ctx cosmossdktypes.Context, req rewardtypes.RefundMsgRequest) (rewardtypes.Refund, bool) {
// 				panic("mock out the GetPendingRefund method")
// 			},
// 			LoggerFunc: func(ctx cosmossdktypes.Context) log.Logger {
// 				panic("mock out the Logger method")
// 			},
// 		}
//
// 		// use mockedRefunder in code that requires rewardtypes.Refunder
// 		// and then make assertions.
//
// 	}
type RefunderMock struct {
	// DeletePendingRefundFunc mocks the DeletePendingRefund method.
	DeletePendingRefundFunc func(ctx cosmossdktypes.Context, req rewardtypes.RefundMsgRequest)

	// GetPendingRefundFunc mocks the GetPendingRefund method.
	GetPendingRefundFunc func(ctx cosmossdktypes.Context, req rewardtypes.RefundMsgRequest) (rewardtypes.Refund, bool)

	// LoggerFunc mocks the Logger method.
	LoggerFunc func(ctx cosmossdktypes.Context) log.Logger

	// calls tracks calls to the methods.
	calls struct {
		// DeletePendingRefund holds details about calls to the DeletePendingRefund method.
		DeletePendingRefund []struct {
			// Ctx is the ctx argument value.
			Ctx cosmossdktypes.Context
			// Req is the req argument value.
			Req rewardtypes.RefundMsgRequest
		}
		// GetPendingRefund holds details about calls to the GetPendingRefund method.
		GetPendingRefund []struct {
			// Ctx is the ctx argument value.
			Ctx cosmossdktypes.Context
			// Req is the req argument value.
			Req rewardtypes.RefundMsgRequest
		}
		// Logger holds details about calls to the Logger method.
		Logger []struct {
			// Ctx is the ctx argument value.
			Ctx cosmossdktypes.Context
		}
	}
	lockDeletePendingRefund sync.RWMutex
	lockGetPendingRefund    sync.RWMutex
	lockLogger              sync.RWMutex
}

// DeletePendingRefund calls DeletePendingRefundFunc.
func (mock *RefunderMock) DeletePendingRefund(ctx cosmossdktypes.Context, req rewardtypes.RefundMsgRequest) {
	if mock.DeletePendingRefundFunc == nil {
		panic("RefunderMock.DeletePendingRefundFunc: method is nil but Refunder.DeletePendingRefund was just called")
	}
	callInfo := struct {
		Ctx cosmossdktypes.Context
		Req rewardtypes.RefundMsgRequest
	}{
		Ctx: ctx,
		Req: req,
	}
	mock.lockDeletePendingRefund.Lock()
	mock.calls.DeletePendingRefund = append(mock.calls.DeletePendingRefund, callInfo)
	mock.lockDeletePendingRefund.Unlock()
	mock.DeletePendingRefundFunc(ctx, req)
}

// DeletePendingRefundCalls gets all the calls that were made to DeletePendingRefund.
// Check the length with:
//     len(mockedRefunder.DeletePendingRefundCalls())
func (mock *RefunderMock) DeletePendingRefundCalls() []struct {
	Ctx cosmossdktypes.Context
	Req rewardtypes.RefundMsgRequest
} {
	var calls []struct {
		Ctx cosmossdktypes.Context
		Req rewardtypes.RefundMsgRequest
	}
	mock.lockDeletePendingRefund.RLock()
	calls = mock.calls.DeletePendingRefund
	mock.lockDeletePendingRefund.RUnlock()
	return calls
}

// GetPendingRefund calls GetPendingRefundFunc.
func (mock *RefunderMock) GetPendingRefund(ctx cosmossdktypes.Context, req rewardtypes.RefundMsgRequest) (rewardtypes.Refund, bool) {
	if mock.GetPendingRefundFunc == nil {
		panic("RefunderMock.GetPendingRefundFunc: method is nil but Refunder.GetPendingRefund was just called")
	}
	callInfo := struct {
		Ctx cosmossdktypes.Context
		Req rewardtypes.RefundMsgRequest
	}{
		Ctx: ctx,
		Req: req,
	}
	mock.lockGetPendingRefund.Lock()
	mock.calls.GetPendingRefund = append(mock.calls.GetPendingRefund, callInfo)
	mock.lockGetPendingRefund.Unlock()
	return mock.GetPendingRefundFunc(ctx, req)
}

// GetPendingRefundCalls gets all the calls that were made to GetPendingRefund.
// Check the length with:
//     len(mockedRefunder.GetPendingRefundCalls())
func (mock *RefunderMock) GetPendingRefundCalls() []struct {
	Ctx cosmossdktypes.Context
	Req rewardtypes.RefundMsgRequest
} {
	var calls []struct {
		Ctx cosmossdktypes.Context
		Req rewardtypes.RefundMsgRequest
	}
	mock.lockGetPendingRefund.RLock()
	calls = mock.calls.GetPendingRefund
	mock.lockGetPendingRefund.RUnlock()
	return calls
}

// Logger calls LoggerFunc.
func (mock *RefunderMock) Logger(ctx cosmossdktypes.Context) log.Logger {
	if mock.LoggerFunc == nil {
		panic("RefunderMock.LoggerFunc: method is nil but Refunder.Logger was just called")
	}
	callInfo := struct {
		Ctx cosmossdktypes.Context
	}{
		Ctx: ctx,
	}
	mock.lockLogger.Lock()
	mock.calls.Logger = append(mock.calls.Logger, callInfo)
	mock.lockLogger.Unlock()
	return mock.LoggerFunc(ctx)
}

// LoggerCalls gets all the calls that were made to Logger.
// Check the length with:
//     len(mockedRefunder.LoggerCalls())
func (mock *RefunderMock) LoggerCalls() []struct {
	Ctx cosmossdktypes.Context
} {
	var calls []struct {
		Ctx cosmossdktypes.Context
	}
	mock.lockLogger.RLock()
	calls = mock.calls.Logger
	mock.lockLogger.RUnlock()
	return calls
}

// Ensure, that NexusMock does implement rewardtypes.Nexus.
// If this is not the case, regenerate this file with moq.
var _ rewardtypes.Nexus = &NexusMock{}

// NexusMock is a mock implementation of rewardtypes.Nexus.
//
// 	func TestSomethingThatUsesNexus(t *testing.T) {
//
// 		// make and configure a mocked rewardtypes.Nexus
// 		mockedNexus := &NexusMock{
// 			GetChainMaintainersFunc: func(ctx cosmossdktypes.Context, chain nexus.Chain) []cosmossdktypes.ValAddress {
// 				panic("mock out the GetChainMaintainers method")
// 			},
// 			GetChainsFunc: func(ctx cosmossdktypes.Context) []nexus.Chain {
// 				panic("mock out the GetChains method")
// 			},
// 			IsChainActivatedFunc: func(ctx cosmossdktypes.Context, chain nexus.Chain) bool {
// 				panic("mock out the IsChainActivated method")
// 			},
// 		}
//
// 		// use mockedNexus in code that requires rewardtypes.Nexus
// 		// and then make assertions.
//
// 	}
type NexusMock struct {
	// GetChainMaintainersFunc mocks the GetChainMaintainers method.
	GetChainMaintainersFunc func(ctx cosmossdktypes.Context, chain nexus.Chain) []cosmossdktypes.ValAddress

	// GetChainsFunc mocks the GetChains method.
	GetChainsFunc func(ctx cosmossdktypes.Context) []nexus.Chain

	// IsChainActivatedFunc mocks the IsChainActivated method.
	IsChainActivatedFunc func(ctx cosmossdktypes.Context, chain nexus.Chain) bool

	// calls tracks calls to the methods.
	calls struct {
		// GetChainMaintainers holds details about calls to the GetChainMaintainers method.
		GetChainMaintainers []struct {
			// Ctx is the ctx argument value.
			Ctx cosmossdktypes.Context
			// Chain is the chain argument value.
			Chain nexus.Chain
		}
		// GetChains holds details about calls to the GetChains method.
		GetChains []struct {
			// Ctx is the ctx argument value.
			Ctx cosmossdktypes.Context
		}
		// IsChainActivated holds details about calls to the IsChainActivated method.
		IsChainActivated []struct {
			// Ctx is the ctx argument value.
			Ctx cosmossdktypes.Context
			// Chain is the chain argument value.
			Chain nexus.Chain
		}
	}
	lockGetChainMaintainers sync.RWMutex
	lockGetChains           sync.RWMutex
	lockIsChainActivated    sync.RWMutex
}

// GetChainMaintainers calls GetChainMaintainersFunc.
func (mock *NexusMock) GetChainMaintainers(ctx cosmossdktypes.Context, chain nexus.Chain) []cosmossdktypes.ValAddress {
	if mock.GetChainMaintainersFunc == nil {
		panic("NexusMock.GetChainMaintainersFunc: method is nil but Nexus.GetChainMaintainers was just called")
	}
	callInfo := struct {
		Ctx   cosmossdktypes.Context
		Chain nexus.Chain
	}{
		Ctx:   ctx,
		Chain: chain,
	}
	mock.lockGetChainMaintainers.Lock()
	mock.calls.GetChainMaintainers = append(mock.calls.GetChainMaintainers, callInfo)
	mock.lockGetChainMaintainers.Unlock()
	return mock.GetChainMaintainersFunc(ctx, chain)
}

// GetChainMaintainersCalls gets all the calls that were made to GetChainMaintainers.
// Check the length with:
//     len(mockedNexus.GetChainMaintainersCalls())
func (mock *NexusMock) GetChainMaintainersCalls() []struct {
	Ctx   cosmossdktypes.Context
	Chain nexus.Chain
} {
	var calls []struct {
		Ctx   cosmossdktypes.Context
		Chain nexus.Chain
	}
	mock.lockGetChainMaintainers.RLock()
	calls = mock.calls.GetChainMaintainers
	mock.lockGetChainMaintainers.RUnlock()
	return calls
}

// GetChains calls GetChainsFunc.
func (mock *NexusMock) GetChains(ctx cosmossdktypes.Context) []nexus.Chain {
	if mock.GetChainsFunc == nil {
		panic("NexusMock.GetChainsFunc: method is nil but Nexus.GetChains was just called")
	}
	callInfo := struct {
		Ctx cosmossdktypes.Context
	}{
		Ctx: ctx,
	}
	mock.lockGetChains.Lock()
	mock.calls.GetChains = append(mock.calls.GetChains, callInfo)
	mock.lockGetChains.Unlock()
	return mock.GetChainsFunc(ctx)
}

// GetChainsCalls gets all the calls that were made to GetChains.
// Check the length with:
//     len(mockedNexus.GetChainsCalls())
func (mock *NexusMock) GetChainsCalls() []struct {
	Ctx cosmossdktypes.Context
} {
	var calls []struct {
		Ctx cosmossdktypes.Context
	}
	mock.lockGetChains.RLock()
	calls = mock.calls.GetChains
	mock.lockGetChains.RUnlock()
	return calls
}

// IsChainActivated calls IsChainActivatedFunc.
func (mock *NexusMock) IsChainActivated(ctx cosmossdktypes.Context, chain nexus.Chain) bool {
	if mock.IsChainActivatedFunc == nil {
		panic("NexusMock.IsChainActivatedFunc: method is nil but Nexus.IsChainActivated was just called")
	}
	callInfo := struct {
		Ctx   cosmossdktypes.Context
		Chain nexus.Chain
	}{
		Ctx:   ctx,
		Chain: chain,
	}
	mock.lockIsChainActivated.Lock()
	mock.calls.IsChainActivated = append(mock.calls.IsChainActivated, callInfo)
	mock.lockIsChainActivated.Unlock()
	return mock.IsChainActivatedFunc(ctx, chain)
}

// IsChainActivatedCalls gets all the calls that were made to IsChainActivated.
// Check the length with:
//     len(mockedNexus.IsChainActivatedCalls())
func (mock *NexusMock) IsChainActivatedCalls() []struct {
	Ctx   cosmossdktypes.Context
	Chain nexus.Chain
} {
	var calls []struct {
		Ctx   cosmossdktypes.Context
		Chain nexus.Chain
	}
	mock.lockIsChainActivated.RLock()
	calls = mock.calls.IsChainActivated
	mock.lockIsChainActivated.RUnlock()
	return calls
}

// Ensure, that MinterMock does implement rewardtypes.Minter.
// If this is not the case, regenerate this file with moq.
var _ rewardtypes.Minter = &MinterMock{}

// MinterMock is a mock implementation of rewardtypes.Minter.
//
// 	func TestSomethingThatUsesMinter(t *testing.T) {
//
// 		// make and configure a mocked rewardtypes.Minter
// 		mockedMinter := &MinterMock{
// 			GetMinterFunc: func(ctx cosmossdktypes.Context) minttypes.Minter {
// 				panic("mock out the GetMinter method")
// 			},
// 			GetParamsFunc: func(ctx cosmossdktypes.Context) minttypes.Params {
// 				panic("mock out the GetParams method")
// 			},
// 			StakingTokenSupplyFunc: func(ctx cosmossdktypes.Context) cosmossdktypes.Int {
// 				panic("mock out the StakingTokenSupply method")
// 			},
// 		}
//
// 		// use mockedMinter in code that requires rewardtypes.Minter
// 		// and then make assertions.
//
// 	}
type MinterMock struct {
	// GetMinterFunc mocks the GetMinter method.
	GetMinterFunc func(ctx cosmossdktypes.Context) minttypes.Minter

	// GetParamsFunc mocks the GetParams method.
	GetParamsFunc func(ctx cosmossdktypes.Context) minttypes.Params

	// StakingTokenSupplyFunc mocks the StakingTokenSupply method.
	StakingTokenSupplyFunc func(ctx cosmossdktypes.Context) cosmossdktypes.Int

	// calls tracks calls to the methods.
	calls struct {
		// GetMinter holds details about calls to the GetMinter method.
		GetMinter []struct {
			// Ctx is the ctx argument value.
			Ctx cosmossdktypes.Context
		}
		// GetParams holds details about calls to the GetParams method.
		GetParams []struct {
			// Ctx is the ctx argument value.
			Ctx cosmossdktypes.Context
		}
		// StakingTokenSupply holds details about calls to the StakingTokenSupply method.
		StakingTokenSupply []struct {
			// Ctx is the ctx argument value.
			Ctx cosmossdktypes.Context
		}
	}
	lockGetMinter          sync.RWMutex
	lockGetParams          sync.RWMutex
	lockStakingTokenSupply sync.RWMutex
}

// GetMinter calls GetMinterFunc.
func (mock *MinterMock) GetMinter(ctx cosmossdktypes.Context) minttypes.Minter {
	if mock.GetMinterFunc == nil {
		panic("MinterMock.GetMinterFunc: method is nil but Minter.GetMinter was just called")
	}
	callInfo := struct {
		Ctx cosmossdktypes.Context
	}{
		Ctx: ctx,
	}
	mock.lockGetMinter.Lock()
	mock.calls.GetMinter = append(mock.calls.GetMinter, callInfo)
	mock.lockGetMinter.Unlock()
	return mock.GetMinterFunc(ctx)
}

// GetMinterCalls gets all the calls that were made to GetMinter.
// Check the length with:
//     len(mockedMinter.GetMinterCalls())
func (mock *MinterMock) GetMinterCalls() []struct {
	Ctx cosmossdktypes.Context
} {
	var calls []struct {
		Ctx cosmossdktypes.Context
	}
	mock.lockGetMinter.RLock()
	calls = mock.calls.GetMinter
	mock.lockGetMinter.RUnlock()
	return calls
}

// GetParams calls GetParamsFunc.
func (mock *MinterMock) GetParams(ctx cosmossdktypes.Context) minttypes.Params {
	if mock.GetParamsFunc == nil {
		panic("MinterMock.GetParamsFunc: method is nil but Minter.GetParams was just called")
	}
	callInfo := struct {
		Ctx cosmossdktypes.Context
	}{
		Ctx: ctx,
	}
	mock.lockGetParams.Lock()
	mock.calls.GetParams = append(mock.calls.GetParams, callInfo)
	mock.lockGetParams.Unlock()
	return mock.GetParamsFunc(ctx)
}

// GetParamsCalls gets all the calls that were made to GetParams.
// Check the length with:
//     len(mockedMinter.GetParamsCalls())
func (mock *MinterMock) GetParamsCalls() []struct {
	Ctx cosmossdktypes.Context
} {
	var calls []struct {
		Ctx cosmossdktypes.Context
	}
	mock.lockGetParams.RLock()
	calls = mock.calls.GetParams
	mock.lockGetParams.RUnlock()
	return calls
}

// StakingTokenSupply calls StakingTokenSupplyFunc.
func (mock *MinterMock) StakingTokenSupply(ctx cosmossdktypes.Context) cosmossdktypes.Int {
	if mock.StakingTokenSupplyFunc == nil {
		panic("MinterMock.StakingTokenSupplyFunc: method is nil but Minter.StakingTokenSupply was just called")
	}
	callInfo := struct {
		Ctx cosmossdktypes.Context
	}{
		Ctx: ctx,
	}
	mock.lockStakingTokenSupply.Lock()
	mock.calls.StakingTokenSupply = append(mock.calls.StakingTokenSupply, callInfo)
	mock.lockStakingTokenSupply.Unlock()
	return mock.StakingTokenSupplyFunc(ctx)
}

// StakingTokenSupplyCalls gets all the calls that were made to StakingTokenSupply.
// Check the length with:
//     len(mockedMinter.StakingTokenSupplyCalls())
func (mock *MinterMock) StakingTokenSupplyCalls() []struct {
	Ctx cosmossdktypes.Context
} {
	var calls []struct {
		Ctx cosmossdktypes.Context
	}
	mock.lockStakingTokenSupply.RLock()
	calls = mock.calls.StakingTokenSupply
	mock.lockStakingTokenSupply.RUnlock()
	return calls
}

// Ensure, that DistributorMock does implement rewardtypes.Distributor.
// If this is not the case, regenerate this file with moq.
var _ rewardtypes.Distributor = &DistributorMock{}

// DistributorMock is a mock implementation of rewardtypes.Distributor.
//
// 	func TestSomethingThatUsesDistributor(t *testing.T) {
//
// 		// make and configure a mocked rewardtypes.Distributor
// 		mockedDistributor := &DistributorMock{
// 			AllocateTokensToValidatorFunc: func(ctx cosmossdktypes.Context, val stakingtypes.ValidatorI, tokens cosmossdktypes.DecCoins)  {
// 				panic("mock out the AllocateTokensToValidator method")
// 			},
// 		}
//
// 		// use mockedDistributor in code that requires rewardtypes.Distributor
// 		// and then make assertions.
//
// 	}
type DistributorMock struct {
	// AllocateTokensToValidatorFunc mocks the AllocateTokensToValidator method.
	AllocateTokensToValidatorFunc func(ctx cosmossdktypes.Context, val stakingtypes.ValidatorI, tokens cosmossdktypes.DecCoins)

	// calls tracks calls to the methods.
	calls struct {
		// AllocateTokensToValidator holds details about calls to the AllocateTokensToValidator method.
		AllocateTokensToValidator []struct {
			// Ctx is the ctx argument value.
			Ctx cosmossdktypes.Context
			// Val is the val argument value.
			Val stakingtypes.ValidatorI
			// Tokens is the tokens argument value.
			Tokens cosmossdktypes.DecCoins
		}
	}
	lockAllocateTokensToValidator sync.RWMutex
}

// AllocateTokensToValidator calls AllocateTokensToValidatorFunc.
func (mock *DistributorMock) AllocateTokensToValidator(ctx cosmossdktypes.Context, val stakingtypes.ValidatorI, tokens cosmossdktypes.DecCoins) {
	if mock.AllocateTokensToValidatorFunc == nil {
		panic("DistributorMock.AllocateTokensToValidatorFunc: method is nil but Distributor.AllocateTokensToValidator was just called")
	}
	callInfo := struct {
		Ctx    cosmossdktypes.Context
		Val    stakingtypes.ValidatorI
		Tokens cosmossdktypes.DecCoins
	}{
		Ctx:    ctx,
		Val:    val,
		Tokens: tokens,
	}
	mock.lockAllocateTokensToValidator.Lock()
	mock.calls.AllocateTokensToValidator = append(mock.calls.AllocateTokensToValidator, callInfo)
	mock.lockAllocateTokensToValidator.Unlock()
	mock.AllocateTokensToValidatorFunc(ctx, val, tokens)
}

// AllocateTokensToValidatorCalls gets all the calls that were made to AllocateTokensToValidator.
// Check the length with:
//     len(mockedDistributor.AllocateTokensToValidatorCalls())
func (mock *DistributorMock) AllocateTokensToValidatorCalls() []struct {
	Ctx    cosmossdktypes.Context
	Val    stakingtypes.ValidatorI
	Tokens cosmossdktypes.DecCoins
} {
	var calls []struct {
		Ctx    cosmossdktypes.Context
		Val    stakingtypes.ValidatorI
		Tokens cosmossdktypes.DecCoins
	}
	mock.lockAllocateTokensToValidator.RLock()
	calls = mock.calls.AllocateTokensToValidator
	mock.lockAllocateTokensToValidator.RUnlock()
	return calls
}

// Ensure, that StakerMock does implement rewardtypes.Staker.
// If this is not the case, regenerate this file with moq.
var _ rewardtypes.Staker = &StakerMock{}

// StakerMock is a mock implementation of rewardtypes.Staker.
//
// 	func TestSomethingThatUsesStaker(t *testing.T) {
//
// 		// make and configure a mocked rewardtypes.Staker
// 		mockedStaker := &StakerMock{
// 			IterateBondedValidatorsByPowerFunc: func(ctx cosmossdktypes.Context, fn func(index int64, validator stakingtypes.ValidatorI) (stop bool))  {
// 				panic("mock out the IterateBondedValidatorsByPower method")
// 			},
// 			PowerReductionFunc: func(ctx cosmossdktypes.Context) cosmossdktypes.Int {
// 				panic("mock out the PowerReduction method")
// 			},
// 			ValidatorFunc: func(ctx cosmossdktypes.Context, addr cosmossdktypes.ValAddress) stakingtypes.ValidatorI {
// 				panic("mock out the Validator method")
// 			},
// 		}
//
// 		// use mockedStaker in code that requires rewardtypes.Staker
// 		// and then make assertions.
//
// 	}
type StakerMock struct {
	// IterateBondedValidatorsByPowerFunc mocks the IterateBondedValidatorsByPower method.
	IterateBondedValidatorsByPowerFunc func(ctx cosmossdktypes.Context, fn func(index int64, validator stakingtypes.ValidatorI) (stop bool))

	// PowerReductionFunc mocks the PowerReduction method.
	PowerReductionFunc func(ctx cosmossdktypes.Context) cosmossdktypes.Int

	// ValidatorFunc mocks the Validator method.
	ValidatorFunc func(ctx cosmossdktypes.Context, addr cosmossdktypes.ValAddress) stakingtypes.ValidatorI

	// calls tracks calls to the methods.
	calls struct {
		// IterateBondedValidatorsByPower holds details about calls to the IterateBondedValidatorsByPower method.
		IterateBondedValidatorsByPower []struct {
			// Ctx is the ctx argument value.
			Ctx cosmossdktypes.Context
			// Fn is the fn argument value.
			Fn func(index int64, validator stakingtypes.ValidatorI) (stop bool)
		}
		// PowerReduction holds details about calls to the PowerReduction method.
		PowerReduction []struct {
			// Ctx is the ctx argument value.
			Ctx cosmossdktypes.Context
		}
		// Validator holds details about calls to the Validator method.
		Validator []struct {
			// Ctx is the ctx argument value.
			Ctx cosmossdktypes.Context
			// Addr is the addr argument value.
			Addr cosmossdktypes.ValAddress
		}
	}
	lockIterateBondedValidatorsByPower sync.RWMutex
	lockPowerReduction                 sync.RWMutex
	lockValidator                      sync.RWMutex
}

// IterateBondedValidatorsByPower calls IterateBondedValidatorsByPowerFunc.
func (mock *StakerMock) IterateBondedValidatorsByPower(ctx cosmossdktypes.Context, fn func(index int64, validator stakingtypes.ValidatorI) (stop bool)) {
	if mock.IterateBondedValidatorsByPowerFunc == nil {
		panic("StakerMock.IterateBondedValidatorsByPowerFunc: method is nil but Staker.IterateBondedValidatorsByPower was just called")
	}
	callInfo := struct {
		Ctx cosmossdktypes.Context
		Fn  func(index int64, validator stakingtypes.ValidatorI) (stop bool)
	}{
		Ctx: ctx,
		Fn:  fn,
	}
	mock.lockIterateBondedValidatorsByPower.Lock()
	mock.calls.IterateBondedValidatorsByPower = append(mock.calls.IterateBondedValidatorsByPower, callInfo)
	mock.lockIterateBondedValidatorsByPower.Unlock()
	mock.IterateBondedValidatorsByPowerFunc(ctx, fn)
}

// IterateBondedValidatorsByPowerCalls gets all the calls that were made to IterateBondedValidatorsByPower.
// Check the length with:
//     len(mockedStaker.IterateBondedValidatorsByPowerCalls())
func (mock *StakerMock) IterateBondedValidatorsByPowerCalls() []struct {
	Ctx cosmossdktypes.Context
	Fn  func(index int64, validator stakingtypes.ValidatorI) (stop bool)
} {
	var calls []struct {
		Ctx cosmossdktypes.Context
		Fn  func(index int64, validator stakingtypes.ValidatorI) (stop bool)
	}
	mock.lockIterateBondedValidatorsByPower.RLock()
	calls = mock.calls.IterateBondedValidatorsByPower
	mock.lockIterateBondedValidatorsByPower.RUnlock()
	return calls
}

// PowerReduction calls PowerReductionFunc.
func (mock *StakerMock) PowerReduction(ctx cosmossdktypes.Context) cosmossdktypes.Int {
	if mock.PowerReductionFunc == nil {
		panic("StakerMock.PowerReductionFunc: method is nil but Staker.PowerReduction was just called")
	}
	callInfo := struct {
		Ctx cosmossdktypes.Context
	}{
		Ctx: ctx,
	}
	mock.lockPowerReduction.Lock()
	mock.calls.PowerReduction = append(mock.calls.PowerReduction, callInfo)
	mock.lockPowerReduction.Unlock()
	return mock.PowerReductionFunc(ctx)
}

// PowerReductionCalls gets all the calls that were made to PowerReduction.
// Check the length with:
//     len(mockedStaker.PowerReductionCalls())
func (mock *StakerMock) PowerReductionCalls() []struct {
	Ctx cosmossdktypes.Context
} {
	var calls []struct {
		Ctx cosmossdktypes.Context
	}
	mock.lockPowerReduction.RLock()
	calls = mock.calls.PowerReduction
	mock.lockPowerReduction.RUnlock()
	return calls
}

// Validator calls ValidatorFunc.
func (mock *StakerMock) Validator(ctx cosmossdktypes.Context, addr cosmossdktypes.ValAddress) stakingtypes.ValidatorI {
	if mock.ValidatorFunc == nil {
		panic("StakerMock.ValidatorFunc: method is nil but Staker.Validator was just called")
	}
	callInfo := struct {
		Ctx  cosmossdktypes.Context
		Addr cosmossdktypes.ValAddress
	}{
		Ctx:  ctx,
		Addr: addr,
	}
	mock.lockValidator.Lock()
	mock.calls.Validator = append(mock.calls.Validator, callInfo)
	mock.lockValidator.Unlock()
	return mock.ValidatorFunc(ctx, addr)
}

// ValidatorCalls gets all the calls that were made to Validator.
// Check the length with:
//     len(mockedStaker.ValidatorCalls())
func (mock *StakerMock) ValidatorCalls() []struct {
	Ctx  cosmossdktypes.Context
	Addr cosmossdktypes.ValAddress
} {
	var calls []struct {
		Ctx  cosmossdktypes.Context
		Addr cosmossdktypes.ValAddress
	}
	mock.lockValidator.RLock()
	calls = mock.calls.Validator
	mock.lockValidator.RUnlock()
	return calls
}

// Ensure, that BankerMock does implement rewardtypes.Banker.
// If this is not the case, regenerate this file with moq.
var _ rewardtypes.Banker = &BankerMock{}

// BankerMock is a mock implementation of rewardtypes.Banker.
//
// 	func TestSomethingThatUsesBanker(t *testing.T) {
//
// 		// make and configure a mocked rewardtypes.Banker
// 		mockedBanker := &BankerMock{
// 			MintCoinsFunc: func(ctx cosmossdktypes.Context, name string, amt cosmossdktypes.Coins) error {
// 				panic("mock out the MintCoins method")
// 			},
// 			SendCoinsFromModuleToAccountFunc: func(ctx cosmossdktypes.Context, senderModule string, recipientAddr cosmossdktypes.AccAddress, amt cosmossdktypes.Coins) error {
// 				panic("mock out the SendCoinsFromModuleToAccount method")
// 			},
// 			SendCoinsFromModuleToModuleFunc: func(ctx cosmossdktypes.Context, senderModule string, recipientModule string, amt cosmossdktypes.Coins) error {
// 				panic("mock out the SendCoinsFromModuleToModule method")
// 			},
// 		}
//
// 		// use mockedBanker in code that requires rewardtypes.Banker
// 		// and then make assertions.
//
// 	}
type BankerMock struct {
	// MintCoinsFunc mocks the MintCoins method.
	MintCoinsFunc func(ctx cosmossdktypes.Context, name string, amt cosmossdktypes.Coins) error

	// SendCoinsFromModuleToAccountFunc mocks the SendCoinsFromModuleToAccount method.
	SendCoinsFromModuleToAccountFunc func(ctx cosmossdktypes.Context, senderModule string, recipientAddr cosmossdktypes.AccAddress, amt cosmossdktypes.Coins) error

	// SendCoinsFromModuleToModuleFunc mocks the SendCoinsFromModuleToModule method.
	SendCoinsFromModuleToModuleFunc func(ctx cosmossdktypes.Context, senderModule string, recipientModule string, amt cosmossdktypes.Coins) error

	// calls tracks calls to the methods.
	calls struct {
		// MintCoins holds details about calls to the MintCoins method.
		MintCoins []struct {
			// Ctx is the ctx argument value.
			Ctx cosmossdktypes.Context
			// Name is the name argument value.
			Name string
			// Amt is the amt argument value.
			Amt cosmossdktypes.Coins
		}
		// SendCoinsFromModuleToAccount holds details about calls to the SendCoinsFromModuleToAccount method.
		SendCoinsFromModuleToAccount []struct {
			// Ctx is the ctx argument value.
			Ctx cosmossdktypes.Context
			// SenderModule is the senderModule argument value.
			SenderModule string
			// RecipientAddr is the recipientAddr argument value.
			RecipientAddr cosmossdktypes.AccAddress
			// Amt is the amt argument value.
			Amt cosmossdktypes.Coins
		}
		// SendCoinsFromModuleToModule holds details about calls to the SendCoinsFromModuleToModule method.
		SendCoinsFromModuleToModule []struct {
			// Ctx is the ctx argument value.
			Ctx cosmossdktypes.Context
			// SenderModule is the senderModule argument value.
			SenderModule string
			// RecipientModule is the recipientModule argument value.
			RecipientModule string
			// Amt is the amt argument value.
			Amt cosmossdktypes.Coins
		}
	}
	lockMintCoins                    sync.RWMutex
	lockSendCoinsFromModuleToAccount sync.RWMutex
	lockSendCoinsFromModuleToModule  sync.RWMutex
}

// MintCoins calls MintCoinsFunc.
func (mock *BankerMock) MintCoins(ctx cosmossdktypes.Context, name string, amt cosmossdktypes.Coins) error {
	if mock.MintCoinsFunc == nil {
		panic("BankerMock.MintCoinsFunc: method is nil but Banker.MintCoins was just called")
	}
	callInfo := struct {
		Ctx  cosmossdktypes.Context
		Name string
		Amt  cosmossdktypes.Coins
	}{
		Ctx:  ctx,
		Name: name,
		Amt:  amt,
	}
	mock.lockMintCoins.Lock()
	mock.calls.MintCoins = append(mock.calls.MintCoins, callInfo)
	mock.lockMintCoins.Unlock()
	return mock.MintCoinsFunc(ctx, name, amt)
}

// MintCoinsCalls gets all the calls that were made to MintCoins.
// Check the length with:
//     len(mockedBanker.MintCoinsCalls())
func (mock *BankerMock) MintCoinsCalls() []struct {
	Ctx  cosmossdktypes.Context
	Name string
	Amt  cosmossdktypes.Coins
} {
	var calls []struct {
		Ctx  cosmossdktypes.Context
		Name string
		Amt  cosmossdktypes.Coins
	}
	mock.lockMintCoins.RLock()
	calls = mock.calls.MintCoins
	mock.lockMintCoins.RUnlock()
	return calls
}

// SendCoinsFromModuleToAccount calls SendCoinsFromModuleToAccountFunc.
func (mock *BankerMock) SendCoinsFromModuleToAccount(ctx cosmossdktypes.Context, senderModule string, recipientAddr cosmossdktypes.AccAddress, amt cosmossdktypes.Coins) error {
	if mock.SendCoinsFromModuleToAccountFunc == nil {
		panic("BankerMock.SendCoinsFromModuleToAccountFunc: method is nil but Banker.SendCoinsFromModuleToAccount was just called")
	}
	callInfo := struct {
		Ctx           cosmossdktypes.Context
		SenderModule  string
		RecipientAddr cosmossdktypes.AccAddress
		Amt           cosmossdktypes.Coins
	}{
		Ctx:           ctx,
		SenderModule:  senderModule,
		RecipientAddr: recipientAddr,
		Amt:           amt,
	}
	mock.lockSendCoinsFromModuleToAccount.Lock()
	mock.calls.SendCoinsFromModuleToAccount = append(mock.calls.SendCoinsFromModuleToAccount, callInfo)
	mock.lockSendCoinsFromModuleToAccount.Unlock()
	return mock.SendCoinsFromModuleToAccountFunc(ctx, senderModule, recipientAddr, amt)
}

// SendCoinsFromModuleToAccountCalls gets all the calls that were made to SendCoinsFromModuleToAccount.
// Check the length with:
//     len(mockedBanker.SendCoinsFromModuleToAccountCalls())
func (mock *BankerMock) SendCoinsFromModuleToAccountCalls() []struct {
	Ctx           cosmossdktypes.Context
	SenderModule  string
	RecipientAddr cosmossdktypes.AccAddress
	Amt           cosmossdktypes.Coins
} {
	var calls []struct {
		Ctx           cosmossdktypes.Context
		SenderModule  string
		RecipientAddr cosmossdktypes.AccAddress
		Amt           cosmossdktypes.Coins
	}
	mock.lockSendCoinsFromModuleToAccount.RLock()
	calls = mock.calls.SendCoinsFromModuleToAccount
	mock.lockSendCoinsFromModuleToAccount.RUnlock()
	return calls
}

// SendCoinsFromModuleToModule calls SendCoinsFromModuleToModuleFunc.
func (mock *BankerMock) SendCoinsFromModuleToModule(ctx cosmossdktypes.Context, senderModule string, recipientModule string, amt cosmossdktypes.Coins) error {
	if mock.SendCoinsFromModuleToModuleFunc == nil {
		panic("BankerMock.SendCoinsFromModuleToModuleFunc: method is nil but Banker.SendCoinsFromModuleToModule was just called")
	}
	callInfo := struct {
		Ctx             cosmossdktypes.Context
		SenderModule    string
		RecipientModule string
		Amt             cosmossdktypes.Coins
	}{
		Ctx:             ctx,
		SenderModule:    senderModule,
		RecipientModule: recipientModule,
		Amt:             amt,
	}
	mock.lockSendCoinsFromModuleToModule.Lock()
	mock.calls.SendCoinsFromModuleToModule = append(mock.calls.SendCoinsFromModuleToModule, callInfo)
	mock.lockSendCoinsFromModuleToModule.Unlock()
	return mock.SendCoinsFromModuleToModuleFunc(ctx, senderModule, recipientModule, amt)
}

// SendCoinsFromModuleToModuleCalls gets all the calls that were made to SendCoinsFromModuleToModule.
// Check the length with:
//     len(mockedBanker.SendCoinsFromModuleToModuleCalls())
func (mock *BankerMock) SendCoinsFromModuleToModuleCalls() []struct {
	Ctx             cosmossdktypes.Context
	SenderModule    string
	RecipientModule string
	Amt             cosmossdktypes.Coins
} {
	var calls []struct {
		Ctx             cosmossdktypes.Context
		SenderModule    string
		RecipientModule string
		Amt             cosmossdktypes.Coins
	}
	mock.lockSendCoinsFromModuleToModule.RLock()
	calls = mock.calls.SendCoinsFromModuleToModule
	mock.lockSendCoinsFromModuleToModule.RUnlock()
	return calls
}
