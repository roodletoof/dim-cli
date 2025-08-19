# What is this?

A super dim library for clis. It is intended to be 5% of the work for 60% of
the results.

# Usage

example:

```go
package main

import (
	"github.com/roodletoof/dim-cli"
	_ "embed"
)

// note: main.go is *this* source file

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
```

The above program supports required positional parameters, optional keyed
parameters, e.g. "--key value", flags "--some-flag", and a simple --help
command.

If anything is wrong with the arguments provided the program will panic with a
pretty descriptive message. I think this is appropriate for most situations.
    
The Help function expects you to provide a string that contains the source code
of where it was called. Use the embed package for this as in the example. It
will print out all the source code between the call and the first instance
where "HELP STOP" is found in the source code, if the --help flag is given to
the program.

The library does not check that you don't provide extra arguments, it only
cares whether the arguments it tries to parse exist.

Positional arguments and keyed arguments are parsed via the json package, so if
you want to provide a list in the example above you would put this in the
argument list: --names '["name1", "name2"]'

This is super scuffed, and I hope you enjoy :)
