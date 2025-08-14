// This program can encrypt a message txt file
// containing numbers using RSA encryption

package main

import (
	"fmt"
	"math/rand"
	"os"
)

func findPrime() uint64 {
	testNum := rand.Uint64()%100 + 10
	for i := 2; uint64(i) < testNum; i++ {
		if testNum%uint64(i) == 0 {
			return findPrime()
		}
	}
	return testNum
}

func encrypt(fileName string, p uint64, q uint64, n uint64, phi uint64) {

}

func decrypt(fileName string) {

}

func main() {
	p := findPrime()
	q := findPrime()
	var n uint64 = uint64(p) * uint64(q)
	var phi uint64 = (uint64(p - 1)) * (uint64(q - 1))

}
