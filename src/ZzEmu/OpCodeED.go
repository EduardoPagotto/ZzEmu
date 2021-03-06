package ZzEmu

func initOpcodeEDMap() {

	// BEGIN of 0xed shifted opcodes
	/* IN B,(C) */
	OpcodeEDMap[0x40] = instrED__IN_B_iC
	/* OUT (C),B */
	OpcodeEDMap[0x41] = instrED__OUT_iC_B
	/* SBC HL,BC */
	OpcodeEDMap[0x42] = instrED__SBC_HL_BC
	/* LD (nnnn),BC */
	OpcodeEDMap[0x43] = instrED__LD_iNNNN_BC
	/* RETI */
	OpcodeEDMap[0x4d] = instrED__RETI
	/* NEG */
	OpcodeEDMap[0x7c] = instrED__NEG
	// Fallthrough cases
	OpcodeEDMap[0x44] = OpcodeEDMap[0x7c]
	OpcodeEDMap[0x4c] = OpcodeEDMap[0x7c]
	OpcodeEDMap[0x54] = OpcodeEDMap[0x7c]
	OpcodeEDMap[0x5c] = OpcodeEDMap[0x7c]
	OpcodeEDMap[0x64] = OpcodeEDMap[0x7c]
	OpcodeEDMap[0x6c] = OpcodeEDMap[0x7c]
	OpcodeEDMap[0x74] = OpcodeEDMap[0x7c]
	/* RETN */
	OpcodeEDMap[0x7d] = instrED__RETN
	// Fallthrough cases
	OpcodeEDMap[0x45] = OpcodeEDMap[0x7d]
	//OpcodeEDMap[0x4d] = OpcodeEDMap[0x7d]
	OpcodeEDMap[0x55] = OpcodeEDMap[0x7d]
	OpcodeEDMap[0x5d] = OpcodeEDMap[0x7d]
	OpcodeEDMap[0x65] = OpcodeEDMap[0x7d]
	OpcodeEDMap[0x6d] = OpcodeEDMap[0x7d]
	OpcodeEDMap[0x75] = OpcodeEDMap[0x7d]
	/* IM 0 */
	OpcodeEDMap[0x6e] = instrED__IM_0
	// Fallthrough cases
	OpcodeEDMap[0x46] = OpcodeEDMap[0x6e]
	OpcodeEDMap[0x4e] = OpcodeEDMap[0x6e]
	OpcodeEDMap[0x66] = OpcodeEDMap[0x6e]
	/* LD I,A */
	OpcodeEDMap[0x47] = instrED__LD_I_A
	/* IN C,(C) */
	OpcodeEDMap[0x48] = instrED__IN_C_iC
	// 	/* OUT (C),C */
	OpcodeEDMap[0x49] = instrED__OUT_iC_C
	/* ADC HL,BC */
	OpcodeEDMap[0x4a] = instrED__ADC_HL_BC
	/* LD BC,(nnnn) */
	OpcodeEDMap[0x4b] = instrED__LD_BC_iNNNN
	/* LD R,A */
	OpcodeEDMap[0x4f] = instrED__LD_R_A
	/* IN D,(C) */
	OpcodeEDMap[0x50] = instrED__IN_D_iC
	/* OUT (C),D */
	OpcodeEDMap[0x51] = instrED__OUT_iC_D
	/* SBC HL,DE */
	OpcodeEDMap[0x52] = instrED__SBC_HL_DE
	/* LD (nnnn),DE */
	OpcodeEDMap[0x53] = instrED__LD_iNNNN_DE
	/* IM 1 */
	OpcodeEDMap[0x76] = instrED__IM_1
	// Fallthrough cases
	OpcodeEDMap[0x56] = OpcodeEDMap[0x76]
	/* LD A,I */
	OpcodeEDMap[0x57] = instrED__LD_A_I
	/* IN E,(C) */
	OpcodeEDMap[0x58] = instrED__IN_E_iC
	/* OUT (C),E */
	OpcodeEDMap[0x59] = instrED__OUT_iC_E
	/* ADC HL,DE */
	OpcodeEDMap[0x5a] = instrED__ADC_HL_DE
	/* LD DE,(nnnn) */
	OpcodeEDMap[0x5b] = instrED__LD_DE_iNNNN
	/* IM 2 */
	OpcodeEDMap[0x7e] = instrED__IM_2
	// Fallthrough cases
	OpcodeEDMap[0x5e] = OpcodeEDMap[0x7e]
	/* LD A,R */
	OpcodeEDMap[0x5f] = instrED__LD_A_R
	/* IN H,(C) */
	OpcodeEDMap[0x60] = instrED__IN_H_iC
	/* OUT (C),H */
	OpcodeEDMap[0x61] = instrED__OUT_iC_H
	/* SBC HL,HL */
	OpcodeEDMap[0x62] = instrED__SBC_HL_HL
	/* LD (nnnn),HL */
	OpcodeEDMap[0x63] = instrED__LD_iNNNN_HL
	/* RRD */
	OpcodeEDMap[0x67] = instrED__RRD
	/* IN L,(C) */
	OpcodeEDMap[0x68] = instrED__IN_L_iC
	/* OUT (C),L */
	OpcodeEDMap[0x69] = instrED__OUT_iC_L
	/* ADC HL,HL */
	OpcodeEDMap[0x6a] = instrED__ADC_HL_HL
	/* LD HL,(nnnn) */
	OpcodeEDMap[0x6b] = instrED__LD_HL_iNNNN
	/* RLD */
	OpcodeEDMap[0x6f] = instrED__RLD
	/* IN F,(C) */
	OpcodeEDMap[0x70] = instrED__IN_F_iC
	/* OUT (C),0 */
	OpcodeEDMap[0x71] = instrED__OUT_iC_0
	/* SBC HL,SP */
	OpcodeEDMap[0x72] = instrED__SBC_HL_SP
	/* LD (nnnn),SP */
	OpcodeEDMap[0x73] = instrED__LD_iNNNN_SP
	/* IN A,(C) */
	OpcodeEDMap[0x78] = instrED__IN_A_iC
	/* OUT (C),A */
	OpcodeEDMap[0x79] = instrED__OUT_iC_A
	/* ADC HL,SP */
	OpcodeEDMap[0x7a] = instrED__ADC_HL_SP
	/* LD SP,(nnnn) */
	OpcodeEDMap[0x7b] = instrED__LD_SP_iNNNN
	/* LDI */
	OpcodeEDMap[0xa0] = instrED__LDI
	/* CPI */
	OpcodeEDMap[0xa1] = instrED__CPI
	/* INI */
	OpcodeEDMap[0xa2] = instrED__INI
	/* OUTI */
	OpcodeEDMap[0xa3] = instrED__OUTI
	/* LDD */
	OpcodeEDMap[0xa8] = instrED__LDD
	/* CPD */
	OpcodeEDMap[0xa9] = instrED__CPD
	/* IND */
	OpcodeEDMap[0xaa] = instrED__IND
	/* OUTD */
	OpcodeEDMap[0xab] = instrED__OUTD
	/* LDIR */
	OpcodeEDMap[0xb0] = instrED__LDIR
	/* CPIR */
	OpcodeEDMap[0xb1] = instrED__CPIR
	/* INIR */
	OpcodeEDMap[0xb2] = instrED__INIR
	/* OTIR */
	OpcodeEDMap[0xb3] = instrED__OTIR
	/* LDDR */
	OpcodeEDMap[0xb8] = instrED__LDDR
	/* CPDR */
	OpcodeEDMap[0xb9] = instrED__CPDR
	/* INDR */
	OpcodeEDMap[0xba] = instrED__INDR
	/* OTDR */
	OpcodeEDMap[0xbb] = instrED__OTDR
	/* slttrap */
	OpcodeEDMap[0xfb] = instrED__SLTTRAP

	// 	// END of 0xed shifted opcodes
}

