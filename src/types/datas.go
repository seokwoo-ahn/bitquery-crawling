package types

type Timestamp struct {
	Time     string
	Unixtime float64
}

type BitcoinData struct {
	Timestamp        Timestamp
	Height           float64
	BlockHash        string
	Difficulty       float64
	TransactionCount float64
}

type EhtereumData struct {
	Timestamp        Timestamp
	Height           float64
	BlockHash        string
	Difficulty       float64
	TransactionCount float64
}
