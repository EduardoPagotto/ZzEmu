package ZzEmu

func initOpCodesDD() {
	// 	// BEGIN of 0xdd shifted opcodes
	// 	/* ADD REGISTER,BC */
	// 	OpcodesMap[SHIFT_0xDD+0x09] = instrDD__ADD_REG_BC
	// 	/* ADD REGISTER,DE */
	// 	OpcodesMap[SHIFT_0xDD+0x19] = instrDD__ADD_REG_DE
	// 	/* LD REGISTER,nnnn */
	// 	OpcodesMap[SHIFT_0xDD+0x21] = instrDD__LD_REG_NNNN
	// 	/* LD (nnnn),REGISTER */
	// 	OpcodesMap[SHIFT_0xDD+0x22] = instrDD__LD_iNNNN_REG
	// 	/* INC REGISTER */
	// 	OpcodesMap[SHIFT_0xDD+0x23] = instrDD__INC_REG
	// 	/* INC REGISTERH */
	// 	OpcodesMap[SHIFT_0xDD+0x24] = instrDD__INC_REGH
	// 	/* DEC REGISTERH */
	// 	OpcodesMap[SHIFT_0xDD+0x25] = instrDD__DEC_REGH
	// 	/* LD REGISTERH,nn */
	// 	OpcodesMap[SHIFT_0xDD+0x26] = instrDD__LD_REGH_NN
	// 	/* ADD REGISTER,REGISTER */
	// 	OpcodesMap[SHIFT_0xDD+0x29] = instrDD__ADD_REG_REG
	// 	/* LD REGISTER,(nnnn) */
	// 	OpcodesMap[SHIFT_0xDD+0x2a] = instrDD__LD_REG_iNNNN
	// 	/* DEC REGISTER */
	// 	OpcodesMap[SHIFT_0xDD+0x2b] = instrDD__DEC_REG
	// 	/* INC REGISTERL */
	// 	OpcodesMap[SHIFT_0xDD+0x2c] = instrDD__INC_REGL
	// 	/* DEC REGISTERL */
	// 	OpcodesMap[SHIFT_0xDD+0x2d] = instrDD__DEC_REGL
	// 	/* LD REGISTERL,nn */
	// 	OpcodesMap[SHIFT_0xDD+0x2e] = instrDD__LD_REGL_NN
	// 	/* INC (REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDD+0x34] = instrDD__INC_iREGpDD
	// 	/* DEC (REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDD+0x35] = instrDD__DEC_iREGpDD
	// 	/* LD (REGISTER+dd),nn */
	// 	OpcodesMap[SHIFT_0xDD+0x36] = instrDD__LD_iREGpDD_NN
	// 	/* ADD REGISTER,SP */
	// 	OpcodesMap[SHIFT_0xDD+0x39] = instrDD__ADD_REG_SP
	// 	/* LD B,REGISTERH */
	// 	OpcodesMap[SHIFT_0xDD+0x44] = instrDD__LD_B_REGH
	// 	/* LD B,REGISTERL */
	// 	OpcodesMap[SHIFT_0xDD+0x45] = instrDD__LD_B_REGL
	// 	/* LD B,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDD+0x46] = instrDD__LD_B_iREGpDD
	// 	/* LD C,REGISTERH */
	// 	OpcodesMap[SHIFT_0xDD+0x4c] = instrDD__LD_C_REGH
	// 	/* LD C,REGISTERL */
	// 	OpcodesMap[SHIFT_0xDD+0x4d] = instrDD__LD_C_REGL
	// 	/* LD C,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDD+0x4e] = instrDD__LD_C_iREGpDD
	// 	/* LD D,REGISTERH */
	// 	OpcodesMap[SHIFT_0xDD+0x54] = instrDD__LD_D_REGH
	// 	/* LD D,REGISTERL */
	// 	OpcodesMap[SHIFT_0xDD+0x55] = instrDD__LD_D_REGL
	// 	/* LD D,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDD+0x56] = instrDD__LD_D_iREGpDD
	// 	/* LD E,REGISTERH */
	// 	OpcodesMap[SHIFT_0xDD+0x5c] = instrDD__LD_E_REGH
	// 	/* LD E,REGISTERL */
	// 	OpcodesMap[SHIFT_0xDD+0x5d] = instrDD__LD_E_REGL
	// 	/* LD E,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDD+0x5e] = instrDD__LD_E_iREGpDD
	// 	/* LD REGISTERH,B */
	// 	OpcodesMap[SHIFT_0xDD+0x60] = instrDD__LD_REGH_B
	// 	/* LD REGISTERH,C */
	// 	OpcodesMap[SHIFT_0xDD+0x61] = instrDD__LD_REGH_C
	// 	/* LD REGISTERH,D */
	// 	OpcodesMap[SHIFT_0xDD+0x62] = instrDD__LD_REGH_D
	// 	/* LD REGISTERH,E */
	// 	OpcodesMap[SHIFT_0xDD+0x63] = instrDD__LD_REGH_E
	// 	/* LD REGISTERH,REGISTERH */
	// 	OpcodesMap[SHIFT_0xDD+0x64] = instrDD__LD_REGH_REGH
	// 	/* LD REGISTERH,REGISTERL */
	// 	OpcodesMap[SHIFT_0xDD+0x65] = instrDD__LD_REGH_REGL
	// 	/* LD H,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDD+0x66] = instrDD__LD_H_iREGpDD
	// 	/* LD REGISTERH,A */
	// 	OpcodesMap[SHIFT_0xDD+0x67] = instrDD__LD_REGH_A
	// 	/* LD REGISTERL,B */
	// 	OpcodesMap[SHIFT_0xDD+0x68] = instrDD__LD_REGL_B
	// 	/* LD REGISTERL,C */
	// 	OpcodesMap[SHIFT_0xDD+0x69] = instrDD__LD_REGL_C
	// 	/* LD REGISTERL,D */
	// 	OpcodesMap[SHIFT_0xDD+0x6a] = instrDD__LD_REGL_D
	// 	/* LD REGISTERL,E */
	// 	OpcodesMap[SHIFT_0xDD+0x6b] = instrDD__LD_REGL_E
	// 	/* LD REGISTERL,REGISTERH */
	// 	OpcodesMap[SHIFT_0xDD+0x6c] = instrDD__LD_REGL_REGH
	// 	/* LD REGISTERL,REGISTERL */
	// 	OpcodesMap[SHIFT_0xDD+0x6d] = instrDD__LD_REGL_REGL
	// 	/* LD L,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDD+0x6e] = instrDD__LD_L_iREGpDD
	// 	/* LD REGISTERL,A */
	// 	OpcodesMap[SHIFT_0xDD+0x6f] = instrDD__LD_REGL_A
	// 	/* LD (REGISTER+dd),B */
	// 	OpcodesMap[SHIFT_0xDD+0x70] = instrDD__LD_iREGpDD_B
	// 	/* LD (REGISTER+dd),C */
	// 	OpcodesMap[SHIFT_0xDD+0x71] = instrDD__LD_iREGpDD_C
	// 	/* LD (REGISTER+dd),D */
	// 	OpcodesMap[SHIFT_0xDD+0x72] = instrDD__LD_iREGpDD_D
	// 	/* LD (REGISTER+dd),E */
	// 	OpcodesMap[SHIFT_0xDD+0x73] = instrDD__LD_iREGpDD_E
	// 	/* LD (REGISTER+dd),H */
	// 	OpcodesMap[SHIFT_0xDD+0x74] = instrDD__LD_iREGpDD_H
	// 	/* LD (REGISTER+dd),L */
	// 	OpcodesMap[SHIFT_0xDD+0x75] = instrDD__LD_iREGpDD_L
	// 	/* LD (REGISTER+dd),A */
	// 	OpcodesMap[SHIFT_0xDD+0x77] = instrDD__LD_iREGpDD_A
	// 	/* LD A,REGISTERH */
	// 	OpcodesMap[SHIFT_0xDD+0x7c] = instrDD__LD_A_REGH
	// 	/* LD A,REGISTERL */
	// 	OpcodesMap[SHIFT_0xDD+0x7d] = instrDD__LD_A_REGL
	// 	/* LD A,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDD+0x7e] = instrDD__LD_A_iREGpDD
	// 	/* ADD A,REGISTERH */
	// 	OpcodesMap[SHIFT_0xDD+0x84] = instrDD__ADD_A_REGH
	// 	/* ADD A,REGISTERL */
	// 	OpcodesMap[SHIFT_0xDD+0x85] = instrDD__ADD_A_REGL
	// 	/* ADD A,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDD+0x86] = instrDD__ADD_A_iREGpDD
	// 	/* ADC A,REGISTERH */
	// 	OpcodesMap[SHIFT_0xDD+0x8c] = instrDD__ADC_A_REGH
	// 	/* ADC A,REGISTERL */
	// 	OpcodesMap[SHIFT_0xDD+0x8d] = instrDD__ADC_A_REGL
	// 	/* ADC A,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDD+0x8e] = instrDD__ADC_A_iREGpDD
	// 	/* SUB A,REGISTERH */
	// 	OpcodesMap[SHIFT_0xDD+0x94] = instrDD__SUB_A_REGH
	// 	/* SUB A,REGISTERL */
	// 	OpcodesMap[SHIFT_0xDD+0x95] = instrDD__SUB_A_REGL
	// 	/* SUB A,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDD+0x96] = instrDD__SUB_A_iREGpDD
	// 	/* SBC A,REGISTERH */
	// 	OpcodesMap[SHIFT_0xDD+0x9c] = instrDD__SBC_A_REGH
	// 	/* SBC A,REGISTERL */
	// 	OpcodesMap[SHIFT_0xDD+0x9d] = instrDD__SBC_A_REGL
	// 	/* SBC A,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDD+0x9e] = instrDD__SBC_A_iREGpDD
	// 	/* AND A,REGISTERH */
	// 	OpcodesMap[SHIFT_0xDD+0xa4] = instrDD__AND_A_REGH
	// 	/* AND A,REGISTERL */
	// 	OpcodesMap[SHIFT_0xDD+0xa5] = instrDD__AND_A_REGL
	// 	/* AND A,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDD+0xa6] = instrDD__AND_A_iREGpDD
	// 	/* XOR A,REGISTERH */
	// 	OpcodesMap[SHIFT_0xDD+0xac] = instrDD__XOR_A_REGH
	// 	/* XOR A,REGISTERL */
	// 	OpcodesMap[SHIFT_0xDD+0xad] = instrDD__XOR_A_REGL
	// 	/* XOR A,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDD+0xae] = instrDD__XOR_A_iREGpDD
	// 	/* OR A,REGISTERH */
	// 	OpcodesMap[SHIFT_0xDD+0xb4] = instrDD__OR_A_REGH
	// 	/* OR A,REGISTERL */
	// 	OpcodesMap[SHIFT_0xDD+0xb5] = instrDD__OR_A_REGL
	// 	/* OR A,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDD+0xb6] = instrDD__OR_A_iREGpDD
	// 	/* CP A,REGISTERH */
	// 	OpcodesMap[SHIFT_0xDD+0xbc] = instrDD__CP_A_REGH
	// 	/* CP A,REGISTERL */
	// 	OpcodesMap[SHIFT_0xDD+0xbd] = instrDD__CP_A_REGL
	// 	/* CP A,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDD+0xbe] = instrDD__CP_A_iREGpDD
	// 	/* shift DDFDCB */
	OpcodeDDMap[0xcb] = instrDD__SHIFT_DDFDCB
	// 	/* POP REGISTER */
	// 	OpcodesMap[SHIFT_0xDD+0xe1] = instrDD__POP_REG
	// 	/* EX (SP),REGISTER */
	// 	OpcodesMap[SHIFT_0xDD+0xe3] = instrDD__EX_iSP_REG
	// 	/* PUSH REGISTER */
	// 	OpcodesMap[SHIFT_0xDD+0xe5] = instrDD__PUSH_REG
	// 	/* JP REGISTER */
	// 	OpcodesMap[SHIFT_0xDD+0xe9] = instrDD__JP_REG
	// 	/* LD SP,REGISTER */
	// 	OpcodesMap[SHIFT_0xDD+0xf9] = instrDD__LD_SP_REG

	// 	// END of 0xdd shifted opcodes

}

