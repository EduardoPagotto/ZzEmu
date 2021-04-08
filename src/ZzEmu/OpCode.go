package ZzEmu

func initOpcodes() {
	OpcodeMap[0x00] = instr__NOP
	OpcodeMap[0x01] = instr__LD_BC_NNNN
	OpcodeMap[0x02] = instr__LD_iBC_A
	OpcodeMap[0x03] = instr__INC_BC
	OpcodeMap[0x04] = instr__INC_B
	OpcodeMap[0x05] = instr__DEC_B
	OpcodeMap[0x06] = instr__LD_B_NN
	OpcodeMap[0x07] = instr__RLCA

	OpcodeMap[0x08] = instr__EX_AF_AF
	// OpcodeMap[0x09] = instr__ADD_HL_BC
	// OpcodeMap[0x0a] = instr__LD_A_iBC
	// OpcodeMap[0x0b] = instr__DEC_BC
	// OpcodeMap[0x0c] = instr__INC_C
	// OpcodeMap[0x0d] = instr__DEC_C
	// OpcodeMap[0x0e] = instr__LD_C_NN
	// OpcodeMap[0x0f] = instr__RRCA
	// OpcodeMap[0x10] = instr__DJNZ_OFFSET
	// OpcodeMap[0x11] = instr__LD_DE_NNNN
	// OpcodeMap[0x12] = instr__LD_iDE_A
	// OpcodeMap[0x13] = instr__INC_DE
	// OpcodeMap[0x14] = instr__INC_D
	// OpcodeMap[0x15] = instr__DEC_D
	// OpcodeMap[0x16] = instr__LD_D_NN
	// OpcodeMap[0x17] = instr__RLA
	// OpcodeMap[0x18] = instr__JR_OFFSET
	// OpcodeMap[0x19] = instr__ADD_HL_DE
	// OpcodeMap[0x1a] = instr__LD_A_iDE
	// OpcodeMap[0x1b] = instr__DEC_DE
	// OpcodeMap[0x1c] = instr__INC_E
	// OpcodeMap[0x1d] = instr__DEC_E
	// OpcodeMap[0x1e] = instr__LD_E_NN
	// OpcodeMap[0x1f] = instr__RRA
	// OpcodeMap[0x20] = instr__JR_NZ_OFFSET
	// OpcodeMap[0x21] = instr__LD_HL_NNNN
	// OpcodeMap[0x22] = instr__LD_iNNNN_HL
	// OpcodeMap[0x23] = instr__INC_HL
	// OpcodeMap[0x24] = instr__INC_H
	// OpcodeMap[0x25] = instr__DEC_H
	// OpcodeMap[0x26] = instr__LD_H_NN
	// OpcodeMap[0x27] = instr__DAA
	// OpcodeMap[0x28] = instr__JR_Z_OFFSET
	// OpcodeMap[0x29] = instr__ADD_HL_HL
	// OpcodeMap[0x2a] = instr__LD_HL_iNNNN
	// OpcodeMap[0x2b] = instr__DEC_HL
	// OpcodeMap[0x2c] = instr__INC_L
	// OpcodeMap[0x2d] = instr__DEC_L
	// OpcodeMap[0x2e] = instr__LD_L_NN
	// OpcodeMap[0x2f] = instr__CPL
	// OpcodeMap[0x30] = instr__JR_NC_OFFSET
	// OpcodeMap[0x31] = instr__LD_SP_NNNN
	// OpcodeMap[0x32] = instr__LD_iNNNN_A
	// OpcodeMap[0x33] = instr__INC_SP
	// OpcodeMap[0x34] = instr__INC_iHL
	// OpcodeMap[0x35] = instr__DEC_iHL
	// OpcodeMap[0x36] = instr__LD_iHL_NN
	// OpcodeMap[0x37] = instr__SCF
	// OpcodeMap[0x38] = instr__JR_C_OFFSET
	// OpcodeMap[0x39] = instr__ADD_HL_SP
	// OpcodeMap[0x3a] = instr__LD_A_iNNNN
	// OpcodeMap[0x3b] = instr__DEC_SP
	// OpcodeMap[0x3c] = instr__INC_A
	// OpcodeMap[0x3d] = instr__DEC_A
	// OpcodeMap[0x3e] = instr__LD_A_NN
	// OpcodeMap[0x3f] = instr__CCF
	// OpcodeMap[0x40] = instr__LD_B_B
	// OpcodeMap[0x41] = instr__LD_B_C
	// OpcodeMap[0x42] = instr__LD_B_D
	// OpcodeMap[0x43] = instr__LD_B_E
	// OpcodeMap[0x44] = instr__LD_B_H
	// OpcodeMap[0x45] = instr__LD_B_L
	// OpcodeMap[0x46] = instr__LD_B_iHL
	// OpcodeMap[0x47] = instr__LD_B_A
	// OpcodeMap[0x48] = instr__LD_C_B
	// OpcodeMap[0x49] = instr__LD_C_C
	// OpcodeMap[0x4a] = instr__LD_C_D
	// OpcodeMap[0x4b] = instr__LD_C_E
	// OpcodeMap[0x4c] = instr__LD_C_H
	// OpcodeMap[0x4d] = instr__LD_C_L
	// OpcodeMap[0x4e] = instr__LD_C_iHL
	// OpcodeMap[0x4f] = instr__LD_C_A
	// OpcodeMap[0x50] = instr__LD_D_B
	// OpcodeMap[0x51] = instr__LD_D_C
	// OpcodeMap[0x52] = instr__LD_D_D
	// OpcodeMap[0x53] = instr__LD_D_E
	// OpcodeMap[0x54] = instr__LD_D_H
	// OpcodeMap[0x55] = instr__LD_D_L
	// OpcodeMap[0x56] = instr__LD_D_iHL
	// OpcodeMap[0x57] = instr__LD_D_A
	// OpcodeMap[0x58] = instr__LD_E_B
	// OpcodeMap[0x59] = instr__LD_E_C
	// OpcodeMap[0x5a] = instr__LD_E_D
	// OpcodeMap[0x5b] = instr__LD_E_E
	// OpcodeMap[0x5c] = instr__LD_E_H
	// OpcodeMap[0x5d] = instr__LD_E_L
	// OpcodeMap[0x5e] = instr__LD_E_iHL
	// OpcodeMap[0x5f] = instr__LD_E_A
	// OpcodeMap[0x60] = instr__LD_H_B
	// OpcodeMap[0x61] = instr__LD_H_C
	// OpcodeMap[0x62] = instr__LD_H_D
	// OpcodeMap[0x63] = instr__LD_H_E
	// OpcodeMap[0x64] = instr__LD_H_H
	// OpcodeMap[0x65] = instr__LD_H_L
	// OpcodeMap[0x66] = instr__LD_H_iHL
	// OpcodeMap[0x67] = instr__LD_H_A
	// OpcodeMap[0x68] = instr__LD_L_B
	// OpcodeMap[0x69] = instr__LD_L_C
	// OpcodeMap[0x6a] = instr__LD_L_D
	// OpcodeMap[0x6b] = instr__LD_L_E
	// OpcodeMap[0x6c] = instr__LD_L_H
	// OpcodeMap[0x6d] = instr__LD_L_L
	// OpcodeMap[0x6e] = instr__LD_L_iHL
	// OpcodeMap[0x6f] = instr__LD_L_A
	// OpcodeMap[0x70] = instr__LD_iHL_B
	// OpcodeMap[0x71] = instr__LD_iHL_C
	// OpcodeMap[0x72] = instr__LD_iHL_D
	// OpcodeMap[0x73] = instr__LD_iHL_E
	// OpcodeMap[0x74] = instr__LD_iHL_H
	// OpcodeMap[0x75] = instr__LD_iHL_L
	// OpcodeMap[0x76] = instr__HALT
	// OpcodeMap[0x77] = instr__LD_iHL_A
	// OpcodeMap[0x78] = instr__LD_A_B
	// OpcodeMap[0x79] = instr__LD_A_C
	// OpcodeMap[0x7a] = instr__LD_A_D
	// OpcodeMap[0x7b] = instr__LD_A_E
	// OpcodeMap[0x7c] = instr__LD_A_H
	// OpcodeMap[0x7d] = instr__LD_A_L
	// OpcodeMap[0x7e] = instr__LD_A_iHL
	// OpcodeMap[0x7f] = instr__LD_A_A
	// OpcodeMap[0x80] = instr__ADD_A_B
	// OpcodeMap[0x81] = instr__ADD_A_C
	// OpcodeMap[0x82] = instr__ADD_A_D
	// OpcodeMap[0x83] = instr__ADD_A_E
	// OpcodeMap[0x84] = instr__ADD_A_H
	// OpcodeMap[0x85] = instr__ADD_A_L
	// OpcodeMap[0x86] = instr__ADD_A_iHL
	// OpcodeMap[0x87] = instr__ADD_A_A
	// OpcodeMap[0x88] = instr__ADC_A_B
	// OpcodeMap[0x89] = instr__ADC_A_C
	// OpcodeMap[0x8a] = instr__ADC_A_D
	// OpcodeMap[0x8b] = instr__ADC_A_E
	// OpcodeMap[0x8c] = instr__ADC_A_H
	// OpcodeMap[0x8d] = instr__ADC_A_L
	// OpcodeMap[0x8e] = instr__ADC_A_iHL
	// OpcodeMap[0x8f] = instr__ADC_A_A
	// OpcodeMap[0x90] = instr__SUB_A_B
	// OpcodeMap[0x91] = instr__SUB_A_C
	// OpcodeMap[0x92] = instr__SUB_A_D
	// OpcodeMap[0x93] = instr__SUB_A_E
	// OpcodeMap[0x94] = instr__SUB_A_H
	// OpcodeMap[0x95] = instr__SUB_A_L
	// OpcodeMap[0x96] = instr__SUB_A_iHL
	// OpcodeMap[0x97] = instr__SUB_A_A
	// OpcodeMap[0x98] = instr__SBC_A_B
	// OpcodeMap[0x99] = instr__SBC_A_C
	// OpcodeMap[0x9a] = instr__SBC_A_D
	// OpcodeMap[0x9b] = instr__SBC_A_E
	// OpcodeMap[0x9c] = instr__SBC_A_H
	// OpcodeMap[0x9d] = instr__SBC_A_L
	// OpcodeMap[0x9e] = instr__SBC_A_iHL
	// OpcodeMap[0x9f] = instr__SBC_A_A
	// OpcodeMap[0xa0] = instr__AND_A_B
	// OpcodeMap[0xa1] = instr__AND_A_C
	// OpcodeMap[0xa2] = instr__AND_A_D
	// OpcodeMap[0xa3] = instr__AND_A_E
	// OpcodeMap[0xa4] = instr__AND_A_H
	// OpcodeMap[0xa5] = instr__AND_A_L
	// OpcodeMap[0xa6] = instr__AND_A_iHL
	// OpcodeMap[0xa7] = instr__AND_A_A
	// OpcodeMap[0xa8] = instr__XOR_A_B
	// OpcodeMap[0xa9] = instr__XOR_A_C
	// OpcodeMap[0xaa] = instr__XOR_A_D
	// OpcodeMap[0xab] = instr__XOR_A_E
	// OpcodeMap[0xac] = instr__XOR_A_H
	// OpcodeMap[0xad] = instr__XOR_A_L
	// OpcodeMap[0xae] = instr__XOR_A_iHL
	// OpcodeMap[0xaf] = instr__XOR_A_A
	// OpcodeMap[0xb0] = instr__OR_A_B
	// OpcodeMap[0xb1] = instr__OR_A_C
	// OpcodeMap[0xb2] = instr__OR_A_D
	// OpcodeMap[0xb3] = instr__OR_A_E
	// OpcodeMap[0xb4] = instr__OR_A_H
	// OpcodeMap[0xb5] = instr__OR_A_L
	// OpcodeMap[0xb6] = instr__OR_A_iHL
	// OpcodeMap[0xb7] = instr__OR_A_A
	// OpcodeMap[0xb8] = instr__CP_B
	// OpcodeMap[0xb9] = instr__CP_C
	// OpcodeMap[0xba] = instr__CP_D
	// OpcodeMap[0xbb] = instr__CP_E
	// OpcodeMap[0xbc] = instr__CP_H
	// OpcodeMap[0xbd] = instr__CP_L
	// OpcodeMap[0xbe] = instr__CP_iHL
	// OpcodeMap[0xbf] = instr__CP_A
	// OpcodeMap[0xc0] = instr__RET_NZ
	// OpcodeMap[0xc1] = instr__POP_BC
	// OpcodeMap[0xc2] = instr__JP_NZ_NNNN
	// OpcodeMap[0xc3] = instr__JP_NNNN
	// OpcodeMap[0xc4] = instr__CALL_NZ_NNNN
	// OpcodeMap[0xc5] = instr__PUSH_BC
	// OpcodeMap[0xc6] = instr__ADD_A_NN
	// OpcodeMap[0xc7] = instr__RST_00
	// OpcodeMap[0xc8] = instr__RET_Z
	// OpcodeMap[0xc9] = instr__RET
	// OpcodeMap[0xca] = instr__JP_Z_NNNN
	OpcodeMap[0xcb] = instr__SHIFT_CB
	// OpcodeMap[0xcc] = instr__CALL_Z_NNNN
	// OpcodeMap[0xcd] = instr__CALL_NNNN
	// OpcodeMap[0xce] = instr__ADC_A_NN
	// OpcodeMap[0xcf] = instr__RST_8
	// OpcodeMap[0xd0] = instr__RET_NC
	// OpcodeMap[0xd1] = instr__POP_DE
	// OpcodeMap[0xd2] = instr__JP_NC_NNNN
	// OpcodeMap[0xd3] = instr__OUT_iNN_A
	// OpcodeMap[0xd4] = instr__CALL_NC_NNNN
	// OpcodeMap[0xd5] = instr__PUSH_DE
	// OpcodeMap[0xd6] = instr__SUB_NN
	// OpcodeMap[0xd7] = instr__RST_10
	// OpcodeMap[0xd8] = instr__RET_C
	// OpcodeMap[0xd9] = instr__EXX
	// OpcodeMap[0xda] = instr__JP_C_NNNN
	// OpcodeMap[0xdb] = instr__IN_A_iNN
	// OpcodeMap[0xdc] = instr__CALL_C_NNNN
	OpcodeMap[0xdd] = instr__SHIFT_DD
	// OpcodeMap[0xde] = instr__SBC_A_NN
	// OpcodeMap[0xdf] = instr__RST_18
	// OpcodeMap[0xe0] = instr__RET_PO
	// OpcodeMap[0xe1] = instr__POP_HL
	// OpcodeMap[0xe2] = instr__JP_PO_NNNN
	// OpcodeMap[0xe3] = instr__EX_iSP_HL
	// OpcodeMap[0xe4] = instr__CALL_PO_NNNN
	// OpcodeMap[0xe5] = instr__PUSH_HL
	// OpcodeMap[0xe6] = instr__AND_NN
	// OpcodeMap[0xe7] = instr__RST_20
	// OpcodeMap[0xe8] = instr__RET_PE
	// OpcodeMap[0xe9] = instr__JP_HL
	// OpcodeMap[0xea] = instr__JP_PE_NNNN
	// OpcodeMap[0xeb] = instr__EX_DE_HL
	// OpcodeMap[0xec] = instr__CALL_PE_NNNN
	OpcodeMap[0xed] = instr__SHIFT_ED
	// OpcodeMap[0xee] = instr__XOR_A_NN
	// OpcodeMap[0xef] = instr__RST_28
	// OpcodeMap[0xf0] = instr__RET_P
	// OpcodeMap[0xf1] = instr__POP_AF
	// OpcodeMap[0xf2] = instr__JP_P_NNNN
	// OpcodeMap[0xf3] = instr__DI
	// OpcodeMap[0xf4] = instr__CALL_P_NNNN
	// OpcodeMap[0xf5] = instr__PUSH_AF
	// OpcodeMap[0xf6] = instr__OR_NN
	// OpcodeMap[0xf7] = instr__RST_30
	// OpcodeMap[0xf8] = instr__RET_M
	// OpcodeMap[0xf9] = instr__LD_SP_HL
	// OpcodeMap[0xfa] = instr__JP_M_NNNN
	// OpcodeMap[0xfb] = instr__EI
	// OpcodeMap[0xfc] = instr__CALL_M_NNNN
	OpcodeMap[0xfd] = instr__SHIFT_FD
	// OpcodeMap[0xfe] = instr__CP_NN
	// OpcodeMap[0xff] = instr__RST_38
}

