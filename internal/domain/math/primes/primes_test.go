package primes_test

import (
	"Pietroski/GolangVectisInitialChallange/internal/domain/math/primes"
	"fmt"
	"testing"
	"time"
)

func TestGenPrimesUpToPos(t *testing.T) {
	upToNum := int64(100_003)
	start := time.Now()
	primeList, listSum := primes.GenPrimesUpToPos(upToNum)
	elapsed := time.Since(start) // M1 Silicon perform in about 69.45 milliseconds in average...
	fmt.Println((*primeList)[upToNum-1], listSum, elapsed)
	// t.Errorf("failed")

	// LargestPrimeFactorOfTheThreeMostRightDigitsOf
	start = time.Now()
	lpf := primes.LargestPrimeFactorOfTheMostRightDigitsOf(primeList, listSum, 3)
	elapsed = time.Since(start) // M1 Silicon perform in about 69.45 milliseconds in average...
	fmt.Println(lpf, elapsed)
	// t.Errorf("failed")
}