// /* ADD ix,BC */
// func instrDD__ADD_REG_BC(z80 *Z80) {
// 	z80.memory.ContendReadNoMreq_loop(z80.IR(), 1, 7)
// 	z80.add16(z80.ix, z80.BC())
// }

// /* ADD ix,DE */
// func instrDD__ADD_REG_DE(z80 *Z80) {
// 	z80.memory.ContendReadNoMreq_loop(z80.IR(), 1, 7)
// 	z80.add16(z80.ix, z80.DE())
// }

// /* LD ix,nnnn */
// func instrDD__LD_REG_NNNN(z80 *Z80) {
// 	b1 := z80.memory.ReadByte(z80.PC())
// 	z80.IncPC(1)
// 	b2 := z80.memory.ReadByte(z80.PC())
// 	z80.IncPC(1)
// 	z80.SetIX(joinBytes(b2, b1))
// }

// /* LD (nnnn),ix */
// func instrDD__LD_iNNNN_REG(z80 *Z80) {
// 	z80.ld16nnrr(z80.IXL, z80.IXH)
// 	// break
// }

// /* INC ix */
// func instrDD__INC_REG(z80 *Z80) {
// 	z80.memory.ContendReadNoMreq_loop(z80.IR(), 1, 2)
// 	z80.IncIX()
// }

