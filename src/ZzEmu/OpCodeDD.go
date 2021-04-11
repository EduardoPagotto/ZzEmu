package ZzEmu

func initOpcodeDDMap() {
	// BEGIN of 0xdd shifted opcodes
	/* ADD REGISTER,BC */
	OpcodeDDCBMap[0x09] = instrDD__ADD_REG_BC
	/* ADD REGISTER,DE */
	OpcodeDDCBMap[0x19] = instrDD__ADD_REG_DE
	/* LD REGISTER,nnnn */
	OpcodeDDCBMap[0x21] = instrDD__LD_REG_NNNN
	/* LD (nnnn),REGISTER */
	OpcodeDDCBMap[0x22] = instrDD__LD_iNNNN_REG
	/* INC REGISTER */
	OpcodeDDCBMap[0x23] = instrDD__INC_REG
	/* INC REGISTERH */
	OpcodeDDCBMap[0x24] = instrDD__INC_REGH
	/* DEC REGISTERH */
	OpcodeDDCBMap[0x25] = instrDD__DEC_REGH
	/* LD REGISTERH,nn */
	OpcodeDDCBMap[0x26] = instrDD__LD_REGH_NN
	/* ADD REGISTER,REGISTER */
	OpcodeDDCBMap[0x29] = instrDD__ADD_REG_REG
	/* LD REGISTER,(nnnn) */
	OpcodeDDCBMap[0x2a] = instrDD__LD_REG_iNNNN
	/* DEC REGISTER */
	OpcodeDDCBMap[0x2b] = instrDD__DEC_REG
	/* INC REGISTERL */
	OpcodeDDCBMap[0x2c] = instrDD__INC_REGL
	/* DEC REGISTERL */
	OpcodeDDCBMap[0x2d] = instrDD__DEC_REGL
	/* LD REGISTERL,nn */
	OpcodeDDCBMap[0x2e] = instrDD__LD_REGL_NN
	/* INC (REGISTER+dd) */
	OpcodeDDCBMap[0x34] = instrDD__INC_iREGpDD
	/* DEC (REGISTER+dd) */
	OpcodeDDCBMap[0x35] = instrDD__DEC_iREGpDD
	/* LD (REGISTER+dd),nn */
	OpcodeDDCBMap[0x36] = instrDD__LD_iREGpDD_NN
	/* ADD REGISTER,SP */
	OpcodeDDCBMap[0x39] = instrDD__ADD_REG_SP
	/* LD B,REGISTERH */
	OpcodeDDCBMap[0x44] = instrDD__LD_B_REGH
	/* LD B,REGISTERL */
	OpcodeDDCBMap[0x45] = instrDD__LD_B_REGL
	/* LD B,(REGISTER+dd) */
	OpcodeDDCBMap[0x46] = instrDD__LD_B_iREGpDD
	/* LD C,REGISTERH */
	OpcodeDDCBMap[0x4c] = instrDD__LD_C_REGH
	/* LD C,REGISTERL */
	OpcodeDDCBMap[0x4d] = instrDD__LD_C_REGL
	/* LD C,(REGISTER+dd) */
	OpcodeDDCBMap[0x4e] = instrDD__LD_C_iREGpDD
	// 	/* LD D,REGISTERH */
	OpcodeDDCBMap[0x54] = instrDD__LD_D_REGH
	/* LD D,REGISTERL */
	OpcodeDDCBMap[0x55] = instrDD__LD_D_REGL
	/* LD D,(REGISTER+dd) */
	OpcodeDDCBMap[0x56] = instrDD__LD_D_iREGpDD
	// 	/* LD E,REGISTERH */
	OpcodeDDCBMap[0x5c] = instrDD__LD_E_REGH
	/* LD E,REGISTERL */
	OpcodeDDCBMap[0x5d] = instrDD__LD_E_REGL
	/* LD E,(REGISTER+dd) */
	OpcodeDDCBMap[0x5e] = instrDD__LD_E_iREGpDD
	// 	/* LD REGISTERH,B */
	OpcodeDDCBMap[0x60] = instrDD__LD_REGH_B
	/* LD REGISTERH,C */
	OpcodeDDCBMap[0x61] = instrDD__LD_REGH_C
	/* LD REGISTERH,D */
	OpcodeDDCBMap[0x62] = instrDD__LD_REGH_D
	/* LD REGISTERH,E */
	OpcodeDDCBMap[0x63] = instrDD__LD_REGH_E
	/* LD REGISTERH,REGISTERH */
	OpcodeDDCBMap[0x64] = instrDD__LD_REGH_REGH
	/* LD REGISTERH,REGISTERL */
	OpcodeDDCBMap[0x65] = instrDD__LD_REGH_REGL
	/* LD H,(REGISTER+dd) */
	OpcodeDDCBMap[0x66] = instrDD__LD_H_iREGpDD
	// 	/* LD REGISTERH,A */
	OpcodeDDCBMap[0x67] = instrDD__LD_REGH_A
	/* LD REGISTERL,B */
	OpcodeDDCBMap[0x68] = instrDD__LD_REGL_B
	/* LD REGISTERL,C */
	OpcodeDDCBMap[0x69] = instrDD__LD_REGL_C
	/* LD REGISTERL,D */
	OpcodeDDCBMap[0x6a] = instrDD__LD_REGL_D
	/* LD REGISTERL,E */
	OpcodeDDCBMap[0x6b] = instrDD__LD_REGL_E
	/* LD REGISTERL,REGISTERH */
	OpcodeDDCBMap[0x6c] = instrDD__LD_REGL_REGH
	/* LD REGISTERL,REGISTERL */
	OpcodeDDCBMap[0x6d] = instrDD__LD_REGL_REGL
	/* LD L,(REGISTER+dd) */
	OpcodeDDCBMap[0x6e] = instrDD__LD_L_iREGpDD
	/* LD REGISTERL,A */
	OpcodeDDCBMap[0x6f] = instrDD__LD_REGL_A
	/* LD (REGISTER+dd),B */
	OpcodeDDCBMap[0x70] = instrDD__LD_iREGpDD_r
	/* LD (REGISTER+dd),C */
	OpcodeDDCBMap[0x71] = instrDD__LD_iREGpDD_r
	/* LD (REGISTER+dd),D */
	OpcodeDDCBMap[0x72] = instrDD__LD_iREGpDD_r
	/* LD (REGISTER+dd),E */
	OpcodeDDCBMap[0x73] = instrDD__LD_iREGpDD_r
	/* LD (REGISTER+dd),H */
	OpcodeDDCBMap[0x74] = instrDD__LD_iREGpDD_r
	/* LD (REGISTER+dd),L */
	OpcodeDDCBMap[0x75] = instrDD__LD_iREGpDD_r
	/* LD (REGISTER+dd),A */
	OpcodeDDCBMap[0x77] = instrDD__LD_iREGpDD_r
	/* LD A,REGISTERH */
	OpcodeDDCBMap[0x7c] = instrDD__LD_A_REGH
	/* LD A,REGISTERL */
	OpcodeDDCBMap[0x7d] = instrDD__LD_A_REGL
	/* LD A,(REGISTER+dd) */
	OpcodeDDCBMap[0x7e] = instrDD__LD_A_iREGpDD
	/* ADD A,REGISTERH */
	OpcodeDDCBMap[0x84] = instrDD__ADD_A_REGH
	/* ADD A,REGISTERL */
	OpcodeDDCBMap[0x85] = instrDD__ADD_A_REGL
	/* ADD A,(REGISTER+dd) */
	OpcodeDDCBMap[0x86] = instrDD__ADD_A_iREGpDD
	/* ADC A,REGISTERH */
	OpcodeDDCBMap[0x8c] = instrDD__ADC_A_REGH
	/* ADC A,REGISTERL */
	OpcodeDDCBMap[0x8d] = instrDD__ADC_A_REGL
	/* ADC A,(REGISTER+dd) */
	OpcodeDDCBMap[0x8e] = instrDD__ADC_A_iREGpDD
	/* SUB A,REGISTERH */
	OpcodeDDCBMap[0x94] = instrDD__SUB_A_REGH
	/* SUB A,REGISTERL */
	OpcodeDDCBMap[0x95] = instrDD__SUB_A_REGL
	/* SUB A,(REGISTER+dd) */
	OpcodeDDCBMap[0x96] = instrDD__SUB_A_iREGpDD
	/* SBC A,REGISTERH */
	OpcodeDDCBMap[0x9c] = instrDD__SBC_A_REGH
	/* SBC A,REGISTERL */
	OpcodeDDCBMap[0x9d] = instrDD__SBC_A_REGL
	/* SBC A,(REGISTER+dd) */
	OpcodeDDCBMap[0x9e] = instrDD__SBC_A_iREGpDD
	/* AND A,REGISTERH */
	OpcodeDDCBMap[0xa4] = instrDD__AND_A_REGH
	/* AND A,REGISTERL */
	OpcodeDDCBMap[0xa5] = instrDD__AND_A_REGL
	/* AND A,(REGISTER+dd) */
	OpcodeDDCBMap[0xa6] = instrDD__AND_A_iREGpDD
	/* XOR A,REGISTERH */
	OpcodeDDCBMap[0xac] = instrDD__XOR_A_REGH
	/* XOR A,REGISTERL */
	OpcodeDDCBMap[0xad] = instrDD__XOR_A_REGL
	/* XOR A,(REGISTER+dd) */
	OpcodeDDCBMap[0xae] = instrDD__XOR_A_iREGpDD
	/* OR A,REGISTERH */
	OpcodeDDCBMap[0xb4] = instrDD__OR_A_REGH
	/* OR A,REGISTERL */
	OpcodeDDCBMap[0xb5] = instrDD__OR_A_REGL
	/* OR A,(REGISTER+dd) */
	OpcodeDDCBMap[0xb6] = instrDD__OR_A_iREGpDD
	/* CP A,REGISTERH */
	OpcodeDDCBMap[0xbc] = instrDD__CP_A_REGH
	/* CP A,REGISTERL */
	OpcodeDDCBMap[0xbd] = instrDD__CP_A_REGL
	/* CP A,(REGISTER+dd) */
	OpcodeDDCBMap[0xbe] = instrDD__CP_A_iREGpDD
	/* shift DDFDCB */
	OpcodeDDMap[0xcb] = instrDD__SHIFT_DDFDCB
	/* POP REGISTER */
	OpcodeDDCBMap[0xe1] = instrDD__POP_REG
	/* EX (SP),REGISTER */
	OpcodeDDCBMap[0xe3] = instrDD__EX_iSP_REG
	/* PUSH REGISTER */
	OpcodeDDCBMap[0xe5] = instrDD__PUSH_REG
	/* JP REGISTER */
	OpcodeDDCBMap[0xe9] = instrDD__JP_REG
	/* LD SP,REGISTER */
	OpcodeDDCBMap[0xf9] = instrDD__LD_SP_REG

	// 	// END of 0xdd shifted opcodes

}

