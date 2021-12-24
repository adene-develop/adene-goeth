package eth

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type CreateFilter struct {
	FromBlock string         `json:"fromBlock"`
	ToBlock   string         `json:"toBlock"`
	Address   common.Address `json:"address"`
}

type FilterChange struct {
	Address          common.Address `json:"address"`
	Topics           []common.Hash  `json:"topics"`
	Data             hexutil.Bytes  `json:"data"`
	BlockNumber      string         `json:"blockNumber"`
	TransactionHash  common.Hash    `json:"transactionHash"`
	TransactionIndex string         `json:"transactionIndex"`
	BlockHash        common.Hash    `json:"blockHash"`
	LogIndex         string         `json:"logIndex"`
	Removed          bool           `json:"removed"`
}

func (f *FilterChange) EventID() common.Hash {
	return f.Topics[0]
}
