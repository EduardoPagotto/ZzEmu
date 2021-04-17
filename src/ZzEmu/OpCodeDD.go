package ZzEmu

func initOpcodeDDMap() {
	// BEGIN of 0xdd shifted opcodes
	/* ADD REGISTER,BC */
	OpcodeDDMap[0x09] = instrDD__ADD_REG_BC
	/* ADD REGISTER,DE */
	OpcodeDDMap[0x19] = instrDD__ADD_REG_DE
	/* LD REGISTER,nnnn */
	OpcodeDDMap[0x21] = instrDD__LD_REG_NNNN
	/* LD (nnnn),REGISTER */
	OpcodeDDMap[0x22] = instrDD__LD_iNNNN_REG
	/* INC REGISTER */
	OpcodeDDMap[0x23] = instrDD__INC_REG
	/* INC REGISTERH */
	OpcodeDDMap[0x24] = instrDD__INC_REGH
	/* DEC REGISTERH */
	OpcodeDDMap[0x25] = instrDD__DEC_REGH
	/* LD REGISTERH,nn */
	OpcodeDDMap[0x26] = instrDD__LD_REGH_NN
	/* ADD REGISTER,REGISTER */
	OpcodeDDMap[0x29] = instrDD__ADD_REG_REG
	/* LD REGISTER,(nnnn) */
	OpcodeDDMap[0x2a] = instrDD__LD_REG_iNNNN
	/* DEC REGISTER */
	OpcodeDDMap[0x2b] = instrDD__DEC_REG
	/* INC REGISTERL */
	OpcodeDDMap[0x2c] = instrDD__INC_REGL
	/* DEC REGISTERL */
	OpcodeDDMap[0x2d] = instrDD__DEC_REGL
	/* LD REGISTERL,nn */
	OpcodeDDMap[0x2e] = instrDD__LD_REGL_NN
	/* INC (REGISTER+dd) */
	OpcodeDDMap[0x34] = instrDD__INC_iREGpDD
	/* DEC (REGISTER+dd) */
	OpcodeDDMap[0x35] = instrDD__DEC_iREGpDD
	/* LD (REGISTER+dd),nn */
	OpcodeDDMap[0x36] = instrDD__LD_iREGpDD_NN
	/* ADD REGISTER,SP */
	OpcodeDDMap[0x39] = instrDD__ADD_REG_SP
	/* LD B,REGISTERH */
	OpcodeDDMap[0x44] = instrDD__LD_B_REGH
	/* LD B,REGISTERL */
	OpcodeDDMap[0x45] = instrDD__LD_B_REGL
	/* LD B,(REGISTER+dd) */
	OpcodeDDMap[0x46] = instrDD__LD_B_iREGpDD
	/* LD C,REGISTERH */
	OpcodeDDMap[0x4c] = instrDD__LD_C_REGH
	/* LD C,REGISTERL */
	OpcodeDDMap[0x4d] = instrDD__LD_C_REGL
	/* LD C,(REGISTER+dd) */
	OpcodeDDMap[0x4e] = instrDD__LD_C_iREGpDD
	// 	/* LD D,REGISTERH */
	OpcodeDDMap[0x54] = instrDD__LD_D_REGH
	/* LD D,REGISTERL */
	OpcodeDDMap[0x55] = instrDD__LD_D_REGL
	/* LD D,(REGISTER+dd) */
	OpcodeDDMap[0x56] = instrDD__LD_D_iREGpDD
	// 	/* LD E,REGISTERH */
	OpcodeDDMap[0x5c] = instrDD__LD_E_REGH
	/* LD E,REGISTERL */
	OpcodeDDMap[0x5d] = instrDD__LD_E_REGL
	/* LD E,(REGISTER+dd) */
	OpcodeDDMap[0x5e] = instrDD__LD_E_iREGpDD
	// 	/* LD REGISTERH,B */
	OpcodeDDMap[0x60] = instrDD__LD_REGH_B
	/* LD REGISTERH,C */
	OpcodeDDMap[0x61] = instrDD__LD_REGH_C
	/* LD REGISTERH,D */
	OpcodeDDMap[0x62] = instrDD__LD_REGH_D
	/* LD REGISTERH,E */
	OpcodeDDMap[0x63] = instrDD__LD_REGH_E
	/* LD REGISTERH,REGISTERH */
	OpcodeDDMap[0x64] = instrDD__LD_REGH_REGH
	/* LD REGISTERH,REGISTERL */
	OpcodeDDMap[0x65] = instrDD__LD_REGH_REGL
	/* LD H,(REGISTER+dd) */
	OpcodeDDMap[0x66] = instrDD__LD_H_iREGpDD
	// 	/* LD REGISTERH,A */
	OpcodeDDMap[0x67] = instrDD__LD_REGH_A
	/* LD REGISTERL,B */
	OpcodeDDMap[0x68] = instrDD__LD_REGL_B
	/* LD REGISTERL,C */
	OpcodeDDMap[0x69] = instrDD__LD_REGL_C
	/* LD REGISTERL,D */
	OpcodeDDMap[0x6a] = instrDD__LD_REGL_D
	/* LD REGISTERL,E */
	OpcodeDDMap[0x6b] = instrDD__LD_REGL_E
	/* LD REGISTERL,REGISTERH */
	OpcodeDDMap[0x6c] = instrDD__LD_REGL_REGH
	/* LD REGISTERL,REGISTERL */
	OpcodeDDMap[0x6d] = instrDD__LD_REGL_REGL
	/* LD L,(REGISTER+dd) */
	OpcodeDDMap[0x6e] = instrDD__LD_L_iREGpDD
	/* LD REGISTERL,A */
	OpcodeDDMap[0x6f] = instrDD__LD_REGL_A
	/* LD (REGISTER+dd),B */
	OpcodeDDMap[0x70] = instrDD__LD_iREGpDD_r
	/* LD (REGISTER+dd),C */
	OpcodeDDMap[0x71] = instrDD__LD_iREGpDD_r
	/* LD (REGISTER+dd),D */
	OpcodeDDMap[0x72] = instrDD__LD_iREGpDD_r
	/* LD (REGISTER+dd),E */
	OpcodeDDMap[0x73] = instrDD__LD_iREGpDD_r
	/* LD (REGISTER+dd),H */
	OpcodeDDMap[0x74] = instrDD__LD_iREGpDD_r
	/* LD (REGISTER+dd),L */
	OpcodeDDMap[0x75] = instrDD__LD_iREGpDD_r
	/* LD (REGISTER+dd),A */
	OpcodeDDMap[0x77] = instrDD__LD_iREGpDD_r
	/* LD A,REGISTERH */
	OpcodeDDMap[0x7c] = instrDD__LD_A_REGH
	/* LD A,REGISTERL */
	OpcodeDDMap[0x7d] = instrDD__LD_A_REGL
	/* LD A,(REGISTER+dd) */
	OpcodeDDMap[0x7e] = instrDD__LD_A_iREGpDD
	/* ADD A,REGISTERH */
	OpcodeDDMap[0x84] = instrDD__ADD_A_REGH
	/* ADD A,REGISTERL */
	OpcodeDDMap[0x85] = instrDD__ADD_A_REGL
	/* ADD A,(REGISTER+dd) */
	OpcodeDDMap[0x86] = instrDD__ADD_A_iREGpDD
	/* ADC A,REGISTERH */
	OpcodeDDMap[0x8c] = instrDD__ADC_A_REGH
	/* ADC A,REGISTERL */
	OpcodeDDMap[0x8d] = instrDD__ADC_A_REGL
	/* ADC A,(REGISTER+dd) */
	OpcodeDDMap[0x8e] = instrDD__ADC_A_iREGpDD
	/* SUB A,REGISTERH */
	OpcodeDDMap[0x94] = instrDD__SUB_A_REGH
	/* SUB A,REGISTERL */
	OpcodeDDMap[0x95] = instrDD__SUB_A_REGL
	/* SUB A,(REGISTER+dd) */
	OpcodeDDMap[0x96] = instrDD__SUB_A_iREGpDD
	/* SBC A,REGISTERH */
	OpcodeDDMap[0x9c] = instrDD__SBC_A_REGH
	/* SBC A,REGISTERL */
	OpcodeDDMap[0x9d] = instrDD__SBC_A_REGL
	/* SBC A,(REGISTER+dd) */
	OpcodeDDMap[0x9e] = instrDD__SBC_A_iREGpDD
	/* AND A,REGISTERH */
	OpcodeDDMap[0xa4] = instrDD__AND_A_REGH
	/* AND A,REGISTERL */
	OpcodeDDMap[0xa5] = instrDD__AND_A_REGL
	/* AND A,(REGISTER+dd) */
	OpcodeDDMap[0xa6] = instrDD__AND_A_iREGpDD
	/* XOR A,REGISTERH */
	OpcodeDDMap[0xac] = instrDD__XOR_A_REGH
	/* XOR A,REGISTERL */
	OpcodeDDMap[0xad] = instrDD__XOR_A_REGL
	/* XOR A,(REGISTER+dd) */
	OpcodeDDMap[0xae] = instrDD__XOR_A_iREGpDD
	/* OR A,REGISTERH */
	OpcodeDDMap[0xb4] = instrDD__OR_A_REGH
	/* OR A,REGISTERL */
	OpcodeDDMap[0xb5] = instrDD__OR_A_REGL
	/* OR A,(REGISTER+dd) */
	OpcodeDDMap[0xb6] = instrDD__OR_A_iREGpDD
	/* CP A,REGISTERH */
	OpcodeDDMap[0xbc] = instrDD__CP_A_REGH
	/* CP A,REGISTERL */
	OpcodeDDMap[0xbd] = instrDD__CP_A_REGL
	/* CP A,(REGISTER+dd) */
	OpcodeDDMap[0xbe] = instrDD__CP_A_iREGpDD
	/* shift DDFDCB */
	OpcodeDDMap[0xcb] = instrDD__SHIFT_DDFDCB
	/* POP REGISTER */
	OpcodeDDMap[0xe1] = instrDD__POP_REG
	/* EX (SP),REGISTER */
	OpcodeDDMap[0xe3] = instrDD__EX_iSP_REG
	/* PUSH REGISTER */
	OpcodeDDMap[0xe5] = instrDD__PUSH_REG
	/* JP REGISTER */
	OpcodeDDMap[0xe9] = instrDD__JP_REG
	/* LD SP,REGISTER */
	OpcodeDDMap[0xf9] = instrDD__LD_SP_REG

	// END of 0xdd shifted opcodes

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
	z.LoadR(&z.IX)
}

