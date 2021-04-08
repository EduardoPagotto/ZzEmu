package ZzEmu

func initOpCodesFD() {

	// 	// BEGIN of 0xfd shifted opcodes
	// 	/* ADD REGISTER,BC */
	// 	OpcodesMap[SHIFT_0xFD+0x09] = instrFD__ADD_REG_BC
	// 	/* ADD REGISTER,DE */
	// 	OpcodesMap[SHIFT_0xFD+0x19] = instrFD__ADD_REG_DE
	// 	/* LD REGISTER,nnnn */
	// 	OpcodesMap[SHIFT_0xFD+0x21] = instrFD__LD_REG_NNNN
	// 	/* LD (nnnn),REGISTER */
	// 	OpcodesMap[SHIFT_0xFD+0x22] = instrFD__LD_iNNNN_REG
	// 	/* INC REGISTER */
	// 	OpcodesMap[SHIFT_0xFD+0x23] = instrFD__INC_REG
	// 	/* INC REGISTERH */
	// 	OpcodesMap[SHIFT_0xFD+0x24] = instrFD__INC_REGH
	// 	/* DEC REGISTERH */
	// 	OpcodesMap[SHIFT_0xFD+0x25] = instrFD__DEC_REGH
	// 	/* LD REGISTERH,nn */
	// 	OpcodesMap[SHIFT_0xFD+0x26] = instrFD__LD_REGH_NN
	// 	/* ADD REGISTER,REGISTER */
	// 	OpcodesMap[SHIFT_0xFD+0x29] = instrFD__ADD_REG_REG
	// 	/* LD REGISTER,(nnnn) */
	// 	OpcodesMap[SHIFT_0xFD+0x2a] = instrFD__LD_REG_iNNNN
	// 	/* DEC REGISTER */
	// 	OpcodesMap[SHIFT_0xFD+0x2b] = instrFD__DEC_REG
	// 	/* INC REGISTERL */
	// 	OpcodesMap[SHIFT_0xFD+0x2c] = instrFD__INC_REGL
	// 	/* DEC REGISTERL */
	// 	OpcodesMap[SHIFT_0xFD+0x2d] = instrFD__DEC_REGL
	// 	/* LD REGISTERL,nn */
	// 	OpcodesMap[SHIFT_0xFD+0x2e] = instrFD__LD_REGL_NN
	// 	/* INC (REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xFD+0x34] = instrFD__INC_iREGpDD
	// 	/* DEC (REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xFD+0x35] = instrFD__DEC_iREGpDD
	// 	/* LD (REGISTER+dd),nn */
	// 	OpcodesMap[SHIFT_0xFD+0x36] = instrFD__LD_iREGpDD_NN
	// 	/* ADD REGISTER,SP */
	// 	OpcodesMap[SHIFT_0xFD+0x39] = instrFD__ADD_REG_SP
	// 	/* LD B,REGISTERH */
	// 	OpcodesMap[SHIFT_0xFD+0x44] = instrFD__LD_B_REGH
	// 	/* LD B,REGISTERL */
	// 	OpcodesMap[SHIFT_0xFD+0x45] = instrFD__LD_B_REGL
	// 	/* LD B,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xFD+0x46] = instrFD__LD_B_iREGpDD
	// 	/* LD C,REGISTERH */
	// 	OpcodesMap[SHIFT_0xFD+0x4c] = instrFD__LD_C_REGH
	// 	/* LD C,REGISTERL */
	// 	OpcodesMap[SHIFT_0xFD+0x4d] = instrFD__LD_C_REGL
	// 	/* LD C,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xFD+0x4e] = instrFD__LD_C_iREGpDD
	// 	/* LD D,REGISTERH */
	// 	OpcodesMap[SHIFT_0xFD+0x54] = instrFD__LD_D_REGH
	// 	/* LD D,REGISTERL */
	// 	OpcodesMap[SHIFT_0xFD+0x55] = instrFD__LD_D_REGL
	// 	/* LD D,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xFD+0x56] = instrFD__LD_D_iREGpDD
	// 	/* LD E,REGISTERH */
	// 	OpcodesMap[SHIFT_0xFD+0x5c] = instrFD__LD_E_REGH
	// 	/* LD E,REGISTERL */
	// 	OpcodesMap[SHIFT_0xFD+0x5d] = instrFD__LD_E_REGL
	// 	/* LD E,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xFD+0x5e] = instrFD__LD_E_iREGpDD
	// 	/* LD REGISTERH,B */
	// 	OpcodesMap[SHIFT_0xFD+0x60] = instrFD__LD_REGH_B
	// 	/* LD REGISTERH,C */
	// 	OpcodesMap[SHIFT_0xFD+0x61] = instrFD__LD_REGH_C
	// 	/* LD REGISTERH,D */
	// 	OpcodesMap[SHIFT_0xFD+0x62] = instrFD__LD_REGH_D
	// 	/* LD REGISTERH,E */
	// 	OpcodesMap[SHIFT_0xFD+0x63] = instrFD__LD_REGH_E
	// 	/* LD REGISTERH,REGISTERH */
	// 	OpcodesMap[SHIFT_0xFD+0x64] = instrFD__LD_REGH_REGH
	// 	/* LD REGISTERH,REGISTERL */
	// 	OpcodesMap[SHIFT_0xFD+0x65] = instrFD__LD_REGH_REGL
	// 	/* LD H,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xFD+0x66] = instrFD__LD_H_iREGpDD
	// 	/* LD REGISTERH,A */
	// 	OpcodesMap[SHIFT_0xFD+0x67] = instrFD__LD_REGH_A
	// 	/* LD REGISTERL,B */
	// 	OpcodesMap[SHIFT_0xFD+0x68] = instrFD__LD_REGL_B
	// 	/* LD REGISTERL,C */
	// 	OpcodesMap[SHIFT_0xFD+0x69] = instrFD__LD_REGL_C
	// 	/* LD REGISTERL,D */
	// 	OpcodesMap[SHIFT_0xFD+0x6a] = instrFD__LD_REGL_D
	// 	/* LD REGISTERL,E */
	// 	OpcodesMap[SHIFT_0xFD+0x6b] = instrFD__LD_REGL_E
	// 	/* LD REGISTERL,REGISTERH */
	// 	OpcodesMap[SHIFT_0xFD+0x6c] = instrFD__LD_REGL_REGH
	// 	/* LD REGISTERL,REGISTERL */
	// 	OpcodesMap[SHIFT_0xFD+0x6d] = instrFD__LD_REGL_REGL
	// 	/* LD L,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xFD+0x6e] = instrFD__LD_L_iREGpDD
	// 	/* LD REGISTERL,A */
	// 	OpcodesMap[SHIFT_0xFD+0x6f] = instrFD__LD_REGL_A
	// 	/* LD (REGISTER+dd),B */
	// 	OpcodesMap[SHIFT_0xFD+0x70] = instrFD__LD_iREGpDD_B
	// 	/* LD (REGISTER+dd),C */
	// 	OpcodesMap[SHIFT_0xFD+0x71] = instrFD__LD_iREGpDD_C
	// 	/* LD (REGISTER+dd),D */
	// 	OpcodesMap[SHIFT_0xFD+0x72] = instrFD__LD_iREGpDD_D
	// 	/* LD (REGISTER+dd),E */
	// 	OpcodesMap[SHIFT_0xFD+0x73] = instrFD__LD_iREGpDD_E
	// 	/* LD (REGISTER+dd),H */
	// 	OpcodesMap[SHIFT_0xFD+0x74] = instrFD__LD_iREGpDD_H
	// 	/* LD (REGISTER+dd),L */
	// 	OpcodesMap[SHIFT_0xFD+0x75] = instrFD__LD_iREGpDD_L
	// 	/* LD (REGISTER+dd),A */
	// 	OpcodesMap[SHIFT_0xFD+0x77] = instrFD__LD_iREGpDD_A
	// 	/* LD A,REGISTERH */
	// 	OpcodesMap[SHIFT_0xFD+0x7c] = instrFD__LD_A_REGH
	// 	/* LD A,REGISTERL */
	// 	OpcodesMap[SHIFT_0xFD+0x7d] = instrFD__LD_A_REGL
	// 	/* LD A,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xFD+0x7e] = instrFD__LD_A_iREGpDD
	// 	/* ADD A,REGISTERH */
	// 	OpcodesMap[SHIFT_0xFD+0x84] = instrFD__ADD_A_REGH
	// 	/* ADD A,REGISTERL */
	// 	OpcodesMap[SHIFT_0xFD+0x85] = instrFD__ADD_A_REGL
	// 	/* ADD A,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xFD+0x86] = instrFD__ADD_A_iREGpDD
	// 	/* ADC A,REGISTERH */
	// 	OpcodesMap[SHIFT_0xFD+0x8c] = instrFD__ADC_A_REGH
	// 	/* ADC A,REGISTERL */
	// 	OpcodesMap[SHIFT_0xFD+0x8d] = instrFD__ADC_A_REGL
	// 	/* ADC A,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xFD+0x8e] = instrFD__ADC_A_iREGpDD
	// 	/* SUB A,REGISTERH */
	// 	OpcodesMap[SHIFT_0xFD+0x94] = instrFD__SUB_A_REGH
	// 	/* SUB A,REGISTERL */
	// 	OpcodesMap[SHIFT_0xFD+0x95] = instrFD__SUB_A_REGL
	// 	/* SUB A,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xFD+0x96] = instrFD__SUB_A_iREGpDD
	// 	/* SBC A,REGISTERH */
	// 	OpcodesMap[SHIFT_0xFD+0x9c] = instrFD__SBC_A_REGH
	// 	/* SBC A,REGISTERL */
	// 	OpcodesMap[SHIFT_0xFD+0x9d] = instrFD__SBC_A_REGL
	// 	/* SBC A,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xFD+0x9e] = instrFD__SBC_A_iREGpDD
	// 	/* AND A,REGISTERH */
	// 	OpcodesMap[SHIFT_0xFD+0xa4] = instrFD__AND_A_REGH
	// 	/* AND A,REGISTERL */
	// 	OpcodesMap[SHIFT_0xFD+0xa5] = instrFD__AND_A_REGL
	// 	/* AND A,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xFD+0xa6] = instrFD__AND_A_iREGpDD
	// 	/* XOR A,REGISTERH */
	// 	OpcodesMap[SHIFT_0xFD+0xac] = instrFD__XOR_A_REGH
	// 	/* XOR A,REGISTERL */
	// 	OpcodesMap[SHIFT_0xFD+0xad] = instrFD__XOR_A_REGL
	// 	/* XOR A,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xFD+0xae] = instrFD__XOR_A_iREGpDD
	// 	/* OR A,REGISTERH */
	// 	OpcodesMap[SHIFT_0xFD+0xb4] = instrFD__OR_A_REGH
	// 	/* OR A,REGISTERL */
	// 	OpcodesMap[SHIFT_0xFD+0xb5] = instrFD__OR_A_REGL
	// 	/* OR A,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xFD+0xb6] = instrFD__OR_A_iREGpDD
	// 	/* CP A,REGISTERH */
	// 	OpcodesMap[SHIFT_0xFD+0xbc] = instrFD__CP_A_REGH
	// 	/* CP A,REGISTERL */
	// 	OpcodesMap[SHIFT_0xFD+0xbd] = instrFD__CP_A_REGL
	// 	/* CP A,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xFD+0xbe] = instrFD__CP_A_iREGpDD
	// 	/* shift DDFDCB */
	// 	OpcodesMap[SHIFT_0xFD+0xcb] = instrFD__SHIFT_DDFDCB
	// 	/* POP REGISTER */
	// 	OpcodesMap[SHIFT_0xFD+0xe1] = instrFD__POP_REG
	// 	/* EX (SP),REGISTER */
	// 	OpcodesMap[SHIFT_0xFD+0xe3] = instrFD__EX_iSP_REG
	// 	/* PUSH REGISTER */
	// 	OpcodesMap[SHIFT_0xFD+0xe5] = instrFD__PUSH_REG
	// 	/* JP REGISTER */
	// 	OpcodesMap[SHIFT_0xFD+0xe9] = instrFD__JP_REG
	// 	/* LD SP,REGISTER */
	// 	OpcodesMap[SHIFT_0xFD+0xf9] = instrFD__LD_SP_REG

	// 	// END of 0xfd shifted opcodes
}

