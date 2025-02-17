package layer2

import (
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/eventbus"
	"github.com/nuvosphere/nudex-voter/internal/pool"
	"github.com/nuvosphere/nudex-voter/internal/state"
	"github.com/nuvosphere/nudex-voter/internal/types"
	log "github.com/sirupsen/logrus"
)

func (l *Layer2Listener) ContractState() *state.ContractState {
	return state.NewContractState(l.db.GetContractDB())
}

func (l *Layer2Listener) SubmitTask(taskResult *db.TaskResult) {
	l.eventBus.Publish(eventbus.EventSubmitTask{}, taskResult)
}

func (l *Layer2Listener) postTask(task any) {
	l.eventBus.Publish(eventbus.EventTask{}, task)
}

func (l *Layer2Listener) PostTask(task any) {
	if value, ok := task.(pool.Task[uint64]); ok {
		l.check(value)
	} else {
		l.postTask(task)
	}
}

func (l *Layer2Listener) check(task any) {
	log.Debugf("check: task: %v", task)

	switch value := task.(type) {
	case *db.CreateWalletTask:
		l.postTask(&db.CreateWallet{
			TaskId:      value.TaskId,
			UserAddress: value.UserAddress,
			State:       db.Created,
			Account:     value.Account,
			Chain:       value.Chain,
			Index:       value.Index,
		})
	case *db.DepositTask:
		token, err := l.ContractState().GetTokenInfo(value.Ticker)
		if err != nil {
			log.Printf("Error getting token for %v: %v", value.Ticker, err)
			return
		}

		_, err = l.ContractState().GetAccount(value.UserAddress)
		if err != nil {
			log.Printf("Error getting account for %v: %v", value.UserAddress, err)
			return
		}

		l.postTask(&db.Deposit{
			Common: db.Common{
				TaskId:          value.TaskId,
				State:           value.State,
				Chain:           value.Chain,
				ChainID:         value.ChainID,
				ContractAddress: token.ContractAddress,
				Decimals:        token.Decimals,
				Symbol:          token.Symbol,
				Amount:          value.Amount,
				TxHash:          value.TxHash,
			},
			ToAddress:   value.DepositAddress,
			BlockHeight: value.BlockHeight,
			TxIndex:     value.TxIndex,
		})
	case *db.WithdrawalTask:
		token, err := l.ContractState().GetTokenInfo(value.Ticker)
		if err != nil {
			log.Printf("Error getting token for %v: %v", value.Ticker, err)
			return
		}

		account, err := l.ContractState().GetAccount(value.UserAddress)
		if err != nil {
			log.Printf("Error getting account for %v: %v", value.UserAddress, err)
			return
		}

		l.postTask(&db.Withdrawal{
			Common: db.Common{
				TaskId:          value.TaskId,
				State:           value.State,
				Chain:           value.Chain,
				ChainID:         value.ChainID,
				ContractAddress: token.ContractAddress,
				Decimals:        token.Decimals,
				Symbol:          token.Symbol,
				Amount:          value.Amount,
			},
			FromAddress:   account.Address,
			WithdrawalFee: token.WithdrawFee,
			ToAddress:     value.ToAddress,
		})

	case *db.ConsolidationTask:
		token, err := l.ContractState().GetTokenInfo(value.Ticker)
		if err != nil {
			log.Printf("Error getting token for %v: %v", value.Ticker, err)
			return
		}

		l.postTask(&db.Consolidation{
			FromAddress: value.FromAddress,
			Common: db.Common{
				TaskId:          value.TaskId,
				State:           value.State,
				Chain:           value.Chain,
				ChainID:         value.ChainID,
				ContractAddress: token.ContractAddress,
				Decimals:        token.Decimals,
				Symbol:          token.Symbol,
				Amount:          value.Amount,
				TxHash:          value.TxHash,
			},
		})

	case *db.TransferTask:
		token, err := l.ContractState().GetTokenInfo(value.Ticker)
		if err != nil {
			log.Printf("Error getting token for %v: %v", value.Ticker, err)
			return
		}

		l.postTask(&db.Transfer{
			Common: db.Common{
				TaskId:          value.TaskId,
				State:           value.State,
				Chain:           value.Chain,
				ChainID:         value.ChainID,
				ContractAddress: token.ContractAddress,
				Decimals:        token.Decimals,
				Symbol:          token.Symbol,
				Amount:          value.Amount,
			},
			FromAddress: value.FromAddress,
			ToAddress:   value.ToAddress,
		})

	case *db.Task:
		l.check(value.DetailTask())
	default:
		panic("unknown task type")
	}
}

func (l *Layer2Listener) GetToken(ticker types.Byte32) (*db.TokenInfo, error) {
	token, err := l.ContractState().GetTokenInfo(ticker)
	if err != nil {
		log.Printf("Error getting token for %v: %v", ticker, err)
		return nil, err
	}

	return token, nil
}
