package contract

import (
	"context"
	"github.com/adene-develop/adene-goeth/eth"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"math/big"
)

// ERC721Enumerable view functions
type ERC721Enumerable interface {
	ERC721
	// TokenOfOwnerByIndex returns a token ID owned by `owner` at a given `index` of its token list.
	// Use along with {balanceOf} to enumerate all of ``owner``'s tokens.
	TokenOfOwnerByIndex(ctx context.Context, owner common.Address, index int64) (tokenID int64, err error)

	// TotalSupply returns the total amount of tokens stored by the contract.
	TotalSupply(ctx context.Context) (total int64, err error)

	// TokenByIndex returns a token ID at a given `index` of all the tokens stored by the contract.
	// Use along with {totalSupply} to enumerate all tokens.
	TokenByIndex(ctx context.Context, index int64) (tokenID int64, err error)
}

func NewERC721Enumerable(client *eth.Client, address common.Address) ERC721Enumerable {
	return &ERC721EnumerableContract{
		ERC721:  NewERC721(client, address),
		address: address,
		client:  client,
	}
}

type ERC721EnumerableContract struct {
	ERC721
	address common.Address
	client  *eth.Client
}

func (e *ERC721EnumerableContract) TokenOfOwnerByIndex(ctx context.Context, owner common.Address, index int64) (int64, error) {
	var result struct {
		TokenID *big.Int
	}

	if err := e.client.CallContractViewFunction(ctx, ERC721EnumerableABI, e.address, &result, "tokenOfOwnerByIndex", owner, big.NewInt(index)); err != nil {
		return 0, errors.Wrap(err, "ERC721EnumerableContract call view `tokenOfOwnerByIndex` function error ")
	}
	return result.TokenID.Int64(), nil
}

func (e *ERC721EnumerableContract) TotalSupply(ctx context.Context) (int64, error) {
	var result struct {
		TotalSupply *big.Int
	}

	if err := e.client.CallContractViewFunction(ctx, ERC721EnumerableABI, e.address, &result, "totalSupply"); err != nil {
		return 0, errors.Wrap(err, "ERC721EnumerableContract call view `totalSupply` function error ")
	}
	return result.TotalSupply.Int64(), nil
}

func (e *ERC721EnumerableContract) TokenByIndex(ctx context.Context, index int64) (tokenID int64, err error) {
	var result struct {
		TokenID *big.Int
	}

	if err := e.client.CallContractViewFunction(ctx, ERC721EnumerableABI, e.address, &result, "tokenByIndex", big.NewInt(index)); err != nil {
		return 0, errors.Wrap(err, "ERC721EnumerableContract call view `tokenByIndex` function error ")
	}
	return result.TokenID.Int64(), nil
}
