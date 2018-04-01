// +build mage

package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/magefile/mage/mg"
	"github.com/mcandre/go-chop"
	"github.com/mcandre/mage-extras"
)

// artifactsPath describes where artifacts are produced.
var artifactsPath = "bin"

// Default references the default build task.
var Default = Test

// Test executes the integration test suite.
func Test() error {
	mg.Deps(Install)

	exampleFilename := "example.txt"

	exampleFileChop, err := os.Open(exampleFilename)

	if err != nil {
		return err
	}

	cmdChop := exec.Command("chop")
	cmdChop.Stdin = bufio.NewReader(exampleFileChop)
	cmdChop.Stdout = os.Stdout
	cmdChop.Stderr = os.Stderr

	if err := cmdChop.Run(); err != nil {
		return err
	}

	exampleFileChomp, err := os.Open(exampleFilename)

	if err != nil {
		return err
	}

	cmdChomp := exec.Command("chomp")
	cmdChomp.Stdin = bufio.NewReader(exampleFileChomp)
	cmdChomp.Stdout = os.Stdout
	cmdChomp.Stderr = os.Stderr

	return cmdChomp.Run()
}

// GoVet runs go tool vet.
func GoVet() error { return mageextras.GoVet("-shadow") }

// GoLint runs golint.
func GoLint() error { return mageextras.GoLint() }

// Gofmt runs gofmt.
func GoFmt() error { return mageextras.GoFmt("-s", "-w") }

// GoImports runs goimports.
func GoImports() error { return mageextras.GoImports("-w") }

// Errcheck runs errcheck.
func Errcheck() error { return mageextras.Errcheck("-blank") }

// Nakedret runs nakedret.
func Nakedret() error { return mageextras.Nakedret("-l", "0") }

// Lint runs the lint suite.
func Lint() error {
	mg.Deps(GoVet)
	mg.Deps(GoLint)
	mg.Deps(GoFmt)
	mg.Deps(GoImports)
	mg.Deps(Errcheck)
	mg.Deps(Nakedret)
	return nil
}

// portBasename labels the artifact basename.
var portBasename = fmt.Sprintf("chop-%s", chop.Version)

// Port archives build artifacts.
func Port() error { mg.Deps(Gox); return mageextras.Archive(portBasename, artifactsPath) }

// Gox cross-compiles Go binaries.
func Gox() error {
	return mageextras.Gox(
		artifactsPath,
		strings.Join(
			[]string{
				portBasename,
				"{{.OS}}",
				"{{.Arch}}",
				"{{.Dir}}",
			},
			mageextras.PathSeparatorString,
		),
	)
}

// Install builds and installs Go applications.
func Install() error { return mageextras.Install() }

// Uninstall deletes installed Go applications.
func Uninstall() error { return mageextras.Uninstall("chop", "chomp") }

// Clean deletes artifacts.
func Clean() error { return os.RemoveAll(artifactsPath) }
