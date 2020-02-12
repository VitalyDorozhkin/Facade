package facade

import "fmt"

type HardDrive struct {
	disk [16]byte
}

func (h *HardDrive) read(position int) byte {

	if position >= len(h.disk) {
		fmt.Printf("HardDrive don't has %d pointer\n", position)
		return 0
	}
	fmt.Printf("HardDrive read from %d\n", position)
	return h.disk[position]
}

func (h *HardDrive) write(position int, data byte) {
	if position >= len(h.disk) {
		fmt.Printf("HardDrive don't has %d pointer\n", position)
		return
	}
	fmt.Printf("HardDrive write %b to %d\n", data, position)
	h.disk[position] = data
}
