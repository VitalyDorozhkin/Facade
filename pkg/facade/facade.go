package facade

import "fmt"

type ComputerFacade struct {
	processor *CPU
	ram       *Ram
	hd        *HardDrive
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

func (c *ComputerFacade) Write(data byte) (position int) {
	position = c.hd.writeFirst(data)
	if position < len(c.hd.disc) {
		fmt.Printf("Data loaded on HardDrive[%d]\n", position)
		return
	}
	fmt.Println("Hard drive is full!")
	return
}

func NewComputerFacade() *ComputerFacade {
	return &ComputerFacade{
		processor: &CPU{},
		ram:       &Ram{4, make([]byte, 4), make([]bool, 4)},
		hd:        &HardDrive{16, make([]byte, 16), make([]bool, 16)},
	}
}
