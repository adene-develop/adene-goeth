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
}

func ParseEvents(filterChanges []*eth.FilterChange, events Events) error {
	if filterChanges == nil || len(filterChanges) == 0 {
		return nil
	}

	if events == nil {
		return errors.New("events is nil")
	}

	contract.

	for i := 0; i < len(filterChanges); i++ {
		switch filterChanges[i].EventID() {
		case SALE2021Q4ABI.Events[EventBoughtName].ID:
			user := common.BytesToAddress(filterChanges[i].Topics[1].Bytes())
			var event struct {
				Level        uint8
				Amount       uint16
				StartTokenId *big.Int
				ToTokenId    *big.Int
			}
			err := SALE2021Q4ABI.UnpackIntoInterface(&event, EventBoughtName, filterChanges[i].Data)
			if err != nil {
				return errors.Wrap(err, "could not unpack bought event data to interface")
			}
			s.boughtListener(user, event.Level, event.Amount, event.StartTokenId.Int64(), event.ToTokenId.Int64())
		default:
		}
	}
	return nil
}