// /* ADD iy,BC */
// func instrFD__ADD_REG_BC(z80 *Z80) {
// 	z80.memory.ContendReadNoMreq_loop(z80.IR(), 1, 7)
// 	z80.add16(z80.iy, z80.BC())
// }

// /* ADD iy,DE */
// func instrFD__ADD_REG_DE(z80 *Z80) {
// 	z80.memory.ContendReadNoMreq_loop(z80.IR(), 1, 7)
// 	z80.add16(z80.iy, z80.DE())
// }

// /* LD iy,nnnn */
// func instrFD__LD_REG_NNNN(z80 *Z80) {
// 	b1 := z80.memory.ReadByte(z80.PC())
// 	z80.IncPC(1)
// 	b2 := z80.memory.ReadByte(z80.PC())
// 	z80.IncPC(1)
// 	z80.SetIY(joinBytes(b2, b1))
// }

// /* LD (nnnn),iy */
// func instrFD__LD_iNNNN_REG(z80 *Z80) {
// 	z80.ld16nnrr(z80.IYL, z80.IYH)
// 	// break
// }

// /* INC iy */
// func instrFD__INC_REG(z80 *Z80) {
// 	z80.memory.ContendReadNoMreq_loop(z80.IR(), 1, 2)
// 	z80.IncIY()
// }

