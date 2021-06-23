package evm

import (
	"context"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	geth "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	tmLog "github.com/tendermint/tendermint/libs/log"

	"github.com/axelarnetwork/axelar-core/cmd/axelard/cmd/vald/broadcast/types"
	"github.com/axelarnetwork/axelar-core/cmd/axelard/cmd/vald/evm/rpc"
	btc "github.com/axelarnetwork/axelar-core/x/bitcoin/types"
	evmTypes "github.com/axelarnetwork/axelar-core/x/evm/types"
	vote "github.com/axelarnetwork/axelar-core/x/vote/exported"
)

// Smart contract event signatures
var (
	ERC20TransferSig        = crypto.Keccak256Hash([]byte("Transfer(address,address,uint256)"))
	ERC20TokenDeploymentSig = crypto.Keccak256Hash([]byte("TokenDeployed(string,address)"))
)

// Mgr manages all communication with Ethereum
type Mgr struct {
	logger      tmLog.Logger
	rpcs        map[string]rpc.Client
	broadcaster types.Broadcaster
	sender      sdk.AccAddress
	cdc         *codec.LegacyAmino
}

// NewMgr returns a new Mgr instance
func NewMgr(rpcs map[string]rpc.Client, broadcaster types.Broadcaster, sender sdk.AccAddress, logger tmLog.Logger, cdc *codec.LegacyAmino) *Mgr {
	return &Mgr{
		rpcs:        rpcs,
		broadcaster: broadcaster,
		sender:      sender,
		logger:      logger.With("listener", "evm"),
		cdc:         cdc,
	}
}

// ProcessNewChain notifies the operator that vald needs to be restarted/udpated for a new chain
func (mgr Mgr) ProcessNewChain(attributes []sdk.Attribute) (err error) {
	chain, nativeAsset, err := parseNewChainParams(attributes)
	if err != nil {
		return sdkerrors.Wrap(err, "Invalid update event")
	}

	mgr.logger.Info(fmt.Sprintf("VALD needs to be updated and restarted for new chain %s with native asset %s", chain, nativeAsset))
	return nil
}

// ProcessChainConfirmation votes on the correctness of an EVM chain token deposit
func (mgr Mgr) ProcessChainConfirmation(attributes []sdk.Attribute) (err error) {
	chain, poll, err := parseChainConfirmationParams(mgr.cdc, attributes)
	if err != nil {
		return sdkerrors.Wrap(err, "EVM chain confirmation failed")
	}

	_, confirmed := mgr.rpcs[strings.ToLower(chain)]

	msg := evmTypes.NewVoteConfirmChainRequest(mgr.sender, chain, poll, confirmed)
	mgr.logger.Debug(fmt.Sprintf("broadcasting vote %v for poll %s", msg.Confirmed, poll.String()))
	return mgr.broadcaster.Broadcast(msg)
}

// ProcessDepositConfirmation votes on the correctness of an EVM chain token deposit
func (mgr Mgr) ProcessDepositConfirmation(attributes []sdk.Attribute) (err error) {
	chain, txID, amount, burnAddr, tokenAddr, confHeight, poll, err := parseDepositConfirmationParams(mgr.cdc, attributes)
	if err != nil {
		return sdkerrors.Wrap(err, "EVM deposit confirmation failed")
	}

	rpc, found := mgr.rpcs[strings.ToLower(chain)]
	if !found {
		return sdkerrors.Wrap(err, fmt.Sprintf("Unable to find an RPC for chain '%s'", chain))
	}

	confirmed := mgr.validate(rpc, txID, confHeight, func(txReceipt *geth.Receipt) bool {
		err = confirmERC20Deposit(txReceipt, amount, burnAddr, tokenAddr)
		if err != nil {
			mgr.logger.Debug(sdkerrors.Wrap(err, "deposit confirmation failed").Error())
			return false
		}
		return true
	})

	msg := evmTypes.NewVoteConfirmDepositRequest(mgr.sender, chain, poll, txID, evmTypes.Address(burnAddr), confirmed)
	mgr.logger.Debug(fmt.Sprintf("broadcasting vote %v for poll %s", msg.Confirmed, poll.String()))
	return mgr.broadcaster.Broadcast(msg)
}

