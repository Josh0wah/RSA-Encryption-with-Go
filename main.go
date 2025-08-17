// This program can encrypt a message txt file
// containing numbers using RSA encryption

package main

import (
	"fmt"
	"math/big"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	switch os.Args[1] {
	case "-e":
		//         file name
		encrypt(os.Args[2])
	case "-d":
		//          file name           d          n
		decrypt(os.Args[2], os.Args[3], os.Args[4])
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
		if (d*e)%phi == 1 {
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

	messageBytes, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	message := string(messageBytes)
	message = strings.Trim(message, "\n")

	toInt, err := strconv.Atoi(message)
	if err != nil {
		panic(err)
	}

	m := big.NewInt(int64(toInt))
	eBig := big.NewInt(int64(e))
	nBig := big.NewInt(int64(n))

	encryptedBig := new(big.Int).Exp(m, eBig, nBig)
	encrypted := encryptedBig.Uint64()

	fEncrypted, err := os.Create("encrypted.txt")
	if err != nil {
		panic(err)
	}
	defer fEncrypted.Close()

	_, err = fEncrypted.WriteString(strconv.FormatUint(encrypted, 10) + "\n")
	if err != nil {
		panic(err)
	}

	fmt.Print("Encryption successful. keys.txt and encrypted.txt have been created.\n")
}

func decrypt(fileName string, d string, n string) {
	dInt, err := strconv.Atoi(d)
	if err != nil {
		panic(err)
	}
	nInt, err := strconv.Atoi(n)
	if err != nil {
		panic(err)
	}

	encryptedBytes, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	encrypted := string(encryptedBytes)
	encrypted = strings.Trim(encrypted, "\n")

	toInt, err := strconv.Atoi(string(encrypted))
	if err != nil {
		panic(err)
	}

	dBig := big.NewInt(int64(dInt))
	nBig := big.NewInt(int64(nInt))

	encryptedBig := big.NewInt(int64(toInt))
	decryptedBig := new(big.Int).Exp(encryptedBig, dBig, nBig)
	decrypted := decryptedBig.Uint64()

	fDecrypted, err := os.Create("decrypted.txt")
	if err != nil {
		panic(err)
	}
	defer fDecrypted.Close()

	_, err = fDecrypted.WriteString(strconv.FormatUint(decrypted, 10) + "\n")
	if err != nil {
		panic(err)
	}

	fmt.Print("Decryption successful. decrypted.txt has been created.\n")
}
