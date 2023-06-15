package eth_utils

import "math/big"

func ToGwei(wei *big.Int) string {
	g := new(big.Float).SetUint64(wei.Uint64())
	g.Quo(g, big.NewFloat(1e9))
	return g.String()
}

func ToEther(wei *big.Int) string {
	e := new(big.Float).SetUint64(wei.Uint64())
	e.Quo(e, big.NewFloat(1e18))
	return e.String()
}

func FromEther(ether string) (*big.Int, bool) {
	e, success := new(big.Float).SetString(ether)
	if !success {
		return nil, success
	}
	e.Mul(e, big.NewFloat(1e18))
	e2, a := e.Int(nil)
	if a != big.Exact {
		return nil, false
	}
	return e2, true
}
