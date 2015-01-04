package main

import (
	"fmt"
	"github.com/obsc/golam/lam"
	"os"
)

func main() {
	lines := lam.Scanner(os.Stdin)
	for {
		line := <-lines
		fmt.Printf("%s", line)
	}
	fmt.Println("end")
}
