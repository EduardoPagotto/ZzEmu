package ZzEmu

type stepInfo int

type CPU struct {
	Memory
	Cycles uint64
	PC     RegisterPair
	SP     RegisterPair
	AF     RegisterPair
	BC     RegisterPair
	DE     RegisterPair
	HL     RegisterPair
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
	// cpu.PC = 0x0000
}

func (cpu *CPU) ret(info *stepInfo) {
	// cpu.push16(cpu.PC - 1)
	// cpu.PC = info.address
	// cpu.PC = 0x0000
}

// func (cpu *CPU) Ready16(address uint16) uint16 {
// 	// cpu.PC = 0x0000
// 	return 0x0000
// }

func (cpu *CPU) Reset() {
	cpu.PC.Set(0x0000)
	cpu.SP.Set(0x0000)
	// TODO ajuste Flags
}
