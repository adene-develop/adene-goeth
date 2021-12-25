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
	return nil
}
