package ZzEmu

func initOpcodesED() {

	// 	// BEGIN of 0xed shifted opcodes
	// 	/* IN B,(C) */
	// 	OpcodesMap[SHIFT_0xED+0x40] = instrED__IN_B_iC
	// 	/* OUT (C),B */
	// 	OpcodesMap[SHIFT_0xED+0x41] = instrED__OUT_iC_B
	// 	/* SBC HL,BC */
	// 	OpcodesMap[SHIFT_0xED+0x42] = instrED__SBC_HL_BC
	// 	/* LD (nnnn),BC */
	// 	OpcodesMap[SHIFT_0xED+0x43] = instrED__LD_iNNNN_BC
	// 	/* NEG */
	// 	OpcodesMap[SHIFT_0xED+0x7c] = instrED__NEG
	// 	// Fallthrough cases
	// 	OpcodesMap[SHIFT_0xED+0x44] = OpcodesMap[SHIFT_0xED+0x7c]
	// 	OpcodesMap[SHIFT_0xED+0x4c] = OpcodesMap[SHIFT_0xED+0x7c]
	// 	OpcodesMap[SHIFT_0xED+0x54] = OpcodesMap[SHIFT_0xED+0x7c]
	// 	OpcodesMap[SHIFT_0xED+0x5c] = OpcodesMap[SHIFT_0xED+0x7c]
	// 	OpcodesMap[SHIFT_0xED+0x64] = OpcodesMap[SHIFT_0xED+0x7c]
	// 	OpcodesMap[SHIFT_0xED+0x6c] = OpcodesMap[SHIFT_0xED+0x7c]
	// 	OpcodesMap[SHIFT_0xED+0x74] = OpcodesMap[SHIFT_0xED+0x7c]
	// 	/* RETN */
	// 	OpcodesMap[SHIFT_0xED+0x7d] = instrED__RETN
	// 	// Fallthrough cases
	// 	OpcodesMap[SHIFT_0xED+0x45] = OpcodesMap[SHIFT_0xED+0x7d]
	// 	OpcodesMap[SHIFT_0xED+0x4d] = OpcodesMap[SHIFT_0xED+0x7d]
	// 	OpcodesMap[SHIFT_0xED+0x55] = OpcodesMap[SHIFT_0xED+0x7d]
	// 	OpcodesMap[SHIFT_0xED+0x5d] = OpcodesMap[SHIFT_0xED+0x7d]
	// 	OpcodesMap[SHIFT_0xED+0x65] = OpcodesMap[SHIFT_0xED+0x7d]
	// 	OpcodesMap[SHIFT_0xED+0x6d] = OpcodesMap[SHIFT_0xED+0x7d]
	// 	OpcodesMap[SHIFT_0xED+0x75] = OpcodesMap[SHIFT_0xED+0x7d]
	// 	/* IM 0 */
	// 	OpcodesMap[SHIFT_0xED+0x6e] = instrED__IM_0
	// 	// Fallthrough cases
	// 	OpcodesMap[SHIFT_0xED+0x46] = OpcodesMap[SHIFT_0xED+0x6e]
	// 	OpcodesMap[SHIFT_0xED+0x4e] = OpcodesMap[SHIFT_0xED+0x6e]
	// 	OpcodesMap[SHIFT_0xED+0x66] = OpcodesMap[SHIFT_0xED+0x6e]
	// 	/* LD I,A */
	// 	OpcodesMap[SHIFT_0xED+0x47] = instrED__LD_I_A
	// 	/* IN C,(C) */
	// 	OpcodesMap[SHIFT_0xED+0x48] = instrED__IN_C_iC
	// 	/* OUT (C),C */
	// 	OpcodesMap[SHIFT_0xED+0x49] = instrED__OUT_iC_C
	// 	/* ADC HL,BC */
	// 	OpcodesMap[SHIFT_0xED+0x4a] = instrED__ADC_HL_BC
	// 	/* LD BC,(nnnn) */
	// 	OpcodesMap[SHIFT_0xED+0x4b] = instrED__LD_BC_iNNNN
	// 	/* LD R,A */
	// 	OpcodesMap[SHIFT_0xED+0x4f] = instrED__LD_R_A
	// 	/* IN D,(C) */
	// 	OpcodesMap[SHIFT_0xED+0x50] = instrED__IN_D_iC
	// 	/* OUT (C),D */
	// 	OpcodesMap[SHIFT_0xED+0x51] = instrED__OUT_iC_D
	// 	/* SBC HL,DE */
	// 	OpcodesMap[SHIFT_0xED+0x52] = instrED__SBC_HL_DE
	// 	/* LD (nnnn),DE */
	// 	OpcodesMap[SHIFT_0xED+0x53] = instrED__LD_iNNNN_DE
	// 	/* IM 1 */
	// 	OpcodesMap[SHIFT_0xED+0x76] = instrED__IM_1
	// 	// Fallthrough cases
	// 	OpcodesMap[SHIFT_0xED+0x56] = OpcodesMap[SHIFT_0xED+0x76]
	// 	/* LD A,I */
	// 	OpcodesMap[SHIFT_0xED+0x57] = instrED__LD_A_I
	// 	/* IN E,(C) */
	// 	OpcodesMap[SHIFT_0xED+0x58] = instrED__IN_E_iC
	// 	/* OUT (C),E */
	// 	OpcodesMap[SHIFT_0xED+0x59] = instrED__OUT_iC_E
	// 	/* ADC HL,DE */
	// 	OpcodesMap[SHIFT_0xED+0x5a] = instrED__ADC_HL_DE
	// 	/* LD DE,(nnnn) */
	// 	OpcodesMap[SHIFT_0xED+0x5b] = instrED__LD_DE_iNNNN
	// 	/* IM 2 */
	// 	OpcodesMap[SHIFT_0xED+0x7e] = instrED__IM_2
	// 	// Fallthrough cases
	// 	OpcodesMap[SHIFT_0xED+0x5e] = OpcodesMap[SHIFT_0xED+0x7e]
	// 	/* LD A,R */
	// 	OpcodesMap[SHIFT_0xED+0x5f] = instrED__LD_A_R
	// 	/* IN H,(C) */
	// 	OpcodesMap[SHIFT_0xED+0x60] = instrED__IN_H_iC
	// 	/* OUT (C),H */
	// 	OpcodesMap[SHIFT_0xED+0x61] = instrED__OUT_iC_H
	// 	/* SBC HL,HL */
	// 	OpcodesMap[SHIFT_0xED+0x62] = instrED__SBC_HL_HL
	// 	/* LD (nnnn),HL */
	// 	OpcodesMap[SHIFT_0xED+0x63] = instrED__LD_iNNNN_HL
	// 	/* RRD */
	// 	OpcodesMap[SHIFT_0xED+0x67] = instrED__RRD
	// 	/* IN L,(C) */
	// 	OpcodesMap[SHIFT_0xED+0x68] = instrED__IN_L_iC
	// 	/* OUT (C),L */
	// 	OpcodesMap[SHIFT_0xED+0x69] = instrED__OUT_iC_L
	// 	/* ADC HL,HL */
	// 	OpcodesMap[SHIFT_0xED+0x6a] = instrED__ADC_HL_HL
	// 	/* LD HL,(nnnn) */
	// 	OpcodesMap[SHIFT_0xED+0x6b] = instrED__LD_HL_iNNNN
	// 	/* RLD */
	// 	OpcodesMap[SHIFT_0xED+0x6f] = instrED__RLD
	// 	/* IN F,(C) */
	// 	OpcodesMap[SHIFT_0xED+0x70] = instrED__IN_F_iC
	// 	/* OUT (C),0 */
	// 	OpcodesMap[SHIFT_0xED+0x71] = instrED__OUT_iC_0
	// 	/* SBC HL,SP */
	// 	OpcodesMap[SHIFT_0xED+0x72] = instrED__SBC_HL_SP
	// 	/* LD (nnnn),SP */
	// 	OpcodesMap[SHIFT_0xED+0x73] = instrED__LD_iNNNN_SP
	// 	/* IN A,(C) */
	// 	OpcodesMap[SHIFT_0xED+0x78] = instrED__IN_A_iC
	// 	/* OUT (C),A */
	// 	OpcodesMap[SHIFT_0xED+0x79] = instrED__OUT_iC_A
	// 	/* ADC HL,SP */
	// 	OpcodesMap[SHIFT_0xED+0x7a] = instrED__ADC_HL_SP
	// 	/* LD SP,(nnnn) */
	// 	OpcodesMap[SHIFT_0xED+0x7b] = instrED__LD_SP_iNNNN
	// 	/* LDI */
	// 	OpcodesMap[SHIFT_0xED+0xa0] = instrED__LDI
	// 	/* CPI */
	// 	OpcodesMap[SHIFT_0xED+0xa1] = instrED__CPI
	// 	/* INI */
	// 	OpcodesMap[SHIFT_0xED+0xa2] = instrED__INI
	// 	/* OUTI */
	// 	OpcodesMap[SHIFT_0xED+0xa3] = instrED__OUTI
	// 	/* LDD */
	// 	OpcodesMap[SHIFT_0xED+0xa8] = instrED__LDD
	// 	/* CPD */
	// 	OpcodesMap[SHIFT_0xED+0xa9] = instrED__CPD
	// 	/* IND */
	// 	OpcodesMap[SHIFT_0xED+0xaa] = instrED__IND
	// 	/* OUTD */
	// 	OpcodesMap[SHIFT_0xED+0xab] = instrED__OUTD
	// 	/* LDIR */
	// 	OpcodesMap[SHIFT_0xED+0xb0] = instrED__LDIR
	// 	/* CPIR */
	// 	OpcodesMap[SHIFT_0xED+0xb1] = instrED__CPIR
	// 	/* INIR */
	// 	OpcodesMap[SHIFT_0xED+0xb2] = instrED__INIR
	// 	/* OTIR */
	// 	OpcodesMap[SHIFT_0xED+0xb3] = instrED__OTIR
	// 	/* LDDR */
	// 	OpcodesMap[SHIFT_0xED+0xb8] = instrED__LDDR
	// 	/* CPDR */
	// 	OpcodesMap[SHIFT_0xED+0xb9] = instrED__CPDR
	// 	/* INDR */
	// 	OpcodesMap[SHIFT_0xED+0xba] = instrED__INDR
	// 	/* OTDR */
	// 	OpcodesMap[SHIFT_0xED+0xbb] = instrED__OTDR
	// 	/* slttrap */
	// 	OpcodesMap[SHIFT_0xED+0xfb] = instrED__SLTTRAP

	// 	// END of 0xed shifted opcodes
}

