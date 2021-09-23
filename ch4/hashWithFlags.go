// This program prints the SHA256 hash of the standard input by default,
// but supports a command line flag to print SHA384 or SHA512 has instead
package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

func main() {
	hashFlag := flag.Int("hash", 256, "type of hash [256/384/512]")
	flag.Parse()
	if *hashFlag != 256 && *hashFlag != 384 && *hashFlag != 512 {
		fmt.Println("Invalid flag value")
		return
	}
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	if *hashFlag == 256 {
		fmt.Printf("%x\n", sha256.Sum256([]byte(str)))
	} else if *hashFlag == 384 {
		fmt.Printf("%x\n", sha512.Sum384([]byte(str)))
	} else if *hashFlag == 512 {
		fmt.Printf("%x\n", sha512.Sum512([]byte(str)))
	}
}
