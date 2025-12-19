package main

import (
	"fmt"

	"github.com/roodletoof/dim-cli/v2"
)

func main() {
	var cli = dimcli.NewCli()

	var fubar = cli.Flag("--fubar")
	if fubar {
		println("called with fubar flag")
	}

	if cli.Command("fu") {
		f, err := cli.Float("--some-float")
		if cli.Help { return }

		if err != nil {
			fmt.Printf("err: %v\n", err)
			return
		}
		fmt.Println(f)
	} else if cli.Command("bar") {
		i, err := cli.Int("--some-int")
		if cli.Help { return }
		if err != nil {
			fmt.Printf("err: %v\n", err)
			return
		}

		fmt.Println(i)
	}

	if cli.Help { return }
}
