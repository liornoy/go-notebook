// CheckDup get a file name and checks if the file has a duplicate lines.
// Main function get file names in command args and prints the file name of
// those who has duplicate lines in them.
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	files := os.Args[1:]
	if len(files) == 0 {
		fmt.Println("no command line args")
		return
	}
	for _, fileName := range files {
		counts := make(map[string]int)
		hasDup := CheckDup(fileName, counts)
		if hasDup {
			fmt.Println(fileName)
		}
	}
}

func CheckDup(fileName string, counts map[string]int) bool {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return false
	}
	for _, line := range strings.Split(string(data), "\n") {
		counts[line]++
		if counts[line] > 1 {
			return true
		}
	}
	return false
}