func instr__NOP(z *Z80, opcode byte) {
	z.Tstates += 4
}

func instr__LD_BC_NNNN(z *Z80, opcode byte) {
	z.Tstates += 10
	z.C = z.Memory.Read(z.pc)
	z.pc++
	z.B = z.Memory.Read(z.pc)
	z.pc++
}

func instr__LD_iBC_A(z *Z80, opcode byte) {
	z.Tstates += 7
	z.Memory.Write(z.BC.Get(), z.A)
}

func instr__INC_BC(z *Z80, opcode byte) {
	z.Tstates += 6
	z.BC.Inc()
}

func instr__INC_B(z *Z80, opcode byte) {
	// TODO testar se flagas funcinam
	z.Tstates += 4
	z.B++
	z.F = (z.F & FLAG_C) | (ternOpB(z.B == 0x80, FLAG_V, 0)) | (ternOpB((z.B&0x0f) != 0, 0, FLAG_H)) | sz53Table[z.B]
}

func instr__DEC_B(z *Z80, opcode byte) {
	z.Tstates += 4
	z.F = (z.F & FLAG_C) | (ternOpB(z.B&0x0f != 0, 0, FLAG_H)) | FLAG_N
	z.B--
	z.F |= (ternOpB(z.B == 0x7f, FLAG_V, 0)) | sz53Table[z.B]
}

