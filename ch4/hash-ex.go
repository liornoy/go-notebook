package main

// this program gets 2 strings from the user
// and hash it using SHA256, then prints how many bits differ
import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"os"
)

func maain() {
	var s1, s2 string
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter first string to hash:")
	s1, _ = reader.ReadString('\n')
	fmt.Println("Enter second string to hash:")
	s2, _ = reader.ReadString('\n')
	c1 := sha256.Sum256([]byte(s1))
	c2 := sha256.Sum256([]byte(s2))
	diffBitsCount := DiffBitCount(c1, c2)
	fmt.Printf("Hashed strings:\n%x\n%x\n", c1, c2)

	fmt.Println("Number of different bits: ", diffBitsCount)
}
func DiffBitCount(x, y [32]byte) int {
	diffBitsCount := 0
	for i := 0; i < len(x); i++ {
		for j := 0; j < 8; j++ {
			xBit := x[i] & (1 << j)
			yBit := y[i] & (1 << j)
			if xBit != yBit {
				diffBitsCount++
			}
		}
	}
	return diffBitsCount
}
