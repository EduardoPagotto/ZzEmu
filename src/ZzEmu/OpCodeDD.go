package ZzEmu

func initOpcodeDDMap() {
	// 	// BEGIN of 0xdd shifted opcodes
	// 	/* ADD REGISTER,BC */
	// 	OpcodeDDCBMap[0x09] = instrDD__ADD_REG_BC
	// 	/* ADD REGISTER,DE */
	// 	OpcodeDDCBMap[0x19] = instrDD__ADD_REG_DE
	// 	/* LD REGISTER,nnnn */
	// 	OpcodeDDCBMap[0x21] = instrDD__LD_REG_NNNN
	/* LD (nnnn),REGISTER */
	OpcodeDDCBMap[0x22] = instrDD__LD_iNNNN_REG
	// 	/* INC REGISTER */
	// 	OpcodeDDCBMap[0x23] = instrDD__INC_REG
	// 	/* INC REGISTERH */
	// 	OpcodeDDCBMap[0x24] = instrDD__INC_REGH
	// 	/* DEC REGISTERH */
	// 	OpcodeDDCBMap[0x25] = instrDD__DEC_REGH
	// 	/* LD REGISTERH,nn */
	// 	OpcodeDDCBMap[0x26] = instrDD__LD_REGH_NN
	// 	/* ADD REGISTER,REGISTER */
	// 	OpcodeDDCBMap[0x29] = instrDD__ADD_REG_REG
	/* LD REGISTER,(nnnn) */
	OpcodeDDCBMap[0x2a] = instrDD__LD_REG_iNNNN
	// 	/* DEC REGISTER */
	// 	OpcodeDDCBMap[0x2b] = instrDD__DEC_REG
	// 	/* INC REGISTERL */
	// 	OpcodeDDCBMap[0x2c] = instrDD__INC_REGL
	// 	/* DEC REGISTERL */
	// 	OpcodeDDCBMap[0x2d] = instrDD__DEC_REGL
	// 	/* LD REGISTERL,nn */
	// 	OpcodeDDCBMap[0x2e] = instrDD__LD_REGL_NN
	// 	/* INC (REGISTER+dd) */
	// 	OpcodeDDCBMap[0x34] = instrDD__INC_iREGpDD
	// 	/* DEC (REGISTER+dd) */
	// 	OpcodeDDCBMap[0x35] = instrDD__DEC_iREGpDD
	// 	/* LD (REGISTER+dd),nn */
	// 	OpcodeDDCBMap[0x36] = instrDD__LD_iREGpDD_NN
	// 	/* ADD REGISTER,SP */
	// 	OpcodeDDCBMap[0x39] = instrDD__ADD_REG_SP
	// 	/* LD B,REGISTERH */
	// 	OpcodeDDCBMap[0x44] = instrDD__LD_B_REGH
	// 	/* LD B,REGISTERL */
	// 	OpcodeDDCBMap[0x45] = instrDD__LD_B_REGL
	// 	/* LD B,(REGISTER+dd) */
	// 	OpcodeDDCBMap[0x46] = instrDD__LD_B_iREGpDD
	// 	/* LD C,REGISTERH */
	// 	OpcodeDDCBMap[0x4c] = instrDD__LD_C_REGH
	// 	/* LD C,REGISTERL */
	// 	OpcodeDDCBMap[0x4d] = instrDD__LD_C_REGL
	// 	/* LD C,(REGISTER+dd) */
	// 	OpcodeDDCBMap[0x4e] = instrDD__LD_C_iREGpDD
	// 	/* LD D,REGISTERH */
	// 	OpcodeDDCBMap[0x54] = instrDD__LD_D_REGH
	// 	/* LD D,REGISTERL */
	// 	OpcodeDDCBMap[0x55] = instrDD__LD_D_REGL
	// 	/* LD D,(REGISTER+dd) */
	// 	OpcodeDDCBMap[0x56] = instrDD__LD_D_iREGpDD
	// 	/* LD E,REGISTERH */
	// 	OpcodeDDCBMap[0x5c] = instrDD__LD_E_REGH
	// 	/* LD E,REGISTERL */
	// 	OpcodeDDCBMap[0x5d] = instrDD__LD_E_REGL
	// 	/* LD E,(REGISTER+dd) */
	// 	OpcodeDDCBMap[0x5e] = instrDD__LD_E_iREGpDD
	// 	/* LD REGISTERH,B */
	// 	OpcodeDDCBMap[0x60] = instrDD__LD_REGH_B
	// 	/* LD REGISTERH,C */
	// 	OpcodeDDCBMap[0x61] = instrDD__LD_REGH_C
	// 	/* LD REGISTERH,D */
	// 	OpcodeDDCBMap[0x62] = instrDD__LD_REGH_D
	// 	/* LD REGISTERH,E */
	// 	OpcodeDDCBMap[0x63] = instrDD__LD_REGH_E
	// 	/* LD REGISTERH,REGISTERH */
	// 	OpcodeDDCBMap[0x64] = instrDD__LD_REGH_REGH
	// 	/* LD REGISTERH,REGISTERL */
	// 	OpcodeDDCBMap[0x65] = instrDD__LD_REGH_REGL
	// 	/* LD H,(REGISTER+dd) */
	// 	OpcodeDDCBMap[0x66] = instrDD__LD_H_iREGpDD
	// 	/* LD REGISTERH,A */
	// 	OpcodeDDCBMap[0x67] = instrDD__LD_REGH_A
	// 	/* LD REGISTERL,B */
	// 	OpcodeDDCBMap[0x68] = instrDD__LD_REGL_B
	// 	/* LD REGISTERL,C */
	// 	OpcodeDDCBMap[0x69] = instrDD__LD_REGL_C
	// 	/* LD REGISTERL,D */
	// 	OpcodeDDCBMap[0x6a] = instrDD__LD_REGL_D
	// 	/* LD REGISTERL,E */
	// 	OpcodeDDCBMap[0x6b] = instrDD__LD_REGL_E
	// 	/* LD REGISTERL,REGISTERH */
	// 	OpcodeDDCBMap[0x6c] = instrDD__LD_REGL_REGH
	// 	/* LD REGISTERL,REGISTERL */
	// 	OpcodeDDCBMap[0x6d] = instrDD__LD_REGL_REGL
	// 	/* LD L,(REGISTER+dd) */
	// 	OpcodeDDCBMap[0x6e] = instrDD__LD_L_iREGpDD
	// 	/* LD REGISTERL,A */
	// 	OpcodeDDCBMap[0x6f] = instrDD__LD_REGL_A
	// 	/* LD (REGISTER+dd),B */
	// 	OpcodeDDCBMap[0x70] = instrDD__LD_iREGpDD_B
	// 	/* LD (REGISTER+dd),C */
	// 	OpcodeDDCBMap[0x71] = instrDD__LD_iREGpDD_C
	// 	/* LD (REGISTER+dd),D */
	// 	OpcodeDDCBMap[0x72] = instrDD__LD_iREGpDD_D
	// 	/* LD (REGISTER+dd),E */
	// 	OpcodeDDCBMap[0x73] = instrDD__LD_iREGpDD_E
	// 	/* LD (REGISTER+dd),H */
	// 	OpcodeDDCBMap[0x74] = instrDD__LD_iREGpDD_H
	// 	/* LD (REGISTER+dd),L */
	// 	OpcodeDDCBMap[0x75] = instrDD__LD_iREGpDD_L
	// 	/* LD (REGISTER+dd),A */
	// 	OpcodeDDCBMap[0x77] = instrDD__LD_iREGpDD_A
	// 	/* LD A,REGISTERH */
	// 	OpcodeDDCBMap[0x7c] = instrDD__LD_A_REGH
	// 	/* LD A,REGISTERL */
	// 	OpcodeDDCBMap[0x7d] = instrDD__LD_A_REGL
	// 	/* LD A,(REGISTER+dd) */
	// 	OpcodeDDCBMap[0x7e] = instrDD__LD_A_iREGpDD
	// 	/* ADD A,REGISTERH */
	// 	OpcodeDDCBMap[0x84] = instrDD__ADD_A_REGH
	// 	/* ADD A,REGISTERL */
	// 	OpcodeDDCBMap[0x85] = instrDD__ADD_A_REGL
	// 	/* ADD A,(REGISTER+dd) */
	// 	OpcodeDDCBMap[0x86] = instrDD__ADD_A_iREGpDD
	// 	/* ADC A,REGISTERH */
	// 	OpcodeDDCBMap[0x8c] = instrDD__ADC_A_REGH
	// 	/* ADC A,REGISTERL */
	// 	OpcodeDDCBMap[0x8d] = instrDD__ADC_A_REGL
	// 	/* ADC A,(REGISTER+dd) */
	// 	OpcodeDDCBMap[0x8e] = instrDD__ADC_A_iREGpDD
	// 	/* SUB A,REGISTERH */
	// 	OpcodeDDCBMap[0x94] = instrDD__SUB_A_REGH
	// 	/* SUB A,REGISTERL */
	// 	OpcodeDDCBMap[0x95] = instrDD__SUB_A_REGL
	// 	/* SUB A,(REGISTER+dd) */
	// 	OpcodeDDCBMap[0x96] = instrDD__SUB_A_iREGpDD
	// 	/* SBC A,REGISTERH */
	// 	OpcodeDDCBMap[0x9c] = instrDD__SBC_A_REGH
	// 	/* SBC A,REGISTERL */
	// 	OpcodeDDCBMap[0x9d] = instrDD__SBC_A_REGL
	// 	/* SBC A,(REGISTER+dd) */
	// 	OpcodeDDCBMap[0x9e] = instrDD__SBC_A_iREGpDD
	// 	/* AND A,REGISTERH */
	// 	OpcodeDDCBMap[0xa4] = instrDD__AND_A_REGH
	// 	/* AND A,REGISTERL */
	// 	OpcodeDDCBMap[0xa5] = instrDD__AND_A_REGL
	// 	/* AND A,(REGISTER+dd) */
	// 	OpcodeDDCBMap[0xa6] = instrDD__AND_A_iREGpDD
	// 	/* XOR A,REGISTERH */
	// 	OpcodeDDCBMap[0xac] = instrDD__XOR_A_REGH
	// 	/* XOR A,REGISTERL */
	// 	OpcodeDDCBMap[0xad] = instrDD__XOR_A_REGL
	// 	/* XOR A,(REGISTER+dd) */
	// 	OpcodeDDCBMap[0xae] = instrDD__XOR_A_iREGpDD
	// 	/* OR A,REGISTERH */
	// 	OpcodeDDCBMap[0xb4] = instrDD__OR_A_REGH
	// 	/* OR A,REGISTERL */
	// 	OpcodeDDCBMap[0xb5] = instrDD__OR_A_REGL
	// 	/* OR A,(REGISTER+dd) */
	// 	OpcodeDDCBMap[0xb6] = instrDD__OR_A_iREGpDD
	// 	/* CP A,REGISTERH */
	// 	OpcodeDDCBMap[0xbc] = instrDD__CP_A_REGH
	// 	/* CP A,REGISTERL */
	// 	OpcodeDDCBMap[0xbd] = instrDD__CP_A_REGL
	// 	/* CP A,(REGISTER+dd) */
	// 	OpcodeDDCBMap[0xbe] = instrDD__CP_A_iREGpDD
	// 	/* shift DDFDCB */
	OpcodeDDMap[0xcb] = instrDD__SHIFT_DDFDCB
	// 	/* POP REGISTER */
	// 	OpcodeDDCBMap[0xe1] = instrDD__POP_REG
	// 	/* EX (SP),REGISTER */
	// 	OpcodeDDCBMap[0xe3] = instrDD__EX_iSP_REG
	// 	/* PUSH REGISTER */
	// 	OpcodeDDCBMap[0xe5] = instrDD__PUSH_REG
	// 	/* JP REGISTER */
	// 	OpcodeDDCBMap[0xe9] = instrDD__JP_REG
	// 	/* LD SP,REGISTER */
	// 	OpcodeDDCBMap[0xf9] = instrDD__LD_SP_REG

	// 	// END of 0xdd shifted opcodes

}