func instr__LD_B_NN(z *Z80, opcode byte) {
	z.Tstates += 7
	z.B = z.Memory.Read(z.pc)
	z.pc++
}

func instr__RLCA(z *Z80, opcode byte) {
	z.Tstates += 4
	z.A = (z.A << 1) | (z.A >> 7)
	z.F = (z.F & (FLAG_P | FLAG_Z | FLAG_S)) | (z.A & (FLAG_C | FLAG_3 | FLAG_5))
}

/* EX AF,AF' */
func instr__EX_AF_AF(z *Z80, opcode byte) {
	z.Tstates += 4
	var olda, oldf = z.A, z.F
	z.A = z.A_
	z.F = z.F_
	z.A_ = olda
	z.F_ = oldf
}

// /* ADD HL,BC */
// func instr__ADD_HL_BC(z *Z80, opcode byte) {
// 	z.memory.ContendReadNoMreq_loop(z.IR(), 1, 7)
// 	z.add16(z.hl, z.BC())
// }

// /* LD A,(BC) */
// func instr__LD_A_iBC(z *Z80, opcode byte) {
// 	z.A = z.memory.ReadByte(z.BC())
// }

// /* DEC BC */
// func instr__DEC_BC(z *Z80, opcode byte) {
// 	z.memory.ContendReadNoMreq_loop(z.IR(), 1, 2)
// 	z.DecBC()
// }

// /* INC C */
// func instr__INC_C(z *Z80, opcode byte) {
// 	z.incC()
// }

// /* DEC C */
// func instr__DEC_C(z *Z80, opcode byte) {
// 	z.decC()
// }

// /* LD C,nn */
// func instr__LD_C_NN(z *Z80, opcode byte) {
// 	z.C = z.memory.ReadByte(z.PC())
// 	z.IncPC(1)
// }

// /* RRCA */
// func instr__RRCA(z *Z80, opcode byte) {
// 	z.F = (z.F & (FLAG_P | FLAG_Z | FLAG_S)) | (z.A & FLAG_C)
// 	z.A = (z.A >> 1) | (z.A << 7)
// 	z.F |= (z.A & (FLAG_3 | FLAG_5))
// }

// /* DJNZ offset */
// func instr__DJNZ_OFFSET(z *Z80, opcode byte) {
// 	z.memory.ContendReadNoMreq(z.IR(), 1)
// 	z.B--
// 	if z.B != 0 {
// 		z.jr()
// 	} else {
// 		z.memory.ContendRead(z.PC(), 3)
// 	}
// 	z.IncPC(1)
// }

// /* LD DE,nnnn */
// func instr__LD_DE_NNNN(z *Z80, opcode byte) {
// 	b1 := z.memory.ReadByte(z.PC())
// 	z.IncPC(1)
// 	b2 := z.memory.ReadByte(z.PC())
// 	z.IncPC(1)
// 	z.SetDE(joinBytes(b2, b1))
// }

// /* LD (DE),A */
// func instr__LD_iDE_A(z *Z80, opcode byte) {
// 	z.memory.WriteByte(z.DE(), z.A)
// }

// /* INC DE */
// func instr__INC_DE(z *Z80, opcode byte) {
// 	z.memory.ContendReadNoMreq_loop(z.IR(), 1, 2)
// 	z.IncDE()
// }

// /* INC D */
// func instr__INC_D(z *Z80, opcode byte) {
// 	z.incD()
// }

// /* DEC D */
// func instr__DEC_D(z *Z80, opcode byte) {
// 	z.decD()
// }

// /* LD D,nn */
// func instr__LD_D_NN(z *Z80, opcode byte) {
// 	z.D = z.memory.ReadByte(z.PC())
// 	z.IncPC(1)
// }

// /* RLA */
// func instr__RLA(z *Z80, opcode byte) {
// 	var bytetemp byte = z.A
// 	z.A = (z.A << 1) | (z.F & FLAG_C)
// 	z.F = (z.F & (FLAG_P | FLAG_Z | FLAG_S)) | (z.A & (FLAG_3 | FLAG_5)) | (bytetemp >> 7)
// }

// /* JR offset */
// func instr__JR_OFFSET(z *Z80, opcode byte) {
// 	z.jr()
// 	z.IncPC(1)
// }

// /* ADD HL,DE */
// func instr__ADD_HL_DE(z *Z80, opcode byte) {
// 	z.memory.ContendReadNoMreq_loop(z.IR(), 1, 7)
// 	z.add16(z.hl, z.DE())
// }

// /* LD A,(DE) */
// func instr__LD_A_iDE(z *Z80, opcode byte) {
// 	z.A = z.memory.ReadByte(z.DE())
// }

// /* DEC DE */
// func instr__DEC_DE(z *Z80, opcode byte) {
// 	z.memory.ContendReadNoMreq_loop(z.IR(), 1, 2)
// 	z.DecDE()
// }

// /* INC E */
// func instr__INC_E(z *Z80, opcode byte) {
// 	z.incE()
// }

// /* DEC E */
// func instr__DEC_E(z *Z80, opcode byte) {
// 	z.decE()
// }

// /* LD E,nn */
// func instr__LD_E_NN(z *Z80, opcode byte) {
// 	z.E = z.memory.ReadByte(z.PC())
// 	z.IncPC(1)
// }

// /* RRA */
// func instr__RRA(z *Z80, opcode byte) {
// 	var bytetemp byte = z.A
// 	z.A = (z.A >> 1) | (z.F << 7)
// 	z.F = (z.F & (FLAG_P | FLAG_Z | FLAG_S)) | (z.A & (FLAG_3 | FLAG_5)) | (bytetemp & FLAG_C)
// }

// /* JR NZ,offset */
// func instr__JR_NZ_OFFSET(z *Z80, opcode byte) {
// 	if (z.F & FLAG_Z) == 0 {
// 		z.jr()
// 	} else {
// 		z.memory.ContendRead(z.PC(), 3)
// 	}
// 	z.IncPC(1)
// }

// /* LD HL,nnnn */
// func instr__LD_HL_NNNN(z *Z80, opcode byte) {
// 	b1 := z.memory.ReadByte(z.PC())
// 	z.IncPC(1)
// 	b2 := z.memory.ReadByte(z.PC())
// 	z.IncPC(1)
// 	z.SetHL(joinBytes(b2, b1))
// }

// /* LD (nnnn),HL */
// func instr__LD_iNNNN_HL(z *Z80, opcode byte) {
// 	z.ld16nnrr(z.L, z.H)
// 	// break
// }

// /* INC HL */
// func instr__INC_HL(z *Z80, opcode byte) {
// 	z.memory.ContendReadNoMreq_loop(z.IR(), 1, 2)
// 	z.IncHL()
// }

// /* INC H */
// func instr__INC_H(z *Z80, opcode byte) {
// 	z.incH()
// }

// /* DEC H */
// func instr__DEC_H(z *Z80, opcode byte) {
// 	z.decH()
// }

// /* LD H,nn */
// func instr__LD_H_NN(z *Z80, opcode byte) {
// 	z.H = z.memory.ReadByte(z.PC())
// 	z.IncPC(1)
// }

