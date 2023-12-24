package asymmetric

import (
	"asymmetric-encr/internal/gears/codec/gob"
	"math/big"
)

type RSAEncryptor struct {
	keys *RSAKeys
}

func NewRSAEncryptor(keys *RSAKeys) (*RSAEncryptor, error) {
	return &RSAEncryptor{
		keys: keys,
	}, nil
}

func (r *RSAEncryptor) Encrypt(msg []byte) ([]byte, error) {
	return gob.SliceToBytes(encryptBytes(msg, r.keys.Pub, r.keys.Mod))
}

func (r *RSAEncryptor) Decrypt(cipheredbytes []byte) ([]byte, error) {
	ciphertext, err := gob.BytesToSlice[*big.Int](cipheredbytes)
	if err != nil {
		return nil, err
	}

	return decryptBytes(ciphertext, r.keys.Priv, r.keys.Mod), nil
}

func (r *RSAEncryptor) PubKey() []byte {
	return r.keys.PubKey()
}

func modExp(base, exponent, modulus *big.Int) *big.Int {
	result := new(big.Int).SetInt64(1)
	result.Mul(result, base)
	return result.Exp(base, exponent, modulus)
}

func encryptBytes(msg []byte, pub, mod *big.Int) []*big.Int {
	result := make([]*big.Int, len(msg))

	for i, b := range msg {
		bInt := int64(b)
		enc := big.NewInt(1).Exp(big.NewInt(bInt), pub, mod)
		result[i] = enc
	}

	return result
}

func decryptBytes(msg []*big.Int, d, n *big.Int) []byte {
	result := make([]byte, 0)
	for _, b := range msg {
		dec := big.NewInt(1).Exp(b, d, n)
		result = append(result, dec.Bytes()...)
	}

	return result
}