// ProcessTokenConfirmation votes on the correctness of an EVM chain token deployment
func (mgr Mgr) ProcessTokenConfirmation(attributes []sdk.Attribute) error {
	chain, txID, gatewayAddr, tokenAddr, symbol, confHeight, poll, err := parseTokenConfirmationParams(mgr.cdc, attributes)
	if err != nil {
		return sdkerrors.Wrap(err, "EVM token deployment confirmation failed")
	}

	rpc, found := mgr.rpcs[strings.ToLower(chain)]
	if !found {
		return sdkerrors.Wrap(err, fmt.Sprintf("Unable to find an RPC for chain '%s'", chain))
	}

	confirmed := mgr.validate(rpc, txID, confHeight, func(txReceipt *geth.Receipt) bool {
		err = confirmERC20TokenDeployment(txReceipt, symbol, gatewayAddr, tokenAddr)
		if err != nil {
			mgr.logger.Debug(sdkerrors.Wrap(err, "token confirmation failed").Error())
			return false
		}
		return true
	})

	msg := evmTypes.NewVoteConfirmTokenRequest(mgr.sender, chain, symbol, poll, txID, confirmed)
	mgr.logger.Debug(fmt.Sprintf("broadcasting vote %v for poll %s", msg.Confirmed, poll.String()))
	return mgr.broadcaster.Broadcast(msg)
}

func parseNewChainParams(attributes []sdk.Attribute) (
	chain string,
	nativeAsset string,
	err error,
) {
	var chainFound, nativeAssetFound bool
	for _, attribute := range attributes {
		switch attribute.Key {
		case evmTypes.AttributeKeyChain:
			chain = attribute.Value
			chainFound = true
		case evmTypes.AttributeKeyNativeAsset:
			nativeAsset = attribute.Value
			nativeAssetFound = true
		default:
		}
	}
	if !chainFound || !nativeAssetFound {
		return "", "", fmt.Errorf("insufficient event attributes")
	}
	return chain, nativeAsset, nil
}

func parseChainConfirmationParams(cdc *codec.LegacyAmino, attributes []sdk.Attribute) (
	chain string,
	poll vote.PollMeta,
	err error,
) {
	var chainFound, pollFound bool
	for _, attribute := range attributes {
		switch attribute.Key {
		case evmTypes.AttributeKeyChain:
			chain = attribute.Value
			chainFound = true
		case evmTypes.AttributeKeyPoll:
			cdc.MustUnmarshalJSON([]byte(attribute.Value), &poll)
			pollFound = true
		default:
		}
	}
	if !chainFound || !pollFound {
		return "", vote.PollMeta{}, fmt.Errorf("insufficient event attributes")
	}
	return chain, poll, nil
}