// /* DAA */
// func instr__DAA(z *Z80, opcode byte) {
// 	var add, carry byte = 0, (z.F & FLAG_C)
// 	if ((z.F & FLAG_H) != 0) || ((z.A & 0x0f) > 9) {
// 		add = 6
// 	}
// 	if (carry != 0) || (z.A > 0x99) {
// 		add |= 0x60
// 	}
// 	if z.A > 0x99 {
// 		carry = FLAG_C
// 	}
// 	if (z.F & FLAG_N) != 0 {
// 		z.sub(add)
// 	} else {
// 		z.add(add)
// 	}
// 	var temp byte = byte(int(z.F) & ^(FLAG_C|FLAG_P)) | carry | parityTable[z.A]
// 	z.F = temp
// }

// /* JR Z,offset */
// func instr__JR_Z_OFFSET(z *Z80, opcode byte) {
// 	if (z.F & FLAG_Z) != 0 {
// 		z.jr()
// 	} else {
// 		z.memory.ContendRead(z.PC(), 3)
// 	}
// 	z.IncPC(1)
// }

// /* ADD HL,HL */
// func instr__ADD_HL_HL(z *Z80, opcode byte) {
// 	z.memory.ContendReadNoMreq_loop(z.IR(), 1, 7)
// 	z.add16(z.hl, z.HL())
// }

// /* LD HL,(nnnn) */
// func instr__LD_HL_iNNNN(z *Z80, opcode byte) {
// 	z.ld16rrnn(&z.L, &z.H)
// 	// break
// }

// /* DEC HL */
// func instr__DEC_HL(z *Z80, opcode byte) {
// 	z.memory.ContendReadNoMreq_loop(z.IR(), 1, 2)
// 	z.DecHL()
// }

// /* INC L */
// func instr__INC_L(z *Z80, opcode byte) {
// 	z.incL()
// }

// /* DEC L */
// func instr__DEC_L(z *Z80, opcode byte) {
// 	z.decL()
// }

// /* LD L,nn */
// func instr__LD_L_NN(z *Z80, opcode byte) {
// 	z.L = z.memory.ReadByte(z.PC())
// 	z.IncPC(1)
// }

// /* CPL */
// func instr__CPL(z *Z80, opcode byte) {
// 	z.A ^= 0xff
// 	z.F = (z.F & (FLAG_C | FLAG_P | FLAG_Z | FLAG_S)) |
// 		(z.A & (FLAG_3 | FLAG_5)) |
// 		(FLAG_N | FLAG_H)
// }

// /* JR NC,offset */
// func instr__JR_NC_OFFSET(z *Z80, opcode byte) {
// 	if (z.F & FLAG_C) == 0 {
// 		z.jr()
// 	} else {
// 		z.memory.ContendRead(z.PC(), 3)
// 	}
// 	z.IncPC(1)
// }

// /* LD SP,nnnn */
// func instr__LD_SP_NNNN(z *Z80, opcode byte) {
// 	b1 := z.memory.ReadByte(z.PC())
// 	z.IncPC(1)
// 	b2 := z.memory.ReadByte(z.PC())
// 	z.IncPC(1)
// 	z.SetSP(joinBytes(b2, b1))
// }

// /* LD (nnnn),A */
// func instr__LD_iNNNN_A(z *Z80, opcode byte) {
// 	var wordtemp uint16 = uint16(z.memory.ReadByte(z.PC()))
// 	z.IncPC(1)
// 	wordtemp |= uint16(z.memory.ReadByte(z.PC())) << 8
// 	z.IncPC(1)
// 	z.memory.WriteByte(wordtemp, z.A)
// }

// /* INC SP */
// func instr__INC_SP(z *Z80, opcode byte) {
// 	z.memory.ContendReadNoMreq_loop(z.IR(), 1, 2)
// 	z.IncSP()
// }

// /* INC (HL) */
// func instr__INC_iHL(z *Z80, opcode byte) {
// 	{
// 		var bytetemp byte = z.memory.ReadByte(z.HL())
// 		z.memory.ContendReadNoMreq(z.HL(), 1)
// 		z.inc(&bytetemp)
// 		z.memory.WriteByte(z.HL(), bytetemp)
// 	}
// }

// /* DEC (HL) */
// func instr__DEC_iHL(z *Z80, opcode byte) {
// 	{
// 		var bytetemp byte = z.memory.ReadByte(z.HL())
// 		z.memory.ContendReadNoMreq(z.HL(), 1)
// 		z.dec(&bytetemp)
// 		z.memory.WriteByte(z.HL(), bytetemp)
// 	}
// }

// /* LD (HL),nn */
// func instr__LD_iHL_NN(z *Z80, opcode byte) {
// 	z.memory.WriteByte(z.HL(), z.memory.ReadByte(z.PC()))
// 	z.IncPC(1)
// }

// /* SCF */
// func instr__SCF(z *Z80, opcode byte) {
// 	z.F = (z.F & (FLAG_P | FLAG_Z | FLAG_S)) |
// 		(z.A & (FLAG_3 | FLAG_5)) |
// 		FLAG_C
// }

// /* JR C,offset */
// func instr__JR_C_OFFSET(z *Z80, opcode byte) {
// 	if (z.F & FLAG_C) != 0 {
// 		z.jr()
// 	} else {
// 		z.memory.ContendRead(z.PC(), 3)
// 	}
// 	z.IncPC(1)
// }

// /* ADD HL,SP */
// func instr__ADD_HL_SP(z *Z80, opcode byte) {
// 	z.memory.ContendReadNoMreq_loop(z.IR(), 1, 7)
// 	z.add16(z.hl, z.SP())
// }

// /* LD A,(nnnn) */
// func instr__LD_A_iNNNN(z *Z80, opcode byte) {
// 	var wordtemp uint16 = uint16(z.memory.ReadByte(z.PC()))
// 	z.IncPC(1)
// 	wordtemp |= uint16(z.memory.ReadByte(z.PC())) << 8
// 	z.IncPC(1)
// 	z.A = z.memory.ReadByte(wordtemp)
// }

// /* DEC SP */
// func instr__DEC_SP(z *Z80, opcode byte) {
// 	z.memory.ContendReadNoMreq_loop(z.IR(), 1, 2)
// 	z.DecSP()
// }

// /* INC A */
// func instr__INC_A(z *Z80, opcode byte) {
// 	z.incA()
// }

// /* DEC A */
// func instr__DEC_A(z *Z80, opcode byte) {
// 	z.decA()
// }

// /* LD A,nn */
// func instr__LD_A_NN(z *Z80, opcode byte) {
// 	z.A = z.memory.ReadByte(z.PC())
// 	z.IncPC(1)
// }

// /* CCF */
// func instr__CCF(z *Z80, opcode byte) {
// 	z.F = (z.F & (FLAG_P | FLAG_Z | FLAG_S)) |
// 		ternOpB((z.F&FLAG_C) != 0, FLAG_H, FLAG_C) |
// 		(z.A & (FLAG_3 | FLAG_5))
// }

// /* LD B,B */
// func instr__LD_B_B(z *Z80, opcode byte) {
// }

// /* LD B,C */
// func instr__LD_B_C(z *Z80, opcode byte) {
// 	z.B = z.C
// }

// /* LD B,D */
// func instr__LD_B_D(z *Z80, opcode byte) {
// 	z.B = z.D
// }

// /* LD B,E */
// func instr__LD_B_E(z *Z80, opcode byte) {
// 	z.B = z.E
// }

// /* LD B,H */
// func instr__LD_B_H(z *Z80, opcode byte) {
// 	z.B = z.H
// }

// /* LD B,L */
// func instr__LD_B_L(z *Z80, opcode byte) {
// 	z.B = z.L
// }

// /* LD B,(HL) */
// func instr__LD_B_iHL(z *Z80, opcode byte) {
// 	z.B = z.memory.ReadByte(z.HL())
// }

// /* LD B,A */
// func instr__LD_B_A(z *Z80, opcode byte) {
// 	z.B = z.A
// }

// /* LD C,B */
// func instr__LD_C_B(z *Z80, opcode byte) {
// 	z.C = z.B
// }

// /* LD C,C */
// func instr__LD_C_C(z *Z80, opcode byte) {
// }

// /* LD C,D */
// func instr__LD_C_D(z *Z80, opcode byte) {
// 	z.C = z.D
// }

// /* LD C,E */
// func instr__LD_C_E(z *Z80, opcode byte) {
// 	z.C = z.E
// }

// /* LD C,H */
// func instr__LD_C_H(z *Z80, opcode byte) {
// 	z.C = z.H
// }

