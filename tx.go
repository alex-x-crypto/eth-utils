package eth_utils

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"time"
)

func SignTx(
	signer *ecdsa.PrivateKey, chainId *big.Int, nonce uint64,
	to *common.Address, value *big.Int, data []byte,
	maxPriorityFeePerGas, maxFeePerGas *big.Int, gasLimit uint64,
) (*types.Transaction, error) {
	baseTx := &types.DynamicFeeTx{
		ChainID: chainId,
		To:      to,
		Nonce:   nonce,
		Value:   value,
		Data:    data,

		GasTipCap: maxPriorityFeePerGas,
		GasFeeCap: maxFeePerGas,

		Gas: gasLimit,
	}

	tx, err := types.SignNewTx(
		signer,
		types.LatestSignerForChainID(chainId),
		baseTx,
	)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func WaitConfirm(ctx context.Context, ec *ethclient.Client, txHash common.Hash, timeout time.Duration) error {
	pending := true
	for pending {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(timeout):
			return errors.New("timeout")
		case <-time.After(time.Second):
			_, isPending, err := ec.TransactionByHash(ctx, txHash)
			if err != nil {
				return err
			}
			if !isPending {
				pending = false // break `for`
			}
		}
	}
	receipt, err := ec.TransactionReceipt(ctx, txHash)
	if err != nil {
		return err
	}
	if receipt.Status == 0 {
		msg := fmt.Sprintf("transaction reverted, hash %s", receipt.TxHash.String())
		return errors.New(msg)
	}
	return nil
}