// /* INC IXH */
// func instrDD__INC_REGH(z80 *Z80) {
// 	z80.incIXH()
// }

// /* DEC IXH */
// func instrDD__DEC_REGH(z80 *Z80) {
// 	z80.decIXH()
// }

// /* LD IXH,nn */
// func instrDD__LD_REGH_NN(z80 *Z80) {
// 	z80.IXH = z80.memory.ReadByte(z80.PC())
// 	z80.IncPC(1)
// }

// /* ADD ix,ix */
// func instrDD__ADD_REG_REG(z80 *Z80) {
// 	z80.memory.ContendReadNoMreq_loop(z80.IR(), 1, 7)
// 	z80.add16(z80.ix, z80.IX())
// }

// /* LD ix,(nnnn) */
// func instrDD__LD_REG_iNNNN(z80 *Z80) {
// 	z80.ld16rrnn(&z80.IXL, &z80.IXH)
// 	// break
// }

// /* DEC ix */
// func instrDD__DEC_REG(z80 *Z80) {
// 	z80.memory.ContendReadNoMreq_loop(z80.IR(), 1, 2)
// 	z80.DecIX()
// }

// /* INC IXL */
// func instrDD__INC_REGL(z80 *Z80) {
// 	z80.incIXL()
// }

// /* DEC IXL */
// func instrDD__DEC_REGL(z80 *Z80) {
// 	z80.decIXL()
// }

