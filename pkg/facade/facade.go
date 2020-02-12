package facade

import "fmt"

type Memory interface {
	read(position int) byte
	write(position int, data byte)
}

type CPU struct {
	pointer int
	cache   byte
}

type HardDrive struct {
	disk [16]byte
}

type Ram struct {
	disk [4]byte
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

/* Facade */
type ComputerFacade struct {
	processor *CPU
	ram       Memory
	hd        Memory
}

func (c *ComputerFacade) Start() int {
	c.processor.freeze()
	c.ram.write(0, c.hd.read(0))
	c.processor.jump(8)
	return c.processor.execute()
}

func (c *ComputerFacade) Swap(ramPos int, drivePos int) {
	c.processor.jump(ramPos)
	c.processor.setCache(c.ram)
	c.processor.jump(drivePos)
	c.processor.loadCache(c.hd)
}

func NewComputerFacade() *ComputerFacade {
	return &ComputerFacade{
		processor: &CPU{},
		ram:       &Ram{},
		hd:        &HardDrive{},
	}
}
