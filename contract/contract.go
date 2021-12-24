package contract

import (
	"github.com/adene-develop/adene-goeth/eth"
	"github.com/ethereum/go-ethereum/common"
)

type Contract interface {
	Address() common.Address
	Client() *eth.Client
}