// /* LD IXL,nn */
// func instrDD__LD_REGL_NN(z80 *Z80) {
// 	z80.IXL = z80.memory.ReadByte(z80.PC())
// 	z80.IncPC(1)
// }

// /* INC (ix+dd) */
// func instrDD__INC_iREGpDD(z80 *Z80) {
// 	var offset byte = z80.memory.ReadByte(z80.PC())
// 	z80.memory.ContendReadNoMreq_loop(z80.PC(), 1, 5)
// 	z80.IncPC(1)
// 	var wordtemp uint16 = z80.IX() + uint16(signExtend(offset))
// 	var bytetemp byte = z80.memory.ReadByte(wordtemp)
// 	z80.memory.ContendReadNoMreq(wordtemp, 1)
// 	z80.inc(&bytetemp)
// 	z80.memory.WriteByte(wordtemp, bytetemp)
// }

// /* DEC (ix+dd) */
// func instrDD__DEC_iREGpDD(z80 *Z80) {
// 	var offset byte = z80.memory.ReadByte(z80.PC())
// 	z80.memory.ContendReadNoMreq_loop(z80.PC(), 1, 5)
// 	z80.IncPC(1)
// 	var wordtemp uint16 = z80.IX() + uint16(signExtend(offset))
// 	var bytetemp byte = z80.memory.ReadByte(wordtemp)
// 	z80.memory.ContendReadNoMreq(wordtemp, 1)
// 	z80.dec(&bytetemp)
// 	z80.memory.WriteByte(wordtemp, bytetemp)
// }

// /* LD (ix+dd),nn */
// func instrDD__LD_iREGpDD_NN(z80 *Z80) {
// 	offset := z80.memory.ReadByte(z80.PC())
// 	z80.IncPC(1)
// 	value := z80.memory.ReadByte(z80.PC())
// 	z80.memory.ContendReadNoMreq_loop(z80.PC(), 1, 2)
// 	z80.IncPC(1)
// 	z80.memory.WriteByte(z80.IX()+uint16(signExtend(offset)), value)
// }

