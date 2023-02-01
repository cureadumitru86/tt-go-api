package logic

import (
	"errors"
	"math/rand"
)

// CalcFib - Calculates, with recursion, the random Nth fibbonacci number between the min and max value associated.
func CalcFib(min int, max int) (int, int, error) {
	if max < min {
		return -1, -1, errors.New("min is larger than max value")
	}
	num := rand.Intn(max-min) + min
	i := fib(num)

	return num, i, nil
}

func fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}
