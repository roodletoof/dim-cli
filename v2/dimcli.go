package dimcli

import (
	"fmt"
	"os"
	"slices"
	"strconv"
)

type Cli struct {
	Args []string
	Help bool
	Indentation int
}

func NewCli() Cli {
	var cli = Cli{
		Args: os.Args[1:],
		Indentation: 0,
	}
	cli.Help = cli.Flag("-h") || cli.Flag("--help")
	return cli
}

func (c *Cli) Command(name string) bool {
	if (c.Help) {
		c.printIndentation()
		fmt.Printf("[ cmd ] %v\n", name)
	}
	if len(c.Args) == 0 {
		return false
	}
	if c.Args[0] == name {
		c.Args = c.Args[1:]
		c.Indentation++
		return true
	}
	return false
}

func (c Cli) Flag(name string) bool {
	if c.Help {
		c.printIndentation()
		fmt.Printf("[ flg ] %v\n", name)
		return false
	}
	return slices.Contains(c.Args, name)
}

func (c Cli) printIndentation() {
	for range c.Indentation {
		fmt.Printf("\t")
	}
}

func (c Cli) String(name string) (string, error) {
	if c.Help {
		c.printIndentation()
		fmt.Printf("[ str ] %v\n", name)
		return "", nil
	}
	var index int
	for i, arg := range c.Args {
		if name == arg {
			index = i+1
			break
		}
	}
	if len(c.Args) <= index {
		return "", fmt.Errorf("No value provided after %v", name)
	}
	return c.Args[index], nil
}

func (c Cli) Int(name string) (int64, error) {
	if c.Help {
		c.printIndentation()
		fmt.Printf("[ int ] %v\n", name)
		return 0, nil
	}
	s, err := c.String(name)
	if err != nil {
		return 0, err
	}
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return i, nil
}

func (c Cli) Float(name string) (float64, error) {
	if c.Help {
		c.printIndentation()
		fmt.Printf("[ flt ] %v\n", name)
		return 0, nil
	}
	s, err := c.String(name)
	if err != nil {
		return 0, err
	}
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, err
	}
	return f, nil
}