// /* INC IYH */
// func instrFD__INC_REGH(z80 *Z80) {
// 	z80.incIYH()
// }

// /* DEC IYH */
// func instrFD__DEC_REGH(z80 *Z80) {
// 	z80.decIYH()
// }

// /* LD IYH,nn */
// func instrFD__LD_REGH_NN(z80 *Z80) {
// 	z80.IYH = z80.memory.ReadByte(z80.PC())
// 	z80.IncPC(1)
// }

// /* ADD iy,iy */
// func instrFD__ADD_REG_REG(z80 *Z80) {
// 	z80.memory.ContendReadNoMreq_loop(z80.IR(), 1, 7)
// 	z80.add16(z80.iy, z80.IY())
// }

// /* LD iy,(nnnn) */
// func instrFD__LD_REG_iNNNN(z80 *Z80) {
// 	z80.ld16rrnn(&z80.IYL, &z80.IYH)
// 	// break
// }

// /* DEC iy */
// func instrFD__DEC_REG(z80 *Z80) {
// 	z80.memory.ContendReadNoMreq_loop(z80.IR(), 1, 2)
// 	z80.DecIY()
// }

// /* INC IYL */
// func instrFD__INC_REGL(z80 *Z80) {
// 	z80.incIYL()
// }

// /* DEC IYL */
// func instrFD__DEC_REGL(z80 *Z80) {
// 	z80.decIYL()
// }