// /* ADD ix,BC */
// func instrDD__ADD_REG_BC(z *Z80, opcode byte) {
// 	z.memory.ContendReadNoMreq_loop(z.IR(), 1, 7)
// 	z.add16(z.ix, z.BC())
// }

// /* ADD ix,DE */
// func instrDD__ADD_REG_DE(z *Z80, opcode byte) {
// 	z.memory.ContendReadNoMreq_loop(z.IR(), 1, 7)
// 	z.add16(z.ix, z.DE())
// }

// /* LD ix,nnnn */
// func instrDD__LD_REG_NNNN(z *Z80, opcode byte) {
// 	b1 := z.memory.ReadByte(z.PC())
// 	z.IncPC(1)
// 	b2 := z.memory.ReadByte(z.PC())
// 	z.IncPC(1)
// 	z.SetIX(joinBytes(b2, b1))
// }

/* LD (nnnn),ix */
func instrDD__LD_iNNNN_REG(z *Z80, opcode byte) {
	z.Tstates += 20
	z.StoreIndex16(z.IXL, z.IXH)
	// break
}

// /* INC ix */
// func instrDD__INC_REG(z *Z80, opcode byte) {
// 	z.memory.ContendReadNoMreq_loop(z.IR(), 1, 2)
// 	z.IncIX()
// }

