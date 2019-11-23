package main

import (
	"github.com/m0rf30/pacur/cmd"
)

func main() {
	err := cmd.Parse()
	if err != nil {
		panic(err)
	}
}
