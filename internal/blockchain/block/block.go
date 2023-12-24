package block

import (
	"asymmetric-encr/internal/blockchain/transactions"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"time"

	"github.com/onrik/gomerkle"
)

type Block struct {
	Header       Header
	Hash         []byte
	Transactions []transactions.Transaction
}

func (b *Block) HexHash() string {
	return hex.EncodeToString(b.Hash)
}

type Header struct {
	PrevHash       []byte
	MerkleRootHash []byte
	Timestamp      uint64
	Target         []byte
	Nonce          uint64
}

func (h *Header) Serialize() []byte {
	buf := append(h.PrevHash, h.MerkleRootHash...)
	buf = append(buf, h.Target...)
	bLen := len(buf)
	buf = append(buf, make([]byte, binary.MaxVarintLen64*2)...)
	binary.PutVarint(buf[bLen:], int64(h.Timestamp))
	binary.PutUvarint(buf[bLen+binary.MaxVarintLen64:], h.Nonce)
	return buf
}

func Genesis(txs []transactions.Transaction) Block {
	return Block{
		Header: Header{
			MerkleRootHash: merkleRootHash(txs),
			Timestamp:      uint64(time.Now().Unix()),
		},
		Transactions: txs,
	}
}

func New(prevHash []byte, txs []transactions.Transaction) Block {
	return Block{
		Header: Header{
			PrevHash:       prevHash,
			MerkleRootHash: merkleRootHash(txs),
			Timestamp:      uint64(time.Now().Unix()),
		},
		Transactions: txs,
	}
}

func merkleRootHash(txs []transactions.Transaction) []byte {
	tree := gomerkle.NewTree(sha256.New())
	for _, tx := range txs {
		tree.AddHash(tx.ID)
	}

	tree.Generate()
	return tree.Root()
}
