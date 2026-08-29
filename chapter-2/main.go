package main

import (
	"fmt"

	"golang.og/chapter-2/util"
)

func main() {
	util.PrintMessage("Hello, Raghav")
	util.PrintMessage(fmt.Sprint(util.Add(2, 3)))
}