func parseDepositConfirmationParams(cdc *codec.LegacyAmino, attributes []sdk.Attribute) (
	chain string,
	txID common.Hash,
	amount sdk.Uint,
	burnAddr, tokenAddr common.Address,
	confHeight uint64,
	poll vote.PollMeta,
	err error,
) {
	var chainFound, txIDFound, amountFound, burnAddrFound, tokenAddrFound, confHeightFound, pollFound bool
	for _, attribute := range attributes {
		switch attribute.Key {
		case evmTypes.AttributeKeyChain:
			chain = attribute.Value
			chainFound = true
		case evmTypes.AttributeKeyTxID:
			txID = common.HexToHash(attribute.Value)
			txIDFound = true
		case evmTypes.AttributeKeyAmount:
			amount, err = sdk.ParseUint(attribute.Value)
			if err != nil {
				return "", common.Hash{}, sdk.Uint{}, common.Address{}, common.Address{}, 0, vote.PollMeta{},
					sdkerrors.Wrap(err, "parsing transfer amount failed")
			}
			amountFound = true
		case evmTypes.AttributeKeyBurnAddress:
			burnAddr = common.HexToAddress(attribute.Value)
			burnAddrFound = true
		case evmTypes.AttributeKeyTokenAddress:
			tokenAddr = common.HexToAddress(attribute.Value)
			tokenAddrFound = true
		case evmTypes.AttributeKeyConfHeight:
			confHeight, err = strconv.ParseUint(attribute.Value, 10, 64)
			if err != nil {
				return "", common.Hash{}, sdk.Uint{}, common.Address{}, common.Address{}, 0, vote.PollMeta{},
					sdkerrors.Wrap(err, "parsing confirmation height failed")
			}
			confHeightFound = true
		case evmTypes.AttributeKeyPoll:
			cdc.MustUnmarshalJSON([]byte(attribute.Value), &poll)
			pollFound = true
		default:
		}
	}
	if !chainFound || !txIDFound || !amountFound || !burnAddrFound || !tokenAddrFound || !confHeightFound || !pollFound {
		return "", common.Hash{}, sdk.Uint{}, common.Address{}, common.Address{}, 0, vote.PollMeta{},
			fmt.Errorf("insufficient event attributes")
	}
	return chain, txID, amount, burnAddr, tokenAddr, confHeight, poll, nil
}

func parseTokenConfirmationParams(cdc *codec.LegacyAmino, attributes []sdk.Attribute) (
	chain string,
	txID common.Hash,
	gatewayAddr, tokenAddr common.Address,
	symbol string,
	confHeight uint64,
	poll vote.PollMeta,
	err error,
) {
	var chainFound, txIDFound, gatewayAddrFound, tokenAddrFound, symbolFound, confHeightFound, pollFound bool
	for _, attribute := range attributes {
		switch attribute.Key {
		case evmTypes.AttributeKeyChain:
			chain = attribute.Value
			chainFound = true
		case evmTypes.AttributeKeyTxID:
			txID = common.HexToHash(attribute.Value)
			txIDFound = true
		case evmTypes.AttributeKeyGatewayAddress:
			gatewayAddr = common.HexToAddress(attribute.Value)
			gatewayAddrFound = true
		case evmTypes.AttributeKeyTokenAddress:
			tokenAddr = common.HexToAddress(attribute.Value)
			tokenAddrFound = true
		case evmTypes.AttributeKeySymbol:
			symbol = attribute.Value
			symbolFound = true
		case evmTypes.AttributeKeyConfHeight:
			h, err := strconv.Atoi(attribute.Value)
			if err != nil {
				return "", common.Hash{}, common.Address{}, common.Address{}, "", 0, vote.PollMeta{},
					sdkerrors.Wrap(err, "parsing confirmation height failed")
			}
			confHeight = uint64(h)
			confHeightFound = true
		case btc.AttributeKeyPoll:
			cdc.MustUnmarshalJSON([]byte(attribute.Value), &poll)
			pollFound = true
		default:
		}
	}
	if !chainFound || !txIDFound || !gatewayAddrFound || !tokenAddrFound || !symbolFound || !confHeightFound || !pollFound {
		return "", common.Hash{}, common.Address{}, common.Address{}, "", 0, vote.PollMeta{},
			fmt.Errorf("insufficient event attributes")
	}
	return chain, txID, gatewayAddr, tokenAddr, symbol, confHeight, poll, nil
}

