package ZzEmu

type stepInfo int

type CPU struct {
	Memory
	Cycles uint64
	AF     RegisterPair
	BC     RegisterPair
	DE     RegisterPair
	HL     RegisterPair
	IX     RegisterPair
	IY     RegisterPair
	PC     RegisterPair
	SP     RegisterPair

	AF_ uint16
	BC_ uint16
	DE_ uint16
	HL_ uint16

	I       byte
	R       byte
	R7      byte
	IFF1    byte
	IFF2    byte
	IM      byte
	halted  byte
	tstates byte

	table [256]func(*stepInfo)
}

func (c *CPU) createTable() {
	c.table = [256]func(*stepInfo){
		c.call, c.ret, //c.rst,
	}
}

func (cpu *CPU) call(info *stepInfo) {
	cpu.PC.SetLo(cpu.Memory.Read(cpu.SP.Get()))
	cpu.SP.Inc()
	cpu.PC.SetHi(cpu.Memory.Read(cpu.SP.Get()))
	cpu.SP.Inc()
}

func (cpu *CPU) ret(info *stepInfo) {
	cpu.PC.Set(cpu.Pop16())
}

// func (cpu *CPU) Ready16(address uint16) uint16 {
// 	// cpu.PC = 0x0000
// 	return 0x0000
// }

func (cpu *CPU) Reset() {
	// TODO ajuste Flags
}

func (cpu *CPU) Push16(value uint16) {
	cpu.SP.Dec()
	cpu.Memory.Write(cpu.SP.Get(), byte(value>>8))
	cpu.SP.Dec()
	cpu.Memory.Write(cpu.SP.Get(), byte(value&0xFF))
}

func (cpu *CPU) Pop16() uint16 {

	valLo := uint16(cpu.Memory.Read(cpu.SP.Get()))
	cpu.SP.Inc()

	valHi := uint16(cpu.Memory.Read(cpu.SP.Get()))
	cpu.SP.Inc()

	return uint16((valHi << 8) | valLo)
}

// func (cpu *CPU) rst(info *stepInfo, value uint16) {
// 	cpu.Push16(cpu.PC.Get())
// 	cpu.PC.Set(value)
// }