/* LD (nnnn),ix */
func instrDD__LD_iNNNN_REG(z *Z80, opcode byte) {
	z.Tstates += 20
	z.StoreIndexR(z.IX)
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
	z.LoadIndexR(&z.IX)
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
	var bytetemp byte = z.bus.ReadMemory(wordtemp)
	//z.Memory.ContendReadNoMreq(wordtemp, 1)
	z.Inc(&bytetemp)
	z.bus.WriteMemory(wordtemp, bytetemp)
}

/* DEC (ix+dd) */
func instrDD__DEC_iREGpDD(z *Z80, opcode byte) {
	z.Tstates += 23
	var offset byte = z.Load8()
	//z.Memory.ContendReadNoMreq_loop(z.pc, 1, 5)
	var wordtemp uint16 = z.IX.Get() + uint16(signExtend(offset))
	var bytetemp byte = z.bus.ReadMemory(wordtemp)
	//z.Memory.ContendReadNoMreq(wordtemp, 1)
	z.Dec(&bytetemp)
	z.bus.WriteMemory(wordtemp, bytetemp)
}

/* LD (ix+dd),nn */
func instrDD__LD_iREGpDD_NN(z *Z80, opcode byte) {
	z.Tstates += 19
	offset := z.Load8()
	value := z.Load8()
	//z.Memory.ContendReadNoMreq_loop(z.pc, 1, 2)
	z.bus.WriteMemory(z.IX.Get()+uint16(signExtend(offset)), value)
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
	z.B = z.bus.ReadMemory(z.IX.Get() + uint16(signExtend(offset)))
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
	z.C = z.bus.ReadMemory(z.IX.Get() + uint16(signExtend(offset)))
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
	z.D = z.bus.ReadMemory(z.IX.Get() + uint16(signExtend(offset)))
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
	z.E = z.bus.ReadMemory(z.IX.Get() + uint16(signExtend(offset)))
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
	z.H = z.bus.ReadMemory(z.IX.Get() + uint16(signExtend(offset)))
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
	z.L = z.bus.ReadMemory(z.IX.Get() + uint16(signExtend(offset)))
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
	z.bus.WriteMemory(z.IX.Get()+uint16(signExtend(offset)), reg)
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
	z.A = z.bus.ReadMemory(z.IX.Get() + uint16(signExtend(offset)))
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
	var bytetemp byte = z.bus.ReadMemory(z.IX.Get() + uint16(signExtend(offset)))
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
	var bytetemp byte = z.bus.ReadMemory(z.IX.Get() + uint16(signExtend(offset)))
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
	var bytetemp byte = z.bus.ReadMemory(z.IX.Get() + uint16(signExtend(offset)))
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
	var bytetemp byte = z.bus.ReadMemory(z.IX.Get() + uint16(signExtend(offset)))
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
	var bytetemp byte = z.bus.ReadMemory(z.IX.Get() + uint16(signExtend(offset)))
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
	var bytetemp byte = z.bus.ReadMemory(z.IX.Get() + uint16(signExtend(offset)))
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
	var offset byte = z.bus.ReadMemory(z.pc)
	//z.Memory.ContendReadNoMreq_loop(z.pc, 1, 5)
	var bytetemp byte = z.bus.ReadMemory(z.IX.Get() + uint16(signExtend(offset)))
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
	var bytetemp byte = z.bus.ReadMemory(z.IX.Get() + uint16(signExtend(offset)))
	z.Cp(bytetemp)
}

