package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Create a big integer with value 2^255
	bigIntValue := new(big.Int).Exp(big.NewInt(2), big.NewInt(256), nil)

	// Take the absolute value
	absoluteValue := new(big.Int).Abs(bigIntValue)

	// Print the result
	fmt.Printf("Original: %s\nOriginal: %s\n", bigIntValue.String(), absoluteValue.String())

}