/* IN B,(C) */
func instrED__IN_B_iC(z *Z80, opcode byte) {
	z.Tstates += 12
	z.in(&z.B, z.BC.Get())
}

/* OUT (C),B */
func instrED__OUT_iC_B(z *Z80, opcode byte) {
	z.Tstates += 12
	z.writePort(z.BC.Get(), z.B)
}

/* SBC HL,BC */
func instrED__SBC_HL_BC(z *Z80, opcode byte) {
	//z.Memory.ContendReadNoMreq_loop(z.IR(), 1, 7)
	z.Tstates += 15
	z.Sbc16(z.BC.Get())
}

/* LD (nnnn),BC */
func instrED__LD_iNNNN_BC(z *Z80, opcode byte) {
	z.Tstates += 20
	z.StoreIndexR(z.BC)
}

func instrED__RETI(z *Z80, opcode byte) {
	// nao existia ???
	z.Tstates += 14
	z.pc = z.Pop()
}

/* NEG */
func instrED__NEG(z *Z80, opcode byte) {
	z.Tstates += 8
	bytetemp := z.A
	z.A = 0
	z.Sub(bytetemp)
}

/* RETN */
func instrED__RETN(z *Z80, opcode byte) {
	z.Tstates += 14
	z.IFF1 = z.IFF2
	z.pc = z.Pop()
}