// /* INC IXH */
// func instrDD__INC_REGH(z *Z80, opcode byte) {
// 	z.incIXH()
// }

// /* DEC IXH */
// func instrDD__DEC_REGH(z *Z80, opcode byte) {
// 	z.decIXH()
// }

// /* LD IXH,nn */
// func instrDD__LD_REGH_NN(z *Z80, opcode byte) {
// 	z.IXH = z.memory.ReadByte(z.PC())
// 	z.IncPC(1)
// }

// /* ADD ix,ix */
// func instrDD__ADD_REG_REG(z *Z80, opcode byte) {
// 	z.memory.ContendReadNoMreq_loop(z.IR(), 1, 7)
// 	z.add16(z.ix, z.IX())
// }

/* LD ix,(nnnn) */
func instrDD__LD_REG_iNNNN(z *Z80, opcode byte) {
	z.Tstates += 20
	z.LoadIndex16(&z.IXL, &z.IXH)
	// break
}

// /* DEC ix */
// func instrDD__DEC_REG(z *Z80, opcode byte) {
// 	z.memory.ContendReadNoMreq_loop(z.IR(), 1, 2)
// 	z.DecIX()
// }

// /* INC IXL */
// func instrDD__INC_REGL(z *Z80, opcode byte) {
// 	z.incIXL()
// }

