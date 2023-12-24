package transactions

type UTXO struct {
	TransactionID []byte
	OutputIdx     int64
	Output        Output
}
