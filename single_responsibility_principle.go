package main

import (
	"fmt"
)

type Command struct {
	Args [][]byte
}

func (c *Command) Encode(x []string) {
	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
		c.Args = append(c.Args, []byte(x[i]))
	}
	fmt.Println(c.Args)
}

func (c *Command) Decode() {
	for i := 0; i < len(c.Args); i++ {
		str := string(c.Args[i])
		fmt.Println("String =", str)
	}
}
