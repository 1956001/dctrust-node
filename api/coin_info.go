package api

import (
	"github.com/kvant-node/core/types"
	"github.com/kvant-node/rpc/lib/types"
)

type CoinInfoResponse struct {
	Name           string `json:"name"`
	Symbol         string `json:"symbol"`
	Volume         string `json:"volume"`
	Crr            uint   `json:"crr"`
	ReserveBalance string `json:"reserve_balance"`
	MaxSupply      string `json:"max_supply"`
}

func CoinInfo(coinSymbol string, height int) (*CoinInfoResponse, error) {
	cState, err := GetStateForHeight(height)
	if err != nil {
		return nil, err
	}

	cState.Lock()
	defer cState.Unlock()

	coin := cState.Coins.GetCoin(types.StrToCoinSymbol(coinSymbol))
	if coin == nil {
		return nil, rpctypes.RPCError{Code: 404, Message: "Coin not found"}
	}

	return &CoinInfoResponse{
		Name:           coin.Name(),
		Symbol:         coin.Symbol().String(),
		Volume:         coin.Volume().String(),
		Crr:            coin.Crr(),
		ReserveBalance: coin.Reserve().String(),
		MaxSupply:      coin.MaxSupply().String(),
	}, nil
}
