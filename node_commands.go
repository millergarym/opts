package opts

import (
	"fmt"
	"os"
)

func (n *node) AddCommand(cmd SubOpts) Opts {
	sub, ok := cmd.(*node)
	if !ok {
		panic("another implementation of opts???")
	}
	n.addCommand(sub)
	return n
}

func (n *node) SubAddCommand(cmd SubOpts) SubOpts {
	sub, ok := cmd.(*node)
	if !ok {
		panic("another implementation of opts???")
	}
	n.addCommand(sub)
	return n
}

func (n *node) addCommand(sub *node) {
	//if still no name, needs to be manually set
	if sub.name == "" {
		n.errorf("cannot add command, please set a Name()")
		return
	}
	if _, exists := n.cmds[sub.name]; exists {
		n.errorf("cannot add command, '%s' already exists", sub.name)
		return
	}
	sub.parent = n
	n.cmds[sub.name] = sub
}

func (n *node) matchedCommand() *node {
	if n.cmd != nil {
		return n.cmd.matchedCommand()
	}
	return n
}

//IsRunnable
func (n *node) IsRunnable() bool {
	ok, _ := n.run(true)
	return ok
}

//Run the parsed configuration
func (n *node) Run() error {
	_, err := n.run(false)
	return err
}

type runner1 interface {
	Run() error
}

type runner2 interface {
	Run()
}

func (n *node) run(test bool) (bool, error) {
	m := n.matchedCommand()
	v := m.val.Addr().Interface()
	r1, ok1 := v.(runner1)
	r2, ok2 := v.(runner2)
	if test {
		return ok1 || ok2, nil
	}
	if ok1 {
		return true, r1.Run()
	}
	if ok2 {
		r2.Run()
		return true, nil
	}
	if len(m.cmds) > 0 {
		//if matched command has no run,
		//but has commands, show help instead
		return false, exitError(m.Help())
	}
	return false, fmt.Errorf("command '%s' is not runnable", m.name)
}

//Run the parsed configuration
func (n *node) RunFatal() {
	if err := n.Run(); err != nil {
		if e, ok := err.(exitError); ok {
			fmt.Fprint(os.Stderr, string(e))
			os.Exit(0)
		}
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}
}