// /* IN B,(C) */
// func instrED__IN_B_iC(z80 *Z80) {
// 	z80.in(&z80.B, z80.BC())
// }

// /* OUT (C),B */
// func instrED__OUT_iC_B(z80 *Z80) {
// 	z80.writePort(z80.BC(), z80.B)
// }

// /* SBC HL,BC */
// func instrED__SBC_HL_BC(z80 *Z80) {
// 	z80.memory.ContendReadNoMreq_loop(z80.IR(), 1, 7)
// 	z80.sbc16(z80.BC())
// }

// /* LD (nnnn),BC */
// func instrED__LD_iNNNN_BC(z80 *Z80) {
// 	z80.ld16nnrr(z80.C, z80.B)
// 	// break
// }

// /* NEG */
// func instrED__NEG(z80 *Z80) {
// 	bytetemp := z80.A
// 	z80.A = 0
// 	z80.sub(bytetemp)
// }

// /* RETN */
// func instrED__RETN(z80 *Z80) {
// 	z80.IFF1 = z80.IFF2
// 	z80.ret()
// }

// /* IM 0 */
// func instrED__IM_0(z80 *Z80) {
// 	z80.IM = 0
// }

// /* LD I,A */
// func instrED__LD_I_A(z80 *Z80) {
// 	z80.memory.ContendReadNoMreq(z80.IR(), 1)
// 	z80.I = z80.A
// }