/* IM 0 */
func instrED__IM_0(z *Z80, opcode byte) {
	z.Tstates += 8
	z.IM = 0
}

/* LD I,A */
func instrED__LD_I_A(z *Z80, opcode byte) {
	z.Tstates += 9
	//z.Memory.ContendReadNoMreq(z.IR(), 1)
	z.I = z.A
}

/* IN C,(C) */
func instrED__IN_C_iC(z *Z80, opcode byte) {
	z.Tstates += 12
	z.in(&z.C, z.BC.Get())
}

/* OUT (C),C */
func instrED__OUT_iC_C(z *Z80, opcode byte) {
	z.Tstates += 12
	z.writePort(z.BC.Get(), z.C)
}

/* ADC HL,BC */
func instrED__ADC_HL_BC(z *Z80, opcode byte) {
	//z.Memory.ContendReadNoMreq_loop(z.IR(), 1, 7)
	z.Tstates += 15
	z.Adc16(z.BC.Get())
}

/* LD BC,(nnnn) */
func instrED__LD_BC_iNNNN(z *Z80, opcode byte) {
	z.Tstates += 20
	z.LoadIndexR(&z.BC)
}

/* LD R,A */
func instrED__LD_R_A(z *Z80, opcode byte) {
	//z.Memory.ContendReadNoMreq(z.IR(), 1)
	/* Keep the RZX instruction counter right */
	z.Tstates += 9
	z.rzxInstructionsOffset += (int(z.R) - int(z.A))
	z.R, z.R7 = uint16(z.A), z.A
}

/* IN D,(C) */
func instrED__IN_D_iC(z *Z80, opcode byte) {
	z.Tstates += 12
	z.in(&z.D, z.BC.Get())
}

/* OUT (C),D */
func instrED__OUT_iC_D(z *Z80, opcode byte) {
	z.Tstates += 12
	z.writePort(z.BC.Get(), z.D)
}

/* SBC HL,DE */
func instrED__SBC_HL_DE(z *Z80, opcode byte) {
	//z.Memory.ContendReadNoMreq_loop(z.IR(), 1, 7)
	z.Tstates += 15
	z.Sbc16(z.DE.Get())
}

/* LD (nnnn),DE */
func instrED__LD_iNNNN_DE(z *Z80, opcode byte) {
	z.Tstates += 20
	z.StoreIndexR(z.DE)
}

/* IM 1 */
func instrED__IM_1(z *Z80, opcode byte) {
	z.Tstates += 8
	z.IM = 1
}

/* LD A,I */
func instrED__LD_A_I(z *Z80, opcode byte) {
	//z.Memory.ContendReadNoMreq(z.IR(), 1)
	z.Tstates += 9
	z.A = z.I
	z.F = (z.F & FLAG_C) | sz53Table[z.A] | ternOpB(z.IFF2 != 0, FLAG_V, 0)
}

/* IN E,(C) */
func instrED__IN_E_iC(z *Z80, opcode byte) {
	z.Tstates += 12
	z.in(&z.E, z.BC.Get())
}

/* OUT (C),E */
func instrED__OUT_iC_E(z *Z80, opcode byte) {
	z.Tstates += 12
	z.writePort(z.BC.Get(), z.E)
}

