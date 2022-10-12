package types

type Timestamp struct {
	Time     string
	Unixtime float64
}

type BitcoinBlockData struct {
	Timestamp        Timestamp
	Height           float64
	BlockHash        string
	Difficulty       float64
	TransactionCount float64
}

type BitcoinTxData struct {
	BitcoinBlockData BitcoinBlockData
	Count            float64
	Hash             string
	InputValue       float64
	OutputValue      float64
}

type EthereumBlockData struct {
	Timestamp        Timestamp
	Height           float64
	BlockHash        string
	Difficulty       float64
	TransactionCount float64
}
