package main

import (
	"github.com/aanno/pacur/v2/command"
)

func main() {
	err := command.Parse()
	if err != nil {
		panic(err)
	}
}
