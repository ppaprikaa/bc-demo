package asymmetric

import (
	"asymmetric-encr/internal/gears/math"
	"math/rand"
)

func generateRandomSetOfPrimes() []int {
	return math.SieveOfEratosthenes(generateMaxPrimeLimit())
}

func generateMaxPrimeLimit() int {
	var maxPrimes int

	for maxPrimes < 75 {
		maxPrimes = rand.Intn(150)
	}

	return maxPrimes
}