// /* ADD ix,SP */
// func instrDD__ADD_REG_SP(z80 *Z80) {
// 	z80.memory.ContendReadNoMreq_loop(z80.IR(), 1, 7)
// 	z80.add16(z80.ix, z80.SP())
// }

// /* LD B,IXH */
// func instrDD__LD_B_REGH(z80 *Z80) {
// 	z80.B = z80.IXH
// }

// /* LD B,IXL */
// func instrDD__LD_B_REGL(z80 *Z80) {
// 	z80.B = z80.IXL
// }

// /* LD B,(ix+dd) */
// func instrDD__LD_B_iREGpDD(z80 *Z80) {
// 	var offset byte = z80.memory.ReadByte(z80.PC())
// 	z80.memory.ContendReadNoMreq_loop(z80.PC(), 1, 5)
// 	z80.IncPC(1)
// 	z80.B = z80.memory.ReadByte(z80.IX() + uint16(signExtend(offset)))
// }

// /* LD C,IXH */
// func instrDD__LD_C_REGH(z80 *Z80) {
// 	z80.C = z80.IXH
// }

// /* LD C,IXL */
// func instrDD__LD_C_REGL(z80 *Z80) {
// 	z80.C = z80.IXL
// }

// /* LD C,(ix+dd) */
// func instrDD__LD_C_iREGpDD(z80 *Z80) {
// 	var offset byte = z80.memory.ReadByte(z80.PC())
// 	z80.memory.ContendReadNoMreq_loop(z80.PC(), 1, 5)
// 	z80.IncPC(1)
// 	z80.C = z80.memory.ReadByte(z80.IX() + uint16(signExtend(offset)))
// }

// /* LD D,IXH */
// func instrDD__LD_D_REGH(z80 *Z80) {
// 	z80.D = z80.IXH
// }

// /* LD D,IXL */
// func instrDD__LD_D_REGL(z80 *Z80) {
// 	z80.D = z80.IXL
// }

// /* LD D,(ix+dd) */
// func instrDD__LD_D_iREGpDD(z80 *Z80) {
// 	var offset byte = z80.memory.ReadByte(z80.PC())
// 	z80.memory.ContendReadNoMreq_loop(z80.PC(), 1, 5)
// 	z80.IncPC(1)
// 	z80.D = z80.memory.ReadByte(z80.IX() + uint16(signExtend(offset)))
// }

// /* LD E,IXH */
// func instrDD__LD_E_REGH(z80 *Z80) {
// 	z80.E = z80.IXH
// }

// /* LD E,IXL */
// func instrDD__LD_E_REGL(z80 *Z80) {
// 	z80.E = z80.IXL
// }

// /* LD E,(ix+dd) */
// func instrDD__LD_E_iREGpDD(z80 *Z80) {
// 	var offset byte = z80.memory.ReadByte(z80.PC())
// 	z80.memory.ContendReadNoMreq_loop(z80.PC(), 1, 5)
// 	z80.IncPC(1)
// 	z80.E = z80.memory.ReadByte(z80.IX() + uint16(signExtend(offset)))
// }

// /* LD IXH,B */
// func instrDD__LD_REGH_B(z80 *Z80) {
// 	z80.IXH = z80.B
// }

// /* LD IXH,C */
// func instrDD__LD_REGH_C(z80 *Z80) {
// 	z80.IXH = z80.C
// }

// /* LD IXH,D */
// func instrDD__LD_REGH_D(z80 *Z80) {
// 	z80.IXH = z80.D
// }

// /* LD IXH,E */
// func instrDD__LD_REGH_E(z80 *Z80) {
// 	z80.IXH = z80.E
// }

// /* LD IXH,IXH */
// func instrDD__LD_REGH_REGH(z80 *Z80) {
// }

// /* LD IXH,IXL */
// func instrDD__LD_REGH_REGL(z80 *Z80) {
// 	z80.IXH = z80.IXL
// }

// /* LD H,(ix+dd) */
// func instrDD__LD_H_iREGpDD(z80 *Z80) {
// 	var offset byte = z80.memory.ReadByte(z80.PC())
// 	z80.memory.ContendReadNoMreq_loop(z80.PC(), 1, 5)
// 	z80.IncPC(1)
// 	z80.H = z80.memory.ReadByte(z80.IX() + uint16(signExtend(offset)))
// }

