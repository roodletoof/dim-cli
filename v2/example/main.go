package main

import "github.com/roodletoof/dim-cli/v2"

func main() {
	println(dimcli.PackagePath())
	println(dimcli.FullTypeNameFor[dimcli.Optional[dimcli.Optional[bool]]]())
}
