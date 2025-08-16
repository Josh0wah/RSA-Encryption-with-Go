// This program can encrypt a message txt file
// containing numbers using RSA encryption

package main

import (
	"fmt"
	"math"
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

func gcd(a, b uint64) uint64 {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func encrypt(fileName string, p uint64, q uint64, n uint64, phi uint64) {

}

func decrypt(fileName string) {

}

func main() {
	p := findPrime()
	q := findPrime()
	for q == p {
		q = findPrime()
	}
	var n uint64 = uint64(p) * uint64(q)
	var phi uint64 = (uint64(p - 1)) * (uint64(q - 1))
	var e uint64
	for e = 2; e < phi; e++ {
		if gcd(e, phi) == 1 {
			break
		}
	}
	var d uint64
	for d = 2; d < phi; d++ {
		if (e*d)%phi == 1 {
			break
		}
	}

	fKeys, err := os.Create("files/keys.txt")
	if err != nil {
		panic(err)
	}
	defer fKeys.Close()

	keys := "Private key (e, n): " + string(e) + " " + string(n) + "\n Public key (d, n): " + string(d) + " " + string(n)
	_, err = fKeys.WriteString(keys)
	if err != nil {
		panic(err)
	}

}