/* ADD ix,BC */
func instrDD__ADD_REG_BC(z *Z80, opcode byte) {
	z.Tstates += 15
	//z.Memory.ContendReadNoMreq_loop(z.IR(), 1, 7)
	z.Add16(z.IX, z.BC.Get())
}

/* ADD ix,DE */
func instrDD__ADD_REG_DE(z *Z80, opcode byte) {
	z.Tstates += 15
	//z.Memory.ContendReadNoMreq_loop(z.IR(), 1, 7)
	z.Add16(z.IX, z.DE.Get())
}

/* LD ix,nnnn */
func instrDD__LD_REG_NNNN(z *Z80, opcode byte) {
	z.Tstates += 14
	val := z.Load16()
	z.IX.Set(val)
}

/* LD (nnnn),ix */
func instrDD__LD_iNNNN_REG(z *Z80, opcode byte) {
	z.Tstates += 20
	z.StoreIndex16(z.IXL, z.IXH)
	// break
}

/* INC ix */
func instrDD__INC_REG(z *Z80, opcode byte) {
	z.Tstates += 8
	//z.Memory.ContendReadNoMreq_loop(z.IR(), 1, 2)
	z.IX.Inc() //z.IncIX()
}

/* INC IXH */
func instrDD__INC_REGH(z *Z80, opcode byte) {
	z.Tstates += 8
	z.Inc(&z.IXH) // incIXH()
}

