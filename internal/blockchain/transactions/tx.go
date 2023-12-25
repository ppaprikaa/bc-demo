package transactions

import (
	"asymmetric-encr/internal/blockchain/amount"
	"asymmetric-encr/internal/gears/hash"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
)

type Transaction struct {
	ID      []byte
	Inputs  []Input
	Outputs []Output
}

type Input struct {
	TransactionID []byte
	OutputIdx     int64
	Sign          []byte
	PubKey        []byte
}

type Output struct {
	Value int
	Addr  []byte
}

func New(utxos []*UTXO, from, to []byte, a amount.Amount) (*Transaction, error) {
	var (
		Tx    = Transaction{}
		total = amount.Amount(0)
	)

	for i, utxo := range utxos {
		if i == 0 || total < a {
			Tx.Inputs = append(Tx.Inputs, Input{
				TransactionID: utxo.TransactionID,
				OutputIdx:     utxo.OutputIdx,
			})
			total += amount.Amount(utxo.Output.Value)
			continue
		}
		break
	}

	if total < a {
		return nil, errors.New("not enough coins")
	}

	Tx.Outputs = append(Tx.Outputs, Output{
		Value: int(a),
		Addr:  to,
	})

	if total > a {
		Tx.Outputs = append(Tx.Outputs, Output{
			Value: int(total - a),
			Addr:  from,
		})
	}

	var err error

	if Tx.ID, err = Tx.Hash(); err != nil {
		return nil, err
	}

	return &Tx, nil
}

func NewCB(to []byte, gas amount.Amount) (*Transaction, error) {
	var (
		Tx   = Transaction{}
		Salt = make([]byte, 32)
		err  error
	)

	if _, err = rand.Read(Salt); err != nil {
		return nil, err
	}

	Tx.Inputs = append(Tx.Inputs, Input{
		OutputIdx: -1,
		PubKey:    Salt,
	})

	Tx.Outputs = append(Tx.Outputs, Output{
		Value: int(gas),
		Addr:  to,
	})

	if Tx.ID, err = Tx.Hash(); err != nil {
		return nil, err
	}

	return &Tx, nil
}

func (t *Transaction) Serialize() ([]byte, error) {
	t.ID = nil
	return json.Marshal(&t)
}

func (t *Transaction) Hash() ([]byte, error) {
	bytes, err := t.Serialize()
	if err != nil {
		return nil, err
	}

	return hash.DoubleSHA256Hash(bytes), nil
}

func (t *Transaction) HexID() string {
	return hex.EncodeToString(t.ID)
}
