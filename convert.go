package eth_utils

import "math/big"

func ToGwei(wei *big.Int) string {
	g := new(big.Float).SetUint64(wei.Uint64())
	g.Quo(g, big.NewFloat(1e9))
	return g.String()
}

func ToEther(wei *big.Int) string {
	g := new(big.Float).SetUint64(wei.Uint64())
	g.Quo(g, big.NewFloat(1e18))
	return g.String()
}
