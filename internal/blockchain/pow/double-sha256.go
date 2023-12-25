package pow

import (
	"asymmetric-encr/internal/blockchain/block"
	"asymmetric-encr/internal/gears/hash"
	"math"
	"math/big"
)

type DoubleSHA256 struct {
	target []byte
}

func NewDoubleSHA256(target []byte) *DoubleSHA256 {
	return &DoubleSHA256{
		target: target,
	}
}

func (s *DoubleSHA256) Solve(b *block.Block) error {
	var candidateNonce uint64
	for candidateNonce < math.MaxUint64 {
		b.Header.Nonce = candidateNonce
		h := hash.DoubleSHA256Hash(b.Header.Serialize())
		hashInt := big.NewInt(0).SetBytes(h)
		targetHashInt := big.NewInt(0).SetBytes(s.target)

		if hashInt.Cmp(targetHashInt) <= 0 {
			b.Hash = h
			b.Header.Target = s.target
			return nil
		}

		candidateNonce++
	}

	return ErrFailedToSolve
}

func (s *DoubleSHA256) Verify(b *block.Block) error {
	target := big.NewInt(0).SetBytes(b.Header.Target)
	hash := big.NewInt(0).SetBytes(b.Hash)

	if hash.Cmp(target) > 0 {
		return block.ErrInvalidBlockHash
	}

	return nil
}
