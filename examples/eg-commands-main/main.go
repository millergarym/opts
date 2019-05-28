package main

import (
	"fmt"

	"github.com/jpillora/opts"
)

type Config struct{}

func main() {
	opts.New(&Config{}).
		AddCommand(
			opts.NewSub(&Foo{}).
				SubAddCommand(
					opts.NewSub(&Bar{}),
				),
		).
		Parse().
		Run()
}

type Foo struct {
	Ping string
	Pong string
}

func (f *Foo) Run() {
	fmt.Printf("foo: %+v\n", f)
}

type Bar struct {
	Zip string
	Zop string
}

func (b *Bar) Run() {
	fmt.Printf("bar: %+v\n", b)
}