/* DEC IXH */
func instrDD__DEC_REGH(z *Z80, opcode byte) {
	z.Tstates += 8
	z.Dec(&z.IXH) //decIXH()
}

/* LD IXH,nn */
func instrDD__LD_REGH_NN(z *Z80, opcode byte) {
	z.Tstates += 14
	z.IXH = z.Load8()
}

/* ADD ix,ix */
func instrDD__ADD_REG_REG(z *Z80, opcode byte) {
	z.Tstates += 15
	// z.Memory.ContendReadNoMreq_loop(z.IR(), 1, 7)
	z.Add16(z.IX, z.IX.Get())
}

/* LD ix,(nnnn) */
func instrDD__LD_REG_iNNNN(z *Z80, opcode byte) {
	z.Tstates += 20
	z.LoadIndex16(&z.IXL, &z.IXH)
	// break
}

/* DEC ix */
func instrDD__DEC_REG(z *Z80, opcode byte) {
	z.Tstates += 10
	//z.Memory.ContendReadNoMreq_loop(z.IR(), 1, 2)
	z.IX.Dec()
}

/* INC IXL */
func instrDD__INC_REGL(z *Z80, opcode byte) {
	z.Tstates += 8
	z.Inc(&z.IXL)
}

/* DEC IXL */
func instrDD__DEC_REGL(z *Z80, opcode byte) {
	z.Tstates += 8
	z.Dec(&z.IXL)
}