// /* DEC IXL */
// func instrDD__DEC_REGL(z *Z80, opcode byte) {
// 	z.decIXL()
// }

// /* LD IXL,nn */
// func instrDD__LD_REGL_NN(z *Z80, opcode byte) {
// 	z.IXL = z.memory.ReadByte(z.PC())
// 	z.IncPC(1)
// }

// /* INC (ix+dd) */
// func instrDD__INC_iREGpDD(z *Z80, opcode byte) {
// 	var offset byte = z.memory.ReadByte(z.PC())
// 	z.memory.ContendReadNoMreq_loop(z.PC(), 1, 5)
// 	z.IncPC(1)
// 	var wordtemp uint16 = z.IX() + uint16(signExtend(offset))
// 	var bytetemp byte = z.memory.ReadByte(wordtemp)
// 	z.memory.ContendReadNoMreq(wordtemp, 1)
// 	z.inc(&bytetemp)
// 	z.memory.WriteByte(wordtemp, bytetemp)
// }

// /* DEC (ix+dd) */
// func instrDD__DEC_iREGpDD(z *Z80, opcode byte) {
// 	var offset byte = z.memory.ReadByte(z.PC())
// 	z.memory.ContendReadNoMreq_loop(z.PC(), 1, 5)
// 	z.IncPC(1)
// 	var wordtemp uint16 = z.IX() + uint16(signExtend(offset))
// 	var bytetemp byte = z.memory.ReadByte(wordtemp)
// 	z.memory.ContendReadNoMreq(wordtemp, 1)
// 	z.dec(&bytetemp)
// 	z.memory.WriteByte(wordtemp, bytetemp)
// }

// /* LD (ix+dd),nn */
// func instrDD__LD_iREGpDD_NN(z *Z80, opcode byte) {
// 	offset := z.memory.ReadByte(z.PC())
// 	z.IncPC(1)
// 	value := z.memory.ReadByte(z.PC())
// 	z.memory.ContendReadNoMreq_loop(z.PC(), 1, 2)
// 	z.IncPC(1)
// 	z.memory.WriteByte(z.IX()+uint16(signExtend(offset)), value)
// }