// /* LD C,L */
// func instr__LD_C_L(z *Z80, opcode byte) {
// 	z.C = z.L
// }

// /* LD C,(HL) */
// func instr__LD_C_iHL(z *Z80, opcode byte) {
// 	z.C = z.memory.ReadByte(z.HL())
// }

// /* LD C,A */
// func instr__LD_C_A(z *Z80, opcode byte) {
// 	z.C = z.A
// }

// /* LD D,B */
// func instr__LD_D_B(z *Z80, opcode byte) {
// 	z.D = z.B
// }

// /* LD D,C */
// func instr__LD_D_C(z *Z80, opcode byte) {
// 	z.D = z.C
// }

// /* LD D,D */
// func instr__LD_D_D(z *Z80, opcode byte) {
// }

// /* LD D,E */
// func instr__LD_D_E(z *Z80, opcode byte) {
// 	z.D = z.E
// }

// /* LD D,H */
// func instr__LD_D_H(z *Z80, opcode byte) {
// 	z.D = z.H
// }

// /* LD D,L */
// func instr__LD_D_L(z *Z80, opcode byte) {
// 	z.D = z.L
// }

// /* LD D,(HL) */
// func instr__LD_D_iHL(z *Z80, opcode byte) {
// 	z.D = z.memory.ReadByte(z.HL())
// }

// /* LD D,A */
// func instr__LD_D_A(z *Z80, opcode byte) {
// 	z.D = z.A
// }

// /* LD E,B */
// func instr__LD_E_B(z *Z80, opcode byte) {
// 	z.E = z.B
// }

// /* LD E,C */
// func instr__LD_E_C(z *Z80, opcode byte) {
// 	z.E = z.C
// }

// /* LD E,D */
// func instr__LD_E_D(z *Z80, opcode byte) {
// 	z.E = z.D
// }

// /* LD E,E */
// func instr__LD_E_E(z *Z80, opcode byte) {
// }

// /* LD E,H */
// func instr__LD_E_H(z *Z80, opcode byte) {
// 	z.E = z.H
// }

// /* LD E,L */
// func instr__LD_E_L(z *Z80, opcode byte) {
// 	z.E = z.L
// }

// /* LD E,(HL) */
// func instr__LD_E_iHL(z *Z80, opcode byte) {
// 	z.E = z.memory.ReadByte(z.HL())
// }

// /* LD E,A */
// func instr__LD_E_A(z *Z80, opcode byte) {
// 	z.E = z.A
// }

// /* LD H,B */
// func instr__LD_H_B(z *Z80, opcode byte) {
// 	z.H = z.B
// }

// /* LD H,C */
// func instr__LD_H_C(z *Z80, opcode byte) {
// 	z.H = z.C
// }

// /* LD H,D */
// func instr__LD_H_D(z *Z80, opcode byte) {
// 	z.H = z.D
// }

// /* LD H,E */
// func instr__LD_H_E(z *Z80, opcode byte) {
// 	z.H = z.E
// }

// /* LD H,H */
// func instr__LD_H_H(z *Z80, opcode byte) {
// }

// /* LD H,L */
// func instr__LD_H_L(z *Z80, opcode byte) {
// 	z.H = z.L
// }

// /* LD H,(HL) */
// func instr__LD_H_iHL(z *Z80, opcode byte) {
// 	z.H = z.memory.ReadByte(z.HL())
// }

// /* LD H,A */
// func instr__LD_H_A(z *Z80, opcode byte) {
// 	z.H = z.A
// }

// /* LD L,B */
// func instr__LD_L_B(z *Z80, opcode byte) {
// 	z.L = z.B
// }

// /* LD L,C */
// func instr__LD_L_C(z *Z80, opcode byte) {
// 	z.L = z.C
// }

// /* LD L,D */
// func instr__LD_L_D(z *Z80, opcode byte) {
// 	z.L = z.D
// }

// /* LD L,E */
// func instr__LD_L_E(z *Z80, opcode byte) {
// 	z.L = z.E
// }

// /* LD L,H */
// func instr__LD_L_H(z *Z80, opcode byte) {
// 	z.L = z.H
// }

// /* LD L,L */
// func instr__LD_L_L(z *Z80, opcode byte) {
// }

// /* LD L,(HL) */
// func instr__LD_L_iHL(z *Z80, opcode byte) {
// 	z.L = z.memory.ReadByte(z.HL())
// }

// /* LD L,A */
// func instr__LD_L_A(z *Z80, opcode byte) {
// 	z.L = z.A
// }

// /* LD (HL),B */
// func instr__LD_iHL_B(z *Z80, opcode byte) {
// 	z.memory.WriteByte(z.HL(), z.B)
// }

// /* LD (HL),C */
// func instr__LD_iHL_C(z *Z80, opcode byte) {
// 	z.memory.WriteByte(z.HL(), z.C)
// }

// /* LD (HL),D */
// func instr__LD_iHL_D(z *Z80, opcode byte) {
// 	z.memory.WriteByte(z.HL(), z.D)
// }

// /* LD (HL),E */
// func instr__LD_iHL_E(z *Z80, opcode byte) {
// 	z.memory.WriteByte(z.HL(), z.E)
// }

// /* LD (HL),H */
// func instr__LD_iHL_H(z *Z80, opcode byte) {
// 	z.memory.WriteByte(z.HL(), z.H)
// }

// /* LD (HL),L */
// func instr__LD_iHL_L(z *Z80, opcode byte) {
// 	z.memory.WriteByte(z.HL(), z.L)
// }

// /* HALT */
// func instr__HALT(z *Z80, opcode byte) {
// 	z.Halted = true
// 	z.DecPC(1)
// 	return
// }

// /* LD (HL),A */
// func instr__LD_iHL_A(z *Z80, opcode byte) {
// 	z.memory.WriteByte(z.HL(), z.A)
// }

// /* LD A,B */
// func instr__LD_A_B(z *Z80, opcode byte) {
// 	z.A = z.B
// }

// /* LD A,C */
// func instr__LD_A_C(z *Z80, opcode byte) {
// 	z.A = z.C
// }

// /* LD A,D */
// func instr__LD_A_D(z *Z80, opcode byte) {
// 	z.A = z.D
// }

// /* LD A,E */
// func instr__LD_A_E(z *Z80, opcode byte) {
// 	z.A = z.E
// }

// /* LD A,H */
// func instr__LD_A_H(z *Z80, opcode byte) {
// 	z.A = z.H
// }

// /* LD A,L */
// func instr__LD_A_L(z *Z80, opcode byte) {
// 	z.A = z.L
// }

// /* LD A,(HL) */
// func instr__LD_A_iHL(z *Z80, opcode byte) {
// 	z.A = z.memory.ReadByte(z.HL())
// }

// /* LD A,A */
// func instr__LD_A_A(z *Z80, opcode byte) {
// }

// /* ADD A,B */
// func instr__ADD_A_B(z *Z80, opcode byte) {
// 	z.add(z.B)
// }

// /* ADD A,C */
// func instr__ADD_A_C(z *Z80, opcode byte) {
// 	z.add(z.C)
// }

// /* ADD A,D */
// func instr__ADD_A_D(z *Z80, opcode byte) {
// 	z.add(z.D)
// }

// /* ADD A,E */
// func instr__ADD_A_E(z *Z80, opcode byte) {
// 	z.add(z.E)
// }

// /* ADD A,H */
// func instr__ADD_A_H(z *Z80, opcode byte) {
// 	z.add(z.H)
// }

// /* ADD A,L */
// func instr__ADD_A_L(z *Z80, opcode byte) {
// 	z.add(z.L)
// }

// /* ADD A,(HL) */
// func instr__ADD_A_iHL(z *Z80, opcode byte) {
// 	var bytetemp byte = z.memory.ReadByte(z.HL())

// 	z.add(bytetemp)
// }

// /* ADD A,A */
// func instr__ADD_A_A(z *Z80, opcode byte) {
// 	z.add(z.A)
// }

// /* ADC A,B */
// func instr__ADC_A_B(z *Z80, opcode byte) {
// 	z.adc(z.B)
// }

// /* ADC A,C */
// func instr__ADC_A_C(z *Z80, opcode byte) {
// 	z.adc(z.C)
// }

// /* ADC A,D */
// func instr__ADC_A_D(z *Z80, opcode byte) {
// 	z.adc(z.D)
// }

// /* ADC A,E */
// func instr__ADC_A_E(z *Z80, opcode byte) {
// 	z.adc(z.E)
// }

