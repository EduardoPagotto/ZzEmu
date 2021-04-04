package ZzEmu

type stepInfo int

type CPU struct {
	//Memory
	Cycles uint64
	PC     uint16
	SP     uint16
	AF     uint16
	BC     uint16
	DE     uint16
	HL     uint16
	table  [256]func(*stepInfo)
}

func (c *CPU) createTable() {
	c.table = [256]func(*stepInfo){
		c.call, c.ret,
	}
}

func (cpu *CPU) call(info *stepInfo) {
	// cpu.push16(cpu.PC - 1)
	// cpu.PC = info.address
	cpu.PC = 0x0000
}

func (cpu *CPU) ret(info *stepInfo) {
	// cpu.push16(cpu.PC - 1)
	// cpu.PC = info.address
	cpu.PC = 0x0000
}

// func NewCPU(console *Console) *CPU {
// 	cpu := CPU{Memory: NewCPUMemory(console)}
// 	cpu.createTable()
// 	cpu.Reset()
// }

func (cpu *CPU) Ready16(address int16) int16 {
	cpu.PC = 0x0000
	return 0x0000
}

func (cpu *CPU) Reset() {
	//cpu.PC = cpu.Ready16(0x0000)
	cpu.SP = 0x0000
	//set of flags
}
