package ZzEmu

func initOpcodeDFMap() {

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
// func instrFD__ADD_REG_BC(z *Z80, opcode byte) {
// 	z.memory.ContendReadNoMreq_loop(z.IR(), 1, 7)
// 	z.add16(z.iy, z.BC())
// }

// /* ADD iy,DE */
// func instrFD__ADD_REG_DE(z *Z80, opcode byte) {
// 	z.memory.ContendReadNoMreq_loop(z.IR(), 1, 7)
// 	z.add16(z.iy, z.DE())
// }

// /* LD iy,nnnn */
// func instrFD__LD_REG_NNNN(z *Z80, opcode byte) {
// 	b1 := z.memory.ReadByte(z.PC())
// 	z.IncPC(1)
// 	b2 := z.memory.ReadByte(z.PC())
// 	z.IncPC(1)
// 	z.SetIY(joinBytes(b2, b1))
// }

// /* LD (nnnn),iy */
// func instrFD__LD_iNNNN_REG(z *Z80, opcode byte) {
// 	z.ld16nnrr(z.IYL, z.IYH)
// 	// break
// }

// /* INC iy */
// func instrFD__INC_REG(z *Z80, opcode byte) {
// 	z.memory.ContendReadNoMreq_loop(z.IR(), 1, 2)
// 	z.IncIY()
// }

// /* INC IYH */
// func instrFD__INC_REGH(z *Z80, opcode byte) {
// 	z.incIYH()
// }

// /* DEC IYH */
// func instrFD__DEC_REGH(z *Z80, opcode byte) {
// 	z.decIYH()
// }

// /* LD IYH,nn */
// func instrFD__LD_REGH_NN(z *Z80, opcode byte) {
// 	z.IYH = z.memory.ReadByte(z.PC())
// 	z.IncPC(1)
// }

// /* ADD iy,iy */
// func instrFD__ADD_REG_REG(z *Z80, opcode byte) {
// 	z.memory.ContendReadNoMreq_loop(z.IR(), 1, 7)
// 	z.add16(z.iy, z.IY())
// }

// /* LD iy,(nnnn) */
// func instrFD__LD_REG_iNNNN(z *Z80, opcode byte) {
// 	z.ld16rrnn(&z.IYL, &z.IYH)
// 	// break
// }

// /* DEC iy */
// func instrFD__DEC_REG(z *Z80, opcode byte) {
// 	z.memory.ContendReadNoMreq_loop(z.IR(), 1, 2)
// 	z.DecIY()
// }

// /* INC IYL */
// func instrFD__INC_REGL(z *Z80, opcode byte) {
// 	z.incIYL()
// }

// /* DEC IYL */
// func instrFD__DEC_REGL(z *Z80, opcode byte) {
// 	z.decIYL()
// }

// /* LD IYL,nn */
// func instrFD__LD_REGL_NN(z *Z80, opcode byte) {
// 	z.IYL = z.memory.ReadByte(z.PC())
// 	z.IncPC(1)
// }

// /* INC (iy+dd) */
// func instrFD__INC_iREGpDD(z *Z80, opcode byte) {
// 	var offset byte = z.memory.ReadByte(z.PC())
// 	z.memory.ContendReadNoMreq_loop(z.PC(), 1, 5)
// 	z.IncPC(1)
// 	var wordtemp uint16 = z.IY() + uint16(signExtend(offset))
// 	var bytetemp byte = z.memory.ReadByte(wordtemp)
// 	z.memory.ContendReadNoMreq(wordtemp, 1)
// 	z.inc(&bytetemp)
// 	z.memory.WriteByte(wordtemp, bytetemp)
// }

// /* DEC (iy+dd) */
// func instrFD__DEC_iREGpDD(z *Z80, opcode byte) {
// 	var offset byte = z.memory.ReadByte(z.PC())
// 	z.memory.ContendReadNoMreq_loop(z.PC(), 1, 5)
// 	z.IncPC(1)
// 	var wordtemp uint16 = z.IY() + uint16(signExtend(offset))
// 	var bytetemp byte = z.memory.ReadByte(wordtemp)
// 	z.memory.ContendReadNoMreq(wordtemp, 1)
// 	z.dec(&bytetemp)
// 	z.memory.WriteByte(wordtemp, bytetemp)
// }

// /* LD (iy+dd),nn */
// func instrFD__LD_iREGpDD_NN(z *Z80, opcode byte) {
// 	offset := z.memory.ReadByte(z.PC())
// 	z.IncPC(1)
// 	value := z.memory.ReadByte(z.PC())
// 	z.memory.ContendReadNoMreq_loop(z.PC(), 1, 2)
// 	z.IncPC(1)
// 	z.memory.WriteByte(z.IY()+uint16(signExtend(offset)), value)
// }

// /* ADD iy,SP */
// func instrFD__ADD_REG_SP(z *Z80, opcode byte) {
// 	z.memory.ContendReadNoMreq_loop(z.IR(), 1, 7)
// 	z.add16(z.iy, z.SP())
// }

// /* LD B,IYH */
// func instrFD__LD_B_REGH(z *Z80, opcode byte) {
// 	z.B = z.IYH
// }

// /* LD B,IYL */
// func instrFD__LD_B_REGL(z *Z80, opcode byte) {
// 	z.B = z.IYL
// }

// /* LD B,(iy+dd) */
// func instrFD__LD_B_iREGpDD(z *Z80, opcode byte) {
// 	var offset byte = z.memory.ReadByte(z.PC())
// 	z.memory.ContendReadNoMreq_loop(z.PC(), 1, 5)
// 	z.IncPC(1)
// 	z.B = z.memory.ReadByte(z.IY() + uint16(signExtend(offset)))
// }

// /* LD C,IYH */
// func instrFD__LD_C_REGH(z *Z80, opcode byte) {
// 	z.C = z.IYH
// }

// /* LD C,IYL */
// func instrFD__LD_C_REGL(z *Z80, opcode byte) {
// 	z.C = z.IYL
// }

// /* LD C,(iy+dd) */
// func instrFD__LD_C_iREGpDD(z *Z80, opcode byte) {
// 	var offset byte = z.memory.ReadByte(z.PC())
// 	z.memory.ContendReadNoMreq_loop(z.PC(), 1, 5)
// 	z.IncPC(1)
// 	z.C = z.memory.ReadByte(z.IY() + uint16(signExtend(offset)))
// }

// /* LD D,IYH */
// func instrFD__LD_D_REGH(z *Z80, opcode byte) {
// 	z.D = z.IYH
// }

// /* LD D,IYL */
// func instrFD__LD_D_REGL(z *Z80, opcode byte) {
// 	z.D = z.IYL
// }

// /* LD D,(iy+dd) */
// func instrFD__LD_D_iREGpDD(z *Z80, opcode byte) {
// 	var offset byte = z.memory.ReadByte(z.PC())
// 	z.memory.ContendReadNoMreq_loop(z.PC(), 1, 5)
// 	z.IncPC(1)
// 	z.D = z.memory.ReadByte(z.IY() + uint16(signExtend(offset)))
// }

// /* LD E,IYH */
// func instrFD__LD_E_REGH(z *Z80, opcode byte) {
// 	z.E = z.IYH
// }

// /* LD E,IYL */
// func instrFD__LD_E_REGL(z *Z80, opcode byte) {
// 	z.E = z.IYL
// }

// /* LD E,(iy+dd) */
// func instrFD__LD_E_iREGpDD(z *Z80, opcode byte) {
// 	var offset byte = z.memory.ReadByte(z.PC())
// 	z.memory.ContendReadNoMreq_loop(z.PC(), 1, 5)
// 	z.IncPC(1)
// 	z.E = z.memory.ReadByte(z.IY() + uint16(signExtend(offset)))
// }

// /* LD IYH,B */
// func instrFD__LD_REGH_B(z *Z80, opcode byte) {
// 	z.IYH = z.B
// }

// /* LD IYH,C */
// func instrFD__LD_REGH_C(z *Z80, opcode byte) {
// 	z.IYH = z.C
// }

// /* LD IYH,D */
// func instrFD__LD_REGH_D(z *Z80, opcode byte) {
// 	z.IYH = z.D
// }

// /* LD IYH,E */
// func instrFD__LD_REGH_E(z *Z80, opcode byte) {
// 	z.IYH = z.E
// }

// /* LD IYH,IYH */
// func instrFD__LD_REGH_REGH(z *Z80, opcode byte) {
// }

// /* LD IYH,IYL */
// func instrFD__LD_REGH_REGL(z *Z80, opcode byte) {
// 	z.IYH = z.IYL
// }

// /* LD H,(iy+dd) */
// func instrFD__LD_H_iREGpDD(z *Z80, opcode byte) {
// 	var offset byte = z.memory.ReadByte(z.PC())
// 	z.memory.ContendReadNoMreq_loop(z.PC(), 1, 5)
// 	z.IncPC(1)
// 	z.H = z.memory.ReadByte(z.IY() + uint16(signExtend(offset)))
// }

// /* LD IYH,A */
// func instrFD__LD_REGH_A(z *Z80, opcode byte) {
// 	z.IYH = z.A
// }

// /* LD IYL,B */
// func instrFD__LD_REGL_B(z *Z80, opcode byte) {
// 	z.IYL = z.B
// }

// /* LD IYL,C */
// func instrFD__LD_REGL_C(z *Z80, opcode byte) {
// 	z.IYL = z.C
// }

// /* LD IYL,D */
// func instrFD__LD_REGL_D(z *Z80, opcode byte) {
// 	z.IYL = z.D
// }

// /* LD IYL,E */
// func instrFD__LD_REGL_E(z *Z80, opcode byte) {
// 	z.IYL = z.E
// }

// /* LD IYL,IYH */
// func instrFD__LD_REGL_REGH(z *Z80, opcode byte) {
// 	z.IYL = z.IYH
// }

// /* LD IYL,IYL */
// func instrFD__LD_REGL_REGL(z *Z80, opcode byte) {
// }

// /* LD L,(iy+dd) */
// func instrFD__LD_L_iREGpDD(z *Z80, opcode byte) {
// 	var offset byte = z.memory.ReadByte(z.PC())
// 	z.memory.ContendReadNoMreq_loop(z.PC(), 1, 5)
// 	z.IncPC(1)
// 	z.L = z.memory.ReadByte(z.IY() + uint16(signExtend(offset)))
// }

// /* LD IYL,A */
// func instrFD__LD_REGL_A(z *Z80, opcode byte) {
// 	z.IYL = z.A
// }

// /* LD (iy+dd),B */
// func instrFD__LD_iREGpDD_B(z *Z80, opcode byte) {
// 	offset := z.memory.ReadByte(z.PC())
// 	z.memory.ContendReadNoMreq_loop(z.PC(), 1, 5)
// 	z.IncPC(1)
// 	z.memory.WriteByte(z.IY()+uint16(signExtend(offset)), z.B)
// }

// /* LD (iy+dd),C */
// func instrFD__LD_iREGpDD_C(z *Z80, opcode byte) {
// 	offset := z.memory.ReadByte(z.PC())
// 	z.memory.ContendReadNoMreq_loop(z.PC(), 1, 5)
// 	z.IncPC(1)
// 	z.memory.WriteByte(z.IY()+uint16(signExtend(offset)), z.C)
// }

// /* LD (iy+dd),D */
// func instrFD__LD_iREGpDD_D(z *Z80, opcode byte) {
// 	offset := z.memory.ReadByte(z.PC())
// 	z.memory.ContendReadNoMreq_loop(z.PC(), 1, 5)
// 	z.IncPC(1)
// 	z.memory.WriteByte(z.IY()+uint16(signExtend(offset)), z.D)
// }

// /* LD (iy+dd),E */
// func instrFD__LD_iREGpDD_E(z *Z80, opcode byte) {
// 	offset := z.memory.ReadByte(z.PC())
// 	z.memory.ContendReadNoMreq_loop(z.PC(), 1, 5)
// 	z.IncPC(1)
// 	z.memory.WriteByte(z.IY()+uint16(signExtend(offset)), z.E)
// }

// /* LD (iy+dd),H */
// func instrFD__LD_iREGpDD_H(z *Z80, opcode byte) {
// 	offset := z.memory.ReadByte(z.PC())
// 	z.memory.ContendReadNoMreq_loop(z.PC(), 1, 5)
// 	z.IncPC(1)
// 	z.memory.WriteByte(z.IY()+uint16(signExtend(offset)), z.H)
// }

// /* LD (iy+dd),L */
// func instrFD__LD_iREGpDD_L(z *Z80, opcode byte) {
// 	offset := z.memory.ReadByte(z.PC())
// 	z.memory.ContendReadNoMreq_loop(z.PC(), 1, 5)
// 	z.IncPC(1)
// 	z.memory.WriteByte(z.IY()+uint16(signExtend(offset)), z.L)
// }

// /* LD (iy+dd),A */
// func instrFD__LD_iREGpDD_A(z *Z80, opcode byte) {
// 	offset := z.memory.ReadByte(z.PC())
// 	z.memory.ContendReadNoMreq_loop(z.PC(), 1, 5)
// 	z.IncPC(1)
// 	z.memory.WriteByte(z.IY()+uint16(signExtend(offset)), z.A)
// }

// /* LD A,IYH */
// func instrFD__LD_A_REGH(z *Z80, opcode byte) {
// 	z.A = z.IYH
// }

// /* LD A,IYL */
// func instrFD__LD_A_REGL(z *Z80, opcode byte) {
// 	z.A = z.IYL
// }

// /* LD A,(iy+dd) */
// func instrFD__LD_A_iREGpDD(z *Z80, opcode byte) {
// 	var offset byte = z.memory.ReadByte(z.PC())
// 	z.memory.ContendReadNoMreq_loop(z.PC(), 1, 5)
// 	z.IncPC(1)
// 	z.A = z.memory.ReadByte(z.IY() + uint16(signExtend(offset)))
// }

// /* ADD A,IYH */
// func instrFD__ADD_A_REGH(z *Z80, opcode byte) {
// 	z.add(z.IYH)
// }

// /* ADD A,IYL */
// func instrFD__ADD_A_REGL(z *Z80, opcode byte) {
// 	z.add(z.IYL)
// }

// /* ADD A,(iy+dd) */
// func instrFD__ADD_A_iREGpDD(z *Z80, opcode byte) {
// 	var offset byte = z.memory.ReadByte(z.PC())
// 	z.memory.ContendReadNoMreq_loop(z.PC(), 1, 5)
// 	z.IncPC(1)
// 	var bytetemp byte = z.memory.ReadByte(z.IY() + uint16(signExtend(offset)))
// 	z.add(bytetemp)
// }

// /* ADC A,IYH */
// func instrFD__ADC_A_REGH(z *Z80, opcode byte) {
// 	z.adc(z.IYH)
// }

// /* ADC A,IYL */
// func instrFD__ADC_A_REGL(z *Z80, opcode byte) {
// 	z.adc(z.IYL)
// }

// /* ADC A,(iy+dd) */
// func instrFD__ADC_A_iREGpDD(z *Z80, opcode byte) {
// 	var offset byte = z.memory.ReadByte(z.PC())
// 	z.memory.ContendReadNoMreq_loop(z.PC(), 1, 5)
// 	z.IncPC(1)
// 	var bytetemp byte = z.memory.ReadByte(z.IY() + uint16(signExtend(offset)))
// 	z.adc(bytetemp)
// }

// /* SUB A,IYH */
// func instrFD__SUB_A_REGH(z *Z80, opcode byte) {
// 	z.sub(z.IYH)
// }

// /* SUB A,IYL */
// func instrFD__SUB_A_REGL(z *Z80, opcode byte) {
// 	z.sub(z.IYL)
// }

// /* SUB A,(iy+dd) */
// func instrFD__SUB_A_iREGpDD(z *Z80, opcode byte) {
// 	var offset byte = z.memory.ReadByte(z.PC())
// 	z.memory.ContendReadNoMreq_loop(z.PC(), 1, 5)
// 	z.IncPC(1)
// 	var bytetemp byte = z.memory.ReadByte(z.IY() + uint16(signExtend(offset)))
// 	z.sub(bytetemp)
// }

// /* SBC A,IYH */
// func instrFD__SBC_A_REGH(z *Z80, opcode byte) {
// 	z.sbc(z.IYH)
// }

// /* SBC A,IYL */
// func instrFD__SBC_A_REGL(z *Z80, opcode byte) {
// 	z.sbc(z.IYL)
// }

// /* SBC A,(iy+dd) */
// func instrFD__SBC_A_iREGpDD(z *Z80, opcode byte) {
// 	var offset byte = z.memory.ReadByte(z.PC())
// 	z.memory.ContendReadNoMreq_loop(z.PC(), 1, 5)
// 	z.IncPC(1)
// 	var bytetemp byte = z.memory.ReadByte(z.IY() + uint16(signExtend(offset)))
// 	z.sbc(bytetemp)
// }

// /* AND A,IYH */
// func instrFD__AND_A_REGH(z *Z80, opcode byte) {
// 	z.and(z.IYH)
// }

// /* AND A,IYL */
// func instrFD__AND_A_REGL(z *Z80, opcode byte) {
// 	z.and(z.IYL)
// }

// /* AND A,(iy+dd) */
// func instrFD__AND_A_iREGpDD(z *Z80, opcode byte) {
// 	var offset byte = z.memory.ReadByte(z.PC())
// 	z.memory.ContendReadNoMreq_loop(z.PC(), 1, 5)
// 	z.IncPC(1)
// 	var bytetemp byte = z.memory.ReadByte(z.IY() + uint16(signExtend(offset)))
// 	z.and(bytetemp)
// }

// /* XOR A,IYH */
// func instrFD__XOR_A_REGH(z *Z80, opcode byte) {
// 	z.xor(z.IYH)
// }

// /* XOR A,IYL */
// func instrFD__XOR_A_REGL(z *Z80, opcode byte) {
// 	z.xor(z.IYL)
// }

// /* XOR A,(iy+dd) */
// func instrFD__XOR_A_iREGpDD(z *Z80, opcode byte) {
// 	var offset byte = z.memory.ReadByte(z.PC())
// 	z.memory.ContendReadNoMreq_loop(z.PC(), 1, 5)
// 	z.IncPC(1)
// 	var bytetemp byte = z.memory.ReadByte(z.IY() + uint16(signExtend(offset)))
// 	z.xor(bytetemp)
// }

// /* OR A,IYH */
// func instrFD__OR_A_REGH(z *Z80, opcode byte) {
// 	z.or(z.IYH)
// }

// /* OR A,IYL */
// func instrFD__OR_A_REGL(z *Z80, opcode byte) {
// 	z.or(z.IYL)
// }

// /* OR A,(iy+dd) */
// func instrFD__OR_A_iREGpDD(z *Z80, opcode byte) {
// 	var offset byte = z.memory.ReadByte(z.PC())
// 	z.memory.ContendReadNoMreq_loop(z.PC(), 1, 5)
// 	z.IncPC(1)
// 	var bytetemp byte = z.memory.ReadByte(z.IY() + uint16(signExtend(offset)))
// 	z.or(bytetemp)
// }

// /* CP A,IYH */
// func instrFD__CP_A_REGH(z *Z80, opcode byte) {
// 	z.cp(z.IYH)
// }

// /* CP A,IYL */
// func instrFD__CP_A_REGL(z *Z80, opcode byte) {
// 	z.cp(z.IYL)
// }

// /* CP A,(iy+dd) */
// func instrFD__CP_A_iREGpDD(z *Z80, opcode byte) {
// 	var offset byte = z.memory.ReadByte(z.PC())
// 	z.memory.ContendReadNoMreq_loop(z.PC(), 1, 5)
// 	z.IncPC(1)
// 	var bytetemp byte = z.memory.ReadByte(z.IY() + uint16(signExtend(offset)))
// 	z.cp(bytetemp)
// }

// /* shift DDFDCB */
// func instrFD__SHIFT_DDFDCB(z *Z80, opcode byte) {
// }

// /* POP iy */
// func instrFD__POP_REG(z *Z80, opcode byte) {
// 	z.IYL, z.IYH = z.pop16()
// }

// /* EX (SP),iy */
// func instrFD__EX_iSP_REG(z *Z80, opcode byte) {
// 	var bytetempl = z.memory.ReadByte(z.SP())
// 	var bytetemph = z.memory.ReadByte(z.SP() + 1)
// 	z.memory.ContendReadNoMreq(z.SP()+1, 1)
// 	z.memory.WriteByte(z.SP()+1, z.IYH)
// 	z.memory.WriteByte(z.SP(), z.IYL)
// 	z.memory.ContendWriteNoMreq_loop(z.SP(), 1, 2)
// 	z.IYL = bytetempl
// 	z.IYH = bytetemph
// }

// /* PUSH iy */
// func instrFD__PUSH_REG(z *Z80, opcode byte) {
// 	z.memory.ContendReadNoMreq(z.IR(), 1)
// 	z.push16(z.IYL, z.IYH)
// }

// /* JP iy */
// func instrFD__JP_REG(z *Z80, opcode byte) {
// 	z.SetPC(z.IY()) /* NB: NOT INDIRECT! */
// }

// /* LD SP,iy */
// func instrFD__LD_SP_REG(z *Z80, opcode byte) {
// 	z.memory.ContendReadNoMreq_loop(z.IR(), 1, 2)
// 	z.SetSP(z.IY())
// }

//--