// /* ADC A,H */
// func instr__ADC_A_H(z *Z80, opcode byte) {
// 	z.adc(z.H)
// }

// /* ADC A,L */
// func instr__ADC_A_L(z *Z80, opcode byte) {
// 	z.adc(z.L)
// }

// /* ADC A,(HL) */
// func instr__ADC_A_iHL(z *Z80, opcode byte) {
// 	var bytetemp byte = z.memory.ReadByte(z.HL())

// 	z.adc(bytetemp)
// }

// /* ADC A,A */
// func instr__ADC_A_A(z *Z80, opcode byte) {
// 	z.adc(z.A)
// }

// /* SUB A,B */
// func instr__SUB_A_B(z *Z80, opcode byte) {
// 	z.sub(z.B)
// }

// /* SUB A,C */
// func instr__SUB_A_C(z *Z80, opcode byte) {
// 	z.sub(z.C)
// }

// /* SUB A,D */
// func instr__SUB_A_D(z *Z80, opcode byte) {
// 	z.sub(z.D)
// }

// /* SUB A,E */
// func instr__SUB_A_E(z *Z80, opcode byte) {
// 	z.sub(z.E)
// }

// /* SUB A,H */
// func instr__SUB_A_H(z *Z80, opcode byte) {
// 	z.sub(z.H)
// }

// /* SUB A,L */
// func instr__SUB_A_L(z *Z80, opcode byte) {
// 	z.sub(z.L)
// }

// /* SUB A,(HL) */
// func instr__SUB_A_iHL(z *Z80, opcode byte) {
// 	var bytetemp byte = z.memory.ReadByte(z.HL())

// 	z.sub(bytetemp)
// }

// /* SUB A,A */
// func instr__SUB_A_A(z *Z80, opcode byte) {
// 	z.sub(z.A)
// }

// /* SBC A,B */
// func instr__SBC_A_B(z *Z80, opcode byte) {
// 	z.sbc(z.B)
// }

// /* SBC A,C */
// func instr__SBC_A_C(z *Z80, opcode byte) {
// 	z.sbc(z.C)
// }

// /* SBC A,D */
// func instr__SBC_A_D(z *Z80, opcode byte) {
// 	z.sbc(z.D)
// }

// /* SBC A,E */
// func instr__SBC_A_E(z *Z80, opcode byte) {
// 	z.sbc(z.E)
// }

// /* SBC A,H */
// func instr__SBC_A_H(z *Z80, opcode byte) {
// 	z.sbc(z.H)
// }

// /* SBC A,L */
// func instr__SBC_A_L(z *Z80, opcode byte) {
// 	z.sbc(z.L)
// }

// /* SBC A,(HL) */
// func instr__SBC_A_iHL(z *Z80, opcode byte) {
// 	var bytetemp byte = z.memory.ReadByte(z.HL())

// 	z.sbc(bytetemp)
// }

// /* SBC A,A */
// func instr__SBC_A_A(z *Z80, opcode byte) {
// 	z.sbc(z.A)
// }

// /* AND A,B */
// func instr__AND_A_B(z *Z80, opcode byte) {
// 	z.and(z.B)
// }

// /* AND A,C */
// func instr__AND_A_C(z *Z80, opcode byte) {
// 	z.and(z.C)
// }

// /* AND A,D */
// func instr__AND_A_D(z *Z80, opcode byte) {
// 	z.and(z.D)
// }

// /* AND A,E */
// func instr__AND_A_E(z *Z80, opcode byte) {
// 	z.and(z.E)
// }

// /* AND A,H */
// func instr__AND_A_H(z *Z80, opcode byte) {
// 	z.and(z.H)
// }

// /* AND A,L */
// func instr__AND_A_L(z *Z80, opcode byte) {
// 	z.and(z.L)
// }

// /* AND A,(HL) */
// func instr__AND_A_iHL(z *Z80, opcode byte) {
// 	var bytetemp byte = z.memory.ReadByte(z.HL())

// 	z.and(bytetemp)
// }

// /* AND A,A */
// func instr__AND_A_A(z *Z80, opcode byte) {
// 	z.and(z.A)
// }

// /* XOR A,B */
// func instr__XOR_A_B(z *Z80, opcode byte) {
// 	z.xor(z.B)
// }

// /* XOR A,C */
// func instr__XOR_A_C(z *Z80, opcode byte) {
// 	z.xor(z.C)
// }

// /* XOR A,D */
// func instr__XOR_A_D(z *Z80, opcode byte) {
// 	z.xor(z.D)
// }

// /* XOR A,E */
// func instr__XOR_A_E(z *Z80, opcode byte) {
// 	z.xor(z.E)
// }

// /* XOR A,H */
// func instr__XOR_A_H(z *Z80, opcode byte) {
// 	z.xor(z.H)
// }

// /* XOR A,L */
// func instr__XOR_A_L(z *Z80, opcode byte) {
// 	z.xor(z.L)
// }

// /* XOR A,(HL) */
// func instr__XOR_A_iHL(z *Z80, opcode byte) {
// 	var bytetemp byte = z.memory.ReadByte(z.HL())

// 	z.xor(bytetemp)
// }

// /* XOR A,A */
// func instr__XOR_A_A(z *Z80, opcode byte) {
// 	z.xor(z.A)
// }

// /* OR A,B */
// func instr__OR_A_B(z *Z80, opcode byte) {
// 	z.or(z.B)
// }

// /* OR A,C */
// func instr__OR_A_C(z *Z80, opcode byte) {
// 	z.or(z.C)
// }

// /* OR A,D */
// func instr__OR_A_D(z *Z80, opcode byte) {
// 	z.or(z.D)
// }

// /* OR A,E */
// func instr__OR_A_E(z *Z80, opcode byte) {
// 	z.or(z.E)
// }

// /* OR A,H */
// func instr__OR_A_H(z *Z80, opcode byte) {
// 	z.or(z.H)
// }

// /* OR A,L */
// func instr__OR_A_L(z *Z80, opcode byte) {
// 	z.or(z.L)
// }

// /* OR A,(HL) */
// func instr__OR_A_iHL(z *Z80, opcode byte) {
// 	var bytetemp byte = z.memory.ReadByte(z.HL())

// 	z.or(bytetemp)
// }

// /* OR A,A */
// func instr__OR_A_A(z *Z80, opcode byte) {
// 	z.or(z.A)
// }

// /* CP B */
// func instr__CP_B(z *Z80, opcode byte) {
// 	z.cp(z.B)
// }

// /* CP C */
// func instr__CP_C(z *Z80, opcode byte) {
// 	z.cp(z.C)
// }

// /* CP D */
// func instr__CP_D(z *Z80, opcode byte) {
// 	z.cp(z.D)
// }

// /* CP E */
// func instr__CP_E(z *Z80, opcode byte) {
// 	z.cp(z.E)
// }

// /* CP H */
// func instr__CP_H(z *Z80, opcode byte) {
// 	z.cp(z.H)
// }

// /* CP L */
// func instr__CP_L(z *Z80, opcode byte) {
// 	z.cp(z.L)
// }

// /* CP (HL) */
// func instr__CP_iHL(z *Z80, opcode byte) {
// 	var bytetemp byte = z.memory.ReadByte(z.HL())

// 	z.cp(bytetemp)
// }

// /* CP A */
// func instr__CP_A(z *Z80, opcode byte) {
// 	z.cp(z.A)
// }

// /* RET NZ */
// func instr__RET_NZ(z *Z80, opcode byte) {
// 	z.memory.ContendReadNoMreq(z.IR(), 1)
// 	if !((z.F & FLAG_Z) != 0) {
// 		z.ret()
// 	}
// }

// /* POP BC */
// func instr__POP_BC(z *Z80, opcode byte) {
// 	z.C, z.B = z.pop16()
// }

// /* JP NZ,nnnn */
// func instr__JP_NZ_NNNN(z *Z80, opcode byte) {
// 	if (z.F & FLAG_Z) == 0 {
// 		z.jp()
// 	} else {
// 		z.memory.ContendRead(z.PC(), 3)
// 		z.memory.ContendRead(z.PC()+1, 3)
// 		z.IncPC(2)
// 	}
// }

// /* JP nnnn */
// func instr__JP_NNNN(z *Z80, opcode byte) {
// 	z.jp()
// }

