package db

import (
	"encoding/hex"

	"github.com/nuvosphere/nudex-voter/internal/layer2/contracts"
	"github.com/nuvosphere/nudex-voter/internal/utils"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
)

var (
	WalletCreationReq = contracts.MethodID(contracts.AccountManagerContractMetaData, "registerNewAddress")
	DepositReq        = contracts.MethodID(contracts.DepositManagerContractMetaData, "recordDeposit")
	WithdrawalReq     = contracts.MethodID(contracts.DepositManagerContractMetaData, "recordWithdrawal")
	ConsolidateReq    = contracts.MethodID(contracts.AssetHandlerContractMetaData, "consolidate")
	TransferReq       = contracts.MethodID(contracts.AssetHandlerContractMetaData, "transfer")
)

func DecodeTaskOfEvent(taskId uint64, context []byte, result string) DetailTask {
	parsedABI, err := contracts.ParseABI(contracts.TaskPayloadContractMetaData.ABI)
	utils.Assert(err)

	method := context[:4]
	context = context[4:]
	log.Debugf("context: %x", context)

	methodId := hex.EncodeToString(method)
	switch methodId {
	case WalletCreationReq:
		request := &contracts.TaskPayloadContractWalletCreationRequest{}
		err = parsedABI.UnpackIntoInterface(request, "WalletCreationRequest", context)
		utils.Assert(err)

		return &CreateWalletTask{
			TaskId:      taskId,
			UserAddress: request.UserAddr,
			Account:     request.Account,
			Chain:       request.Chain,
			Index:       request.Index,
		}
	case DepositReq:
		request := &contracts.TaskPayloadContractDepositRequest{}
		err = parsedABI.UnpackIntoInterface(request, "DepositRequest", context)
		utils.Assert(err)

		return &DepositTask{
			TaskId:         taskId,
			State:          Created,
			UserAddress:    request.UserAddress,
			DepositAddress: request.DepositAddress,
			Amount:         decimal.NewFromBigInt(request.Amount, 0),
			Chain:          0,
			ChainID:        request.ChainId,
			Ticker:         request.Ticker,
			TxHash:         request.TxHash,
			TxIndex:        request.LogIndex.Uint64(),
			BlockHeight:    request.BlockHeight.Uint64(),
		}

	case WithdrawalReq:
		request := &contracts.TaskPayloadContractWithdrawalRequest{}
		err = parsedABI.UnpackIntoInterface(request, "WithdrawalRequest", context)
		utils.Assert(err)

		return &WithdrawalTask{
			TaskId:      taskId,
			State:       Created,
			UserAddress: request.UserAddress,
			ToAddress:   request.ToAddress,
			Amount:      decimal.NewFromBigInt(request.Amount, 0),
			Chain:       0,
			ChainID:     request.ChainId,
			Ticker:      request.Ticker,
			TxHash:      result,
		}
	case ConsolidateReq:
		request := &contracts.TaskPayloadContractConsolidationRequest{}
		err = parsedABI.UnpackIntoInterface(request, "WithdrawalRequest", context)
		utils.Assert(err)

		return &ConsolidationTask{
			TaskId:      taskId,
			State:       Created,
			FromAddress: request.FromAddress,
			Amount:      decimal.NewFromBigInt(request.Amount, 0),
			Chain:       0,
			ChainID:     request.ChainId,
			Ticker:      request.Ticker,
			TxHash:      result,
		}

	case TransferReq:
		request := &contracts.TaskPayloadContractTransferRequest{}
		err = parsedABI.UnpackIntoInterface(request, "TransferRequest", context)
		utils.Assert(err)

		return &TransferTask{
			TaskId:      taskId,
			State:       Created,
			FromAddress: request.FromAddress,
			ToAddress:   request.ToAddress,
			Amount:      decimal.NewFromBigInt(request.Amount, 0),
			Chain:       0,
			ChainID:     request.ChainId,
			Ticker:      request.Ticker,
			TxHash:      result,
		}
	}

	panic("DecodeTask: error task type")
}

func DecodeTask(ty int, context []byte, result string) DetailTask {
	switch ty {
	case TaskTypeCreateWallet:
		request := &contracts.AccountManagerContractRequestRegisterAddress{}
		contracts.UnpackEvent(contracts.AccountManagerContractMetaData, "RequestRegisterAddress", request, context)

		return &CreateWalletTask{
			TaskId:      request.TaskId,
			UserAddress: request.TaskParam.UserAddr,
			Account:     request.TaskParam.Account,
			Chain:       request.TaskParam.Chain,
			Index:       request.TaskParam.Index,
		}
	case TaskTypeDeposit:
		request := &contracts.DepositManagerContractRequestDeposit{}
		contracts.UnpackEvent(contracts.DepositManagerContractMetaData, "RequestDeposit", request, context)

		return &DepositTask{
			TaskId:         request.TaskId,
			State:          Created,
			UserAddress:    request.DepositInfo.UserAddress,
			DepositAddress: request.DepositInfo.DepositAddress,
			Amount:         decimal.NewFromBigInt(request.DepositInfo.Amount, 0),
			Chain:          0,
			ChainID:        request.DepositInfo.ChainId,
			Ticker:         request.DepositInfo.Ticker,
			TxHash:         request.DepositInfo.TxHash,
			TxIndex:        request.DepositInfo.LogIndex.Uint64(),
			BlockHeight:    request.DepositInfo.BlockHeight.Uint64(),
		}

	case TaskTypeWithdrawal:
		request := &contracts.DepositManagerContractRequestWithdrawal{}
		contracts.UnpackEvent(contracts.DepositManagerContractMetaData, "RequestWithdrawal", request, context)

		return &WithdrawalTask{
			TaskId:      request.TaskId,
			State:       Created,
			UserAddress: request.WithdrawalInfo.UserAddress,
			ToAddress:   request.WithdrawalInfo.ToAddress,
			Amount:      decimal.NewFromBigInt(request.WithdrawalInfo.Amount, 0),
			Chain:       0,
			ChainID:     request.WithdrawalInfo.ChainId,
			Ticker:      request.WithdrawalInfo.Ticker,
			TxHash:      result,
		}
	case TaskTypeConsolidation:
		request := &contracts.AssetHandlerContractRequestConsolidate{}
		contracts.UnpackEvent(contracts.AssetHandlerContractMetaData, "RequestConsolidate", request, context)

		return &ConsolidationTask{
			TaskId:      request.TaskId,
			State:       Created,
			FromAddress: request.Param.FromAddress,
			Amount:      decimal.NewFromBigInt(request.Param.Amount, 0),
			Chain:       0,
			ChainID:     request.Param.ChainId,
			Ticker:      request.Param.Ticker,
			TxHash:      result,
		}

	case TaskTypeTransfer:
		request := &contracts.AssetHandlerContractRequestTransfer{}
		contracts.UnpackEvent(contracts.AssetHandlerContractMetaData, "RequestTransfer", request, context)

		return &TransferTask{
			TaskId:      request.TaskId,
			State:       Created,
			FromAddress: request.Param.FromAddress,
			ToAddress:   request.Param.ToAddress,
			Amount:      decimal.NewFromBigInt(request.Param.Amount, 0),
			Chain:       0,
			ChainID:     request.Param.ChainId,
			Ticker:      request.Param.Ticker,
			TxHash:      result,
		}
	default:
		panic("DecodeTask: error task type")
	}
}