/* LD IXL,nn */
func instrDD__LD_REGL_NN(z *Z80, opcode byte) {
	z.Tstates += 14
	z.IXL = z.Load8()
}

/* INC (ix+dd) */
func instrDD__INC_iREGpDD(z *Z80, opcode byte) {
	z.Tstates += 23
	var offset byte = z.Load8()
	//z.Memory.ContendReadNoMreq_loop(z.pc, 1, 5)
	var wordtemp uint16 = z.IX.Get() + uint16(signExtend(offset))
	var bytetemp byte = z.Memory.Read(wordtemp)
	//z.Memory.ContendReadNoMreq(wordtemp, 1)
	z.Inc(&bytetemp)
	z.Memory.Write(wordtemp, bytetemp)
}

/* DEC (ix+dd) */
func instrDD__DEC_iREGpDD(z *Z80, opcode byte) {
	z.Tstates += 23
	var offset byte = z.Load8()
	//z.Memory.ContendReadNoMreq_loop(z.pc, 1, 5)
	var wordtemp uint16 = z.IX.Get() + uint16(signExtend(offset))
	var bytetemp byte = z.Memory.Read(wordtemp)
	//z.Memory.ContendReadNoMreq(wordtemp, 1)
	z.Dec(&bytetemp)
	z.Memory.Write(wordtemp, bytetemp)
}

/* LD (ix+dd),nn */
func instrDD__LD_iREGpDD_NN(z *Z80, opcode byte) {
	z.Tstates += 19
	offset := z.Load8()
	value := z.Load8()
	//z.Memory.ContendReadNoMreq_loop(z.pc, 1, 2)
	z.Memory.Write(z.IX.Get()+uint16(signExtend(offset)), value)
}

/* ADD ix,SP */
func instrDD__ADD_REG_SP(z *Z80, opcode byte) {
	z.Tstates += 15
	//z.Memory.ContendReadNoMreq_loop(z.IR(), 1, 7)
	z.Add16(z.IX, z.sp)
}

/* LD B,IXH */
func instrDD__LD_B_REGH(z *Z80, opcode byte) {
	z.Tstates += 8
	z.B = z.IXH
}

/* LD B,IXL */
func instrDD__LD_B_REGL(z *Z80, opcode byte) {
	z.Tstates += 8
	z.B = z.IXL
}

/* LD B,(ix+dd) */
func instrDD__LD_B_iREGpDD(z *Z80, opcode byte) {
	z.Tstates += 19
	var offset byte = z.Load8()
	//z.Memory.ContendReadNoMreq_loop(z.pc, 1, 5)
	z.B = z.Memory.Read(z.IX.Get() + uint16(signExtend(offset)))
}

/* LD C,IXH */
func instrDD__LD_C_REGH(z *Z80, opcode byte) {
	z.Tstates += 8
	z.C = z.IXH
}

/* LD C,IXL */
func instrDD__LD_C_REGL(z *Z80, opcode byte) {
	z.Tstates += 8
	z.C = z.IXL
}

/* LD C,(ix+dd) */
func instrDD__LD_C_iREGpDD(z *Z80, opcode byte) {
	z.Tstates += 19
	var offset byte = z.Load8()
	//z.Memory.ContendReadNoMreq_loop(z.pc, 1, 5)
	z.C = z.Memory.Read(z.IX.Get() + uint16(signExtend(offset)))
}

/* LD D,IXH */
func instrDD__LD_D_REGH(z *Z80, opcode byte) {
	z.Tstates += 8
	z.D = z.IXH
}