// /* LD IYL,nn */
// func instrFD__LD_REGL_NN(z80 *Z80) {
// 	z80.IYL = z80.memory.ReadByte(z80.PC())
// 	z80.IncPC(1)
// }

// /* INC (iy+dd) */
// func instrFD__INC_iREGpDD(z80 *Z80) {
// 	var offset byte = z80.memory.ReadByte(z80.PC())
// 	z80.memory.ContendReadNoMreq_loop(z80.PC(), 1, 5)
// 	z80.IncPC(1)
// 	var wordtemp uint16 = z80.IY() + uint16(signExtend(offset))
// 	var bytetemp byte = z80.memory.ReadByte(wordtemp)
// 	z80.memory.ContendReadNoMreq(wordtemp, 1)
// 	z80.inc(&bytetemp)
// 	z80.memory.WriteByte(wordtemp, bytetemp)
// }

// /* DEC (iy+dd) */
// func instrFD__DEC_iREGpDD(z80 *Z80) {
// 	var offset byte = z80.memory.ReadByte(z80.PC())
// 	z80.memory.ContendReadNoMreq_loop(z80.PC(), 1, 5)
// 	z80.IncPC(1)
// 	var wordtemp uint16 = z80.IY() + uint16(signExtend(offset))
// 	var bytetemp byte = z80.memory.ReadByte(wordtemp)
// 	z80.memory.ContendReadNoMreq(wordtemp, 1)
// 	z80.dec(&bytetemp)
// 	z80.memory.WriteByte(wordtemp, bytetemp)
// }

// /* LD (iy+dd),nn */
// func instrFD__LD_iREGpDD_NN(z80 *Z80) {
// 	offset := z80.memory.ReadByte(z80.PC())
// 	z80.IncPC(1)
// 	value := z80.memory.ReadByte(z80.PC())
// 	z80.memory.ContendReadNoMreq_loop(z80.PC(), 1, 2)
// 	z80.IncPC(1)
// 	z80.memory.WriteByte(z80.IY()+uint16(signExtend(offset)), value)
// }

