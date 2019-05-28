package foo

import (
	"fmt"

	"github.com/jpillora/opts"
	"github.com/jpillora/opts-examples/eg-commands-register/bar"
)

func New() opts.SubOpts {
	c := cmd{}
	//default name for a subcommand is its package name ("foo")
	return opts.NewSub(&c).SubAddCommand(bar.New())
}

type cmd struct {
	Ping string
	Pong string
}

func (f *cmd) Run() error {
	fmt.Printf("foo: %+v\n", f)
	return nil
}
