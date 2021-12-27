package adene

import (
	"github.com/adene-develop/adene-goeth/contract"
	"github.com/adene-develop/adene-goeth/eth"
)

type Events interface {
	contract.ERC20Events
	contract.OwnableEvents

	MinTokensBeforeSwapUpdated(minTokensBeforeSwap int64)
	SwapAndLiquifyEnabledUpdated(enabled bool)
	SwapAndLiquify(tokensSwapped, ethReceived, tokensIntoLiquidity int64)
}

func ParseEvents(filterChanges []*eth.FilterChange, events Events) error {
	if err := contract.ParseERC20Events(filterChanges, events); err != nil {
		return err
	}
	if err := contract.ParseOwnableEvents(filterChanges, events); err != nil {
		return err
	}

	// TODO implement other events
	return nil
}