// /* ADD ix,SP */
// func instrDD__ADD_REG_SP(z *Z80, opcode byte) {
// 	z.memory.ContendReadNoMreq_loop(z.IR(), 1, 7)
// 	z.add16(z.ix, z.SP())
// }

// /* LD B,IXH */
// func instrDD__LD_B_REGH(z *Z80, opcode byte) {
// 	z.B = z.IXH
// }

// /* LD B,IXL */
// func instrDD__LD_B_REGL(z *Z80, opcode byte) {
// 	z.B = z.IXL
// }

// /* LD B,(ix+dd) */
// func instrDD__LD_B_iREGpDD(z *Z80, opcode byte) {
// 	var offset byte = z.memory.ReadByte(z.PC())
// 	z.memory.ContendReadNoMreq_loop(z.PC(), 1, 5)
// 	z.IncPC(1)
// 	z.B = z.memory.ReadByte(z.IX() + uint16(signExtend(offset)))
// }

// /* LD C,IXH */
// func instrDD__LD_C_REGH(z *Z80, opcode byte) {
// 	z.C = z.IXH
// }

// /* LD C,IXL */
// func instrDD__LD_C_REGL(z *Z80, opcode byte) {
// 	z.C = z.IXL
// }

// /* LD C,(ix+dd) */
// func instrDD__LD_C_iREGpDD(z *Z80, opcode byte) {
// 	var offset byte = z.memory.ReadByte(z.PC())
// 	z.memory.ContendReadNoMreq_loop(z.PC(), 1, 5)
// 	z.IncPC(1)
// 	z.C = z.memory.ReadByte(z.IX() + uint16(signExtend(offset)))
// }

// /* LD D,IXH */
// func instrDD__LD_D_REGH(z *Z80, opcode byte) {
// 	z.D = z.IXH
// }

// /* LD D,IXL */
// func instrDD__LD_D_REGL(z *Z80, opcode byte) {
// 	z.D = z.IXL
// }

// /* LD D,(ix+dd) */
// func instrDD__LD_D_iREGpDD(z *Z80, opcode byte) {
// 	var offset byte = z.memory.ReadByte(z.PC())
// 	z.memory.ContendReadNoMreq_loop(z.PC(), 1, 5)
// 	z.IncPC(1)
// 	z.D = z.memory.ReadByte(z.IX() + uint16(signExtend(offset)))
// }

// /* LD E,IXH */
// func instrDD__LD_E_REGH(z *Z80, opcode byte) {
// 	z.E = z.IXH
// }

// /* LD E,IXL */
// func instrDD__LD_E_REGL(z *Z80, opcode byte) {
// 	z.E = z.IXL
// }

// /* LD E,(ix+dd) */
// func instrDD__LD_E_iREGpDD(z *Z80, opcode byte) {
// 	var offset byte = z.memory.ReadByte(z.PC())
// 	z.memory.ContendReadNoMreq_loop(z.PC(), 1, 5)
// 	z.IncPC(1)
// 	z.E = z.memory.ReadByte(z.IX() + uint16(signExtend(offset)))
// }

// /* LD IXH,B */
// func instrDD__LD_REGH_B(z *Z80, opcode byte) {
// 	z.IXH = z.B
// }

// /* LD IXH,C */
// func instrDD__LD_REGH_C(z *Z80, opcode byte) {
// 	z.IXH = z.C
// }

// /* LD IXH,D */
// func instrDD__LD_REGH_D(z *Z80, opcode byte) {
// 	z.IXH = z.D
// }

// /* LD IXH,E */
// func instrDD__LD_REGH_E(z *Z80, opcode byte) {
// 	z.IXH = z.E
// }

// /* LD IXH,IXH */
// func instrDD__LD_REGH_REGH(z *Z80, opcode byte) {
// }

// /* LD IXH,IXL */
// func instrDD__LD_REGH_REGL(z *Z80, opcode byte) {
// 	z.IXH = z.IXL
// }

// /* LD H,(ix+dd) */
// func instrDD__LD_H_iREGpDD(z *Z80, opcode byte) {
// 	var offset byte = z.memory.ReadByte(z.PC())
// 	z.memory.ContendReadNoMreq_loop(z.PC(), 1, 5)
// 	z.IncPC(1)
// 	z.H = z.memory.ReadByte(z.IX() + uint16(signExtend(offset)))
// }