/* ADC HL,DE */
func instrED__ADC_HL_DE(z *Z80, opcode byte) {
	//z.Memory.ContendReadNoMreq_loop(z.IR(), 1, 7)
	z.Tstates += 15
	z.Adc16(z.DE.Get())
}

/* LD DE,(nnnn) */
func instrED__LD_DE_iNNNN(z *Z80, opcode byte) {
	z.Tstates += 20
	z.LoadIndexR(&z.DE)
}

/* IM 2 */
func instrED__IM_2(z *Z80, opcode byte) {
	z.Tstates += 8
	z.IM = 2
}

/* LD A,R */
func instrED__LD_A_R(z *Z80, opcode byte) {
	//z.Memory.ContendReadNoMreq(z.IR(), 1)
	z.Tstates += 9
	z.A = byte(z.R&0x7f) | (z.R7 & 0x80)
	z.F = (z.F & FLAG_C) | sz53Table[z.A] | ternOpB(z.IFF2 != 0, FLAG_V, 0)
}

/* IN H,(C) */
func instrED__IN_H_iC(z *Z80, opcode byte) {
	z.Tstates += 12
	z.in(&z.H, z.BC.Get())
}

/* OUT (C),H */
func instrED__OUT_iC_H(z *Z80, opcode byte) {
	z.Tstates += 12
	z.writePort(z.BC.Get(), z.H)
}

/* SBC HL,HL */
func instrED__SBC_HL_HL(z *Z80, opcode byte) {
	//z.Memory.ContendReadNoMreq_loop(z.IR(), 1, 7)
	z.Tstates += 15
	z.Sbc16(z.HL.Get())
}

/* LD (nnnn),HL */
func instrED__LD_iNNNN_HL(z *Z80, opcode byte) {
	z.Tstates += 20
	z.StoreIndexR(z.HL)
}

/* RRD */
func instrED__RRD(z *Z80, opcode byte) {
	z.Tstates += 18
	var bytetemp byte = z.bus.ReadMemory(z.HL.Get())
	//z.Memory.ContendReadNoMreq_loop(z.HL.Get(), 1, 4)
	z.bus.WriteMemory(z.HL.Get(), (z.A<<4)|(bytetemp>>4))
	z.A = (z.A & 0xf0) | (bytetemp & 0x0f)
	z.F = (z.F & FLAG_C) | sz53pTable[z.A]
}

/* IN L,(C) */
func instrED__IN_L_iC(z *Z80, opcode byte) {
	z.Tstates += 12
	z.in(&z.L, z.BC.Get())
}

/* OUT (C),L */
func instrED__OUT_iC_L(z *Z80, opcode byte) {
	z.Tstates += 12
	z.writePort(z.BC.Get(), z.L)
}

/* ADC HL,HL */
func instrED__ADC_HL_HL(z *Z80, opcode byte) {
	z.Tstates += 15
	//z.Memory.ContendReadNoMreq_loop(z.IR(), 1, 7)
	z.Adc16(z.HL.Get())
}

/* LD HL,(nnnn) */
func instrED__LD_HL_iNNNN(z *Z80, opcode byte) {
	z.Tstates += 20
	z.LoadIndexR(&z.HL)
}

/* RLD */
func instrED__RLD(z *Z80, opcode byte) {
	z.Tstates += 18
	var bytetemp byte = z.bus.ReadMemory(z.HL.Get())
	//z.Memory.ContendReadNoMreq_loop(z.HL.Get(), 1, 4)
	z.bus.WriteMemory(z.HL.Get(), (bytetemp<<4)|(z.A&0x0f))
	z.A = (z.A & 0xf0) | (bytetemp >> 4)
	z.F = (z.F & FLAG_C) | sz53pTable[z.A]
}

/* IN F,(C) */
func instrED__IN_F_iC(z *Z80, opcode byte) {
	z.Tstates += 12
	var bytetemp byte
	z.in(&bytetemp, z.BC.Get())
}

/* OUT (C),0 */
func instrED__OUT_iC_0(z *Z80, opcode byte) {
	z.Tstates += 12
	z.writePort(z.BC.Get(), 0)
}