// /* LD IXH,A */
// func instrDD__LD_REGH_A(z80 *Z80) {
// 	z80.IXH = z80.A
// }

// /* LD IXL,B */
// func instrDD__LD_REGL_B(z80 *Z80) {
// 	z80.IXL = z80.B
// }

// /* LD IXL,C */
// func instrDD__LD_REGL_C(z80 *Z80) {
// 	z80.IXL = z80.C
// }

// /* LD IXL,D */
// func instrDD__LD_REGL_D(z80 *Z80) {
// 	z80.IXL = z80.D
// }

// /* LD IXL,E */
// func instrDD__LD_REGL_E(z80 *Z80) {
// 	z80.IXL = z80.E
// }

// /* LD IXL,IXH */
// func instrDD__LD_REGL_REGH(z80 *Z80) {
// 	z80.IXL = z80.IXH
// }

// /* LD IXL,IXL */
// func instrDD__LD_REGL_REGL(z80 *Z80) {
// }

// /* LD L,(ix+dd) */
// func instrDD__LD_L_iREGpDD(z80 *Z80) {
// 	var offset byte = z80.memory.ReadByte(z80.PC())
// 	z80.memory.ContendReadNoMreq_loop(z80.PC(), 1, 5)
// 	z80.IncPC(1)
// 	z80.L = z80.memory.ReadByte(z80.IX() + uint16(signExtend(offset)))
// }

// /* LD IXL,A */
// func instrDD__LD_REGL_A(z80 *Z80) {
// 	z80.IXL = z80.A
// }

// /* LD (ix+dd),B */
// func instrDD__LD_iREGpDD_B(z80 *Z80) {
// 	offset := z80.memory.ReadByte(z80.PC())
// 	z80.memory.ContendReadNoMreq_loop(z80.PC(), 1, 5)
// 	z80.IncPC(1)
// 	z80.memory.WriteByte(z80.IX()+uint16(signExtend(offset)), z80.B)
// }

// /* LD (ix+dd),C */
// func instrDD__LD_iREGpDD_C(z80 *Z80) {
// 	offset := z80.memory.ReadByte(z80.PC())
// 	z80.memory.ContendReadNoMreq_loop(z80.PC(), 1, 5)
// 	z80.IncPC(1)
// 	z80.memory.WriteByte(z80.IX()+uint16(signExtend(offset)), z80.C)
// }

// /* LD (ix+dd),D */
// func instrDD__LD_iREGpDD_D(z80 *Z80) {
// 	offset := z80.memory.ReadByte(z80.PC())
// 	z80.memory.ContendReadNoMreq_loop(z80.PC(), 1, 5)
// 	z80.IncPC(1)
// 	z80.memory.WriteByte(z80.IX()+uint16(signExtend(offset)), z80.D)
// }

// /* LD (ix+dd),E */
// func instrDD__LD_iREGpDD_E(z80 *Z80) {
// 	offset := z80.memory.ReadByte(z80.PC())
// 	z80.memory.ContendReadNoMreq_loop(z80.PC(), 1, 5)
// 	z80.IncPC(1)
// 	z80.memory.WriteByte(z80.IX()+uint16(signExtend(offset)), z80.E)
// }

// /* LD (ix+dd),H */
// func instrDD__LD_iREGpDD_H(z80 *Z80) {
// 	offset := z80.memory.ReadByte(z80.PC())
// 	z80.memory.ContendReadNoMreq_loop(z80.PC(), 1, 5)
// 	z80.IncPC(1)
// 	z80.memory.WriteByte(z80.IX()+uint16(signExtend(offset)), z80.H)
// }

// /* LD (ix+dd),L */
// func instrDD__LD_iREGpDD_L(z80 *Z80) {
// 	offset := z80.memory.ReadByte(z80.PC())
// 	z80.memory.ContendReadNoMreq_loop(z80.PC(), 1, 5)
// 	z80.IncPC(1)
// 	z80.memory.WriteByte(z80.IX()+uint16(signExtend(offset)), z80.L)
// }

// /* LD (ix+dd),A */
// func instrDD__LD_iREGpDD_A(z80 *Z80) {
// 	offset := z80.memory.ReadByte(z80.PC())
// 	z80.memory.ContendReadNoMreq_loop(z80.PC(), 1, 5)
// 	z80.IncPC(1)
// 	z80.memory.WriteByte(z80.IX()+uint16(signExtend(offset)), z80.A)
// }

