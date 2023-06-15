package eth_utils

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
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
