package facade

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
