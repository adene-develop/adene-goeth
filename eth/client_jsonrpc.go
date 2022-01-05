package eth

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"strconv"
)

func (c *Client) EthNewFilter(ctx context.Context, fromBlock string, toBlock string, addresses []common.Address, topics [][]common.Hash) (string, error) {
	var filterID string

	err := c.RpcClient().CallContext(ctx, &filterID, "eth_newFilter", toNewFilterArg(fromBlock, toBlock, addresses, topics))
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

func toNewFilterArg(fromBlock string, toBlock string, addresses []common.Address, topics [][]common.Hash) interface{} {
	arg := map[string]interface{}{}
	fromBlockID, err := strconv.ParseInt(fromBlock, 10, 64)
	if err == nil {
		arg["fromBlock"] = fromBlockID
	} else {
		arg["fromBlock"] = fromBlock
	}

	toBlockID, err := strconv.ParseInt(toBlock, 10, 64)
	if err != nil {
		arg["toBlock"] = toBlock
	} else {
		arg["toBlock"] = toBlockID
	}

	if addresses != nil && len(addresses) > 0 {
		if len(addresses) == 1 {
			arg["address"] = addresses[0]
		} else {
			arg["address"] = addresses
		}
	}

	if topics != nil && len(topics) > 0 {
		argTopics := make([]interface{}, 0)
		for i := 0; i < len(topics); i++ {
			if topics[i] == nil || len(topics[i]) == 0 {
				argTopics = append(argTopics, "null")
			} else if len(topics[i]) == 1 {
				argTopics = append(argTopics, topics[i][0])
			} else {
				argTopics = append(argTopics, topics[i])
			}
		}
		arg["topics"] = argTopics
	}

	return arg
}