// /* LD IXH,A */
// func instrDD__LD_REGH_A(z *Z80, opcode byte) {
// 	z.IXH = z.A
// }

// /* LD IXL,B */
// func instrDD__LD_REGL_B(z *Z80, opcode byte) {
// 	z.IXL = z.B
// }

// /* LD IXL,C */
// func instrDD__LD_REGL_C(z *Z80, opcode byte) {
// 	z.IXL = z.C
// }

// /* LD IXL,D */
// func instrDD__LD_REGL_D(z *Z80, opcode byte) {
// 	z.IXL = z.D
// }

// /* LD IXL,E */
// func instrDD__LD_REGL_E(z *Z80, opcode byte) {
// 	z.IXL = z.E
// }

// /* LD IXL,IXH */
// func instrDD__LD_REGL_REGH(z *Z80, opcode byte) {
// 	z.IXL = z.IXH
// }

// /* LD IXL,IXL */
// func instrDD__LD_REGL_REGL(z *Z80, opcode byte) {
// }

// /* LD L,(ix+dd) */
// func instrDD__LD_L_iREGpDD(z *Z80, opcode byte) {
// 	var offset byte = z.memory.ReadByte(z.PC())
// 	z.memory.ContendReadNoMreq_loop(z.PC(), 1, 5)
// 	z.IncPC(1)
// 	z.L = z.memory.ReadByte(z.IX() + uint16(signExtend(offset)))
// }

// /* LD IXL,A */
// func instrDD__LD_REGL_A(z *Z80, opcode byte) {
// 	z.IXL = z.A
// }

// /* LD (ix+dd),B */
// func instrDD__LD_iREGpDD_B(z *Z80, opcode byte) {
// 	offset := z.memory.ReadByte(z.PC())
// 	z.memory.ContendReadNoMreq_loop(z.PC(), 1, 5)
// 	z.IncPC(1)
// 	z.memory.WriteByte(z.IX()+uint16(signExtend(offset)), z.B)
// }

// /* LD (ix+dd),C */
// func instrDD__LD_iREGpDD_C(z *Z80, opcode byte) {
// 	offset := z.memory.ReadByte(z.PC())
// 	z.memory.ContendReadNoMreq_loop(z.PC(), 1, 5)
// 	z.IncPC(1)
// 	z.memory.WriteByte(z.IX()+uint16(signExtend(offset)), z.C)
// }

// /* LD (ix+dd),D */
// func instrDD__LD_iREGpDD_D(z *Z80, opcode byte) {
// 	offset := z.memory.ReadByte(z.PC())
// 	z.memory.ContendReadNoMreq_loop(z.PC(), 1, 5)
// 	z.IncPC(1)
// 	z.memory.WriteByte(z.IX()+uint16(signExtend(offset)), z.D)
// }

// /* LD (ix+dd),E */
// func instrDD__LD_iREGpDD_E(z *Z80, opcode byte) {
// 	offset := z.memory.ReadByte(z.PC())
// 	z.memory.ContendReadNoMreq_loop(z.PC(), 1, 5)
// 	z.IncPC(1)
// 	z.memory.WriteByte(z.IX()+uint16(signExtend(offset)), z.E)
// }

// /* LD (ix+dd),H */
// func instrDD__LD_iREGpDD_H(z *Z80, opcode byte) {
// 	offset := z.memory.ReadByte(z.PC())
// 	z.memory.ContendReadNoMreq_loop(z.PC(), 1, 5)
// 	z.IncPC(1)
// 	z.memory.WriteByte(z.IX()+uint16(signExtend(offset)), z.H)
// }

// /* LD (ix+dd),L */
// func instrDD__LD_iREGpDD_L(z *Z80, opcode byte) {
// 	offset := z.memory.ReadByte(z.PC())
// 	z.memory.ContendReadNoMreq_loop(z.PC(), 1, 5)
// 	z.IncPC(1)
// 	z.memory.WriteByte(z.IX()+uint16(signExtend(offset)), z.L)
// }

// /* LD (ix+dd),A */
// func instrDD__LD_iREGpDD_A(z *Z80, opcode byte) {
// 	offset := z.memory.ReadByte(z.PC())
// 	z.memory.ContendReadNoMreq_loop(z.PC(), 1, 5)
// 	z.IncPC(1)
// 	z.memory.WriteByte(z.IX()+uint16(signExtend(offset)), z.A)
// }

