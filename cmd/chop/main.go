package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/mcandre/go-chop"
)

func versionBanner() {
	fmt.Printf("%s %s \n", os.Args[0], chop.Version)
	os.Exit(0)
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "-v" {
		versionBanner()
	}

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		fmt.Println(chop.Chop(scanner.Text()))
	}
}