func (mgr Mgr) validate(rpc rpc.Client, txID common.Hash, confHeight uint64, validateLogs func(txReceipt *geth.Receipt) bool) bool {
	blockNumber, err := rpc.BlockNumber(context.Background())
	if err != nil {
		mgr.logger.Debug(sdkerrors.Wrap(err, "checking block number failed").Error())
		// TODO: this error is not the caller's fault, so we should implement a retry here instead of voting against
		return false
	}
	txReceipt, err := rpc.TransactionReceipt(context.Background(), txID)
	if err != nil {
		mgr.logger.Debug(sdkerrors.Wrap(err, "transaction receipt call failed").Error())
		return false
	}

	if !isTxFinalized(txReceipt, blockNumber, confHeight) {
		mgr.logger.Debug(fmt.Sprintf("transaction %s does not have enough confirmations yet", txReceipt.TxHash.String()))
		return false
	}

	return validateLogs(txReceipt)
}

func confirmERC20Deposit(txReceipt *geth.Receipt, amount sdk.Uint, burnAddr common.Address, tokenAddr common.Address) error {
	actualAmount := sdk.ZeroUint()
	for _, log := range txReceipt.Logs {
		/* Event is not related to the token */
		if log.Address != tokenAddr {
			continue
		}

		to, transferAmount, err := decodeERC20TransferEvent(log)
		/* Event is not an ERC20 transfer */
		if err != nil {
			continue
		}

		/* Transfer isn't sent to burner */
		if to != burnAddr {
			continue
		}

		actualAmount = actualAmount.Add(transferAmount)
	}

	if !actualAmount.Equal(amount) {
		return fmt.Errorf("given deposit amount: %d, actual amount: %d", amount.Uint64(), actualAmount.Uint64())
	}

	return nil
}

func confirmERC20TokenDeployment(txReceipt *geth.Receipt, expectedSymbol string, gatewayAddr, expectedAddr common.Address) error {
	for _, log := range txReceipt.Logs {
		// Event is not emitted by the axelar gateway
		if log.Address != gatewayAddr {
			continue
		}

		// Event is not for a ERC20 token deployment
		symbol, tokenAddr, err := decodeERC20TokenDeploymentEvent(log)
		if err != nil {
			continue
		}

		// Symbol does not match
		if symbol != expectedSymbol {
			continue
		}

		// token address does not match
		if tokenAddr != expectedAddr {
			continue
		}

		// if we reach this point, it means that the log matches what we want to verify,
		// so the function can return with no error
		return nil
	}

	return fmt.Errorf("failed to confirm token deployment for symbol '%s' at contract address '%s'", expectedSymbol, expectedAddr.String())
}

func isTxFinalized(txReceipt *geth.Receipt, blockNumber uint64, confirmationHeight uint64) bool {
	return blockNumber-txReceipt.BlockNumber.Uint64()+1 >= confirmationHeight
}

func decodeERC20TransferEvent(log *geth.Log) (common.Address, sdk.Uint, error) {

	if len(log.Topics) != 3 || log.Topics[0] != ERC20TransferSig {
		return common.Address{}, sdk.Uint{}, fmt.Errorf("log is not an ERC20 transfer")
	}

	to := common.BytesToAddress(log.Topics[2][:])
	amount := new(big.Int)
	amount.SetBytes(log.Data[:32])

	return to, sdk.NewUintFromBigInt(amount), nil
}

func decodeERC20TokenDeploymentEvent(log *geth.Log) (string, common.Address, error) {
	if len(log.Topics) != 1 || log.Topics[0] != ERC20TokenDeploymentSig {
		return "", common.Address{}, fmt.Errorf("event is not for an ERC20 token deployment")
	}

	// Decode the data field
	stringType, err := abi.NewType("string", "string", nil)
	if err != nil {
		return "", common.Address{}, err
	}
	addressType, err := abi.NewType("address", "address", nil)
	if err != nil {
		return "", common.Address{}, err
	}
	packedArgs := abi.Arguments{{Type: stringType}, {Type: addressType}}
	args, err := packedArgs.Unpack(log.Data)
	if err != nil {
		return "", common.Address{}, err
	}

	return args[0].(string), args[1].(common.Address), nil
}