// /* ADD iy,SP */
// func instrFD__ADD_REG_SP(z80 *Z80) {
// 	z80.memory.ContendReadNoMreq_loop(z80.IR(), 1, 7)
// 	z80.add16(z80.iy, z80.SP())
// }

// /* LD B,IYH */
// func instrFD__LD_B_REGH(z80 *Z80) {
// 	z80.B = z80.IYH
// }

// /* LD B,IYL */
// func instrFD__LD_B_REGL(z80 *Z80) {
// 	z80.B = z80.IYL
// }

// /* LD B,(iy+dd) */
// func instrFD__LD_B_iREGpDD(z80 *Z80) {
// 	var offset byte = z80.memory.ReadByte(z80.PC())
// 	z80.memory.ContendReadNoMreq_loop(z80.PC(), 1, 5)
// 	z80.IncPC(1)
// 	z80.B = z80.memory.ReadByte(z80.IY() + uint16(signExtend(offset)))
// }

// /* LD C,IYH */
// func instrFD__LD_C_REGH(z80 *Z80) {
// 	z80.C = z80.IYH
// }

// /* LD C,IYL */
// func instrFD__LD_C_REGL(z80 *Z80) {
// 	z80.C = z80.IYL
// }

// /* LD C,(iy+dd) */
// func instrFD__LD_C_iREGpDD(z80 *Z80) {
// 	var offset byte = z80.memory.ReadByte(z80.PC())
// 	z80.memory.ContendReadNoMreq_loop(z80.PC(), 1, 5)
// 	z80.IncPC(1)
// 	z80.C = z80.memory.ReadByte(z80.IY() + uint16(signExtend(offset)))
// }

// /* LD D,IYH */
// func instrFD__LD_D_REGH(z80 *Z80) {
// 	z80.D = z80.IYH
// }

// /* LD D,IYL */
// func instrFD__LD_D_REGL(z80 *Z80) {
// 	z80.D = z80.IYL
// }

// /* LD D,(iy+dd) */
// func instrFD__LD_D_iREGpDD(z80 *Z80) {
// 	var offset byte = z80.memory.ReadByte(z80.PC())
// 	z80.memory.ContendReadNoMreq_loop(z80.PC(), 1, 5)
// 	z80.IncPC(1)
// 	z80.D = z80.memory.ReadByte(z80.IY() + uint16(signExtend(offset)))
// }

// /* LD E,IYH */
// func instrFD__LD_E_REGH(z80 *Z80) {
// 	z80.E = z80.IYH
// }

// /* LD E,IYL */
// func instrFD__LD_E_REGL(z80 *Z80) {
// 	z80.E = z80.IYL
// }

// /* LD E,(iy+dd) */
// func instrFD__LD_E_iREGpDD(z80 *Z80) {
// 	var offset byte = z80.memory.ReadByte(z80.PC())
// 	z80.memory.ContendReadNoMreq_loop(z80.PC(), 1, 5)
// 	z80.IncPC(1)
// 	z80.E = z80.memory.ReadByte(z80.IY() + uint16(signExtend(offset)))
// }

// /* LD IYH,B */
// func instrFD__LD_REGH_B(z80 *Z80) {
// 	z80.IYH = z80.B
// }

// /* LD IYH,C */
// func instrFD__LD_REGH_C(z80 *Z80) {
// 	z80.IYH = z80.C
// }

// /* LD IYH,D */
// func instrFD__LD_REGH_D(z80 *Z80) {
// 	z80.IYH = z80.D
// }

// /* LD IYH,E */
// func instrFD__LD_REGH_E(z80 *Z80) {
// 	z80.IYH = z80.E
// }

// /* LD IYH,IYH */
// func instrFD__LD_REGH_REGH(z80 *Z80) {
// }

// /* LD IYH,IYL */
// func instrFD__LD_REGH_REGL(z80 *Z80) {
// 	z80.IYH = z80.IYL
// }

// /* LD H,(iy+dd) */
// func instrFD__LD_H_iREGpDD(z80 *Z80) {
// 	var offset byte = z80.memory.ReadByte(z80.PC())
// 	z80.memory.ContendReadNoMreq_loop(z80.PC(), 1, 5)
// 	z80.IncPC(1)
// 	z80.H = z80.memory.ReadByte(z80.IY() + uint16(signExtend(offset)))
// }