/* SBC HL,SP */
func instrED__SBC_HL_SP(z *Z80, opcode byte) {
	//z.Memory.ContendReadNoMreq_loop(z.IR(), 1, 7)
	z.Tstates += 15
	z.Sbc16(z.sp)
}

/* LD (nnnn),SP */
func instrED__LD_iNNNN_SP(z *Z80, opcode byte) {
	z.Tstates += 20
	sph, spl := splitWord(z.sp)

	ldtemp := z.Load16()
	z.bus.WriteMemory(ldtemp, spl)
	ldtemp++
	z.bus.WriteMemory(ldtemp, sph)
}

/* IN A,(C) */
func instrED__IN_A_iC(z *Z80, opcode byte) {
	z.Tstates += 12
	z.in(&z.A, z.BC.Get())
}

/* OUT (C),A */
func instrED__OUT_iC_A(z *Z80, opcode byte) {
	z.Tstates += 12
	z.writePort(z.BC.Get(), z.A)
}

/* ADC HL,SP */
func instrED__ADC_HL_SP(z *Z80, opcode byte) {
	z.Tstates += 15
	//z.Memory.ContendReadNoMreq_loop(z.IR(), 1, 7)
	z.Adc16(z.sp)
}

/* LD SP,(nnnn) */
func instrED__LD_SP_iNNNN(z *Z80, opcode byte) {
	z.Tstates += 20
	ldtemp := z.Load16()
	spl := z.bus.ReadMemory(ldtemp)
	ldtemp++
	sph := z.bus.ReadMemory(ldtemp)
	z.sp = joinBytes(sph, spl)
}

/* LDI */
func instrED__LDI(z *Z80, opcode byte) {
	z.Tstates += 16
	var bytetemp byte = z.bus.ReadMemory(z.HL.Get())
	z.BC.Dec() //DecBC()
	z.bus.WriteMemory(z.DE.Get(), bytetemp)
	//z.Memory.ContendWriteNoMreq_loop(z.DE.Get(), 1, 2)
	z.DE.Inc() //IncDE()
	z.HL.Inc() //IncHL()
	bytetemp += z.A
	z.F = (z.F & (FLAG_C | FLAG_Z | FLAG_S)) |
		ternOpB(z.BC.Get() != 0, FLAG_V, 0) |
		(bytetemp & FLAG_3) |
		ternOpB((bytetemp&0x02) != 0, FLAG_5, 0)
}

/* CPI */
func instrED__CPI(z *Z80, opcode byte) {
	z.Tstates += 16
	var value byte = z.bus.ReadMemory(z.HL.Get())
	var bytetemp byte = z.A - value
	var lookup byte = ((z.A & 0x08) >> 3) | (((value) & 0x08) >> 2) | ((bytetemp & 0x08) >> 1)
	//z.Memory.ContendReadNoMreq_loop(z.HL.Get(), 1, 5)
	z.HL.Inc() //IncHL()
	z.BC.Dec() //DecBC()
	z.F = (z.F & FLAG_C) | ternOpB(z.BC.Get() != 0, FLAG_V|FLAG_N, FLAG_N) | halfcarrySubTable[lookup] | ternOpB(bytetemp != 0, 0, FLAG_Z) | (bytetemp & FLAG_S)
	if (z.F & FLAG_H) != 0 {
		bytetemp--
	}
	z.F |= (bytetemp & FLAG_3) | ternOpB((bytetemp&0x02) != 0, FLAG_5, 0)
}

/* INI */
func instrED__INI(z *Z80, opcode byte) {
	z.Tstates += 16
	//z.Memory.ContendReadNoMreq(z.IR(), 1)
	var initemp byte = z.readPort(z.BC.Get())
	z.bus.WriteMemory(z.HL.Get(), initemp)

	z.B--
	z.HL.Inc() //IncHL()
	var initemp2 byte = initemp + z.C + 1
	z.F = ternOpB((initemp&0x80) != 0, FLAG_N, 0) |
		ternOpB(initemp2 < initemp, FLAG_H|FLAG_C, 0) |
		ternOpB(parityTable[(initemp2&0x07)^z.B] != 0, FLAG_P, 0) |
		sz53Table[z.B]
}