// /* LD A,IXH */
// func instrDD__LD_A_REGH(z80 *Z80) {
// 	z80.A = z80.IXH
// }

// /* LD A,IXL */
// func instrDD__LD_A_REGL(z80 *Z80) {
// 	z80.A = z80.IXL
// }

// /* LD A,(ix+dd) */
// func instrDD__LD_A_iREGpDD(z80 *Z80) {
// 	var offset byte = z80.memory.ReadByte(z80.PC())
// 	z80.memory.ContendReadNoMreq_loop(z80.PC(), 1, 5)
// 	z80.IncPC(1)
// 	z80.A = z80.memory.ReadByte(z80.IX() + uint16(signExtend(offset)))
// }

// /* ADD A,IXH */
// func instrDD__ADD_A_REGH(z80 *Z80) {
// 	z80.add(z80.IXH)
// }

// /* ADD A,IXL */
// func instrDD__ADD_A_REGL(z80 *Z80) {
// 	z80.add(z80.IXL)
// }

// /* ADD A,(ix+dd) */
// func instrDD__ADD_A_iREGpDD(z80 *Z80) {
// 	var offset byte = z80.memory.ReadByte(z80.PC())
// 	z80.memory.ContendReadNoMreq_loop(z80.PC(), 1, 5)
// 	z80.IncPC(1)
// 	var bytetemp byte = z80.memory.ReadByte(z80.IX() + uint16(signExtend(offset)))
// 	z80.add(bytetemp)
// }

// /* ADC A,IXH */
// func instrDD__ADC_A_REGH(z80 *Z80) {
// 	z80.adc(z80.IXH)
// }

// /* ADC A,IXL */
// func instrDD__ADC_A_REGL(z80 *Z80) {
// 	z80.adc(z80.IXL)
// }

// /* ADC A,(ix+dd) */
// func instrDD__ADC_A_iREGpDD(z80 *Z80) {
// 	var offset byte = z80.memory.ReadByte(z80.PC())
// 	z80.memory.ContendReadNoMreq_loop(z80.PC(), 1, 5)
// 	z80.IncPC(1)
// 	var bytetemp byte = z80.memory.ReadByte(z80.IX() + uint16(signExtend(offset)))
// 	z80.adc(bytetemp)
// }

// /* SUB A,IXH */
// func instrDD__SUB_A_REGH(z80 *Z80) {
// 	z80.sub(z80.IXH)
// }

// /* SUB A,IXL */
// func instrDD__SUB_A_REGL(z80 *Z80) {
// 	z80.sub(z80.IXL)
// }

// /* SUB A,(ix+dd) */
// func instrDD__SUB_A_iREGpDD(z80 *Z80) {
// 	var offset byte = z80.memory.ReadByte(z80.PC())
// 	z80.memory.ContendReadNoMreq_loop(z80.PC(), 1, 5)
// 	z80.IncPC(1)
// 	var bytetemp byte = z80.memory.ReadByte(z80.IX() + uint16(signExtend(offset)))
// 	z80.sub(bytetemp)
// }

// /* SBC A,IXH */
// func instrDD__SBC_A_REGH(z80 *Z80) {
// 	z80.sbc(z80.IXH)
// }

// /* SBC A,IXL */
// func instrDD__SBC_A_REGL(z80 *Z80) {
// 	z80.sbc(z80.IXL)
// }

// /* SBC A,(ix+dd) */
// func instrDD__SBC_A_iREGpDD(z80 *Z80) {
// 	var offset byte = z80.memory.ReadByte(z80.PC())
// 	z80.memory.ContendReadNoMreq_loop(z80.PC(), 1, 5)
// 	z80.IncPC(1)
// 	var bytetemp byte = z80.memory.ReadByte(z80.IX() + uint16(signExtend(offset)))
// 	z80.sbc(bytetemp)
// }

// /* AND A,IXH */
// func instrDD__AND_A_REGH(z80 *Z80) {
// 	z80.and(z80.IXH)
// }

// /* AND A,IXL */
// func instrDD__AND_A_REGL(z80 *Z80) {
// 	z80.and(z80.IXL)
// }

