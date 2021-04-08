package ZzEmu

// The flags
const FLAG_C = 0x01
const FLAG_N = 0x02
const FLAG_P = 0x04
const FLAG_V = FLAG_P
const FLAG_3 = 0x08
const FLAG_H = 0x10
const FLAG_5 = 0x20
const FLAG_Z = 0x40
const FLAG_S = 0x80

// type TsValid struct {
// }

var (
	OpcodeMap     [255]func(z *Z80, opcode byte)
	OpcodeCBMap   [255]func(z *Z80, opcode byte)
	OpcodeDDMap   [255]func(z *Z80, opcode byte)
	OpcodeDDCBMap [255]func(z *Z80, opcode byte)
	OpcodeDFMap   [255]func(z *Z80, opcode byte)
	OpcodeEDMap   [255]func(z *Z80, opcode byte)
)

type Z80 struct {
	Memory Memory
	Cycles uint64

	A, F, B, C, D, E, H, L         byte
	A_, F_, B_, C_, D_, E_, H_, L_ byte
	IXH, IXL, IYH, IYL             byte
	I, IFF1, IFF2, IM              byte

	R7 byte // The highest bit (bit 7) of the R register
	R  uint16

	pc uint16
	sp uint16

	AF, BC, DE, HL     Register16
	IX, IY             Register16
	AF_, BC_, DE_, HL_ Register16

	Tstates             uint16
	Halted              bool
	interruptsEnabledAt int
}

// creates a new Z80 instance.
func NewZ80(memory Memory) *Z80 {
	var z80 *Z80 = new(Z80)
	z80.Memory = memory

	z80.AF = Register16{&z80.A, &z80.F}
	z80.BC = Register16{&z80.B, &z80.C}
	z80.DE = Register16{&z80.D, &z80.E}
	z80.HL = Register16{&z80.H, &z80.L}

	z80.IX = Register16{&z80.IXH, &z80.IXL}
	z80.IY = Register16{&z80.IYH, &z80.IYL}

	z80.AF_ = Register16{&z80.A_, &z80.F_} // ??
	z80.BC_ = Register16{&z80.B_, &z80.C_}
	z80.DE_ = Register16{&z80.D_, &z80.E_}
	z80.HL_ = Register16{&z80.H_, &z80.L_}

	init_table_sz53()

	initOpcodes()
	initOpcodeCBMap()
	initOpcodeDDMap()
	initOpcodeDDCBMap()
	initOpcodeDFMap()
	initOpcodeEDMap()

	z80.Reset()
	return z80
}

//-- Auxiliares

func (z80 *Z80) Reset() {
	z80.A, z80.F, z80.B, z80.C, z80.D, z80.E, z80.H, z80.L = 0, 0, 0, 0, 0, 0, 0, 0
	z80.A_, z80.F_, z80.B_, z80.C_, z80.D_, z80.E_, z80.H_, z80.L_ = 0, 0, 0, 0, 0, 0, 0, 0
	z80.IXH, z80.IXL, z80.IYH, z80.IYL = 0, 0, 0, 0
	z80.sp, z80.I, z80.R, z80.R7, z80.pc, z80.IFF1, z80.IFF2, z80.IM = 0, 0, 0, 0, 0, 0, 0, 0

	z80.Tstates = 0
	z80.Halted = false
	z80.interruptsEnabledAt = 0
}

func (z80 *Z80) Push8(value byte) {
	z80.sp--
	z80.Memory.Write(z80.sp, value)
}

func (z80 *Z80) Pop8() byte {
	val := z80.Memory.Read(z80.sp)
	z80.sp++
	return val
}

func (z80 *Z80) Push16(value uint16) {
	high, low := splitWord(value)
	z80.sp--
	z80.Memory.Write(z80.sp, high)
	z80.sp--
	z80.Memory.Write(z80.sp, low)
}

func (z80 *Z80) Pop16() uint16 {
	valLo := z80.Memory.Read(z80.sp)
	z80.sp++
	valHi := z80.Memory.Read(z80.sp)
	z80.sp++
	return joinBytes(valHi, valLo)
}

// Execute a single instruction at the program counter.
func (z80 *Z80) DoOpcode() {
	// z80.Tstates += 4
	opcode := z80.Memory.Read(z80.pc)
	z80.R = (z80.R + 1) & 0x7f
	z80.pc++

	OpcodeMap[opcode](z80, opcode)
}

/*
Carrega uint16 seguinte do PC
*/
func (z80 *Z80) LdAddrLittleEndian() uint16 {
	ldtemp := uint16(z80.Memory.Read(z80.pc))
	z80.pc++
	ldtemp |= uint16(z80.Memory.Read(z80.pc)) << 8
	z80.pc++
	return ldtemp
}

func (z80 *Z80) LoadByteFromPC() byte {
	val := z80.Memory.Read(z80.pc)
	z80.pc++
	return val
}

/* Le posicao de memoria INDEXADO
regl: ponteiro do registro Low
regh: ponteiro do registro High
*/
func (z80 *Z80) ld16rrnn(regl, regh *byte) {
	ldtemp := z80.LdAddrLittleEndian()
	*regl = z80.Memory.Read(ldtemp)
	ldtemp++
	*regh = z80.Memory.Read(ldtemp)
}

func (z80 *Z80) PtrRegisterByte(opcode byte) *byte {
	r := opcode & 0x07
	switch r {
	case 0:
		return &z80.B
	case 1:
		return &z80.C
	case 2:
		return &z80.D
	case 3:
		return &z80.E
	case 4:
		return &z80.H
	case 5:
		return &z80.L
	case 6:
		return nil
	case 7:
		return &z80.A
	}

	return nil
}
