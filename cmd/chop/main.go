package main

import (
	"bufio"
	"fmt"
	"github.com/mcandre/go-chop"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		fmt.Println(chop.Chop(scanner.Text()))
	}
}