// /* IN C,(C) */
// func instrED__IN_C_iC(z80 *Z80) {
// 	z80.in(&z80.C, z80.BC())
// }

// /* OUT (C),C */
// func instrED__OUT_iC_C(z80 *Z80) {
// 	z80.writePort(z80.BC(), z80.C)
// }

// /* ADC HL,BC */
// func instrED__ADC_HL_BC(z80 *Z80) {
// 	z80.memory.ContendReadNoMreq_loop(z80.IR(), 1, 7)
// 	z80.adc16(z80.BC())
// }

// /* LD BC,(nnnn) */
// func instrED__LD_BC_iNNNN(z80 *Z80) {
// 	z80.ld16rrnn(&z80.C, &z80.B)
// 	// break
// }

// /* LD R,A */
// func instrED__LD_R_A(z80 *Z80) {
// 	z80.memory.ContendReadNoMreq(z80.IR(), 1)
// 	/* Keep the RZX instruction counter right */
// 	z80.rzxInstructionsOffset += (int(z80.R) - int(z80.A))
// 	z80.R, z80.R7 = uint16(z80.A), z80.A
// }

// /* IN D,(C) */
// func instrED__IN_D_iC(z80 *Z80) {
// 	z80.in(&z80.D, z80.BC())
// }

// /* OUT (C),D */
// func instrED__OUT_iC_D(z80 *Z80) {
// 	z80.writePort(z80.BC(), z80.D)
// }

