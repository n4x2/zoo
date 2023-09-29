// Package command provides a simple structure for single command.
package command

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

// HandlerFunc used to define the function that handles a specific
// task for a command.
type HandlerFunc func(fs *flag.FlagSet, args []string) error

// Command represents a command.
type Command struct {
	N  string            // Name of the command.
	B  string            // Brief explanation of the command.
	U  string            // Usage information for the command.
	E  map[string]string // Examples of the command.
	Fs *flag.FlagSet     // Flag set for the command.
	Fn HandlerFunc       // Function to execute the command's task.
}

// Help displays usage information, flag details, and examples for
// the command (if any).
func (c *Command) Help() {
	if c.U != "" {
		fmt.Fprintln(os.Stderr, "Usage:")
		fmt.Fprintf(os.Stderr, "  %s\n\n", c.U)
	}

	fmt.Fprintln(os.Stderr, "Flags:")
	c.Fs.PrintDefaults()

	if len(c.E) != 0 {
		fmt.Fprintln(os.Stderr, "\nExamples:")
		for k, v := range c.E {
			fmt.Fprintf(os.Stderr, "  %s\n\t%s\n", k, v)
		}
	}
	fmt.Fprintln(os.Stderr, "\nUse '--help' to show this text.")
}

// AddExample adds an example 'e' with its associated value 'v' to
// the command. If 'e' or 'v' is empty, it returns an error.
func (c *Command) AddExample(e, v string) error {
	if e == "" {
		return errors.New("example name must be not empty")
	}

	if v == "" {
		return errors.New("example value must be not empty")
	}

	c.E[e] = v

	return nil
}

// New returns a new command with the specified name 'n', 'b' a brief
// explanation of the command, its usage 'u', and a [HandlerFunc] 'h'
// to handle a specific task.
func New(n, b, u string, h HandlerFunc) *Command {
	var c = Command{
		N:  n,
		B:  b,
		U:  u,
		Fs: flag.NewFlagSet(n, flag.ExitOnError),
		Fn: h,
		E:  make(map[string]string),
	}

	c.Fs.Usage = func() { c.Help() }

	return &c
}
