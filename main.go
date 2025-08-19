package dimcli

import (
	"slices"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"runtime"
	"strings"
)

// Help takes the string of the file that it is used (use the embed
// package), and searches for the text "HELP STOP" in the code. Then it prints
// the lines of code between.
func Help(stringOfThisFile string) {
	if !slices.Contains(os.Args, "--help") {
		return
	}
	println()

	var _, _, lineNumBegin, ok = runtime.Caller(1)
	if !ok {
		panic("could not get caller information")
	}
	var lines = strings.Split(stringOfThisFile, "\n")
	for i := range lines { // platform independent
		lines[i] = strings.TrimSuffix(lines[i], "\r")
	}
	lines = lines[lineNumBegin:]
	for _, line := range lines {
		if strings.Contains(line, "HELP STOP") {
			break
		}
		println(line)
	}
	println()
	os.Exit(0)
}

var position = 1
func GetPositional[T any]() (out T) {
	var t = reflect.TypeFor[T]()
	if len(os.Args) <= position {
		var bytes, err = json.Marshal(out)
		if err != nil {
			panic(
				fmt.Sprintf(
					"type %s cannot be marshalled",
					t.Name(),
				),
			)
		}
		panic(
			fmt.Sprintf(
				"missing positional arg %d of type %s\nexample: %s",
				position,
				t.Name(),
				string(bytes),
			),
		)
	}

	var err = json.Unmarshal([]byte(os.Args[position]), &out)
	if err != nil {
		panic(
			fmt.Sprintf(
				"failed parsing %s into arg %d of type %s",
				os.Args[position],
				position,
				t.Name(),
			),
		)
	}

	position++
	return out
}

func GetKeyed[T any](name string, otherwise T) (out T) {
	var key = fmt.Sprintf("--%s", name)
	var index = 0
	for i, value := range os.Args {
		if value == key {
			index = i
			break
		}
	}
	if index == 0 {
		return otherwise
	}
	if (index+1) >= len(os.Args) {
		var t = reflect.TypeFor[T]()

		panic(
			fmt.Sprintf(
				"did not get value for keyed arg --%s.\nexpected value of type %s, kind %s",
				name,
				t.Name(),
				t.Kind(),
			),
		)
	}
	var err = json.Unmarshal(
		[]byte(os.Args[index+1]),
		&out,
	)
	if err != nil {
		panic(err.Error())
	}
	return out
}

func Flag(name string) bool {
	return slices.Contains(os.Args, fmt.Sprintf("--%s", name))
}
