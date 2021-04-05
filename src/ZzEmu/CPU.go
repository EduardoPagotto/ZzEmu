package ZzEmu

type stepInfo struct {
	address uint16
	// value   byte
}

type CPU struct {
	Memory
	Cycles uint64

	A, F, B, C, D, E, H, L         byte
	A_, F_, B_, C_, D_, E_, H_, L_ byte
	IXH, IXL, IYH, IYL             byte
	I, IFF1, IFF2, IM              byte

	R7 byte // The highest bit (bit 7) of the R register
	R  uint16

	PC uint16
	SP uint16

	AF, BC, DE, HL     Register16
	IX, IY             Register16
	AF_, BC_, DE_, HL_ Register16

	// halted  byte
	// tstates byte

	table [256]func(*stepInfo)
}

// NewZ80 creates a new CPU instance.
func NewCPU(memory Memory) *CPU {

	CPU := CPU{Memory: memory}

	CPU.AF = Register16{&CPU.A, &CPU.F}
	CPU.BC = Register16{&CPU.B, &CPU.C}
	CPU.DE = Register16{&CPU.D, &CPU.E}
	CPU.HL = Register16{&CPU.H, &CPU.L}

	CPU.IX = Register16{&CPU.IXH, &CPU.IXL}
	CPU.IY = Register16{&CPU.IYH, &CPU.IYL}

	CPU.AF_ = Register16{&CPU.A_, &CPU.F_} // ??
	CPU.BC_ = Register16{&CPU.B_, &CPU.C_}
	CPU.DE_ = Register16{&CPU.D_, &CPU.E_}
	CPU.HL_ = Register16{&CPU.H_, &CPU.L_}

	CPU.createTable()
	CPU.Reset()
	return &CPU
}

func (c *CPU) createTable() {
	c.table = [256]func(*stepInfo){
		c.call, c.ret, c.rst,
	}
}

func (cpu *CPU) call(info *stepInfo) {

	addLo := cpu.Memory.Read(cpu.PC)
	cpu.PC++
	addHi := cpu.Memory.Read(cpu.PC)
	cpu.PC++

	cpu.Push16(cpu.PC)
	cpu.PC = joinBytes(addHi, addLo)
}

func (cpu *CPU) ret(info *stepInfo) {
	cpu.PC = cpu.Pop16()
}

func (cpu *CPU) rst(info *stepInfo) {
	cpu.Push16(cpu.PC)
	cpu.PC = info.address
}

//-- Auxiliares

func (cpu *CPU) Ready16(address uint16) uint16 {
	valLo := cpu.Memory.Read(address)
	valHi := cpu.Memory.Read(address + 1)
	return joinBytes(valHi, valLo)
}

func (cpu *CPU) Reset() {
	// TODO ajuste Flags
}

func (cpu *CPU) PushSplited(high, low byte) {
	cpu.SP--
	cpu.Memory.Write(cpu.SP, high)
	cpu.SP--
	cpu.Memory.Write(cpu.SP, low)
}

func (cpu *CPU) PopSlited() (byte, byte) {
	valLo := cpu.Memory.Read(cpu.SP)
	cpu.SP++
	valHi := cpu.Memory.Read(cpu.SP)
	cpu.SP++
	return valHi, valLo
}

func (cpu *CPU) Push16(value uint16) {
	//cpu.PushSplited(splitWord(value))
	high, low := splitWord(value)
	cpu.PushSplited(high, low)
}

func (cpu *CPU) Pop16() uint16 {
	valHi, valLo := cpu.PopSlited()
	return joinBytes(valHi, valLo)
}