/* shift DDFDCB */
func instrDD__SHIFT_DDFDCB(z *Z80, opcode byte) {
	//z.Tstates += 4
	addr := z.IX.Get() + uint16(signExtend(z.Load8()))
	opcode2 := z.Load8()
	z.R++
	OpcodeDDCBMap[opcode2](z, opcode2, addr)
}

/* POP ix */
func instrDD__POP_REG(z *Z80, opcode byte) {
	z.Tstates += 14
	z.PopR(&z.IX)
}

/* EX (SP),ix */
func instrDD__EX_iSP_REG(z *Z80, opcode byte) {
	z.Tstates += 23
	var bytetempl = z.bus.ReadMemory(z.sp)
	var bytetemph = z.bus.ReadMemory(z.sp + 1)

	//z.Memory.ContendReadNoMreq(z.SP()+1, 1)
	z.bus.WriteMemory(z.sp+1, z.IXH)
	z.bus.WriteMemory(z.sp, z.IXL)
	//z.Memory.ContendWriteNoMreq_loop(z.SP(), 1, 2)
	z.IXL = bytetempl
	z.IXH = bytetemph
}

/* PUSH ix */
func instrDD__PUSH_REG(z *Z80, opcode byte) {
	z.Tstates += 15
	//z.Memory.ContendReadNoMreq(z.IR(), 1)
	z.PushR(z.IX)
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
