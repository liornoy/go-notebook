package main
//This program prints the command line args and
//compare between string concatenation and the string.Join function
import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	ex1()
	ex2()
}

func fastPrint() {
	defer timeTrack(time.Now(), "fast print")
	fmt.Println(strings.Join(os.Args, " "))
}

func slowPrint() {
	defer timeTrack(time.Now(), "slow print")
	var s, space string
	for _, arg := range os.Args {
		s += space + arg 
		space = " "
	}
	fmt.Println(s)
}

func ex1() {
	for i, arg := range os.Args {
		fmt.Println(i, ": ", arg)
	}
}

func ex2() {
	slowPrint()
	fastPrint()
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", name, elapsed)
}