// /* LD IYH,A */
// func instrFD__LD_REGH_A(z80 *Z80) {
// 	z80.IYH = z80.A
// }

// /* LD IYL,B */
// func instrFD__LD_REGL_B(z80 *Z80) {
// 	z80.IYL = z80.B
// }

// /* LD IYL,C */
// func instrFD__LD_REGL_C(z80 *Z80) {
// 	z80.IYL = z80.C
// }

// /* LD IYL,D */
// func instrFD__LD_REGL_D(z80 *Z80) {
// 	z80.IYL = z80.D
// }

// /* LD IYL,E */
// func instrFD__LD_REGL_E(z80 *Z80) {
// 	z80.IYL = z80.E
// }

// /* LD IYL,IYH */
// func instrFD__LD_REGL_REGH(z80 *Z80) {
// 	z80.IYL = z80.IYH
// }

// /* LD IYL,IYL */
// func instrFD__LD_REGL_REGL(z80 *Z80) {
// }

// /* LD L,(iy+dd) */
// func instrFD__LD_L_iREGpDD(z80 *Z80) {
// 	var offset byte = z80.memory.ReadByte(z80.PC())
// 	z80.memory.ContendReadNoMreq_loop(z80.PC(), 1, 5)
// 	z80.IncPC(1)
// 	z80.L = z80.memory.ReadByte(z80.IY() + uint16(signExtend(offset)))
// }

// /* LD IYL,A */
// func instrFD__LD_REGL_A(z80 *Z80) {
// 	z80.IYL = z80.A
// }

// /* LD (iy+dd),B */
// func instrFD__LD_iREGpDD_B(z80 *Z80) {
// 	offset := z80.memory.ReadByte(z80.PC())
// 	z80.memory.ContendReadNoMreq_loop(z80.PC(), 1, 5)
// 	z80.IncPC(1)
// 	z80.memory.WriteByte(z80.IY()+uint16(signExtend(offset)), z80.B)
// }

// /* LD (iy+dd),C */
// func instrFD__LD_iREGpDD_C(z80 *Z80) {
// 	offset := z80.memory.ReadByte(z80.PC())
// 	z80.memory.ContendReadNoMreq_loop(z80.PC(), 1, 5)
// 	z80.IncPC(1)
// 	z80.memory.WriteByte(z80.IY()+uint16(signExtend(offset)), z80.C)
// }

// /* LD (iy+dd),D */
// func instrFD__LD_iREGpDD_D(z80 *Z80) {
// 	offset := z80.memory.ReadByte(z80.PC())
// 	z80.memory.ContendReadNoMreq_loop(z80.PC(), 1, 5)
// 	z80.IncPC(1)
// 	z80.memory.WriteByte(z80.IY()+uint16(signExtend(offset)), z80.D)
// }

// /* LD (iy+dd),E */
// func instrFD__LD_iREGpDD_E(z80 *Z80) {
// 	offset := z80.memory.ReadByte(z80.PC())
// 	z80.memory.ContendReadNoMreq_loop(z80.PC(), 1, 5)
// 	z80.IncPC(1)
// 	z80.memory.WriteByte(z80.IY()+uint16(signExtend(offset)), z80.E)
// }

// /* LD (iy+dd),H */
// func instrFD__LD_iREGpDD_H(z80 *Z80) {
// 	offset := z80.memory.ReadByte(z80.PC())
// 	z80.memory.ContendReadNoMreq_loop(z80.PC(), 1, 5)
// 	z80.IncPC(1)
// 	z80.memory.WriteByte(z80.IY()+uint16(signExtend(offset)), z80.H)
// }

// /* LD (iy+dd),L */
// func instrFD__LD_iREGpDD_L(z80 *Z80) {
// 	offset := z80.memory.ReadByte(z80.PC())
// 	z80.memory.ContendReadNoMreq_loop(z80.PC(), 1, 5)
// 	z80.IncPC(1)
// 	z80.memory.WriteByte(z80.IY()+uint16(signExtend(offset)), z80.L)
// }

// /* LD (iy+dd),A */
// func instrFD__LD_iREGpDD_A(z80 *Z80) {
// 	offset := z80.memory.ReadByte(z80.PC())
// 	z80.memory.ContendReadNoMreq_loop(z80.PC(), 1, 5)
// 	z80.IncPC(1)
// 	z80.memory.WriteByte(z80.IY()+uint16(signExtend(offset)), z80.A)
// }

