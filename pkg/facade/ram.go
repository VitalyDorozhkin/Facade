package facade

import "fmt"

type Ram struct {
	Size int
	disc []byte
	head []bool
}

func (r *Ram) free(position int) {
	if position < len(r.head) {
		r.head[position] = false
	}
}

func (r *Ram) cleanAll() {
	for i := range r.head {
		r.head[i] = false
	}
}

func (r *Ram) read(position int) byte {
	if position >= len(r.disc) {
		fmt.Printf("Ram don't has %d pointer\n", position)
		return 0
	}
	fmt.Printf("Ram read from %d\n", position)
	return r.disc[position]
}

func (r *Ram) write(position int, data byte) {
	if position >= len(r.disc) {
		fmt.Printf("Ram don't has %d pointer\n", position)
		return
	}
	fmt.Printf("Ram write %b to %d\n", data, position)
	r.disc[position] = data
}
