package eth

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

func (c *Client) EthNewFilter(ctx context.Context, fromBlock string, toBlock string, address common.Address) (string, error) {
	var filterID string

	err := c.RpcClient().CallContext(ctx, &filterID, "eth_newFilter", toNewFilterArg(fromBlock, toBlock, address))
	if err != nil {
		return "", errors.Wrap(err, "call rpc error")
	}
	return filterID, nil
}

func (c *Client) EthGetFilterChanges(ctx context.Context, filterID string) ([]*FilterChange, error) {
	var changes []*FilterChange

	err := c.RpcClient().CallContext(ctx, &changes, "eth_getFilterChanges", filterID)
	if err != nil {
		return nil, errors.Wrap(err, "call rpc error")
	}
	return changes, nil
}

func (c *Client) EthUninstallFilter(ctx context.Context, filterID string) error {
	var removed bool
	err := c.RpcClient().CallContext(ctx, &removed, "eth_uninstallFilter", filterID)
	if err != nil {
		return errors.Wrap(err, "call uninstall filter error")
	}
	if !removed {
		return errors.New("remove filter failed")
	}
	return nil
}

func toNewFilterArg(fromBlock string, toBlock string, address common.Address) interface{} {
	arg := map[string]interface{}{
		"fromBlock": fromBlock,
		"toBlock":   toBlock,
		"address":   address,
	}
	return arg
}