// /* LD A,IYH */
// func instrFD__LD_A_REGH(z80 *Z80) {
// 	z80.A = z80.IYH
// }

// /* LD A,IYL */
// func instrFD__LD_A_REGL(z80 *Z80) {
// 	z80.A = z80.IYL
// }

// /* LD A,(iy+dd) */
// func instrFD__LD_A_iREGpDD(z80 *Z80) {
// 	var offset byte = z80.memory.ReadByte(z80.PC())
// 	z80.memory.ContendReadNoMreq_loop(z80.PC(), 1, 5)
// 	z80.IncPC(1)
// 	z80.A = z80.memory.ReadByte(z80.IY() + uint16(signExtend(offset)))
// }

// /* ADD A,IYH */
// func instrFD__ADD_A_REGH(z80 *Z80) {
// 	z80.add(z80.IYH)
// }

// /* ADD A,IYL */
// func instrFD__ADD_A_REGL(z80 *Z80) {
// 	z80.add(z80.IYL)
// }

// /* ADD A,(iy+dd) */
// func instrFD__ADD_A_iREGpDD(z80 *Z80) {
// 	var offset byte = z80.memory.ReadByte(z80.PC())
// 	z80.memory.ContendReadNoMreq_loop(z80.PC(), 1, 5)
// 	z80.IncPC(1)
// 	var bytetemp byte = z80.memory.ReadByte(z80.IY() + uint16(signExtend(offset)))
// 	z80.add(bytetemp)
// }

// /* ADC A,IYH */
// func instrFD__ADC_A_REGH(z80 *Z80) {
// 	z80.adc(z80.IYH)
// }

// /* ADC A,IYL */
// func instrFD__ADC_A_REGL(z80 *Z80) {
// 	z80.adc(z80.IYL)
// }

// /* ADC A,(iy+dd) */
// func instrFD__ADC_A_iREGpDD(z80 *Z80) {
// 	var offset byte = z80.memory.ReadByte(z80.PC())
// 	z80.memory.ContendReadNoMreq_loop(z80.PC(), 1, 5)
// 	z80.IncPC(1)
// 	var bytetemp byte = z80.memory.ReadByte(z80.IY() + uint16(signExtend(offset)))
// 	z80.adc(bytetemp)
// }

// /* SUB A,IYH */
// func instrFD__SUB_A_REGH(z80 *Z80) {
// 	z80.sub(z80.IYH)
// }

// /* SUB A,IYL */
// func instrFD__SUB_A_REGL(z80 *Z80) {
// 	z80.sub(z80.IYL)
// }

// /* SUB A,(iy+dd) */
// func instrFD__SUB_A_iREGpDD(z80 *Z80) {
// 	var offset byte = z80.memory.ReadByte(z80.PC())
// 	z80.memory.ContendReadNoMreq_loop(z80.PC(), 1, 5)
// 	z80.IncPC(1)
// 	var bytetemp byte = z80.memory.ReadByte(z80.IY() + uint16(signExtend(offset)))
// 	z80.sub(bytetemp)
// }

// /* SBC A,IYH */
// func instrFD__SBC_A_REGH(z80 *Z80) {
// 	z80.sbc(z80.IYH)
// }

// /* SBC A,IYL */
// func instrFD__SBC_A_REGL(z80 *Z80) {
// 	z80.sbc(z80.IYL)
// }

// /* SBC A,(iy+dd) */
// func instrFD__SBC_A_iREGpDD(z80 *Z80) {
// 	var offset byte = z80.memory.ReadByte(z80.PC())
// 	z80.memory.ContendReadNoMreq_loop(z80.PC(), 1, 5)
// 	z80.IncPC(1)
// 	var bytetemp byte = z80.memory.ReadByte(z80.IY() + uint16(signExtend(offset)))
// 	z80.sbc(bytetemp)
// }

// /* AND A,IYH */
// func instrFD__AND_A_REGH(z80 *Z80) {
// 	z80.and(z80.IYH)
// }

// /* AND A,IYL */
// func instrFD__AND_A_REGL(z80 *Z80) {
// 	z80.and(z80.IYL)
// }

