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
	OpcodeMap     [256]func(z *Z80, opcode byte)
	OpcodeCBMap   [256]func(z *Z80, opcode byte)
	OpcodeDDMap   [256]func(z *Z80, opcode byte)
	OpcodeDDCBMap [256]func(z *Z80, opcode byte)
	OpcodeDFMap   [256]func(z *Z80, opcode byte)
	OpcodeEDMap   [256]func(z *Z80, opcode byte)
)

type Z80 struct {
	Memory Memory
	Port   PortAccessor
	//Cycles uint64

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
func NewZ80(memory Memory, port PortAccessor) *Z80 {
	var z80 *Z80 = new(Z80)
	z80.Memory = memory
	z80.Port = port

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

//--- flow controll

func (z80 *Z80) Call() {
	z80.Tstates += 17
	newpc := z80.LdAddrLittleEndian()
	z80.Push16(z80.pc)
	z80.pc = newpc
}

func (z80 *Z80) Jr() {
	z80.Tstates += 12
	var jrtemp int16 = signExtend(z80.Memory.Read(z80.pc))
	z80.pc += uint16(jrtemp)
	z80.pc++
}

func (z80 *Z80) Rst(value byte) {
	z80.Push16(z80.pc)
	z80.pc = uint16(value)
}

//-- Logic

func (z80 *Z80) And(value byte) {
	z80.A &= value
	z80.F = FLAG_H | sz53pTable[z80.A]
}

func (z80 *Z80) Xor(value byte) {
	z80.A ^= value
	z80.F = sz53pTable[z80.A]
}

func (z80 *Z80) Or(value byte) {
	z80.A |= value
	z80.F = sz53pTable[z80.A]
}

func (z80 *Z80) Cp(value byte) {
	var cptemp uint16 = uint16(z80.A) - uint16(value)
	var lookup byte = ((z80.A & 0x88) >> 3) | ((value & 0x88) >> 2) | byte((cptemp&0x88)>>1)
	z80.F = ternOpB((cptemp&0x100) != 0, FLAG_C, ternOpB(cptemp != 0, 0, FLAG_Z)) | FLAG_N | halfcarrySubTable[lookup&0x07] | overflowSubTable[lookup>>4] | (value & (FLAG_3 | FLAG_5)) | byte(cptemp&FLAG_S)
}

//-- aritmetric

func (z80 *Z80) Add(value byte) {
	var addtemp uint = uint(z80.A) + uint(value)
	var lookup byte = ((z80.A & 0x88) >> 3) | ((value & 0x88) >> 2) | byte((addtemp&0x88)>>1)
	z80.A = byte(addtemp)
	z80.F = ternOpB(addtemp&0x100 != 0, FLAG_C, 0) | halfcarryAddTable[lookup&0x07] | overflowAddTable[lookup>>4] | sz53Table[z80.A]
}

func (z80 *Z80) Adc(value byte) {
	var adctemp uint16 = uint16(z80.A) + uint16(value) + (uint16(z80.F) & FLAG_C)
	var lookup byte = byte(((uint16(z80.A) & 0x88) >> 3) | ((uint16(value) & 0x88) >> 2) | ((uint16(adctemp) & 0x88) >> 1))
	z80.A = byte(adctemp)
	z80.F = ternOpB((adctemp&0x100) != 0, FLAG_C, 0) | halfcarryAddTable[lookup&0x07] | overflowAddTable[lookup>>4] | sz53Table[z80.A]
}

func (z80 *Z80) Sub(value byte) {
	var subtemp uint16 = uint16(z80.A) - uint16(value)
	var lookup byte = ((z80.A & 0x88) >> 3) | ((value & 0x88) >> 2) | byte((subtemp&0x88)>>1)
	z80.A = byte(subtemp)
	z80.F = ternOpB(subtemp&0x100 != 0, FLAG_C, 0) | FLAG_N |
		halfcarrySubTable[lookup&0x07] | overflowSubTable[lookup>>4] |
		sz53Table[z80.A]
}

func (z80 *Z80) Sbc(value byte) {
	var sbctemp uint16 = uint16(z80.A) - uint16(value) - (uint16(z80.F) & FLAG_C)
	var lookup byte = ((z80.A & 0x88) >> 3) | ((value & 0x88) >> 2) | byte((sbctemp&0x88)>>1)
	z80.A = byte(sbctemp)
	z80.F = ternOpB((sbctemp&0x100) != 0, FLAG_C, 0) | FLAG_N | halfcarrySubTable[lookup&0x07] | overflowSubTable[lookup>>4] | sz53Table[z80.A]
}

func (z80 *Z80) Inc(value *byte) { // TODO merge com IncR
	*value++
	z80.F = (z80.F & FLAG_C) | ternOpB(*value == 0x80, FLAG_V, 0) | ternOpB((*value&0x0f) != 0, 0, FLAG_H) | sz53Table[(*value)]
}

func (z80 *Z80) Dec(value *byte) { // TODO merge com DecR
	z80.F = (z80.F & FLAG_C) | ternOpB((*value&0x0f) != 0, 0, FLAG_H) | FLAG_N
	*value--
	z80.F |= ternOpB(*value == 0x7f, FLAG_V, 0) | sz53Table[*value]
}

func (z80 *Z80) IncR(opcode byte) {
	var ptrReg *byte = z80.GetPrtRegisterValByte(opcode)
	(*ptrReg)++
	z80.F = (z80.F & FLAG_C) | (ternOpB((*ptrReg) == 0x80, FLAG_V, 0)) | (ternOpB(((*ptrReg)&0x0f) != 0, 0, FLAG_H)) | sz53Table[(*ptrReg)]
}

func (z80 *Z80) DecR(opcode byte) {

	var ptrReg *byte = z80.GetPrtRegisterValByte(opcode)
	z80.F = (z80.F & FLAG_C) | (ternOpB((*ptrReg)&0x0f != 0, 0, FLAG_H)) | FLAG_N
	(*ptrReg)--
	z80.F |= (ternOpB((*ptrReg) == 0x7f, FLAG_V, 0)) | sz53Table[(*ptrReg)]
}

//-- memory

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

func (z80 *Z80) ld16nnrr(regl, regh byte) {
	ldtemp := z80.LdAddrLittleEndian()
	z80.Memory.Write(ldtemp, regl)
	ldtemp++
	z80.Memory.Write(ldtemp, regh)
}

//--- IO

func (z80 *Z80) readPort(address uint16) byte {
	return z80.Port.ReadPort(address)
}

func (z80 *Z80) writePort(address uint16, b byte) {
	z80.Port.WritePort(address, b)
}

//-- register select to opcode

func (z80 *Z80) GetRegisterValByte(opcode byte) byte {
	r := opcode & 0x07
	switch r {
	case 0:
		return z80.B
	case 1:
		return z80.C
	case 2:
		return z80.D
	case 3:
		return z80.E
	case 4:
		return z80.H
	case 5:
		return z80.L
	case 7:
		return z80.A
	}

	return 0
}

func (z80 *Z80) GetPrtRegisterValByte(opcode byte) *byte {
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
	case 7:
		return &z80.A
	}

	return nil
}

//--- entry point

// Execute a single instruction at the program counter.
func (z80 *Z80) DoOpcode() {
	// z80.Tstates += 4
	opcode := z80.Memory.Read(z80.pc)
	z80.R = (z80.R + 1) & 0x7f
	z80.pc++

	OpcodeMap[opcode](z80, opcode)
}