// /* SBC HL,DE */
// func instrED__SBC_HL_DE(z80 *Z80) {
// 	z80.memory.ContendReadNoMreq_loop(z80.IR(), 1, 7)
// 	z80.sbc16(z80.DE())
// }

// /* LD (nnnn),DE */
// func instrED__LD_iNNNN_DE(z80 *Z80) {
// 	z80.ld16nnrr(z80.E, z80.D)
// 	// break
// }

// /* IM 1 */
// func instrED__IM_1(z80 *Z80) {
// 	z80.IM = 1
// }

// /* LD A,I */
// func instrED__LD_A_I(z80 *Z80) {
// 	z80.memory.ContendReadNoMreq(z80.IR(), 1)
// 	z80.A = z80.I
// 	z80.F = (z80.F & FLAG_C) | sz53Table[z80.A] | ternOpB(z80.IFF2 != 0, FLAG_V, 0)
// }

// /* IN E,(C) */
// func instrED__IN_E_iC(z80 *Z80) {
// 	z80.in(&z80.E, z80.BC())
// }

// /* OUT (C),E */
// func instrED__OUT_iC_E(z80 *Z80) {
// 	z80.writePort(z80.BC(), z80.E)
// }

// /* ADC HL,DE */
// func instrED__ADC_HL_DE(z80 *Z80) {
// 	z80.memory.ContendReadNoMreq_loop(z80.IR(), 1, 7)
// 	z80.adc16(z80.DE())
// }

// /* LD DE,(nnnn) */
// func instrED__LD_DE_iNNNN(z80 *Z80) {
// 	z80.ld16rrnn(&z80.E, &z80.D)
// 	// break
// }

// /* IM 2 */
// func instrED__IM_2(z80 *Z80) {
// 	z80.IM = 2
// }

// /* LD A,R */
// func instrED__LD_A_R(z80 *Z80) {
// 	z80.memory.ContendReadNoMreq(z80.IR(), 1)
// 	z80.A = byte(z80.R&0x7f) | (z80.R7 & 0x80)
// 	z80.F = (z80.F & FLAG_C) | sz53Table[z80.A] | ternOpB(z80.IFF2 != 0, FLAG_V, 0)
// }

// /* IN H,(C) */
// func instrED__IN_H_iC(z80 *Z80) {
// 	z80.in(&z80.H, z80.BC())
// }

// /* OUT (C),H */
// func instrED__OUT_iC_H(z80 *Z80) {
// 	z80.writePort(z80.BC(), z80.H)
// }

// /* SBC HL,HL */
// func instrED__SBC_HL_HL(z80 *Z80) {
// 	z80.memory.ContendReadNoMreq_loop(z80.IR(), 1, 7)
// 	z80.sbc16(z80.HL())
// }

// /* LD (nnnn),HL */
// func instrED__LD_iNNNN_HL(z80 *Z80) {
// 	z80.ld16nnrr(z80.L, z80.H)
// 	// break
// }

// /* RRD */
// func instrED__RRD(z80 *Z80) {
// 	var bytetemp byte = z80.memory.ReadByte(z80.HL())
// 	z80.memory.ContendReadNoMreq_loop(z80.HL(), 1, 4)
// 	z80.memory.WriteByte(z80.HL(), (z80.A<<4)|(bytetemp>>4))
// 	z80.A = (z80.A & 0xf0) | (bytetemp & 0x0f)
// 	z80.F = (z80.F & FLAG_C) | sz53pTable[z80.A]
// }

// /* IN L,(C) */
// func instrED__IN_L_iC(z80 *Z80) {
// 	z80.in(&z80.L, z80.BC())
// }

// /* OUT (C),L */
// func instrED__OUT_iC_L(z80 *Z80) {
// 	z80.writePort(z80.BC(), z80.L)
// }

// /* ADC HL,HL */
// func instrED__ADC_HL_HL(z80 *Z80) {
// 	z80.memory.ContendReadNoMreq_loop(z80.IR(), 1, 7)
// 	z80.adc16(z80.HL())
// }

