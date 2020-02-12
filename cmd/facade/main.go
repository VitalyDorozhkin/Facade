package main

import (
	"../../pkg/facade"
	"fmt"
)

func main() {
	computer := facade.NewComputerFacade()
	computer.Start()
	position := computer.Write(11)
	fmt.Printf("position: %d", position)
}
