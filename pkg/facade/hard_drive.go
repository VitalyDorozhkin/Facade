package facade

import "fmt"

type HardDrive struct {
	size int
	disc []byte
	head []bool
}

func (h *HardDrive) free(position int) {
	if position < len(h.head) {
		h.head[position] = false
	}
}

func (h *HardDrive) cleanAll() {
	for i := range h.head {
		h.head[i] = false
	}
}

func (h *HardDrive) read(position int) byte {

	if position >= len(h.disc) {
		fmt.Printf("HardDrive don't has %d pointer\n", position)
		return 0
	}
	fmt.Printf("HardDrive read from %d\n", position)
	return h.disc[position]
}

func (h *HardDrive) write(position int, data byte) {
	if position >= len(h.disc) {
		fmt.Printf("HardDrive don't has %d pointer\n", position)
		return
	}
	fmt.Printf("HardDrive write %b to %d\n", data, position)
	h.disc[position] = data
}

func (h *HardDrive) writeFirst(data byte) (i int) {
	for i = 0; i < len(h.disc); i++ {
		if !h.head[i] {
			h.disc[i] = data
			h.head[i] = true
			break
		}
	}
	return i
}