// /* AND A,(ix+dd) */
// func instrDD__AND_A_iREGpDD(z80 *Z80) {
// 	var offset byte = z80.memory.ReadByte(z80.PC())
// 	z80.memory.ContendReadNoMreq_loop(z80.PC(), 1, 5)
// 	z80.IncPC(1)
// 	var bytetemp byte = z80.memory.ReadByte(z80.IX() + uint16(signExtend(offset)))
// 	z80.and(bytetemp)
// }

// /* XOR A,IXH */
// func instrDD__XOR_A_REGH(z80 *Z80) {
// 	z80.xor(z80.IXH)
// }

// /* XOR A,IXL */
// func instrDD__XOR_A_REGL(z80 *Z80) {
// 	z80.xor(z80.IXL)
// }

// /* XOR A,(ix+dd) */
// func instrDD__XOR_A_iREGpDD(z80 *Z80) {
// 	var offset byte = z80.memory.ReadByte(z80.PC())
// 	z80.memory.ContendReadNoMreq_loop(z80.PC(), 1, 5)
// 	z80.IncPC(1)
// 	var bytetemp byte = z80.memory.ReadByte(z80.IX() + uint16(signExtend(offset)))
// 	z80.xor(bytetemp)
// }

// /* OR A,IXH */
// func instrDD__OR_A_REGH(z80 *Z80) {
// 	z80.or(z80.IXH)
// }

// /* OR A,IXL */
// func instrDD__OR_A_REGL(z80 *Z80) {
// 	z80.or(z80.IXL)
// }

// /* OR A,(ix+dd) */
// func instrDD__OR_A_iREGpDD(z80 *Z80) {
// 	var offset byte = z80.memory.ReadByte(z80.PC())
// 	z80.memory.ContendReadNoMreq_loop(z80.PC(), 1, 5)
// 	z80.IncPC(1)
// 	var bytetemp byte = z80.memory.ReadByte(z80.IX() + uint16(signExtend(offset)))
// 	z80.or(bytetemp)
// }

// /* CP A,IXH */
// func instrDD__CP_A_REGH(z80 *Z80) {
// 	z80.cp(z80.IXH)
// }

// /* CP A,IXL */
// func instrDD__CP_A_REGL(z80 *Z80) {
// 	z80.cp(z80.IXL)
// }

// /* CP A,(ix+dd) */
// func instrDD__CP_A_iREGpDD(z80 *Z80) {
// 	var offset byte = z80.memory.ReadByte(z80.PC())
// 	z80.memory.ContendReadNoMreq_loop(z80.PC(), 1, 5)
// 	z80.IncPC(1)
// 	var bytetemp byte = z80.memory.ReadByte(z80.IX() + uint16(signExtend(offset)))
// 	z80.cp(bytetemp)
// }

// /* shift DDFDCB */
func instrDD__SHIFT_DDFDCB(z *Z80, opcode byte) {
	OpcodeDDCBMap[opcode](z, opcode)
}

// /* POP ix */
// func instrDD__POP_REG(z80 *Z80) {
// 	z80.IXL, z80.IXH = z80.pop16()
// }

// /* EX (SP),ix */
// func instrDD__EX_iSP_REG(z80 *Z80) {
// 	var bytetempl = z80.memory.ReadByte(z80.SP())
// 	var bytetemph = z80.memory.ReadByte(z80.SP() + 1)
// 	z80.memory.ContendReadNoMreq(z80.SP()+1, 1)
// 	z80.memory.WriteByte(z80.SP()+1, z80.IXH)
// 	z80.memory.WriteByte(z80.SP(), z80.IXL)
// 	z80.memory.ContendWriteNoMreq_loop(z80.SP(), 1, 2)
// 	z80.IXL = bytetempl
// 	z80.IXH = bytetemph
// }

// /* PUSH ix */
// func instrDD__PUSH_REG(z80 *Z80) {
// 	z80.memory.ContendReadNoMreq(z80.IR(), 1)
// 	z80.push16(z80.IXL, z80.IXH)
// }

// /* JP ix */
// func instrDD__JP_REG(z80 *Z80) {
// 	z80.SetPC(z80.IX()) /* NB: NOT INDIRECT! */
// }

// /* LD SP,ix */
// func instrDD__LD_SP_REG(z80 *Z80) {
// 	z80.memory.ContendReadNoMreq_loop(z80.IR(), 1, 2)
// 	z80.SetSP(z80.IX())
// }