// /* CALL NZ,nnnn */
// func instr__CALL_NZ_NNNN(z *Z80, opcode byte) {
// 	if (z.F & FLAG_Z) == 0 {
// 		z.call()
// 	} else {
// 		z.memory.ContendRead(z.PC(), 3)
// 		z.memory.ContendRead(z.PC()+1, 3)
// 		z.IncPC(2)
// 	}
// }

// /* PUSH BC */
// func instr__PUSH_BC(z *Z80, opcode byte) {
// 	z.memory.ContendReadNoMreq(z.IR(), 1)
// 	z.push16(z.C, z.B)
// }

// /* ADD A,nn */
// func instr__ADD_A_NN(z *Z80, opcode byte) {
// 	var bytetemp byte = z.memory.ReadByte(z.PC())
// 	z.IncPC(1)
// 	z.add(bytetemp)
// }

// /* RST 00 */
// func instr__RST_00(z *Z80, opcode byte) {
// 	z.memory.ContendReadNoMreq(z.IR(), 1)
// 	z.rst(0x00)
// }

// /* RET Z */
// func instr__RET_Z(z *Z80, opcode byte) {
// 	z.memory.ContendReadNoMreq(z.IR(), 1)
// 	if (z.F & FLAG_Z) != 0 {
// 		z.ret()
// 	}
// }

// /* RET */
// func instr__RET(z *Z80, opcode byte) {
// 	z.ret()
// }

// /* JP Z,nnnn */
// func instr__JP_Z_NNNN(z *Z80, opcode byte) {
// 	if (z.F & FLAG_Z) != 0 {
// 		z.jp()
// 	} else {
// 		z.memory.ContendRead(z.PC(), 3)
// 		z.memory.ContendRead(z.PC()+1, 3)
// 		z.IncPC(2)
// 	}
// }

/* shift CB */
func instr__SHIFT_CB(z *Z80, opcode byte) {
	OpcodeCBMap[opcode](z, opcode)
}

// /* CALL Z,nnnn */
// func instr__CALL_Z_NNNN(z *Z80, opcode byte) {
// 	if (z.F & FLAG_Z) != 0 {
// 		z.call()
// 	} else {
// 		z.memory.ContendRead(z.PC(), 3)
// 		z.memory.ContendRead(z.PC()+1, 3)
// 		z.IncPC(2)
// 	}
// }

// /* CALL nnnn */
// func instr__CALL_NNNN(z *Z80, opcode byte) {
// 	z.call()
// }

// /* ADC A,nn */
// func instr__ADC_A_NN(z *Z80, opcode byte) {
// 	var bytetemp byte = z.memory.ReadByte(z.PC())
// 	z.IncPC(1)
// 	z.adc(bytetemp)
// }

// /* RST 8 */
// func instr__RST_8(z *Z80, opcode byte) {
// 	z.memory.ContendReadNoMreq(z.IR(), 1)
// 	z.rst(0x8)
// }

// /* RET NC */
// func instr__RET_NC(z *Z80, opcode byte) {
// 	z.memory.ContendReadNoMreq(z.IR(), 1)
// 	if !((z.F & FLAG_C) != 0) {
// 		z.ret()
// 	}
// }

// /* POP DE */
// func instr__POP_DE(z *Z80, opcode byte) {
// 	z.E, z.D = z.pop16()
// }

// /* JP NC,nnnn */
// func instr__JP_NC_NNNN(z *Z80, opcode byte) {
// 	if (z.F & FLAG_C) == 0 {
// 		z.jp()
// 	} else {
// 		z.memory.ContendRead(z.PC(), 3)
// 		z.memory.ContendRead(z.PC()+1, 3)
// 		z.IncPC(2)
// 	}
// }

// /* OUT (nn),A */
// func instr__OUT_iNN_A(z *Z80, opcode byte) {
// 	var outtemp uint16 = uint16(z.memory.ReadByte(z.PC())) + (uint16(z.A) << 8)
// 	z.IncPC(1)
// 	z.writePort(outtemp, z.A)
// }

// /* CALL NC,nnnn */
// func instr__CALL_NC_NNNN(z *Z80, opcode byte) {
// 	if (z.F & FLAG_C) == 0 {
// 		z.call()
// 	} else {
// 		z.memory.ContendRead(z.PC(), 3)
// 		z.memory.ContendRead(z.PC()+1, 3)
// 		z.IncPC(2)
// 	}
// }

// /* PUSH DE */
// func instr__PUSH_DE(z *Z80, opcode byte) {
// 	z.memory.ContendReadNoMreq(z.IR(), 1)
// 	z.push16(z.E, z.D)
// }

// /* SUB nn */
// func instr__SUB_NN(z *Z80, opcode byte) {
// 	var bytetemp byte = z.memory.ReadByte(z.PC())
// 	z.IncPC(1)
// 	z.sub(bytetemp)
// }

// /* RST 10 */
// func instr__RST_10(z *Z80, opcode byte) {
// 	z.memory.ContendReadNoMreq(z.IR(), 1)
// 	z.rst(0x10)
// }

// /* RET C */
// func instr__RET_C(z *Z80, opcode byte) {
// 	z.memory.ContendReadNoMreq(z.IR(), 1)
// 	if (z.F & FLAG_C) != 0 {
// 		z.ret()
// 	}
// }

// /* EXX */
// func instr__EXX(z *Z80, opcode byte) {
// 	var wordtemp uint16 = z.BC()
// 	z.SetBC(z.BC_())
// 	z.SetBC_(wordtemp)

// 	wordtemp = z.DE()
// 	z.SetDE(z.DE_())
// 	z.SetDE_(wordtemp)

// 	wordtemp = z.HL()
// 	z.SetHL(z.HL_())
// 	z.SetHL_(wordtemp)
// }

// /* JP C,nnnn */
// func instr__JP_C_NNNN(z *Z80, opcode byte) {
// 	if (z.F & FLAG_C) != 0 {
// 		z.jp()
// 	} else {
// 		z.memory.ContendRead(z.PC(), 3)
// 		z.memory.ContendRead(z.PC()+1, 3)
// 		z.IncPC(2)
// 	}
// }

// /* IN A,(nn) */
// func instr__IN_A_iNN(z *Z80, opcode byte) {
// 	var intemp uint16 = uint16(z.memory.ReadByte(z.PC())) + (uint16(z.A) << 8)
// 	z.IncPC(1)
// 	z.A = z.readPort(intemp)
// }

// /* CALL C,nnnn */
// func instr__CALL_C_NNNN(z *Z80, opcode byte) {
// 	if (z.F & FLAG_C) != 0 {
// 		z.call()
// 	} else {
// 		z.memory.ContendRead(z.PC(), 3)
// 		z.memory.ContendRead(z.PC()+1, 3)
// 		z.IncPC(2)
// 	}
// }

// /* shift DD */
func instr__SHIFT_DD(z *Z80, opcode byte) {
	OpcodeDDMap[opcode](z, opcode)
}

// /* SBC A,nn */
// func instr__SBC_A_NN(z *Z80, opcode byte) {
// 	var bytetemp byte = z.memory.ReadByte(z.PC())
// 	z.IncPC(1)
// 	z.sbc(bytetemp)
// }

// /* RST 18 */
// func instr__RST_18(z *Z80, opcode byte) {
// 	z.memory.ContendReadNoMreq(z.IR(), 1)
// 	z.rst(0x18)
// }

// /* RET PO */
// func instr__RET_PO(z *Z80, opcode byte) {
// 	z.memory.ContendReadNoMreq(z.IR(), 1)
// 	if !((z.F & FLAG_P) != 0) {
// 		z.ret()
// 	}
// }

// /* POP HL */
// func instr__POP_HL(z *Z80, opcode byte) {
// 	z.L, z.H = z.pop16()
// }

// /* JP PO,nnnn */
// func instr__JP_PO_NNNN(z *Z80, opcode byte) {
// 	if (z.F & FLAG_P) == 0 {
// 		z.jp()
// 	} else {
// 		z.memory.ContendRead(z.PC(), 3)
// 		z.memory.ContendRead(z.PC()+1, 3)
// 		z.IncPC(2)
// 	}
// }

// /* EX (SP),HL */
// func instr__EX_iSP_HL(z *Z80, opcode byte) {
// 	var bytetempl = z.memory.ReadByte(z.SP())
// 	var bytetemph = z.memory.ReadByte(z.SP() + 1)
// 	z.memory.ContendReadNoMreq(z.SP()+1, 1)
// 	z.memory.WriteByte(z.SP()+1, z.H)
// 	z.memory.WriteByte(z.SP(), z.L)
// 	z.memory.ContendWriteNoMreq_loop(z.SP(), 1, 2)
// 	z.L = bytetempl
// 	z.H = bytetemph
// }

