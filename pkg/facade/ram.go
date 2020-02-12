package facade

import "fmt"

type Ram struct {
	disk [4]byte
}

func (r *Ram) read(position int) byte {
	if position >= len(r.disk) {
		fmt.Printf("Ram don't has %d pointer\n", position)
		return 0
	}
	fmt.Printf("Ram read from %d\n", position)
	return r.disk[position]
}

func (r *Ram) write(position int, data byte) {
	if position >= len(r.disk) {
		fmt.Printf("Ram don't has %d pointer\n", position)
		return
	}
	fmt.Printf("Ram write %b to %d\n", data, position)
	r.disk[position] = data
}
