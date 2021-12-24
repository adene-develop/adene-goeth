package icon721

import (
	"github.com/adene-develop/adene-goeth/contract"
	"github.com/adene-develop/adene-goeth/eth"
)

type Events interface {
	contract.ERC721Events
	contract.OwnableEvents
}

func ParseEvents(filterChanges []*eth.FilterChange, events Events) {

}