/* LD D,IXL */
func instrDD__LD_D_REGL(z *Z80, opcode byte) {
	z.Tstates += 8
	z.D = z.IXL
}

/* LD D,(ix+dd) */
func instrDD__LD_D_iREGpDD(z *Z80, opcode byte) {
	z.Tstates += 19
	var offset byte = z.Load8()
	//z.Memory.ContendReadNoMreq_loop(z.pc, 1, 5)
	z.D = z.Memory.Read(z.IX.Get() + uint16(signExtend(offset)))
}

/* LD E,IXH */
func instrDD__LD_E_REGH(z *Z80, opcode byte) {
	z.Tstates += 8
	z.E = z.IXH
}

/* LD E,IXL */
func instrDD__LD_E_REGL(z *Z80, opcode byte) {
	z.Tstates += 8
	z.E = z.IXL
}

/* LD E,(ix+dd) */
func instrDD__LD_E_iREGpDD(z *Z80, opcode byte) {
	z.Tstates += 19
	var offset byte = z.Load8()
	//z.Memory.ContendReadNoMreq_loop(z.pc, 1, 5)
	z.E = z.Memory.Read(z.IX.Get() + uint16(signExtend(offset)))
}

/* LD IXH,B */
func instrDD__LD_REGH_B(z *Z80, opcode byte) {
	z.Tstates += 8
	z.IXH = z.B
}

/* LD IXH,C */
func instrDD__LD_REGH_C(z *Z80, opcode byte) {
	z.Tstates += 8
	z.IXH = z.C
}

/* LD IXH,D */
func instrDD__LD_REGH_D(z *Z80, opcode byte) {
	z.Tstates += 8
	z.IXH = z.D
}

/* LD IXH,E */
func instrDD__LD_REGH_E(z *Z80, opcode byte) {
	z.Tstates += 8
	z.IXH = z.E
}

/* LD IXH,IXH */
func instrDD__LD_REGH_REGH(z *Z80, opcode byte) {
	z.Tstates += 8
}

/* LD IXH,IXL */
func instrDD__LD_REGH_REGL(z *Z80, opcode byte) {
	z.Tstates += 8
	z.IXH = z.IXL
}

/* LD H,(ix+dd) */
func instrDD__LD_H_iREGpDD(z *Z80, opcode byte) {
	z.Tstates += 19
	var offset byte = z.Load8()
	//z.Memory.ContendReadNoMreq_loop(z.pc, 1, 5)
	z.H = z.Memory.Read(z.IX.Get() + uint16(signExtend(offset)))
}

/* LD IXH,A */
func instrDD__LD_REGH_A(z *Z80, opcode byte) {
	z.Tstates += 8
	z.IXH = z.A
}

/* LD IXL,B */
func instrDD__LD_REGL_B(z *Z80, opcode byte) {
	z.Tstates += 8
	z.IXL = z.B
}

/* LD IXL,C */
func instrDD__LD_REGL_C(z *Z80, opcode byte) {
	z.Tstates += 8
	z.IXL = z.C
}

/* LD IXL,D */
func instrDD__LD_REGL_D(z *Z80, opcode byte) {
	z.Tstates += 8
	z.IXL = z.D
}

/* LD IXL,E */
func instrDD__LD_REGL_E(z *Z80, opcode byte) {
	z.Tstates += 8
	z.IXL = z.E
}

/* LD IXL,IXH */
func instrDD__LD_REGL_REGH(z *Z80, opcode byte) {
	z.Tstates += 8
	z.IXL = z.IXH
}

/* LD IXL,IXL */
func instrDD__LD_REGL_REGL(z *Z80, opcode byte) {
	z.Tstates += 8
}

/* LD L,(ix+dd) */
func instrDD__LD_L_iREGpDD(z *Z80, opcode byte) {
	z.Tstates += 19
	var offset byte = z.Load8()
	//z.Memory.ContendReadNoMreq_loop(z.pc, 1, 5)
	z.L = z.Memory.Read(z.IX.Get() + uint16(signExtend(offset)))
}

/* LD IXL,A */
func instrDD__LD_REGL_A(z *Z80, opcode byte) {
	z.Tstates += 8
	z.IXL = z.A
}

