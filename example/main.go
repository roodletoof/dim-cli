package main

import (
	"github.com/roodletoof/dim-cli"
	_ "embed"
)

//go:embed main.go
var thisFile string

func main() {
	dimcli.Help(thisFile)
	var age = dimcli.GetPositional[int]()
	var names = dimcli.GetKeyed(
		"names",
		[]string{"Alice", "Bob"},
	)
	var printBar = dimcli.Flag("print-bar")
	// HELP STOP

	println(age)
	for _, name := range names {
		println(name)
	}
	if printBar {
		println("bar")
	}
}
