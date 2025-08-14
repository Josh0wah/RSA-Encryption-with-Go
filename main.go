// This program can encrypt a message txt file
// containing numbers using RSA encryption

package main

import (
	"fmt"
	"math/rand"
	"os"
)

func findPrime() uint {
	testNum := rand.Int()%100 + 10
	for i := 2; i < testNum; i++ {
		if testNum%i == 0 {
			return findPrime()
		}
	}
	return uint(testNum)
}

func encrypt(fileName string, p uint, q uint, n uint64, phi uint64) {

}

func decrypt(fileName string) {

}

func main() {
	p := findPrime()
	q := findPrime()
	var n uint64 = uint64(p) * uint64(q)
	var phi uint64 = (uint64(p - 1)) * (uint64(q - 1))

}
