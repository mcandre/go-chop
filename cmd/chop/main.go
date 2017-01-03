// Package main provides a chop executable.
package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/mcandre/go-chop"
)

// versionBanner prints a command line-accessible version number.
func versionBanner() {
	fmt.Println(os.Args[0], chop.Version)
	os.Exit(0)
}

// main is an entrypoint for launching this application.
func main() {
	if len(os.Args) > 1 && os.Args[1] == "-v" {
		versionBanner()
	}

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		fmt.Println(chop.Chop(scanner.Text()))
	}
}