// /* LD HL,(nnnn) */
// func instrED__LD_HL_iNNNN(z80 *Z80) {
// 	z80.ld16rrnn(&z80.L, &z80.H)
// 	// break
// }

// /* RLD */
// func instrED__RLD(z80 *Z80) {
// 	var bytetemp byte = z80.memory.ReadByte(z80.HL())
// 	z80.memory.ContendReadNoMreq_loop(z80.HL(), 1, 4)
// 	z80.memory.WriteByte(z80.HL(), (bytetemp<<4)|(z80.A&0x0f))
// 	z80.A = (z80.A & 0xf0) | (bytetemp >> 4)
// 	z80.F = (z80.F & FLAG_C) | sz53pTable[z80.A]
// }

// /* IN F,(C) */
// func instrED__IN_F_iC(z80 *Z80) {
// 	var bytetemp byte
// 	z80.in(&bytetemp, z80.BC())
// }

// /* OUT (C),0 */
// func instrED__OUT_iC_0(z80 *Z80) {
// 	z80.writePort(z80.BC(), 0)
// }

// /* SBC HL,SP */
// func instrED__SBC_HL_SP(z80 *Z80) {
// 	z80.memory.ContendReadNoMreq_loop(z80.IR(), 1, 7)
// 	z80.sbc16(z80.SP())
// }

// /* LD (nnnn),SP */
// func instrED__LD_iNNNN_SP(z80 *Z80) {
// 	sph, spl := splitWord(z80.sp)
// 	z80.ld16nnrr(spl, sph)
// 	// break
// }

// /* IN A,(C) */
// func instrED__IN_A_iC(z80 *Z80) {
// 	z80.in(&z80.A, z80.BC())
// }

// /* OUT (C),A */
// func instrED__OUT_iC_A(z80 *Z80) {
// 	z80.writePort(z80.BC(), z80.A)
// }

// /* ADC HL,SP */
// func instrED__ADC_HL_SP(z80 *Z80) {
// 	z80.memory.ContendReadNoMreq_loop(z80.IR(), 1, 7)
// 	z80.adc16(z80.SP())
// }

// /* LD SP,(nnnn) */
// func instrED__LD_SP_iNNNN(z80 *Z80) {
// 	sph, spl := splitWord(z80.SP())
// 	z80.ld16rrnn(&spl, &sph)
// 	z80.SetSP(joinBytes(sph, spl))
// 	// break
// }

// /* LDI */
// func instrED__LDI(z80 *Z80) {
// 	var bytetemp byte = z80.memory.ReadByte(z80.HL())
// 	z80.DecBC()
// 	z80.memory.WriteByte(z80.DE(), bytetemp)
// 	z80.memory.ContendWriteNoMreq_loop(z80.DE(), 1, 2)
// 	z80.IncDE()
// 	z80.IncHL()
// 	bytetemp += z80.A
// 	z80.F = (z80.F & (FLAG_C | FLAG_Z | FLAG_S)) |
// 		ternOpB(z80.BC() != 0, FLAG_V, 0) |
// 		(bytetemp & FLAG_3) |
// 		ternOpB((bytetemp&0x02) != 0, FLAG_5, 0)
// }

// /* CPI */
// func instrED__CPI(z80 *Z80) {
// 	var value byte = z80.memory.ReadByte(z80.HL())
// 	var bytetemp byte = z80.A - value
// 	var lookup byte = ((z80.A & 0x08) >> 3) | (((value) & 0x08) >> 2) | ((bytetemp & 0x08) >> 1)
// 	z80.memory.ContendReadNoMreq_loop(z80.HL(), 1, 5)
// 	z80.IncHL()
// 	z80.DecBC()
// 	z80.F = (z80.F & FLAG_C) | ternOpB(z80.BC() != 0, FLAG_V|FLAG_N, FLAG_N) | halfcarrySubTable[lookup] | ternOpB(bytetemp != 0, 0, FLAG_Z) | (bytetemp & FLAG_S)
// 	if (z80.F & FLAG_H) != 0 {
// 		bytetemp--
// 	}
// 	z80.F |= (bytetemp & FLAG_3) | ternOpB((bytetemp&0x02) != 0, FLAG_5, 0)
// }