/* LD (ix+dd),r */
func instrDD__LD_iREGpDD_r(z *Z80, opcode byte) {
	z.Tstates += 19
	offset := z.Load8()
	//z.Memory.ContendReadNoMreq_loop(z.pc, 1, 5)
	reg := z.GetRegisterValByte(opcode)
	z.Memory.Write(z.IX.Get()+uint16(signExtend(offset)), reg)
}

/* LD A,IXH */
func instrDD__LD_A_REGH(z *Z80, opcode byte) {
	z.Tstates += 8
	z.A = z.IXH
}

/* LD A,IXL */
func instrDD__LD_A_REGL(z *Z80, opcode byte) {
	z.Tstates += 8
	z.A = z.IXL
}

/* LD A,(ix+dd) */
func instrDD__LD_A_iREGpDD(z *Z80, opcode byte) {
	z.Tstates += 19
	var offset byte = z.Load8()
	//z.Memory.ContendReadNoMreq_loop(z.pc, 1, 5)
	z.A = z.Memory.Read(z.IX.Get() + uint16(signExtend(offset)))
}

/* ADD A,IXH */
func instrDD__ADD_A_REGH(z *Z80, opcode byte) {
	z.Tstates += 8
	z.Add(z.IXH)
}

/* ADD A,IXL */
func instrDD__ADD_A_REGL(z *Z80, opcode byte) {
	z.Tstates += 8
	z.Add(z.IXL)
}

/* ADD A,(ix+dd) */
func instrDD__ADD_A_iREGpDD(z *Z80, opcode byte) {
	z.Tstates += 19
	var offset byte = z.Load8()
	//z.Memory.ContendReadNoMreq_loop(z.pc, 1, 5)
	var bytetemp byte = z.Memory.Read(z.IX.Get() + uint16(signExtend(offset)))
	z.Add(bytetemp)
}

/* ADC A,IXH */
func instrDD__ADC_A_REGH(z *Z80, opcode byte) {
	z.Tstates += 8
	z.Adc(z.IXH)
}

/* ADC A,IXL */
func instrDD__ADC_A_REGL(z *Z80, opcode byte) {
	z.Tstates += 8
	z.Adc(z.IXL)
}

/* ADC A,(ix+dd) */
func instrDD__ADC_A_iREGpDD(z *Z80, opcode byte) {
	z.Tstates += 19
	var offset byte = z.Load8()
	//z.Memory.ContendReadNoMreq_loop(z.pc, 1, 5)
	var bytetemp byte = z.Memory.Read(z.IX.Get() + uint16(signExtend(offset)))
	z.Adc(bytetemp)
}

/* SUB A,IXH */
func instrDD__SUB_A_REGH(z *Z80, opcode byte) {
	z.Tstates += 8
	z.Sub(z.IXH)
}

/* SUB A,IXL */
func instrDD__SUB_A_REGL(z *Z80, opcode byte) {
	z.Tstates += 8
	z.Sub(z.IXL)
}

/* SUB A,(ix+dd) */
func instrDD__SUB_A_iREGpDD(z *Z80, opcode byte) {
	z.Tstates += 19
	var offset byte = z.Load8()
	//z.Memory.ContendReadNoMreq_loop(z.pc, 1, 5)
	var bytetemp byte = z.Memory.Read(z.IX.Get() + uint16(signExtend(offset)))
	z.Sub(bytetemp)
}

/* SBC A,IXH */
func instrDD__SBC_A_REGH(z *Z80, opcode byte) {
	z.Tstates += 8
	z.Sbc(z.IXH)
}

/* SBC A,IXL */
func instrDD__SBC_A_REGL(z *Z80, opcode byte) {
	z.Tstates += 8
	z.Sbc(z.IXL)
}

/* SBC A,(ix+dd) */
func instrDD__SBC_A_iREGpDD(z *Z80, opcode byte) {
	z.Tstates += 19
	var offset byte = z.Load8()
	//z.Memory.ContendReadNoMreq_loop(z.pc, 1, 5)
	var bytetemp byte = z.Memory.Read(z.IX.Get() + uint16(signExtend(offset)))
	z.Sbc(bytetemp)
}

/* AND A,IXH */
func instrDD__AND_A_REGH(z *Z80, opcode byte) {
	z.Tstates += 8
	z.And(z.IXH)
}

/* AND A,IXL */
func instrDD__AND_A_REGL(z *Z80, opcode byte) {
	z.Tstates += 8
	z.And(z.IXL)
}