/* OUTI */
func instrED__OUTI(z *Z80, opcode byte) {
	z.Tstates += 16
	//z.Memory.ContendReadNoMreq(z.IR(), 1)
	var outitemp byte = z.bus.ReadMemory(z.HL.Get())
	z.B-- /* This does happen first, despite what the specs say */
	z.writePort(z.BC.Get(), outitemp)

	z.HL.Inc() //IncHL()
	var outitemp2 byte = outitemp + z.L
	z.F = ternOpB((outitemp&0x80) != 0, FLAG_N, 0) |
		ternOpB(outitemp2 < outitemp, FLAG_H|FLAG_C, 0) |
		ternOpB(parityTable[(outitemp2&0x07)^z.B] != 0, FLAG_P, 0) |
		sz53Table[z.B]
}

/* LDD */
func instrED__LDD(z *Z80, opcode byte) {
	z.Tstates += 16
	var bytetemp byte = z.bus.ReadMemory(z.HL.Get())
	z.BC.Dec() //DecBC()
	z.bus.WriteMemory(z.DE.Get(), bytetemp)
	//z.Memory.ContendWriteNoMreq_loop(z.DE.Get(), 1, 2)
	z.DE.Dec() //DecDE()
	z.HL.Dec() //DecHL()
	bytetemp += z.A
	z.F = (z.F & (FLAG_C | FLAG_Z | FLAG_S)) |
		ternOpB(z.BC.Get() != 0, FLAG_V, 0) |
		(bytetemp & FLAG_3) |
		ternOpB((bytetemp&0x02) != 0, FLAG_5, 0)
}

/* CPD */
func instrED__CPD(z *Z80, opcode byte) {
	z.Tstates += 16
	var value byte = z.bus.ReadMemory(z.HL.Get())
	var bytetemp byte = z.A - value
	var lookup byte = ((z.A & 0x08) >> 3) | (((value) & 0x08) >> 2) | ((bytetemp & 0x08) >> 1)
	//z.Memory.ContendReadNoMreq_loop(z.HL.Get(), 1, 5)
	z.HL.Dec() //DecHL()
	z.BC.Dec() // DecBC()
	z.F = (z.F & FLAG_C) | ternOpB(z.BC.Get() != 0, FLAG_V|FLAG_N, FLAG_N) | halfcarrySubTable[lookup] | ternOpB(bytetemp != 0, 0, FLAG_Z) | (bytetemp & FLAG_S)
	if (z.F & FLAG_H) != 0 {
		bytetemp--
	}
	z.F |= (bytetemp & FLAG_3) | ternOpB((bytetemp&0x02) != 0, FLAG_5, 0)
}

/* IND */
func instrED__IND(z *Z80, opcode byte) {
	z.Tstates += 16
	//z.Memory.ContendReadNoMreq(z.IR(), 1)
	var initemp byte = z.readPort(z.BC.Get())
	z.bus.WriteMemory(z.HL.Get(), initemp)

	z.B--
	z.HL.Dec()
	var initemp2 byte = initemp + z.C - 1
	z.F = ternOpB((initemp&0x80) != 0, FLAG_N, 0) |
		ternOpB(initemp2 < initemp, FLAG_H|FLAG_C, 0) |
		ternOpB(parityTable[(initemp2&0x07)^z.B] != 0, FLAG_P, 0) |
		sz53Table[z.B]
}

/* OUTD */
func instrED__OUTD(z *Z80, opcode byte) {
	z.Tstates += 16
	//z.Memory.ContendReadNoMreq(z.IR(), 1)
	var outitemp byte = z.bus.ReadMemory(z.HL.Get())
	z.B-- /* This does happen first, despite what the specs say */
	z.writePort(z.BC.Get(), outitemp)

	z.HL.Dec()
	var outitemp2 byte = outitemp + z.L
	z.F = ternOpB((outitemp&0x80) != 0, FLAG_N, 0) |
		ternOpB(outitemp2 < outitemp, FLAG_H|FLAG_C, 0) |
		ternOpB(parityTable[(outitemp2&0x07)^z.B] != 0, FLAG_P, 0) |
		sz53Table[z.B]
}

