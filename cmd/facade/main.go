package main

import (
	"../../pkg/facade"
)

func main() {
	computer := facade.NewComputerFacade()
	computer.Start()
}