// /* INI */
// func instrED__INI(z80 *Z80) {
// 	z80.memory.ContendReadNoMreq(z80.IR(), 1)
// 	var initemp byte = z80.readPort(z80.BC())
// 	z80.memory.WriteByte(z80.HL(), initemp)

// 	z80.B--
// 	z80.IncHL()
// 	var initemp2 byte = initemp + z80.C + 1
// 	z80.F = ternOpB((initemp&0x80) != 0, FLAG_N, 0) |
// 		ternOpB(initemp2 < initemp, FLAG_H|FLAG_C, 0) |
// 		ternOpB(parityTable[(initemp2&0x07)^z80.B] != 0, FLAG_P, 0) |
// 		sz53Table[z80.B]
// }

// /* OUTI */
// func instrED__OUTI(z80 *Z80) {
// 	z80.memory.ContendReadNoMreq(z80.IR(), 1)
// 	var outitemp byte = z80.memory.ReadByte(z80.HL())
// 	z80.B-- /* This does happen first, despite what the specs say */
// 	z80.writePort(z80.BC(), outitemp)

// 	z80.IncHL()
// 	var outitemp2 byte = outitemp + z80.L
// 	z80.F = ternOpB((outitemp&0x80) != 0, FLAG_N, 0) |
// 		ternOpB(outitemp2 < outitemp, FLAG_H|FLAG_C, 0) |
// 		ternOpB(parityTable[(outitemp2&0x07)^z80.B] != 0, FLAG_P, 0) |
// 		sz53Table[z80.B]
// }

// /* LDD */
// func instrED__LDD(z80 *Z80) {
// 	var bytetemp byte = z80.memory.ReadByte(z80.HL())
// 	z80.DecBC()
// 	z80.memory.WriteByte(z80.DE(), bytetemp)
// 	z80.memory.ContendWriteNoMreq_loop(z80.DE(), 1, 2)
// 	z80.DecDE()
// 	z80.DecHL()
// 	bytetemp += z80.A
// 	z80.F = (z80.F & (FLAG_C | FLAG_Z | FLAG_S)) |
// 		ternOpB(z80.BC() != 0, FLAG_V, 0) |
// 		(bytetemp & FLAG_3) |
// 		ternOpB((bytetemp&0x02) != 0, FLAG_5, 0)
// }

// /* CPD */
// func instrED__CPD(z80 *Z80) {
// 	var value byte = z80.memory.ReadByte(z80.HL())
// 	var bytetemp byte = z80.A - value
// 	var lookup byte = ((z80.A & 0x08) >> 3) | (((value) & 0x08) >> 2) | ((bytetemp & 0x08) >> 1)
// 	z80.memory.ContendReadNoMreq_loop(z80.HL(), 1, 5)
// 	z80.DecHL()
// 	z80.DecBC()
// 	z80.F = (z80.F & FLAG_C) | ternOpB(z80.BC() != 0, FLAG_V|FLAG_N, FLAG_N) | halfcarrySubTable[lookup] | ternOpB(bytetemp != 0, 0, FLAG_Z) | (bytetemp & FLAG_S)
// 	if (z80.F & FLAG_H) != 0 {
// 		bytetemp--
// 	}
// 	z80.F |= (bytetemp & FLAG_3) | ternOpB((bytetemp&0x02) != 0, FLAG_5, 0)
// }

// /* IND */
// func instrED__IND(z80 *Z80) {
// 	z80.memory.ContendReadNoMreq(z80.IR(), 1)
// 	var initemp byte = z80.readPort(z80.BC())
// 	z80.memory.WriteByte(z80.HL(), initemp)

// 	z80.B--
// 	z80.DecHL()
// 	var initemp2 byte = initemp + z80.C - 1
// 	z80.F = ternOpB((initemp&0x80) != 0, FLAG_N, 0) |
// 		ternOpB(initemp2 < initemp, FLAG_H|FLAG_C, 0) |
// 		ternOpB(parityTable[(initemp2&0x07)^z80.B] != 0, FLAG_P, 0) |
// 		sz53Table[z80.B]
// }

// /* OUTD */
// func instrED__OUTD(z80 *Z80) {
// 	z80.memory.ContendReadNoMreq(z80.IR(), 1)
// 	var outitemp byte = z80.memory.ReadByte(z80.HL())
// 	z80.B-- /* This does happen first, despite what the specs say */
// 	z80.writePort(z80.BC(), outitemp)

