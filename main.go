package main

import (
	"github.com/pacur/pacur/cmd"
)

func main() {
	err := cmd.Parse()
	if err != nil {
		panic(err)
	}
}