/* LDIR */
func instrED__LDIR(z *Z80, opcode byte) {
	z.Tstates += 16
	var bytetemp byte = z.bus.ReadMemory(z.HL.Get())
	z.bus.WriteMemory(z.DE.Get(), bytetemp)
	//z.Memory.ContendWriteNoMreq_loop(z.DE.Get(), 1, 2)
	z.BC.Dec()
	bytetemp += z.A
	z.F = (z.F & (FLAG_C | FLAG_Z | FLAG_S)) | ternOpB(z.BC.Get() != 0, FLAG_V, 0) | (bytetemp & FLAG_3) | ternOpB((bytetemp&0x02 != 0), FLAG_5, 0)
	if z.BC.Get() != 0 {
		//z.Memory.ContendWriteNoMreq_loop(z.DE.Get(), 1, 5)
		z.pc -= 2      //DecPC(2)
		z.Tstates += 5 // Testar
	}
	z.HL.Inc()
	z.DE.Inc()
}

/* CPIR */
func instrED__CPIR(z *Z80, opcode byte) {
	z.Tstates += 16
	var value byte = z.bus.ReadMemory(z.HL.Get())
	var bytetemp byte = z.A - value
	var lookup byte = ((z.A & 0x08) >> 3) | (((value) & 0x08) >> 2) | ((bytetemp & 0x08) >> 1)
	//z.Memory.ContendReadNoMreq_loop(z.HL.Get(), 1, 5)
	z.BC.Dec()
	z.F = (z.F & FLAG_C) | (ternOpB(z.BC.Get() != 0, (FLAG_V | FLAG_N), FLAG_N)) | halfcarrySubTable[lookup] | (ternOpB(bytetemp != 0, 0, FLAG_Z)) | (bytetemp & FLAG_S)
	if (z.F & FLAG_H) != 0 {
		bytetemp--
	}
	z.F |= (bytetemp & FLAG_3) | ternOpB((bytetemp&0x02) != 0, FLAG_5, 0)
	if (z.F & (FLAG_V | FLAG_Z)) == FLAG_V {
		//z.Memory.ContendReadNoMreq_loop(z.HL.Get(), 1, 5)
		z.pc -= 2      //z.pc -= 2
		z.Tstates += 5 // Testar
	}
	z.HL.Inc()
}

/* INIR */
func instrED__INIR(z *Z80, opcode byte) {
	z.Tstates += 16
	//z.Memory.ContendReadNoMreq(z.IR(), 1)
	var initemp byte = z.readPort(z.BC.Get())
	z.bus.WriteMemory(z.HL.Get(), initemp)

	z.B--
	var initemp2 byte = initemp + z.C + 1
	z.F = ternOpB(initemp&0x80 != 0, FLAG_N, 0) |
		ternOpB(initemp2 < initemp, FLAG_H|FLAG_C, 0) |
		ternOpB(parityTable[(initemp2&0x07)^z.B] != 0, FLAG_P, 0) |
		sz53Table[z.B]

	if z.B != 0 {
		//z.Memory.ContendWriteNoMreq_loop(z.HL.Get(), 1, 5)
		z.pc -= 2
		z.Tstates += 5
	}
	z.HL.Inc()
}

/* OTIR */
func instrED__OTIR(z *Z80, opcode byte) {
	z.Tstates += 16
	//z.Memory.ContendReadNoMreq(z.IR(), 1)
	var outitemp byte = z.bus.ReadMemory(z.HL.Get())
	z.B-- /* This does happen first, despite what the specs say */
	z.writePort(z.BC.Get(), outitemp)

	z.HL.Inc()
	var outitemp2 byte = outitemp + z.L
	z.F = ternOpB((outitemp&0x80) != 0, FLAG_N, 0) |
		ternOpB(outitemp2 < outitemp, FLAG_H|FLAG_C, 0) |
		ternOpB(parityTable[(outitemp2&0x07)^z.B] != 0, FLAG_P, 0) |
		sz53Table[z.B]

	if z.B != 0 {
		//z.Memory.ContendReadNoMreq_loop(z.BC.Get(), 1, 5)
		z.pc -= 2
		z.Tstates += 5
	}
}

