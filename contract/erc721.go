package contract

import (
	"context"
	"github.com/adene-develop/adene-goeth/eth"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"math/big"
)

// ERC721 Standard view functions
type ERC721 interface {
	Contract
	// Name returns the token collection name.
	Name(ctx context.Context) (string, error)

	// Symbol returns the token collection symbol.
	Symbol(ctx context.Context) (string, error)

	// TokenURI returns the Uniform Resource Identifier (URI) for `tokenId` token.
	TokenURI(ctx context.Context, tokenID int64) (string, error)

	// BalanceOf returns the number of tokens in ``owner``'s account.
	BalanceOf(ctx context.Context, owner common.Address) (balance int64, err error)

	// OwnerOf returns the owner of the `tokenId` token.
	OwnerOf(ctx context.Context, tokenID int64) (owner common.Address, err error)

	// GetApproved returns the account approved for `tokenId` token.
	GetApproved(ctx context.Context, tokenID int64) (operator common.Address, err error)

	// IsApprovedForAll returns if the `operator` is allowed to manage all of the assets of `owner`.
	IsApprovedForAll(ctx context.Context, owner common.Address, operator common.Address) (bool, error)
}

func NewERC721(client *eth.Client, address common.Address) ERC721 {
	return &ERC721Contract{
		address: address,
		client:  client,
	}
}

type ERC721Contract struct {
	address common.Address
	client  *eth.Client
}

func (e *ERC721Contract) Address() common.Address {
	return e.address
}

func (e *ERC721Contract) Client() *eth.Client {
	return e.client
}

func (e *ERC721Contract) Name(ctx context.Context) (string, error) {
	var result struct {
		Name string
	}
	if err := e.client.CallContractViewFunction(ctx, ERC721ABI, e.address, &result, "name"); err != nil {
		return "", errors.Wrap(err, "ERC721Contract call view `name` error")
	}
	return result.Name, nil
}

func (e *ERC721Contract) Symbol(ctx context.Context) (string, error) {
	var result struct {
		Symbol string
	}
	if err := e.client.CallContractViewFunction(ctx, ERC721ABI, e.address, &result, "symbol"); err != nil {
		return "", errors.Wrap(err, "ERC721Contract call view `symbol` error")
	}

	return result.Symbol, nil
}

func (e *ERC721Contract) TokenURI(ctx context.Context, tokenID int64) (string, error) {
	var result struct {
		TokenURI string
	}
	if err := e.client.CallContractViewFunction(ctx, ERC721ABI, e.address, &result, "tokenURI", big.NewInt(tokenID)); err != nil {
		return "", errors.Wrap(err, "ERC721Contract call view `tokenURI` error")
	}
	return result.TokenURI, nil
}

func (e *ERC721Contract) BalanceOf(ctx context.Context, owner common.Address) (int64, error) {
	var result struct {
		Balance *big.Int
	}
	if err := e.client.CallContractViewFunction(ctx, ERC721ABI, e.address, &result, "balanceOf", owner); err != nil {
		return 0, errors.Wrap(err, "ERC721Contract call view `balanceOf` error")
	}

	return result.Balance.Int64(), nil
}

func (e *ERC721Contract) OwnerOf(ctx context.Context, tokenID int64) (common.Address, error) {
	var result struct {
		Owner common.Address
	}
	if err := e.client.CallContractViewFunction(ctx, ERC721ABI, e.address, &result, "ownerOf", big.NewInt(tokenID)); err != nil {
		return common.Address{}, errors.Wrap(err, "ERC721Contract call view `ownerOf` error")
	}
	return result.Owner, nil
}

func (e *ERC721Contract) GetApproved(ctx context.Context, tokenID int64) (common.Address, error) {
	var result struct {
		Operator common.Address
	}
	if err := e.client.CallContractViewFunction(ctx, ERC721ABI, e.address, &result, "getApproved", big.NewInt(tokenID)); err != nil {
		return common.Address{}, errors.Wrap(err, "ERC721Contract call view `getApproved` error")
	}
	return result.Operator, nil
}

func (e *ERC721Contract) IsApprovedForAll(ctx context.Context, owner common.Address, operator common.Address) (bool, error) {
	var result struct {
		IsApprovedForAll bool
	}
	if err := e.client.CallContractViewFunction(ctx, ERC721ABI, e.address, &result, "isApprovedForAll", owner, operator); err != nil {
		return false, errors.Wrap(err, "ERC721Contract call view `isApprovedForAll` error")
	}
	return result.IsApprovedForAll, nil
}

type ERC721Events interface {
	// Transfer emitted when `tokenId` token is transferred from `from` to `to`.
	Transfer(from common.Address, to common.Address, tokenID int64)

	// Approval emitted when `owner` enables `approved` to manage the `tokenId` token.
	Approval(owner common.Address, approved common.Address, tokenID int64)

	// ApprovalForAll emitted when `owner` enables or disables (`approved`) `operator` to manage all of its assets.
	ApprovalForAll(owner common.Address, operator common.Address, approved bool)
}

func ParseERC721Events(filterChanges []*eth.FilterChange, events ERC721Events) error {
	for i := 0; i < len(filterChanges); i++ {
		switch filterChanges[i].EventID() {
		case ERC721ABI.Events["Transfer"].ID:
			if err := parseERC721TransferEvent(filterChanges[i], events); err != nil {
				return errors.Wrap(err, "ParseERC721Events parse transfer event")
			}
		case ERC721ABI.Events["Approval"].ID:
			if err := parseERC721ApprovalEvent(filterChanges[i], events); err != nil {
				return errors.Wrap(err, "ParseERC721Events parse approval event")
			}

		default:
		}
	}
	return nil
}

func parseERC721TransferEvent(change *eth.FilterChange, events ERC721Events) error {
	if change.Topics == nil || len(change.Topics) < 4 {
		return errors.New("invalid topics")
	}
	from := common.BytesToAddress(change.Topics[1].Bytes())
	to := common.BytesToAddress(change.Topics[2].Bytes())
	tokenID := new(big.Int).SetBytes(change.Topics[3].Bytes()).Int64()
	events.Transfer(from, to, tokenID)
	return nil
}

func parseERC721ApprovalEvent(change *eth.FilterChange, events ERC721Events) error {
	if change.Topics == nil || len(change.Topics) < 4 {
		return errors.New("invalid topics")
	}
	from := common.BytesToAddress(change.Topics[1].Bytes())
	to := common.BytesToAddress(change.Topics[2].Bytes())
	tokenID := new(big.Int).SetBytes(change.Topics[3].Bytes()).Int64()
	events.Approval(from, to, tokenID)
	return nil
}

func parseERC721ApprovalForAllEvent(change *eth.FilterChange, events ERC721Events) error {
	if change.Topics == nil || len(change.Topics) < 3 {
		return errors.New("invalid topics")
	}
	from := common.BytesToAddress(change.Topics[1].Bytes())
	to := common.BytesToAddress(change.Topics[2].Bytes())
	var r struct {
		Approved bool
	}
	err := ERC721ABI.UnpackIntoInterface(&r, "ApprovalForAll", change.Data)
	if err != nil {
		return err
	}

	events.ApprovalForAll(from, to, r.Approved)
	return nil
}