// /* LD A,IXH */
// func instrDD__LD_A_REGH(z *Z80, opcode byte) {
// 	z.A = z.IXH
// }

// /* LD A,IXL */
// func instrDD__LD_A_REGL(z *Z80, opcode byte) {
// 	z.A = z.IXL
// }

// /* LD A,(ix+dd) */
// func instrDD__LD_A_iREGpDD(z *Z80, opcode byte) {
// 	var offset byte = z.memory.ReadByte(z.PC())
// 	z.memory.ContendReadNoMreq_loop(z.PC(), 1, 5)
// 	z.IncPC(1)
// 	z.A = z.memory.ReadByte(z.IX() + uint16(signExtend(offset)))
// }

// /* ADD A,IXH */
// func instrDD__ADD_A_REGH(z *Z80, opcode byte) {
// 	z.add(z.IXH)
// }

// /* ADD A,IXL */
// func instrDD__ADD_A_REGL(z *Z80, opcode byte) {
// 	z.add(z.IXL)
// }

// /* ADD A,(ix+dd) */
// func instrDD__ADD_A_iREGpDD(z *Z80, opcode byte) {
// 	var offset byte = z.memory.ReadByte(z.PC())
// 	z.memory.ContendReadNoMreq_loop(z.PC(), 1, 5)
// 	z.IncPC(1)
// 	var bytetemp byte = z.memory.ReadByte(z.IX() + uint16(signExtend(offset)))
// 	z.add(bytetemp)
// }

// /* ADC A,IXH */
// func instrDD__ADC_A_REGH(z *Z80, opcode byte) {
// 	z.adc(z.IXH)
// }

// /* ADC A,IXL */
// func instrDD__ADC_A_REGL(z *Z80, opcode byte) {
// 	z.adc(z.IXL)
// }

// /* ADC A,(ix+dd) */
// func instrDD__ADC_A_iREGpDD(z *Z80, opcode byte) {
// 	var offset byte = z.memory.ReadByte(z.PC())
// 	z.memory.ContendReadNoMreq_loop(z.PC(), 1, 5)
// 	z.IncPC(1)
// 	var bytetemp byte = z.memory.ReadByte(z.IX() + uint16(signExtend(offset)))
// 	z.adc(bytetemp)
// }

// /* SUB A,IXH */
// func instrDD__SUB_A_REGH(z *Z80, opcode byte) {
// 	z.sub(z.IXH)
// }

// /* SUB A,IXL */
// func instrDD__SUB_A_REGL(z *Z80, opcode byte) {
// 	z.sub(z.IXL)
// }

// /* SUB A,(ix+dd) */
// func instrDD__SUB_A_iREGpDD(z *Z80, opcode byte) {
// 	var offset byte = z.memory.ReadByte(z.PC())
// 	z.memory.ContendReadNoMreq_loop(z.PC(), 1, 5)
// 	z.IncPC(1)
// 	var bytetemp byte = z.memory.ReadByte(z.IX() + uint16(signExtend(offset)))
// 	z.sub(bytetemp)
// }

// /* SBC A,IXH */
// func instrDD__SBC_A_REGH(z *Z80, opcode byte) {
// 	z.sbc(z.IXH)
// }

// /* SBC A,IXL */
// func instrDD__SBC_A_REGL(z *Z80, opcode byte) {
// 	z.sbc(z.IXL)
// }

// /* SBC A,(ix+dd) */
// func instrDD__SBC_A_iREGpDD(z *Z80, opcode byte) {
// 	var offset byte = z.memory.ReadByte(z.PC())
// 	z.memory.ContendReadNoMreq_loop(z.PC(), 1, 5)
// 	z.IncPC(1)
// 	var bytetemp byte = z.memory.ReadByte(z.IX() + uint16(signExtend(offset)))
// 	z.sbc(bytetemp)
// }

// /* AND A,IXH */
// func instrDD__AND_A_REGH(z *Z80, opcode byte) {
// 	z.and(z.IXH)
// }

// /* AND A,IXL */
// func instrDD__AND_A_REGL(z *Z80, opcode byte) {
// 	z.and(z.IXL)
// }