// /* CALL PO,nnnn */
// func instr__CALL_PO_NNNN(z *Z80, opcode byte) {
// 	if (z.F & FLAG_P) == 0 {
// 		z.call()
// 	} else {
// 		z.memory.ContendRead(z.PC(), 3)
// 		z.memory.ContendRead(z.PC()+1, 3)
// 		z.IncPC(2)
// 	}
// }

// /* PUSH HL */
// func instr__PUSH_HL(z *Z80, opcode byte) {
// 	z.memory.ContendReadNoMreq(z.IR(), 1)
// 	z.push16(z.L, z.H)
// }

// /* AND nn */
// func instr__AND_NN(z *Z80, opcode byte) {
// 	var bytetemp byte = z.memory.ReadByte(z.PC())
// 	z.IncPC(1)
// 	z.and(bytetemp)
// }

// /* RST 20 */
// func instr__RST_20(z *Z80, opcode byte) {
// 	z.memory.ContendReadNoMreq(z.IR(), 1)
// 	z.rst(0x20)
// }

// /* RET PE */
// func instr__RET_PE(z *Z80, opcode byte) {
// 	z.memory.ContendReadNoMreq(z.IR(), 1)
// 	if (z.F & FLAG_P) != 0 {
// 		z.ret()
// 	}
// }

// /* JP HL */
// func instr__JP_HL(z *Z80, opcode byte) {
// 	z.SetPC(z.HL()) /* NB: NOT INDIRECT! */
// }

// /* JP PE,nnnn */
// func instr__JP_PE_NNNN(z *Z80, opcode byte) {
// 	if (z.F & FLAG_P) != 0 {
// 		z.jp()
// 	} else {
// 		z.memory.ContendRead(z.PC(), 3)
// 		z.memory.ContendRead(z.PC()+1, 3)
// 		z.IncPC(2)
// 	}
// }

// /* EX DE,HL */
// func instr__EX_DE_HL(z *Z80, opcode byte) {
// 	var wordtemp uint16 = z.DE()
// 	z.SetDE(z.HL())
// 	z.SetHL(wordtemp)
// }

// /* CALL PE,nnnn */
// func instr__CALL_PE_NNNN(z *Z80, opcode byte) {
// 	if (z.F & FLAG_P) != 0 {
// 		z.call()
// 	} else {
// 		z.memory.ContendRead(z.PC(), 3)
// 		z.memory.ContendRead(z.PC()+1, 3)
// 		z.IncPC(2)
// 	}
// }

// /* shift ED */
func instr__SHIFT_ED(z *Z80, opcode byte) {
	OpcodeEDMap[opcode](z, opcode)
}

// /* XOR A,nn */
// func instr__XOR_A_NN(z *Z80, opcode byte) {
// 	var bytetemp byte = z.memory.ReadByte(z.PC())
// 	z.IncPC(1)
// 	z.xor(bytetemp)
// }

// /* RST 28 */
// func instr__RST_28(z *Z80, opcode byte) {
// 	z.memory.ContendReadNoMreq(z.IR(), 1)
// 	z.rst(0x28)
// }

// /* RET P */
// func instr__RET_P(z *Z80, opcode byte) {
// 	z.memory.ContendReadNoMreq(z.IR(), 1)
// 	if !((z.F & FLAG_S) != 0) {
// 		z.ret()
// 	}
// }

// /* POP AF */
// func instr__POP_AF(z *Z80, opcode byte) {
// 	z.F, z.A = z.pop16()
// }

// /* JP P,nnnn */
// func instr__JP_P_NNNN(z *Z80, opcode byte) {
// 	if (z.F & FLAG_S) == 0 {
// 		z.jp()
// 	} else {
// 		z.memory.ContendRead(z.PC(), 3)
// 		z.memory.ContendRead(z.PC()+1, 3)
// 		z.IncPC(2)
// 	}
// }

// /* DI */
// func instr__DI(z *Z80, opcode byte) {
// 	z.IFF1, z.IFF2 = 0, 0
// }

// /* CALL P,nnnn */
// func instr__CALL_P_NNNN(z *Z80, opcode byte) {
// 	if (z.F & FLAG_S) == 0 {
// 		z.call()
// 	} else {
// 		z.memory.ContendRead(z.PC(), 3)
// 		z.memory.ContendRead(z.PC()+1, 3)
// 		z.IncPC(2)
// 	}
// }

// /* PUSH AF */
// func instr__PUSH_AF(z *Z80, opcode byte) {
// 	z.memory.ContendReadNoMreq(z.IR(), 1)
// 	z.push16(z.F, z.A)
// }

// /* OR nn */
// func instr__OR_NN(z *Z80, opcode byte) {
// 	var bytetemp byte = z.memory.ReadByte(z.PC())
// 	z.IncPC(1)
// 	z.or(bytetemp)
// }

// /* RST 30 */
// func instr__RST_30(z *Z80, opcode byte) {
// 	z.memory.ContendReadNoMreq(z.IR(), 1)
// 	z.rst(0x30)
// }

// /* RET M */
// func instr__RET_M(z *Z80, opcode byte) {
// 	z.memory.ContendReadNoMreq(z.IR(), 1)
// 	if (z.F & FLAG_S) != 0 {
// 		z.ret()
// 	}
// }

// /* LD SP,HL */
// func instr__LD_SP_HL(z *Z80, opcode byte) {
// 	z.memory.ContendReadNoMreq_loop(z.IR(), 1, 2)
// 	z.SetSP(z.HL())
// }

// /* JP M,nnnn */
// func instr__JP_M_NNNN(z *Z80, opcode byte) {
// 	if (z.F & FLAG_S) != 0 {
// 		z.jp()
// 	} else {
// 		z.memory.ContendRead(z.PC(), 3)
// 		z.memory.ContendRead(z.PC()+1, 3)
// 		z.IncPC(2)
// 	}
// }

// /* EI */
// func instr__EI(z *Z80, opcode byte) {
// 	/* Interrupts are not accepted immediately after an EI, but are
// 	   accepted after the next instruction */
// 	z.IFF1, z.IFF2 = 1, 1
// 	z.interruptsEnabledAt = int(z.Tstates)
// 	// eventAdd(z.Tstates + 1, z80InterruptEvent)
// }

// /* CALL M,nnnn */
// func instr__CALL_M_NNNN(z *Z80, opcode byte) {
// 	if (z.F & FLAG_S) != 0 {
// 		z.call()
// 	} else {
// 		z.memory.ContendRead(z.PC(), 3)
// 		z.memory.ContendRead(z.PC()+1, 3)
// 		z.IncPC(2)
// 	}
// }

// /* shift FD */
func instr__SHIFT_FD(z *Z80, opcode byte) {
	OpcodeDFMap[opcode](z, opcode)
}

// /* CP nn */
// func instr__CP_NN(z *Z80, opcode byte) {
// 	var bytetemp byte = z.memory.ReadByte(z.PC())
// 	z.IncPC(1)
// 	z.cp(bytetemp)
// }

// /* RST 38 */
// func instr__RST_38(z *Z80, opcode byte) {
// 	z.memory.ContendReadNoMreq(z.IR(), 1)
// 	z.rst(0x38)
// }

// func (z *Z80, opcode byte) call(info *stepInfo) {

// 	addLo := z.Memory.Read(z.pc)
// 	z.pc++
// 	addHi := z.Memory.Read(z.pc)
// 	z.pc++

// 	z.Push16(z.pc)
// 	z.pc = joinBytes(addHi, addLo)
// }

// func (z *Z80, opcode byte) jp(info *stepInfo) {
// 	jptemp := z.pc
// 	pcl := z.Memory.Read(jptemp)
// 	jptemp++
// 	pch := z.Memory.Read(jptemp)
// 	z.pc = joinBytes(pch, pcl)
// }

// func (z *Z80, opcode byte) jr(info *stepInfo) {
// 	var jrtemp int16 = signExtend(z.Memory.Read(z.pc))
// 	z.pc += uint16(jrtemp)
// }

// func (z *Z80, opcode byte) ret(info *stepInfo) {
// 	z.pc = z.Pop16()
// }

// func (z *Z80, opcode byte) rst(info *stepInfo) {
// 	z.Push16(z.pc)
// 	z.pc = info.address
// }
