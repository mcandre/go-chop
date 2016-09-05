package main

import (
	"bufio"
	"fmt"
	"os"
	"github.com/mcandre/go-chop"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		fmt.Println(chop.Chop(scanner.Text()))
	}
}
