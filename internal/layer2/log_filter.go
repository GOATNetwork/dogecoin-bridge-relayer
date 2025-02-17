package layer2

import (
	"context"
	"errors"
	"reflect"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/nuvosphere/nudex-voter/internal/db"
	"github.com/nuvosphere/nudex-voter/internal/layer2/contracts"
	vtypes "github.com/nuvosphere/nudex-voter/internal/types"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (l *Layer2Listener) processLogs(vLog types.Log) {
	method, ok := l.addressBind[vLog.Address]
	if ok {
		err := method(vLog)
		if err != nil {
			log.Errorf("call %s processing log: %v", reflect.TypeOf(method).Name(), err)
		}
	}
}

func (l *Layer2Listener) processVotingLog(vLog types.Log) error {
	// save current submitter
	var (
		submitterChosen db.SubmitterChosen
		submitter       string
	)

	eventName := ""

	switch vLog.Topics[0] {
	case contracts.SubmitterChosenTopic:
		eventName = SubmitterChosen
		submitterChosenEvent := contracts.VotingManagerContractSubmitterChosen{}
		contracts.UnpackEventLog(contracts.VotingManagerContractMetaData, &submitterChosenEvent, eventName, vLog)
		submitter = submitterChosenEvent.NewSubmitter.Hex()
	case contracts.SubmitterRotationRequestedTopic:
		eventName = SubmitterRotationRequested
		submitterChosenEvent := contracts.VotingManagerContractSubmitterRotationRequested{}
		contracts.UnpackEventLog(contracts.VotingManagerContractMetaData, &submitterChosenEvent, eventName, vLog)
		submitter = submitterChosenEvent.CurrentSubmitter.Hex()
	default:
		return errors.New("invalid topic")
	}

	submitterChosen.Submitter = submitter
	submitterChosen.BlockNumber = vLog.BlockNumber
	submitterChosen.LogIndex = l.LogIndex(eventName, vLog)

	result := l.db.GetSystemDB().Create(&submitterChosen)
	if result.RowsAffected > 0 {
		l.PostTask(submitterChosen)
	}

	return result.Error
}

func (l *Layer2Listener) processTaskLog(vLog types.Log) error {
	switch vLog.Topics[0] {
	case contracts.TaskSubmittedTopic:
		taskSubmitted := contracts.TaskManagerContractTaskSubmitted{}
		contracts.UnpackEventLog(contracts.TaskManagerContractMetaData, &taskSubmitted, TaskSubmitted, vLog)
		actualTask := db.DecodeTaskOfEvent(taskSubmitted.TaskId, taskSubmitted.CallData, "")
		baseTask := &db.Task{
			ID:        taskSubmitted.TaskId,
			Ty:        actualTask.Type(),
			ChainID:   actualTask.ChainId(),
			State:     db.Created,
			Context:   taskSubmitted.CallData,
			Submitter: taskSubmitted.Submitter.String(),
			LogIndex:  l.LogIndex(TaskSubmitted, vLog),
		}

		err := l.db.GetContractDB().Transaction(func(tx *gorm.DB) error {
			err1 := tx.Create(baseTask).Error
			err2 := tx.Create(baseTask.DetailTask()).Error

			return errors.Join(err1, err2)
		})
		if err != nil {
			return err
		}

		l.PostTask(actualTask)

	case contracts.TaskUpdatedTopic:
		taskUpdated := contracts.TaskManagerContractTaskUpdated{}
		contracts.UnpackEventLog(contracts.TaskManagerContractMetaData, &taskUpdated, TaskUpdated, vLog)

		task := db.Task{}

		err := l.db.GetContractDB().Transaction(func(tx *gorm.DB) error {
			taskErr := tx.
				Model(&task).
				Clauses(clause.Returning{}).
				Where("id = ?", taskUpdated.TaskId).
				Updates(&db.Task{
					State:  int(taskUpdated.State),
					Result: taskUpdated.Result,
				}).
				Error

			queryErr := tx.
				Model(&task).
				Preload(clause.Associations).
				Where("id = ?", taskUpdated.TaskId).
				Last(&task).
				Error

			taskUpdatedEvent := &db.TaskUpdatedEvent{
				ID:         taskUpdated.TaskId,
				Submitter:  taskUpdated.Submitter.Hex(),
				UpdateTime: taskUpdated.UpdateTime.Int64(),
				State:      taskUpdated.State,
				Result:     taskUpdated.Result,
				LogIndex:   l.LogIndex(TaskUpdated, vLog),
			}
			err := tx.Save(taskUpdatedEvent).Error

			return errors.Join(taskErr, queryErr, err)
		})
		if err != nil {
			return err
		}

		l.PostTask(task.DetailTask())
	case contracts.NIP20TokenEventMintbTopic:
		mintb := contracts.InscriptionContractNIP20TokenEventMintb{}
		contracts.UnpackEventLog(contracts.InscriptionContractMetaData, &mintb, NIP20TokenMintbEvent, vLog)

		mintbEvent := &db.InscriptionMintb{
			Recipient: mintb.Recipient.Hex(),
			Ticker:    mintb.Ticker,
			Amount:    decimal.NewFromBigInt(mintb.Amount, 0),
			LogIndex:  l.LogIndex(NIP20TokenMintbEvent, vLog),
		}

		return l.db.GetContractDB().Create(mintbEvent).Error
	case contracts.NIP20TokenEventBurnbTopic:
		burnb := contracts.InscriptionContractNIP20TokenEventBurnb{}
		contracts.UnpackEventLog(contracts.InscriptionContractMetaData, &burnb, NIP20TokenBurnbEvent, vLog)

		burnbEvent := &db.InscriptionBurnb{
			From:     burnb.From.Hex(),
			Ticker:   burnb.Ticker,
			Amount:   decimal.NewFromBigInt(burnb.Amount, 0),
			LogIndex: l.LogIndex(NIP20TokenBurnbEvent, vLog),
		}

		return l.db.GetContractDB().Create(burnbEvent).Error
	default:
		return errors.New("invalid topic")
	}

	return nil
}

func (l *Layer2Listener) LogIndex(eventName string, vlog types.Log) db.LogIndex {
	chainID := l.ChainID(context.Background())

	if eventName == "" {
		eventName = vlog.Topics[0].String()
	}

	return db.LogIndex{
		Address:     vlog.Address,
		EventName:   eventName,
		Log:         &vlog,
		TxHash:      vlog.TxHash,
		ChainId:     chainID.Uint64(),
		BlockNumber: vlog.BlockNumber,
		LogIndex:    uint64(vlog.Index),
	}
}

func (l *Layer2Listener) processAccountLog(vLog types.Log) error {
	if vLog.Topics[0] == contracts.AddressRegisteredTopic {
		addressRegistered := contracts.AccountManagerContractAddressRegistered{}
		contracts.UnpackEventLog(contracts.AccountManagerContractMetaData, &addressRegistered, AddressRegistered, vLog)
		account := db.Account{
			UserAddress: addressRegistered.UserAddr,
			Account:     addressRegistered.Account.Uint64(),
			Chain:       addressRegistered.Chain,
			Index:       uint32(addressRegistered.Index.Uint64()),
			Address:     addressRegistered.NewAddress,
			LogIndex:    l.LogIndex(AddressRegistered, vLog),
		}

		return l.db.GetContractDB().Create(&account).Error
	}

	return nil
}

func (l *Layer2Listener) processParticipantLog(vLog types.Log) error {
	var (
		participantEvent *db.ParticipantEvent
		err              error
	)

	switch vLog.Topics[0] {
	case contracts.ParticipantAddedTopic:
		eventParticipantAdded := contracts.ParticipantManagerContractParticipantAdded{}
		contracts.UnpackEventLog(contracts.ParticipantManagerContractMetaData, &eventParticipantAdded, ParticipantAdded, vLog)
		newParticipant := eventParticipantAdded.Participant
		// save locked relayer member from db
		participant := db.Participant{Address: newParticipant.String()}
		err = l.db.
			GetSystemDB().Transaction(func(tx *gorm.DB) error {
			err1 := tx.FirstOrCreate(&participant, "address = ?", participant.Address).Error
			participantEvent = &db.ParticipantEvent{
				EventName:   ParticipantAdded,
				Address:     participant.Address,
				BlockNumber: vLog.BlockNumber,
				LogIndex:    l.LogIndex(ParticipantAdded, vLog),
			}
			err2 := tx.Create(participantEvent).Error

			return errors.Join(err1, err2)
		})
	case contracts.ParticipantRemovedTopic:
		participantRemovedEvent := contracts.ParticipantManagerContractParticipantRemoved{}
		contracts.UnpackEventLog(contracts.ParticipantManagerContractMetaData, &participantRemovedEvent, ParticipantRemoved, vLog)
		removedParticipant := participantRemovedEvent.Participant.Hex()

		err = l.db.
			GetSystemDB().Transaction(func(tx *gorm.DB) error {
			removedErr := tx.
				Where("address = ?", removedParticipant).
				Delete(&db.Participant{}).
				Error
			participantEvent = &db.ParticipantEvent{
				EventName:   ParticipantRemoved,
				Address:     removedParticipant,
				BlockNumber: vLog.BlockNumber,
				LogIndex:    l.LogIndex(ParticipantRemoved, vLog),
			}
			vlogErr := tx.Create(participantEvent).Error

			return errors.Join(removedErr, vlogErr)
		})
	case contracts.ParticipantsResetTopic:
		participantResetEvent := contracts.ParticipantManagerContractParticipantsReset{}
		contracts.UnpackEventLog(contracts.ParticipantManagerContractMetaData, &participantResetEvent, ParticipantReset, vLog)
		resetParticipant := participantResetEvent.Participants

		err = l.db.
			GetSystemDB().Transaction(func(tx *gorm.DB) error {
			removedErr := tx.
				Session(&gorm.Session{AllowGlobalUpdate: true}).
				Delete(&db.Participant{}).
				Error
			saveErr := tx.Create(lo.Map(resetParticipant, func(item common.Address, index int) db.Participant {
				return db.Participant{Address: item.String()}
			})).Error

			participantEvent = &db.ParticipantEvent{
				EventName:   ParticipantReset,
				Address:     strings.Join(lo.Map(resetParticipant, func(item common.Address, index int) string { return item.String() }), ","),
				BlockNumber: vLog.BlockNumber,
				LogIndex:    l.LogIndex(ParticipantReset, vLog),
			}
			vlogErr := tx.Create(participantEvent).Error

			return errors.Join(removedErr, saveErr, vlogErr)
		})

	default:
		return errors.New("invalid topic")
	}

	if err != nil {
		return err
	}

	if participantEvent != nil {
		l.PostTask(participantEvent.ParticipantEvent())
		log.Infof("Participant %s: %s", participantEvent.EventName, participantEvent.Address)
	}

	return nil
}

func (l *Layer2Listener) processDepositLog(vLog types.Log) error {
	switch vLog.Topics[0] {
	case contracts.DepositRecordedTopic:
		depositRecorded := contracts.DepositManagerContractDepositRecorded{}
		contracts.UnpackEventLog(contracts.DepositManagerContractMetaData, &depositRecorded, DepositRecorded, vLog)

		err := l.db.GetContractDB().Transaction(func(tx *gorm.DB) error {
			depositRecord := db.DepositRecord{
				DepositAddress: strings.ToLower(depositRecorded.DepositAddress),
				Amount:         decimal.NewFromBigInt(depositRecorded.Amount, 0),
				ChainId:        depositRecorded.ChainId,
				TxHash:         depositRecorded.TxHash,
				BlockHeight:    depositRecorded.BlockHeight.Uint64(),
				LogTxIndex:     depositRecorded.LogIndex.Uint64(),
				LogIndex:       l.LogIndex(DepositRecorded, vLog),
			}
			err1 := tx.Create(&depositRecord).Error

			tokenInfo := &db.TokenInfo{}
			tokenErr := tx.Model(tokenInfo).Where("ticker = ?", depositRecorded.Ticker).First(&tokenInfo).Error
			addressBalance := db.AddressBalance{
				Address: depositRecorded.DepositAddress,
				Token:   tokenInfo.ContractAddress,
				Amount:  decimal.NewFromBigInt(depositRecorded.Amount, 0),
				ChainId: depositRecorded.ChainId,
			}

			err2 := tx.Clauses(clause.OnConflict{
				Columns: []clause.Column{
					{Name: "address"},
					{Name: "token"},
					{Name: "chain_id"},
				},
				DoUpdates: clause.Assignments(map[string]interface{}{
					"amount": gorm.Expr("amount + ?", addressBalance.Amount),
				}),
			}).Create(&addressBalance).Error

			return errors.Join(err1, tokenErr, err2)
		})

		return err
	case contracts.WithdrawalRecordedTopic:
		withdrawalRecorded := contracts.DepositManagerContractWithdrawalRecorded{}
		contracts.UnpackEventLog(contracts.DepositManagerContractMetaData, &withdrawalRecorded, WithdrawalRecorded, vLog)

		err := l.db.GetContractDB().Transaction(func(tx *gorm.DB) error {
			account := &db.Account{}
			accountErr := tx.Model(account).Where("user_address = ?", withdrawalRecorded.UserAddress).First(&account).Error

			withdrawalRecord := db.WithdrawalRecord{
				ChainId:        withdrawalRecorded.ChainId,
				UserAddress:    withdrawalRecorded.UserAddress,
				DepositAddress: account.Address,
				ToAddress:      withdrawalRecorded.ToAddress,
				Amount:         decimal.NewFromBigInt(withdrawalRecorded.Amount, 0),
				TxHash:         withdrawalRecorded.TxHash,
				LogIndex:       l.LogIndex(WithdrawalRecorded, vLog),
			}
			err1 := tx.Create(&withdrawalRecord).Error
			err2 := tx.Model(&db.AddressBalance{}).
				Where("chain_id = ? AND address = ?", withdrawalRecorded.ChainId, account.Account).
				Update("amount", gorm.Expr("amount - ?", decimal.NewFromBigInt(withdrawalRecorded.Amount, 0))).
				Error

			return errors.Join(accountErr, err1, err2)
		})

		return err
	}

	return nil
}

func (l *Layer2Listener) processAssetLog(vLog types.Log) error {
	switch vLog.Topics[0] {
	case contracts.AssetListedTopic:
		event := contracts.AssetHandlerContractAssetListed{}
		contracts.UnpackEventLog(contracts.AssetHandlerContractMetaData, &event, AssetListed, vLog)

		asset := &db.Asset{
			Ticker:            event.Ticker,
			Decimals:          event.AssetParam.Decimals,
			DepositEnabled:    event.AssetParam.DepositEnabled,
			WithdrawalEnabled: event.AssetParam.WithdrawalEnabled,
			MinDepositAmount:  event.AssetParam.MinDepositAmount.Uint64(),
			MinWithdrawAmount: event.AssetParam.MinWithdrawAmount.Uint64(),
			AssetAlias:        event.AssetParam.AssetAlias,
		}

		return l.infoDB().Create(asset).Error
	case contracts.AssetUpdatedTopic:
		event := contracts.AssetHandlerContractAssetUpdated{}
		contracts.UnpackEventLog(contracts.AssetHandlerContractMetaData, &event, AssetUpdated, vLog)

		return l.infoDB().Where("ticker = ?", vtypes.Byte32(event.Ticker)).Updates(&db.Asset{
			Ticker:            event.Ticker,
			Decimals:          event.AssetParam.Decimals,
			DepositEnabled:    event.AssetParam.DepositEnabled,
			WithdrawalEnabled: event.AssetParam.WithdrawalEnabled,
			MinDepositAmount:  event.AssetParam.MinDepositAmount.Uint64(),
			MinWithdrawAmount: event.AssetParam.MinWithdrawAmount.Uint64(),
			AssetAlias:        event.AssetParam.AssetAlias,
		}).Error
	case contracts.AssetDelistedTopic:
		event := contracts.AssetHandlerContractAssetDelisted{}
		contracts.UnpackEventLog(contracts.AssetHandlerContractMetaData, &event, AssetDelisted, vLog)

		return l.infoDB().Where("ticker = ?", vtypes.Byte32(event.Ticker)).Delete(&db.Asset{}).Error
	case contracts.ResetLinkedTokenTopic:
		event := contracts.AssetHandlerContractResetLinkedToken{}
		contracts.UnpackEventLog(contracts.AssetHandlerContractMetaData, &event, "ResetLinkedToken", vLog)

		return l.infoDB().
			Model(&db.TokenInfo{}).
			Where("ticker = ?", event.Ticker).
			Update("is_ticker", false).
			Error
	case contracts.TokenSwitchTopic:
		event := contracts.AssetHandlerContractTokenSwitchIterator{}
		contracts.UnpackEventLog(contracts.AssetHandlerContractMetaData, &event, "ResetLinkedToken", vLog)

		return l.infoDB().
			Model(&db.TokenInfo{}).
			Where("ticker = ? AND chain_id = ?", event.Event.Ticker, event.Event.ChainId).
			Update("is_ticker", event.Event.IsActive).
			Error
	case contracts.LinkTokenTopic:
		event := contracts.AssetHandlerContractLinkToken{}
		contracts.UnpackEventLog(contracts.AssetHandlerContractMetaData, &event, "LinkToken", vLog)
		tokenInfos := lo.Map(event.Tokens, func(token contracts.TokenInfo, index int) db.TokenInfo {
			return db.TokenInfo{
				ChainId:         token.ChainId,
				Ticker:          event.Ticker,
				IsActive:        token.IsActive,
				AssetType:       0, // todo
				Decimals:        token.Decimals,
				ContractAddress: token.ContractAddress,
				Symbol:          token.Symbol,
				WithdrawFee:     decimal.NewFromBigInt(token.WithdrawFee, 0),
			}
		})

		return l.infoDB().Create(tokenInfos).Error

	default:
		return errors.New("invalid topic")
	}
}

func (l *Layer2Listener) infoDB() *gorm.DB {
	return l.db.GetContractDB()
}
