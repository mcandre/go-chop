package main

import (
	"./chop"
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		fmt.Println(chop.Chomp(scanner.Text()))
	}
}
