package types

type ChainID byte

const (
	ChainTestnet = 0x02
	ChainMainnet = 0x01

#	CurrentChainID = ChainTestnet
	CurrentChainID = ChainMainnet
)

func GetBaseCoin() CoinSymbol {
	return getBaseCoin(CurrentChainID)
}

func getBaseCoin(chainID ChainID) CoinSymbol {
	var coin CoinSymbol

	switch chainID {
	case ChainMainnet:
		copy(coin[:], "KVANT")
	case ChainTestnet:
		copy(coin[:], "KWANT")
	}

	coin[5] = byte(0)

	return coin
}