// /* AND A,(ix+dd) */
// func instrDD__AND_A_iREGpDD(z *Z80, opcode byte) {
// 	var offset byte = z.memory.ReadByte(z.PC())
// 	z.memory.ContendReadNoMreq_loop(z.PC(), 1, 5)
// 	z.IncPC(1)
// 	var bytetemp byte = z.memory.ReadByte(z.IX() + uint16(signExtend(offset)))
// 	z.and(bytetemp)
// }

// /* XOR A,IXH */
// func instrDD__XOR_A_REGH(z *Z80, opcode byte) {
// 	z.xor(z.IXH)
// }

// /* XOR A,IXL */
// func instrDD__XOR_A_REGL(z *Z80, opcode byte) {
// 	z.xor(z.IXL)
// }

// /* XOR A,(ix+dd) */
// func instrDD__XOR_A_iREGpDD(z *Z80, opcode byte) {
// 	var offset byte = z.memory.ReadByte(z.PC())
// 	z.memory.ContendReadNoMreq_loop(z.PC(), 1, 5)
// 	z.IncPC(1)
// 	var bytetemp byte = z.memory.ReadByte(z.IX() + uint16(signExtend(offset)))
// 	z.xor(bytetemp)
// }

// /* OR A,IXH */
// func instrDD__OR_A_REGH(z *Z80, opcode byte) {
// 	z.or(z.IXH)
// }

// /* OR A,IXL */
// func instrDD__OR_A_REGL(z *Z80, opcode byte) {
// 	z.or(z.IXL)
// }

// /* OR A,(ix+dd) */
// func instrDD__OR_A_iREGpDD(z *Z80, opcode byte) {
// 	var offset byte = z.memory.ReadByte(z.PC())
// 	z.memory.ContendReadNoMreq_loop(z.PC(), 1, 5)
// 	z.IncPC(1)
// 	var bytetemp byte = z.memory.ReadByte(z.IX() + uint16(signExtend(offset)))
// 	z.or(bytetemp)
// }

// /* CP A,IXH */
// func instrDD__CP_A_REGH(z *Z80, opcode byte) {
// 	z.cp(z.IXH)
// }

// /* CP A,IXL */
// func instrDD__CP_A_REGL(z *Z80, opcode byte) {
// 	z.cp(z.IXL)
// }

// /* CP A,(ix+dd) */
// func instrDD__CP_A_iREGpDD(z *Z80, opcode byte) {
// 	var offset byte = z.memory.ReadByte(z.PC())
// 	z.memory.ContendReadNoMreq_loop(z.PC(), 1, 5)
// 	z.IncPC(1)
// 	var bytetemp byte = z.memory.ReadByte(z.IX() + uint16(signExtend(offset)))
// 	z.cp(bytetemp)
// }

// /* shift DDFDCB */
func instrDD__SHIFT_DDFDCB(z *Z80, opcode byte) {
	opcode2 := z.Load8()
	z.R++
	OpcodeDDCBMap[opcode2](z, opcode2)
}

// /* POP ix */
// func instrDD__POP_REG(z *Z80, opcode byte) {
// 	z.IXL, z.IXH = z.pop16()
// }

// /* EX (SP),ix */
// func instrDD__EX_iSP_REG(z *Z80, opcode byte) {
// 	var bytetempl = z.memory.ReadByte(z.SP())
// 	var bytetemph = z.memory.ReadByte(z.SP() + 1)
// 	z.memory.ContendReadNoMreq(z.SP()+1, 1)
// 	z.memory.WriteByte(z.SP()+1, z.IXH)
// 	z.memory.WriteByte(z.SP(), z.IXL)
// 	z.memory.ContendWriteNoMreq_loop(z.SP(), 1, 2)
// 	z.IXL = bytetempl
// 	z.IXH = bytetemph
// }

// /* PUSH ix */
// func instrDD__PUSH_REG(z *Z80, opcode byte) {
// 	z.memory.ContendReadNoMreq(z.IR(), 1)
// 	z.push16(z.IXL, z.IXH)
// }

// /* JP ix */
// func instrDD__JP_REG(z *Z80, opcode byte) {
// 	z.SetPC(z.IX()) /* NB: NOT INDIRECT! */
// }

// /* LD SP,ix */
// func instrDD__LD_SP_REG(z *Z80, opcode byte) {
// 	z.memory.ContendReadNoMreq_loop(z.IR(), 1, 2)
// 	z.SetSP(z.IX())
// }