// /* AND A,(iy+dd) */
// func instrFD__AND_A_iREGpDD(z80 *Z80) {
// 	var offset byte = z80.memory.ReadByte(z80.PC())
// 	z80.memory.ContendReadNoMreq_loop(z80.PC(), 1, 5)
// 	z80.IncPC(1)
// 	var bytetemp byte = z80.memory.ReadByte(z80.IY() + uint16(signExtend(offset)))
// 	z80.and(bytetemp)
// }

// /* XOR A,IYH */
// func instrFD__XOR_A_REGH(z80 *Z80) {
// 	z80.xor(z80.IYH)
// }

// /* XOR A,IYL */
// func instrFD__XOR_A_REGL(z80 *Z80) {
// 	z80.xor(z80.IYL)
// }

// /* XOR A,(iy+dd) */
// func instrFD__XOR_A_iREGpDD(z80 *Z80) {
// 	var offset byte = z80.memory.ReadByte(z80.PC())
// 	z80.memory.ContendReadNoMreq_loop(z80.PC(), 1, 5)
// 	z80.IncPC(1)
// 	var bytetemp byte = z80.memory.ReadByte(z80.IY() + uint16(signExtend(offset)))
// 	z80.xor(bytetemp)
// }

// /* OR A,IYH */
// func instrFD__OR_A_REGH(z80 *Z80) {
// 	z80.or(z80.IYH)
// }

// /* OR A,IYL */
// func instrFD__OR_A_REGL(z80 *Z80) {
// 	z80.or(z80.IYL)
// }

// /* OR A,(iy+dd) */
// func instrFD__OR_A_iREGpDD(z80 *Z80) {
// 	var offset byte = z80.memory.ReadByte(z80.PC())
// 	z80.memory.ContendReadNoMreq_loop(z80.PC(), 1, 5)
// 	z80.IncPC(1)
// 	var bytetemp byte = z80.memory.ReadByte(z80.IY() + uint16(signExtend(offset)))
// 	z80.or(bytetemp)
// }

// /* CP A,IYH */
// func instrFD__CP_A_REGH(z80 *Z80) {
// 	z80.cp(z80.IYH)
// }

// /* CP A,IYL */
// func instrFD__CP_A_REGL(z80 *Z80) {
// 	z80.cp(z80.IYL)
// }

// /* CP A,(iy+dd) */
// func instrFD__CP_A_iREGpDD(z80 *Z80) {
// 	var offset byte = z80.memory.ReadByte(z80.PC())
// 	z80.memory.ContendReadNoMreq_loop(z80.PC(), 1, 5)
// 	z80.IncPC(1)
// 	var bytetemp byte = z80.memory.ReadByte(z80.IY() + uint16(signExtend(offset)))
// 	z80.cp(bytetemp)
// }

// /* shift DDFDCB */
// func instrFD__SHIFT_DDFDCB(z80 *Z80) {
// }

// /* POP iy */
// func instrFD__POP_REG(z80 *Z80) {
// 	z80.IYL, z80.IYH = z80.pop16()
// }

// /* EX (SP),iy */
// func instrFD__EX_iSP_REG(z80 *Z80) {
// 	var bytetempl = z80.memory.ReadByte(z80.SP())
// 	var bytetemph = z80.memory.ReadByte(z80.SP() + 1)
// 	z80.memory.ContendReadNoMreq(z80.SP()+1, 1)
// 	z80.memory.WriteByte(z80.SP()+1, z80.IYH)
// 	z80.memory.WriteByte(z80.SP(), z80.IYL)
// 	z80.memory.ContendWriteNoMreq_loop(z80.SP(), 1, 2)
// 	z80.IYL = bytetempl
// 	z80.IYH = bytetemph
// }

// /* PUSH iy */
// func instrFD__PUSH_REG(z80 *Z80) {
// 	z80.memory.ContendReadNoMreq(z80.IR(), 1)
// 	z80.push16(z80.IYL, z80.IYH)
// }

// /* JP iy */
// func instrFD__JP_REG(z80 *Z80) {
// 	z80.SetPC(z80.IY()) /* NB: NOT INDIRECT! */
// }

// /* LD SP,iy */
// func instrFD__LD_SP_REG(z80 *Z80) {
// 	z80.memory.ContendReadNoMreq_loop(z80.IR(), 1, 2)
// 	z80.SetSP(z80.IY())
// }

//--
