package main

import (
	"fmt"

	"github.com/rmasci/toolkit"
)

func main() {
	var tools toolkit.Tools
	s := tools.RandomStringSource(10)
	fmt.Println("Random string", s)
}