// 	z80.DecHL()
// 	var outitemp2 byte = outitemp + z80.L
// 	z80.F = ternOpB((outitemp&0x80) != 0, FLAG_N, 0) |
// 		ternOpB(outitemp2 < outitemp, FLAG_H|FLAG_C, 0) |
// 		ternOpB(parityTable[(outitemp2&0x07)^z80.B] != 0, FLAG_P, 0) |
// 		sz53Table[z80.B]
// }

// /* LDIR */
// func instrED__LDIR(z80 *Z80) {
// 	var bytetemp byte = z80.memory.ReadByte(z80.HL())
// 	z80.memory.WriteByte(z80.DE(), bytetemp)
// 	z80.memory.ContendWriteNoMreq_loop(z80.DE(), 1, 2)
// 	z80.DecBC()
// 	bytetemp += z80.A
// 	z80.F = (z80.F & (FLAG_C | FLAG_Z | FLAG_S)) | ternOpB(z80.BC() != 0, FLAG_V, 0) | (bytetemp & FLAG_3) | ternOpB((bytetemp&0x02 != 0), FLAG_5, 0)
// 	if z80.BC() != 0 {
// 		z80.memory.ContendWriteNoMreq_loop(z80.DE(), 1, 5)
// 		z80.DecPC(2)
// 	}
// 	z80.IncHL()
// 	z80.IncDE()
// }

// /* CPIR */
// func instrED__CPIR(z80 *Z80) {
// 	var value byte = z80.memory.ReadByte(z80.HL())
// 	var bytetemp byte = z80.A - value
// 	var lookup byte = ((z80.A & 0x08) >> 3) | (((value) & 0x08) >> 2) | ((bytetemp & 0x08) >> 1)
// 	z80.memory.ContendReadNoMreq_loop(z80.HL(), 1, 5)
// 	z80.DecBC()
// 	z80.F = (z80.F & FLAG_C) | (ternOpB(z80.BC() != 0, (FLAG_V | FLAG_N), FLAG_N)) | halfcarrySubTable[lookup] | (ternOpB(bytetemp != 0, 0, FLAG_Z)) | (bytetemp & FLAG_S)
// 	if (z80.F & FLAG_H) != 0 {
// 		bytetemp--
// 	}
// 	z80.F |= (bytetemp & FLAG_3) | ternOpB((bytetemp&0x02) != 0, FLAG_5, 0)
// 	if (z80.F & (FLAG_V | FLAG_Z)) == FLAG_V {
// 		z80.memory.ContendReadNoMreq_loop(z80.HL(), 1, 5)
// 		z80.DecPC(2)
// 	}
// 	z80.IncHL()
// }

// /* INIR */
// func instrED__INIR(z80 *Z80) {
// 	z80.memory.ContendReadNoMreq(z80.IR(), 1)
// 	var initemp byte = z80.readPort(z80.BC())
// 	z80.memory.WriteByte(z80.HL(), initemp)

// 	z80.B--
// 	var initemp2 byte = initemp + z80.C + 1
// 	z80.F = ternOpB(initemp&0x80 != 0, FLAG_N, 0) |
// 		ternOpB(initemp2 < initemp, FLAG_H|FLAG_C, 0) |
// 		ternOpB(parityTable[(initemp2&0x07)^z80.B] != 0, FLAG_P, 0) |
// 		sz53Table[z80.B]

// 	if z80.B != 0 {
// 		z80.memory.ContendWriteNoMreq_loop(z80.HL(), 1, 5)
// 		z80.DecPC(2)
// 	}
// 	z80.IncHL()
// }

// /* OTIR */
// func instrED__OTIR(z80 *Z80) {
// 	z80.memory.ContendReadNoMreq(z80.IR(), 1)
// 	var outitemp byte = z80.memory.ReadByte(z80.HL())
// 	z80.B-- /* This does happen first, despite what the specs say */
// 	z80.writePort(z80.BC(), outitemp)

// 	z80.IncHL()
// 	var outitemp2 byte = outitemp + z80.L
// 	z80.F = ternOpB((outitemp&0x80) != 0, FLAG_N, 0) |
// 		ternOpB(outitemp2 < outitemp, FLAG_H|FLAG_C, 0) |
// 		ternOpB(parityTable[(outitemp2&0x07)^z80.B] != 0, FLAG_P, 0) |
// 		sz53Table[z80.B]

