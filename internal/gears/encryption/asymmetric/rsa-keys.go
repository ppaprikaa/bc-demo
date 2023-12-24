package asymmetric

import (
	"asymmetric-encr/internal/gears/math"
	"context"
	"encoding/json"
	"errors"
	"math/big"
	"math/rand"
	"os"
)

type RSAKeys struct {
	Pub  *big.Int `json:"pub"`
	Priv *big.Int `json:"priv"`
	Mod  *big.Int `json:"mod"`
}

func NewRSAKeys(ctx context.Context) (*RSAKeys, error) {
	Pub, Priv, Mod := generateKeys(ctx)
	return &RSAKeys{
		Pub:  Pub,
		Priv: Priv,
		Mod:  Mod,
	}, nil
}

func (r *RSAKeys) Set(pub, priv, mod []byte) {
	r.Pub.SetBytes(pub)
	r.Priv.SetBytes(priv)
	r.Mod.SetBytes(mod)
}

func (r *RSAKeys) PubKey() []byte {
	return r.Pub.Bytes()
}

func (r *RSAKeys) SaveToFile(filepath string) error {
	keysBytes, err := r.Serialize()
	if err != nil {
		return err
	}

	if err = os.WriteFile(filepath, keysBytes, 0644); err != nil {
		return err
	}

	return nil
}

func (r *RSAKeys) ReadFromFile(filepath string) error {
	keysBytes, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}

	if err = r.Deserialize(keysBytes); err != nil {
		return err
	}

	return nil
}

func (r *RSAKeys) Serialize() ([]byte, error) {
	return json.Marshal(r)
}

func (r *RSAKeys) Deserialize(keys []byte) error {
	return json.Unmarshal(keys, &r)
}

func generateKeys(ctx context.Context) (Pub, Priv, Mod *big.Int) {
	var err error

	phi := big.NewInt(1)
	Mod = big.NewInt(1)
	for {
		p, q := generateRSAPrimes()
		phi = phi.Mul(big.NewInt(p.Int64()-1), big.NewInt(q.Int64()-1))
		Mod = Mod.Mul(p, q)
		Pub = extractRSAPubKey(phi)
		Priv, err = extractRSAPrivKey(p, q, Pub, phi)
		if err != nil {
			continue
		}

		if p.Int64() < Priv.Int64() && q.Int64() < Priv.Int64() && Priv.Int64() < phi.Int64() && Priv.Int64() != Pub.Int64() {
			return Pub, Priv, Mod
		}
	}
}

func extractRSAPubKey(phi *big.Int) *big.Int {
	primesBeforePhi := math.SieveOfEratosthenes(int(phi.Int64()))
	return big.NewInt(int64(primesBeforePhi[len(primesBeforePhi)-1]))
}

func extractRSAPrivKey(p, q, e, phi *big.Int) (*big.Int, error) {
	gcd, x, _ := math.ExtendedGCD(int(e.Int64()), int(phi.Int64()))
	var d int64 = 1

	if gcd == 1 {
		d = int64(x) % phi.Int64()
		return big.NewInt(int64(d)), nil
	}

	return nil, errors.New("invalid phi")
}

func generateRSAPrimes() (P, Q *big.Int) {
	primes := generateRandomSetOfPrimes()

	p := primes[rand.Intn(len(primes))]
	q := p

	for q == p {
		q = primes[rand.Intn(len(primes))]
	}
	return big.NewInt(int64(p)), big.NewInt(int64(q))
}
