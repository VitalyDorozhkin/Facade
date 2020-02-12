package facade

import "fmt"

type CPU struct {
	pointer int
	cache   byte
}

func (c *CPU) freeze() {
	fmt.Println("CPU freeze")
	c.pointer = 0
}

func (c *CPU) setCache(memory Memory) {
	fmt.Println("CPU write cache")
	c.cache = memory.read(c.pointer)
}

func (c *CPU) loadCache(memory Memory) {
	fmt.Println("CPU load cache to memory")
	memory.write(c.pointer, c.cache)
}

func (c *CPU) jump(position int) {
	fmt.Printf("CPU jump to %d\n", position)
	c.pointer = position
}

func (c *CPU) execute() int {
	fmt.Println("CPU execute")
	return c.pointer
}
