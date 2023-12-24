package math

func SieveOfEratosthenes(n int) []int {
	primes := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		primes[i] = true
	}

	for p := 2; p*p <= n; p++ {
		if primes[p] == true {
			for i := p * p; i <= n; i += p {
				primes[i] = false
			}
		}
	}

	var primeNumbers []int
	for p := 2; p <= n; p++ {
		if primes[p] {
			primeNumbers = append(primeNumbers, p)
		}
	}

	return primeNumbers
}
