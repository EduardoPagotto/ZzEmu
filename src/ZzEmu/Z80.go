package ZzEmu

type stepInfo struct {
	address uint16
	// value   byte
}

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

	table [256]func(*stepInfo)
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

	z80.createTable()
	z80.Reset()
	return z80
}

func (z *Z80) createTable() {
	z.table = [256]func(*stepInfo){
		z.call, z.ret, z.rst, z.jp, z.jr,
	}
}

func (z80 *Z80) call(info *stepInfo) {

	addLo := z80.Memory.Read(z80.pc)
	z80.pc++
	addHi := z80.Memory.Read(z80.pc)
	z80.pc++

	z80.Push16(z80.pc)
	z80.pc = joinBytes(addHi, addLo)
}

func (z80 *Z80) jp(info *stepInfo) {
	jptemp := z80.pc
	pcl := z80.Memory.Read(jptemp)
	jptemp++
	pch := z80.Memory.Read(jptemp)
	z80.pc = joinBytes(pch, pcl)
}

func (z80 *Z80) jr(info *stepInfo) {
	var jrtemp int16 = signExtend(z80.Memory.Read(z80.pc))
	z80.pc += uint16(jrtemp)
}

func (z80 *Z80) ret(info *stepInfo) {
	z80.pc = z80.Pop16()
}

func (z80 *Z80) rst(info *stepInfo) {
	z80.Push16(z80.pc)
	z80.pc = info.address
}

//-- Auxiliares

// func (z80 *Z80) Ready16(address uint16) uint16 {
// 	valLo := z80.Memory.Read(address)
// 	valHi := z80.Memory.Read(address + 1)
// 	return joinBytes(valHi, valLo)
// }

func (z80 *Z80) Reset() {
	z80.A, z80.F, z80.B, z80.C, z80.D, z80.E, z80.H, z80.L = 0, 0, 0, 0, 0, 0, 0, 0
	z80.A_, z80.F_, z80.B_, z80.C_, z80.D_, z80.E_, z80.H_, z80.L_ = 0, 0, 0, 0, 0, 0, 0, 0
	z80.IXH, z80.IXL, z80.IYH, z80.IYL = 0, 0, 0, 0
	z80.sp, z80.I, z80.R, z80.R7, z80.pc, z80.IFF1, z80.IFF2, z80.IM = 0, 0, 0, 0, 0, 0, 0, 0

	z80.Tstates = 0
	z80.Halted = false
	z80.interruptsEnabledAt = 0
}

func (z80 *Z80) PushSplited(high, low byte) {
	z80.sp--
	z80.Memory.Write(z80.sp, high)
	z80.sp--
	z80.Memory.Write(z80.sp, low)
}

func (z80 *Z80) PopSlited() (byte, byte) {
	valLo := z80.Memory.Read(z80.sp)
	z80.sp++
	valHi := z80.Memory.Read(z80.sp)
	z80.sp++
	return valHi, valLo
}

func (z80 *Z80) Push16(value uint16) {
	high, low := splitWord(value)
	z80.PushSplited(high, low)
}

func (z80 *Z80) Pop16() uint16 {
	valHi, valLo := z80.PopSlited()
	return joinBytes(valHi, valLo)
}

// func (z80 *Z80) contend(address uint16, time uint16) {
// 	z80.Tstates += z80.Memory.contendMem(address, z80.Tstates, time)
// }

// func (z80 *Z80) contendIO(address uint16, time uint16) {
// 	z80.Tstates += z80.Memory.contendIO(address, z80.Tstates, time)
// }

// func (z80 *Z80) contendLoop(address uint16, time uint16, repet int) {

// 	for i := 0; i < repet; i++ {
// 		z80.contend(address, time)
// 	}

// }
