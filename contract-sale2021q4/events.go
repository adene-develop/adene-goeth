package sale2021q4

import (
	"github.com/adene-develop/adene-goeth/contract"
	"github.com/adene-develop/adene-goeth/eth"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"math/big"
)

const EventBoughtName = "Bought"

type Events interface {
	contract.PauseableEvents
	contract.OwnableEvents

	// Bought emitted when `user` buy `amount` of box with level `level`
	// the list of bought tokens start at id `startTokenID` and end at id `toTokenID`
	Bought(user common.Address, level BoxLevel, amount int, startTokenID, toTokenID int64)
}

func ParseEvents(filterChanges []*eth.FilterChange, events Events) error {
	if filterChanges == nil || len(filterChanges) == 0 {
		return nil
	}

	if events == nil {
		return errors.New("events is nil")
	}

	err := contract.ParsePauseableEvents(filterChanges, events)
	if err != nil {
		return errors.Wrap(err, "sale2021q4 parse pauseable events error")
	}

	err = contract.ParseOwnableEvents(filterChanges, events)
	if err != nil {
		return errors.Wrap(err, "sale2021q4 parse ownable events error")
	}

	for i := 0; i < len(filterChanges); i++ {
		switch filterChanges[i].EventID() {
		case ABI.Events[EventBoughtName].ID:
			if err = parseBoughtEvent(filterChanges[i], events); err != nil {
				return err
			}
		default:
		}
	}
	return nil
}

func parseBoughtEvent(change *eth.FilterChange, events Events) error {
	user := common.BytesToAddress(change.Topics[1].Bytes())
	var event struct {
		Level        uint8
		Amount       uint16
		StartTokenId *big.Int
		ToTokenId    *big.Int
	}
	err := ABI.UnpackIntoInterface(&event, EventBoughtName, change.Data)
	if err != nil {
		return errors.Wrap(err, "could not unpack bought event data to interface")
	}

	events.Bought(user, BoxLevel(event.Level), int(event.Amount), event.StartTokenId.Int64(), event.ToTokenId.Int64())
	return nil
}
