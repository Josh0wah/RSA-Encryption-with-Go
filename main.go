// This program can encrypt a message txt file
// containing numbers using RSA encryption

package main

import (
	"math"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	switch os.Args[1] {
	case "-e":
		encrypt(string(os.Args[2]))
	case "-d":
		decrypt(string(os.Args[2]), os.Args[3], os.Args[4])
	default:
		panic("Invalid command")
	}
}

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

func encrypt(fileName string) {
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

	fKeys, err := os.Create("keys.txt")
	if err != nil {
		panic(err)
	}
	defer fKeys.Close()

	keys := "Private key (e, n): " + strconv.FormatUint(e, 10) + " " + strconv.FormatUint(n, 10) + "\nPublic key (d, n): " + strconv.FormatUint(d, 10) + " " + strconv.FormatUint(n, 10) + "\n"
	_, err = fKeys.WriteString(keys)
	if err != nil {
		panic(err)
	}

	message, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	toInt, _ := strconv.Atoi(string(message))
	encrypted := uint64(math.Pow(float64(toInt), float64(e))) % n

	fEncrypted, err := os.Create("encrypted.txt")
	if err != nil {
		panic(err)
	}
	defer fEncrypted.Close()

	_, err = fEncrypted.WriteString(strconv.FormatUint(encrypted, 10) + "\n")
	if err != nil {
		panic(err)
	}
}

func decrypt(fileName string, d string, n string) {
	dInt, _ := strconv.Atoi(d)
	nInt, _ := strconv.Atoi(n)

	encrypted, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	toInt, _ := strconv.Atoi(string(encrypted))
	decrypted := uint64(math.Pow(float64(toInt), float64(dInt))) % uint64(nInt)

	fDecrypted, err := os.Create("decrypted.txt")
	if err != nil {
		panic(err)
	}
	defer fDecrypted.Close()

	_, err = fDecrypted.WriteString(strconv.FormatUint(decrypted, 10) + "\n")
	if err != nil {
		panic(err)
	}
}
