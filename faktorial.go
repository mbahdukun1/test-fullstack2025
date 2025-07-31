package main

import (
	"fmt"
	"math"
)

func factorial(n int) int64 {
	if n == 0 {
		return 1
	}
	result := int64(1)
	for i := 2; i <= n; i++ {
		result *= int64(i)
	}
	return result
}


func compute(n int) int64 {
	num := float64(factorial(n))
	denom := math.Pow(2, float64(n))
	return int64(math.Ceil(num / denom))
}


func main() {
	for i := 0; i <= 10; i++ {
		fmt.Printf("f(%d) = %d\n", i, compute(i))
	}
}