/* AND A,(ix+dd) */
func instrDD__AND_A_iREGpDD(z *Z80, opcode byte) {
	z.Tstates += 19
	var offset byte = z.Load8()
	//z.Memory.ContendReadNoMreq_loop(z.pc, 1, 5)
	var bytetemp byte = z.Memory.Read(z.IX.Get() + uint16(signExtend(offset)))
	z.And(bytetemp)
}

/* XOR A,IXH */
func instrDD__XOR_A_REGH(z *Z80, opcode byte) {
	z.Tstates += 8
	z.Xor(z.IXH)
}

/* XOR A,IXL */
func instrDD__XOR_A_REGL(z *Z80, opcode byte) {
	z.Tstates += 8
	z.Xor(z.IXL)
}

/* XOR A,(ix+dd) */
func instrDD__XOR_A_iREGpDD(z *Z80, opcode byte) {
	z.Tstates += 19
	var offset byte = z.Load8()
	//z.Memory.ContendReadNoMreq_loop(z.pc, 1, 5)
	var bytetemp byte = z.Memory.Read(z.IX.Get() + uint16(signExtend(offset)))
	z.Xor(bytetemp)
}

/* OR A,IXH */
func instrDD__OR_A_REGH(z *Z80, opcode byte) {
	z.Tstates += 8
	z.Or(z.IXH)
}

/* OR A,IXL */
func instrDD__OR_A_REGL(z *Z80, opcode byte) {
	z.Tstates += 8
	z.Or(z.IXL)
}

/* OR A,(ix+dd) */
func instrDD__OR_A_iREGpDD(z *Z80, opcode byte) {
	z.Tstates += 19
	var offset byte = z.Memory.Read(z.pc)
	//z.Memory.ContendReadNoMreq_loop(z.pc, 1, 5)
	var bytetemp byte = z.Memory.Read(z.IX.Get() + uint16(signExtend(offset)))
	z.Or(bytetemp)
}

/* CP A,IXH */
func instrDD__CP_A_REGH(z *Z80, opcode byte) {
	z.Tstates += 8
	z.Cp(z.IXH)
}

/* CP A,IXL */
func instrDD__CP_A_REGL(z *Z80, opcode byte) {
	z.Tstates += 8
	z.Cp(z.IXL)
}

/* CP A,(ix+dd) */
func instrDD__CP_A_iREGpDD(z *Z80, opcode byte) {
	z.Tstates += 19
	var offset byte = z.Load8()
	//z.Memory.ContendReadNoMreq_loop(z.pc, 1, 5)
	var bytetemp byte = z.Memory.Read(z.IX.Get() + uint16(signExtend(offset)))
	z.Cp(bytetemp)
}

// /* shift DDFDCB */
func instrDD__SHIFT_DDFDCB(z *Z80, opcode byte) {
	z.Tstates += 4
	opcode2 := z.Load8()
	z.R++
	OpcodeDDCBMap[opcode2](z, opcode2)
}

/* POP ix */
func instrDD__POP_REG(z *Z80, opcode byte) {
	z.Tstates += 14
	z.IXH = z.Pop8()
	z.IXL = z.Pop8()
}

/* EX (SP),ix */
func instrDD__EX_iSP_REG(z *Z80, opcode byte) {
	z.Tstates += 23
	var bytetempl = z.Memory.Read(z.sp)
	var bytetemph = z.Memory.Read(z.sp + 1)

	//z.Memory.ContendReadNoMreq(z.SP()+1, 1)
	z.Memory.Write(z.sp+1, z.IXH)
	z.Memory.Write(z.sp, z.IXL)
	//z.Memory.ContendWriteNoMreq_loop(z.SP(), 1, 2)
	z.IXL = bytetempl
	z.IXH = bytetemph
}

/* PUSH ix */
func instrDD__PUSH_REG(z *Z80, opcode byte) {
	z.Tstates += 15
	//z.Memory.ContendReadNoMreq(z.IR(), 1)
	z.Push16(z.IX.Get())
}

/* JP ix */
func instrDD__JP_REG(z *Z80, opcode byte) {
	z.Tstates += 8
	z.pc = z.IX.Get()
}

/* LD SP,ix */
func instrDD__LD_SP_REG(z *Z80, opcode byte) {
	z.Tstates += 10
	//z.Memory.ContendReadNoMreq_loop(z.IR(), 1, 2)
	z.sp = z.IX.Get()
}