/* LDDR */
func instrED__LDDR(z *Z80, opcode byte) {
	z.Tstates += 16
	var bytetemp byte = z.bus.ReadMemory(z.HL.Get())
	z.bus.WriteMemory(z.DE.Get(), bytetemp)
	//z.Memory.ContendWriteNoMreq_loop(z.DE.Get(), 1, 2)
	z.BC.Dec()
	bytetemp += z.A
	z.F = (z.F & (FLAG_C | FLAG_Z | FLAG_S)) | ternOpB(z.BC.Get() != 0, FLAG_V, 0) | (bytetemp & FLAG_3) | ternOpB((bytetemp&0x02 != 0), FLAG_5, 0)
	if z.BC.Get() != 0 {
		//z.Memory.ContendWriteNoMreq_loop(z.DE.Get(), 1, 5)
		z.pc -= 2
		z.Tstates += 5
	}
	z.HL.Dec()
	z.DE.Dec()
}

/* CPDR */
func instrED__CPDR(z *Z80, opcode byte) {
	z.Tstates += 16
	var value byte = z.bus.ReadMemory(z.HL.Get())
	var bytetemp byte = z.A - value
	var lookup byte = ((z.A & 0x08) >> 3) | (((value) & 0x08) >> 2) | ((bytetemp & 0x08) >> 1)
	//z.Memory.ContendReadNoMreq_loop(z.HL.Get(), 1, 5)
	z.BC.Dec()
	z.F = (z.F & FLAG_C) | (ternOpB(z.BC.Get() != 0, (FLAG_V | FLAG_N), FLAG_N)) | halfcarrySubTable[lookup] | (ternOpB(bytetemp != 0, 0, FLAG_Z)) | (bytetemp & FLAG_S)
	if (z.F & FLAG_H) != 0 {
		bytetemp--
	}
	z.F |= (bytetemp & FLAG_3) | ternOpB((bytetemp&0x02) != 0, FLAG_5, 0)
	if (z.F & (FLAG_V | FLAG_Z)) == FLAG_V {
		//z.Memory.ContendReadNoMreq_loop(z.HL.Get(), 1, 5)
		z.pc -= 2
		z.Tstates += 5
	}
	z.HL.Dec()
}

/* INDR */
func instrED__INDR(z *Z80, opcode byte) {
	z.Tstates += 16
	//z.Memory.ContendReadNoMreq(z.IR(), 1)
	var initemp byte = z.readPort(z.BC.Get())
	z.bus.WriteMemory(z.HL.Get(), initemp)

	z.B--
	var initemp2 byte = initemp + z.C - 1
	z.F = ternOpB(initemp&0x80 != 0, FLAG_N, 0) |
		ternOpB(initemp2 < initemp, FLAG_H|FLAG_C, 0) |
		ternOpB(parityTable[(initemp2&0x07)^z.B] != 0, FLAG_P, 0) |
		sz53Table[z.B]

	if z.B != 0 {
		//z.Memory.ContendWriteNoMreq_loop(z.HL.Get(), 1, 5)
		z.pc -= 2
		z.Tstates += 5
	}
	z.HL.Dec()
}

/* OTDR */
func instrED__OTDR(z *Z80, opcode byte) {
	z.Tstates += 16
	//z.Memory.ContendReadNoMreq(z.IR(), 1)
	var outitemp byte = z.bus.ReadMemory(z.HL.Get())
	z.B-- /* This does happen first, despite what the specs say */
	z.writePort(z.BC.Get(), outitemp)

	z.HL.Dec()
	var outitemp2 byte = outitemp + z.L
	z.F = ternOpB((outitemp&0x80) != 0, FLAG_N, 0) |
		ternOpB(outitemp2 < outitemp, FLAG_H|FLAG_C, 0) |
		ternOpB(parityTable[(outitemp2&0x07)^z.B] != 0, FLAG_P, 0) |
		sz53Table[z.B]

	if z.B != 0 {
		//z.Memory.ContendReadNoMreq_loop(z.BC.Get(), 1, 5)
		z.pc -= 2
		z.Tstates += 5
	}
}

/* slttrap */
func instrED__SLTTRAP(z *Z80, opcode byte) {
	z.Tstates += 4 //????
	z.sltTrap(int16(z.HL.Get()), z.A)
}
