package bar

import (
	"fmt"

	"github.com/jpillora/opts"
)

type cmd struct {
	Zip string
	Zop string
}

func New() opts.SubOpts {
	c := cmd{}
	//default name for a subcommand is its package name ("bar")
	return opts.NewSub(&c)
}

func (b *cmd) Run() error {
	fmt.Printf("bar: %+v\n", b)
	return nil
}