// 	if z80.B != 0 {
// 		z80.memory.ContendReadNoMreq_loop(z80.BC(), 1, 5)
// 		z80.DecPC(2)
// 	}
// }

// /* LDDR */
// func instrED__LDDR(z80 *Z80) {
// 	var bytetemp byte = z80.memory.ReadByte(z80.HL())
// 	z80.memory.WriteByte(z80.DE(), bytetemp)
// 	z80.memory.ContendWriteNoMreq_loop(z80.DE(), 1, 2)
// 	z80.DecBC()
// 	bytetemp += z80.A
// 	z80.F = (z80.F & (FLAG_C | FLAG_Z | FLAG_S)) | ternOpB(z80.BC() != 0, FLAG_V, 0) | (bytetemp & FLAG_3) | ternOpB((bytetemp&0x02 != 0), FLAG_5, 0)
// 	if z80.BC() != 0 {
// 		z80.memory.ContendWriteNoMreq_loop(z80.DE(), 1, 5)
// 		z80.DecPC(2)
// 	}
// 	z80.DecHL()
// 	z80.DecDE()
// }

// /* CPDR */
// func instrED__CPDR(z80 *Z80) {
// 	var value byte = z80.memory.ReadByte(z80.HL())
// 	var bytetemp byte = z80.A - value
// 	var lookup byte = ((z80.A & 0x08) >> 3) | (((value) & 0x08) >> 2) | ((bytetemp & 0x08) >> 1)
// 	z80.memory.ContendReadNoMreq_loop(z80.HL(), 1, 5)
// 	z80.DecBC()
// 	z80.F = (z80.F & FLAG_C) | (ternOpB(z80.BC() != 0, (FLAG_V | FLAG_N), FLAG_N)) | halfcarrySubTable[lookup] | (ternOpB(bytetemp != 0, 0, FLAG_Z)) | (bytetemp & FLAG_S)
// 	if (z80.F & FLAG_H) != 0 {
// 		bytetemp--
// 	}
// 	z80.F |= (bytetemp & FLAG_3) | ternOpB((bytetemp&0x02) != 0, FLAG_5, 0)
// 	if (z80.F & (FLAG_V | FLAG_Z)) == FLAG_V {
// 		z80.memory.ContendReadNoMreq_loop(z80.HL(), 1, 5)
// 		z80.DecPC(2)
// 	}
// 	z80.DecHL()
// }

// /* INDR */
// func instrED__INDR(z80 *Z80) {
// 	z80.memory.ContendReadNoMreq(z80.IR(), 1)
// 	var initemp byte = z80.readPort(z80.BC())
// 	z80.memory.WriteByte(z80.HL(), initemp)

// 	z80.B--
// 	var initemp2 byte = initemp + z80.C - 1
// 	z80.F = ternOpB(initemp&0x80 != 0, FLAG_N, 0) |
// 		ternOpB(initemp2 < initemp, FLAG_H|FLAG_C, 0) |
// 		ternOpB(parityTable[(initemp2&0x07)^z80.B] != 0, FLAG_P, 0) |
// 		sz53Table[z80.B]

// 	if z80.B != 0 {
// 		z80.memory.ContendWriteNoMreq_loop(z80.HL(), 1, 5)
// 		z80.DecPC(2)
// 	}
// 	z80.DecHL()
// }

// /* OTDR */
// func instrED__OTDR(z80 *Z80) {
// 	z80.memory.ContendReadNoMreq(z80.IR(), 1)
// 	var outitemp byte = z80.memory.ReadByte(z80.HL())
// 	z80.B-- /* This does happen first, despite what the specs say */
// 	z80.writePort(z80.BC(), outitemp)

// 	z80.DecHL()
// 	var outitemp2 byte = outitemp + z80.L
// 	z80.F = ternOpB((outitemp&0x80) != 0, FLAG_N, 0) |
// 		ternOpB(outitemp2 < outitemp, FLAG_H|FLAG_C, 0) |
// 		ternOpB(parityTable[(outitemp2&0x07)^z80.B] != 0, FLAG_P, 0) |
// 		sz53Table[z80.B]

// 	if z80.B != 0 {
// 		z80.memory.ContendReadNoMreq_loop(z80.BC(), 1, 5)
// 		z80.DecPC(2)
// 	}
// }

// /* slttrap */
// func instrED__SLTTRAP(z80 *Z80) {
// 	z80.sltTrap(int16(z80.HL()), z80.A)
// }
