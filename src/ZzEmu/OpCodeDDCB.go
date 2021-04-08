package ZzEmu

func initOpCodesDDCB() {

	// 	// BEGIN of 0xddfdcb shifted opcodes
	// 	/* LD B,RLC (REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x00] = instrDDCB__LD_B_RLC_iREGpDD
	// 	/* LD C,RLC (REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x01] = instrDDCB__LD_C_RLC_iREGpDD
	// 	/* LD D,RLC (REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x02] = instrDDCB__LD_D_RLC_iREGpDD
	// 	/* LD E,RLC (REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x03] = instrDDCB__LD_E_RLC_iREGpDD
	// 	/* LD H,RLC (REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x04] = instrDDCB__LD_H_RLC_iREGpDD
	// 	/* LD L,RLC (REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x05] = instrDDCB__LD_L_RLC_iREGpDD
	// 	/* RLC (REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x06] = instrDDCB__RLC_iREGpDD
	// 	/* LD A,RLC (REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x07] = instrDDCB__LD_A_RLC_iREGpDD
	// 	/* LD B,RRC (REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x08] = instrDDCB__LD_B_RRC_iREGpDD
	// 	/* LD C,RRC (REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x09] = instrDDCB__LD_C_RRC_iREGpDD
	// 	/* LD D,RRC (REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x0a] = instrDDCB__LD_D_RRC_iREGpDD
	// 	/* LD E,RRC (REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x0b] = instrDDCB__LD_E_RRC_iREGpDD
	// 	/* LD H,RRC (REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x0c] = instrDDCB__LD_H_RRC_iREGpDD
	// 	/* LD L,RRC (REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x0d] = instrDDCB__LD_L_RRC_iREGpDD
	// 	/* RRC (REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x0e] = instrDDCB__RRC_iREGpDD
	// 	/* LD A,RRC (REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x0f] = instrDDCB__LD_A_RRC_iREGpDD
	// 	/* LD B,RL (REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x10] = instrDDCB__LD_B_RL_iREGpDD
	// 	/* LD C,RL (REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x11] = instrDDCB__LD_C_RL_iREGpDD
	// 	/* LD D,RL (REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x12] = instrDDCB__LD_D_RL_iREGpDD
	// 	/* LD E,RL (REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x13] = instrDDCB__LD_E_RL_iREGpDD
	// 	/* LD H,RL (REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x14] = instrDDCB__LD_H_RL_iREGpDD
	// 	/* LD L,RL (REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x15] = instrDDCB__LD_L_RL_iREGpDD
	// 	/* RL (REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x16] = instrDDCB__RL_iREGpDD
	// 	/* LD A,RL (REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x17] = instrDDCB__LD_A_RL_iREGpDD
	// 	/* LD B,RR (REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x18] = instrDDCB__LD_B_RR_iREGpDD
	// 	/* LD C,RR (REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x19] = instrDDCB__LD_C_RR_iREGpDD
	// 	/* LD D,RR (REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x1a] = instrDDCB__LD_D_RR_iREGpDD
	// 	/* LD E,RR (REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x1b] = instrDDCB__LD_E_RR_iREGpDD
	// 	/* LD H,RR (REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x1c] = instrDDCB__LD_H_RR_iREGpDD
	// 	/* LD L,RR (REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x1d] = instrDDCB__LD_L_RR_iREGpDD
	// 	/* RR (REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x1e] = instrDDCB__RR_iREGpDD
	// 	/* LD A,RR (REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x1f] = instrDDCB__LD_A_RR_iREGpDD
	// 	/* LD B,SLA (REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x20] = instrDDCB__LD_B_SLA_iREGpDD
	// 	/* LD C,SLA (REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x21] = instrDDCB__LD_C_SLA_iREGpDD
	// 	/* LD D,SLA (REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x22] = instrDDCB__LD_D_SLA_iREGpDD
	// 	/* LD E,SLA (REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x23] = instrDDCB__LD_E_SLA_iREGpDD
	// 	/* LD H,SLA (REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x24] = instrDDCB__LD_H_SLA_iREGpDD
	// 	/* LD L,SLA (REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x25] = instrDDCB__LD_L_SLA_iREGpDD
	// 	/* SLA (REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x26] = instrDDCB__SLA_iREGpDD
	// 	/* LD A,SLA (REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x27] = instrDDCB__LD_A_SLA_iREGpDD
	// 	/* LD B,SRA (REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x28] = instrDDCB__LD_B_SRA_iREGpDD
	// 	/* LD C,SRA (REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x29] = instrDDCB__LD_C_SRA_iREGpDD
	// 	/* LD D,SRA (REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x2a] = instrDDCB__LD_D_SRA_iREGpDD
	// 	/* LD E,SRA (REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x2b] = instrDDCB__LD_E_SRA_iREGpDD
	// 	/* LD H,SRA (REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x2c] = instrDDCB__LD_H_SRA_iREGpDD
	// 	/* LD L,SRA (REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x2d] = instrDDCB__LD_L_SRA_iREGpDD
	// 	/* SRA (REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x2e] = instrDDCB__SRA_iREGpDD
	// 	/* LD A,SRA (REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x2f] = instrDDCB__LD_A_SRA_iREGpDD
	// 	/* LD B,SLL (REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x30] = instrDDCB__LD_B_SLL_iREGpDD
	// 	/* LD C,SLL (REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x31] = instrDDCB__LD_C_SLL_iREGpDD
	// 	/* LD D,SLL (REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x32] = instrDDCB__LD_D_SLL_iREGpDD
	// 	/* LD E,SLL (REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x33] = instrDDCB__LD_E_SLL_iREGpDD
	// 	/* LD H,SLL (REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x34] = instrDDCB__LD_H_SLL_iREGpDD
	// 	/* LD L,SLL (REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x35] = instrDDCB__LD_L_SLL_iREGpDD
	// 	/* SLL (REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x36] = instrDDCB__SLL_iREGpDD
	// 	/* LD A,SLL (REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x37] = instrDDCB__LD_A_SLL_iREGpDD
	// 	/* LD B,SRL (REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x38] = instrDDCB__LD_B_SRL_iREGpDD
	// 	/* LD C,SRL (REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x39] = instrDDCB__LD_C_SRL_iREGpDD
	// 	/* LD D,SRL (REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x3a] = instrDDCB__LD_D_SRL_iREGpDD
	// 	/* LD E,SRL (REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x3b] = instrDDCB__LD_E_SRL_iREGpDD
	// 	/* LD H,SRL (REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x3c] = instrDDCB__LD_H_SRL_iREGpDD
	// 	/* LD L,SRL (REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x3d] = instrDDCB__LD_L_SRL_iREGpDD
	// 	/* SRL (REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x3e] = instrDDCB__SRL_iREGpDD
	// 	/* LD A,SRL (REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x3f] = instrDDCB__LD_A_SRL_iREGpDD
	// 	/* BIT 0,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x47] = instrDDCB__BIT_0_iREGpDD
	// 	// Fallthrough cases
	// 	OpcodesMap[SHIFT_0xDDCB+0x40] = OpcodesMap[SHIFT_0xDDCB+0x47]
	// 	OpcodesMap[SHIFT_0xDDCB+0x41] = OpcodesMap[SHIFT_0xDDCB+0x47]
	// 	OpcodesMap[SHIFT_0xDDCB+0x42] = OpcodesMap[SHIFT_0xDDCB+0x47]
	// 	OpcodesMap[SHIFT_0xDDCB+0x43] = OpcodesMap[SHIFT_0xDDCB+0x47]
	// 	OpcodesMap[SHIFT_0xDDCB+0x44] = OpcodesMap[SHIFT_0xDDCB+0x47]
	// 	OpcodesMap[SHIFT_0xDDCB+0x45] = OpcodesMap[SHIFT_0xDDCB+0x47]
	// 	OpcodesMap[SHIFT_0xDDCB+0x46] = OpcodesMap[SHIFT_0xDDCB+0x47]
	// 	/* BIT 1,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x4f] = instrDDCB__BIT_1_iREGpDD
	// 	// Fallthrough cases
	// 	OpcodesMap[SHIFT_0xDDCB+0x48] = OpcodesMap[SHIFT_0xDDCB+0x4f]
	// 	OpcodesMap[SHIFT_0xDDCB+0x49] = OpcodesMap[SHIFT_0xDDCB+0x4f]
	// 	OpcodesMap[SHIFT_0xDDCB+0x4a] = OpcodesMap[SHIFT_0xDDCB+0x4f]
	// 	OpcodesMap[SHIFT_0xDDCB+0x4b] = OpcodesMap[SHIFT_0xDDCB+0x4f]
	// 	OpcodesMap[SHIFT_0xDDCB+0x4c] = OpcodesMap[SHIFT_0xDDCB+0x4f]
	// 	OpcodesMap[SHIFT_0xDDCB+0x4d] = OpcodesMap[SHIFT_0xDDCB+0x4f]
	// 	OpcodesMap[SHIFT_0xDDCB+0x4e] = OpcodesMap[SHIFT_0xDDCB+0x4f]
	// 	/* BIT 2,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x57] = instrDDCB__BIT_2_iREGpDD
	// 	// Fallthrough cases
	// 	OpcodesMap[SHIFT_0xDDCB+0x50] = OpcodesMap[SHIFT_0xDDCB+0x57]
	// 	OpcodesMap[SHIFT_0xDDCB+0x51] = OpcodesMap[SHIFT_0xDDCB+0x57]
	// 	OpcodesMap[SHIFT_0xDDCB+0x52] = OpcodesMap[SHIFT_0xDDCB+0x57]
	// 	OpcodesMap[SHIFT_0xDDCB+0x53] = OpcodesMap[SHIFT_0xDDCB+0x57]
	// 	OpcodesMap[SHIFT_0xDDCB+0x54] = OpcodesMap[SHIFT_0xDDCB+0x57]
	// 	OpcodesMap[SHIFT_0xDDCB+0x55] = OpcodesMap[SHIFT_0xDDCB+0x57]
	// 	OpcodesMap[SHIFT_0xDDCB+0x56] = OpcodesMap[SHIFT_0xDDCB+0x57]
	// 	/* BIT 3,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x5f] = instrDDCB__BIT_3_iREGpDD
	// 	// Fallthrough cases
	// 	OpcodesMap[SHIFT_0xDDCB+0x58] = OpcodesMap[SHIFT_0xDDCB+0x5f]
	// 	OpcodesMap[SHIFT_0xDDCB+0x59] = OpcodesMap[SHIFT_0xDDCB+0x5f]
	// 	OpcodesMap[SHIFT_0xDDCB+0x5a] = OpcodesMap[SHIFT_0xDDCB+0x5f]
	// 	OpcodesMap[SHIFT_0xDDCB+0x5b] = OpcodesMap[SHIFT_0xDDCB+0x5f]
	// 	OpcodesMap[SHIFT_0xDDCB+0x5c] = OpcodesMap[SHIFT_0xDDCB+0x5f]
	// 	OpcodesMap[SHIFT_0xDDCB+0x5d] = OpcodesMap[SHIFT_0xDDCB+0x5f]
	// 	OpcodesMap[SHIFT_0xDDCB+0x5e] = OpcodesMap[SHIFT_0xDDCB+0x5f]
	// 	/* BIT 4,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x67] = instrDDCB__BIT_4_iREGpDD
	// 	// Fallthrough cases
	// 	OpcodesMap[SHIFT_0xDDCB+0x60] = OpcodesMap[SHIFT_0xDDCB+0x67]
	// 	OpcodesMap[SHIFT_0xDDCB+0x61] = OpcodesMap[SHIFT_0xDDCB+0x67]
	// 	OpcodesMap[SHIFT_0xDDCB+0x62] = OpcodesMap[SHIFT_0xDDCB+0x67]
	// 	OpcodesMap[SHIFT_0xDDCB+0x63] = OpcodesMap[SHIFT_0xDDCB+0x67]
	// 	OpcodesMap[SHIFT_0xDDCB+0x64] = OpcodesMap[SHIFT_0xDDCB+0x67]
	// 	OpcodesMap[SHIFT_0xDDCB+0x65] = OpcodesMap[SHIFT_0xDDCB+0x67]
	// 	OpcodesMap[SHIFT_0xDDCB+0x66] = OpcodesMap[SHIFT_0xDDCB+0x67]
	// 	/* BIT 5,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x6f] = instrDDCB__BIT_5_iREGpDD
	// 	// Fallthrough cases
	// 	OpcodesMap[SHIFT_0xDDCB+0x68] = OpcodesMap[SHIFT_0xDDCB+0x6f]
	// 	OpcodesMap[SHIFT_0xDDCB+0x69] = OpcodesMap[SHIFT_0xDDCB+0x6f]
	// 	OpcodesMap[SHIFT_0xDDCB+0x6a] = OpcodesMap[SHIFT_0xDDCB+0x6f]
	// 	OpcodesMap[SHIFT_0xDDCB+0x6b] = OpcodesMap[SHIFT_0xDDCB+0x6f]
	// 	OpcodesMap[SHIFT_0xDDCB+0x6c] = OpcodesMap[SHIFT_0xDDCB+0x6f]
	// 	OpcodesMap[SHIFT_0xDDCB+0x6d] = OpcodesMap[SHIFT_0xDDCB+0x6f]
	// 	OpcodesMap[SHIFT_0xDDCB+0x6e] = OpcodesMap[SHIFT_0xDDCB+0x6f]
	// 	/* BIT 6,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x77] = instrDDCB__BIT_6_iREGpDD
	// 	// Fallthrough cases
	// 	OpcodesMap[SHIFT_0xDDCB+0x70] = OpcodesMap[SHIFT_0xDDCB+0x77]
	// 	OpcodesMap[SHIFT_0xDDCB+0x71] = OpcodesMap[SHIFT_0xDDCB+0x77]
	// 	OpcodesMap[SHIFT_0xDDCB+0x72] = OpcodesMap[SHIFT_0xDDCB+0x77]
	// 	OpcodesMap[SHIFT_0xDDCB+0x73] = OpcodesMap[SHIFT_0xDDCB+0x77]
	// 	OpcodesMap[SHIFT_0xDDCB+0x74] = OpcodesMap[SHIFT_0xDDCB+0x77]
	// 	OpcodesMap[SHIFT_0xDDCB+0x75] = OpcodesMap[SHIFT_0xDDCB+0x77]
	// 	OpcodesMap[SHIFT_0xDDCB+0x76] = OpcodesMap[SHIFT_0xDDCB+0x77]
	// 	/* BIT 7,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x7f] = instrDDCB__BIT_7_iREGpDD
	// 	// Fallthrough cases
	// 	OpcodesMap[SHIFT_0xDDCB+0x78] = OpcodesMap[SHIFT_0xDDCB+0x7f]
	// 	OpcodesMap[SHIFT_0xDDCB+0x79] = OpcodesMap[SHIFT_0xDDCB+0x7f]
	// 	OpcodesMap[SHIFT_0xDDCB+0x7a] = OpcodesMap[SHIFT_0xDDCB+0x7f]
	// 	OpcodesMap[SHIFT_0xDDCB+0x7b] = OpcodesMap[SHIFT_0xDDCB+0x7f]
	// 	OpcodesMap[SHIFT_0xDDCB+0x7c] = OpcodesMap[SHIFT_0xDDCB+0x7f]
	// 	OpcodesMap[SHIFT_0xDDCB+0x7d] = OpcodesMap[SHIFT_0xDDCB+0x7f]
	// 	OpcodesMap[SHIFT_0xDDCB+0x7e] = OpcodesMap[SHIFT_0xDDCB+0x7f]
	// 	/* LD B,RES 0,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x80] = instrDDCB__LD_B_RES_0_iREGpDD
	// 	/* LD C,RES 0,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x81] = instrDDCB__LD_C_RES_0_iREGpDD
	// 	/* LD D,RES 0,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x82] = instrDDCB__LD_D_RES_0_iREGpDD
	// 	/* LD E,RES 0,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x83] = instrDDCB__LD_E_RES_0_iREGpDD
	// 	/* LD H,RES 0,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x84] = instrDDCB__LD_H_RES_0_iREGpDD
	// 	/* LD L,RES 0,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x85] = instrDDCB__LD_L_RES_0_iREGpDD
	// 	/* RES 0,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x86] = instrDDCB__RES_0_iREGpDD
	// 	/* LD A,RES 0,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x87] = instrDDCB__LD_A_RES_0_iREGpDD
	// 	/* LD B,RES 1,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x88] = instrDDCB__LD_B_RES_1_iREGpDD
	// 	/* LD C,RES 1,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x89] = instrDDCB__LD_C_RES_1_iREGpDD
	// 	/* LD D,RES 1,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x8a] = instrDDCB__LD_D_RES_1_iREGpDD
	// 	/* LD E,RES 1,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x8b] = instrDDCB__LD_E_RES_1_iREGpDD
	// 	/* LD H,RES 1,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x8c] = instrDDCB__LD_H_RES_1_iREGpDD
	// 	/* LD L,RES 1,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x8d] = instrDDCB__LD_L_RES_1_iREGpDD
	// 	/* RES 1,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x8e] = instrDDCB__RES_1_iREGpDD
	// 	/* LD A,RES 1,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x8f] = instrDDCB__LD_A_RES_1_iREGpDD
	// 	/* LD B,RES 2,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x90] = instrDDCB__LD_B_RES_2_iREGpDD
	// 	/* LD C,RES 2,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x91] = instrDDCB__LD_C_RES_2_iREGpDD
	// 	/* LD D,RES 2,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x92] = instrDDCB__LD_D_RES_2_iREGpDD
	// 	/* LD E,RES 2,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x93] = instrDDCB__LD_E_RES_2_iREGpDD
	// 	/* LD H,RES 2,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x94] = instrDDCB__LD_H_RES_2_iREGpDD
	// 	/* LD L,RES 2,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x95] = instrDDCB__LD_L_RES_2_iREGpDD
	// 	/* RES 2,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x96] = instrDDCB__RES_2_iREGpDD
	// 	/* LD A,RES 2,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x97] = instrDDCB__LD_A_RES_2_iREGpDD
	// 	/* LD B,RES 3,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x98] = instrDDCB__LD_B_RES_3_iREGpDD
	// 	/* LD C,RES 3,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x99] = instrDDCB__LD_C_RES_3_iREGpDD
	// 	/* LD D,RES 3,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x9a] = instrDDCB__LD_D_RES_3_iREGpDD
	// 	/* LD E,RES 3,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x9b] = instrDDCB__LD_E_RES_3_iREGpDD
	// 	/* LD H,RES 3,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x9c] = instrDDCB__LD_H_RES_3_iREGpDD
	// 	/* LD L,RES 3,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x9d] = instrDDCB__LD_L_RES_3_iREGpDD
	// 	/* RES 3,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x9e] = instrDDCB__RES_3_iREGpDD
	// 	/* LD A,RES 3,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0x9f] = instrDDCB__LD_A_RES_3_iREGpDD
	// 	/* LD B,RES 4,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xa0] = instrDDCB__LD_B_RES_4_iREGpDD
	// 	/* LD C,RES 4,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xa1] = instrDDCB__LD_C_RES_4_iREGpDD
	// 	/* LD D,RES 4,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xa2] = instrDDCB__LD_D_RES_4_iREGpDD
	// 	/* LD E,RES 4,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xa3] = instrDDCB__LD_E_RES_4_iREGpDD
	// 	/* LD H,RES 4,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xa4] = instrDDCB__LD_H_RES_4_iREGpDD
	// 	/* LD L,RES 4,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xa5] = instrDDCB__LD_L_RES_4_iREGpDD
	// 	/* RES 4,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xa6] = instrDDCB__RES_4_iREGpDD
	// 	/* LD A,RES 4,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xa7] = instrDDCB__LD_A_RES_4_iREGpDD
	// 	/* LD B,RES 5,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xa8] = instrDDCB__LD_B_RES_5_iREGpDD
	// 	/* LD C,RES 5,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xa9] = instrDDCB__LD_C_RES_5_iREGpDD
	// 	/* LD D,RES 5,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xaa] = instrDDCB__LD_D_RES_5_iREGpDD
	// 	/* LD E,RES 5,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xab] = instrDDCB__LD_E_RES_5_iREGpDD
	// 	/* LD H,RES 5,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xac] = instrDDCB__LD_H_RES_5_iREGpDD
	// 	/* LD L,RES 5,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xad] = instrDDCB__LD_L_RES_5_iREGpDD
	// 	/* RES 5,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xae] = instrDDCB__RES_5_iREGpDD
	// 	/* LD A,RES 5,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xaf] = instrDDCB__LD_A_RES_5_iREGpDD
	// 	/* LD B,RES 6,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xb0] = instrDDCB__LD_B_RES_6_iREGpDD
	// 	/* LD C,RES 6,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xb1] = instrDDCB__LD_C_RES_6_iREGpDD
	// 	/* LD D,RES 6,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xb2] = instrDDCB__LD_D_RES_6_iREGpDD
	// 	/* LD E,RES 6,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xb3] = instrDDCB__LD_E_RES_6_iREGpDD
	// 	/* LD H,RES 6,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xb4] = instrDDCB__LD_H_RES_6_iREGpDD
	// 	/* LD L,RES 6,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xb5] = instrDDCB__LD_L_RES_6_iREGpDD
	// 	/* RES 6,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xb6] = instrDDCB__RES_6_iREGpDD
	// 	/* LD A,RES 6,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xb7] = instrDDCB__LD_A_RES_6_iREGpDD
	// 	/* LD B,RES 7,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xb8] = instrDDCB__LD_B_RES_7_iREGpDD
	// 	/* LD C,RES 7,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xb9] = instrDDCB__LD_C_RES_7_iREGpDD
	// 	/* LD D,RES 7,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xba] = instrDDCB__LD_D_RES_7_iREGpDD
	// 	/* LD E,RES 7,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xbb] = instrDDCB__LD_E_RES_7_iREGpDD
	// 	/* LD H,RES 7,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xbc] = instrDDCB__LD_H_RES_7_iREGpDD
	// 	/* LD L,RES 7,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xbd] = instrDDCB__LD_L_RES_7_iREGpDD
	// 	/* RES 7,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xbe] = instrDDCB__RES_7_iREGpDD
	// 	/* LD A,RES 7,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xbf] = instrDDCB__LD_A_RES_7_iREGpDD
	// 	/* LD B,SET 0,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xc0] = instrDDCB__LD_B_SET_0_iREGpDD
	// 	/* LD C,SET 0,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xc1] = instrDDCB__LD_C_SET_0_iREGpDD
	// 	/* LD D,SET 0,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xc2] = instrDDCB__LD_D_SET_0_iREGpDD
	// 	/* LD E,SET 0,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xc3] = instrDDCB__LD_E_SET_0_iREGpDD
	// 	/* LD H,SET 0,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xc4] = instrDDCB__LD_H_SET_0_iREGpDD
	// 	/* LD L,SET 0,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xc5] = instrDDCB__LD_L_SET_0_iREGpDD
	// 	/* SET 0,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xc6] = instrDDCB__SET_0_iREGpDD
	// 	/* LD A,SET 0,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xc7] = instrDDCB__LD_A_SET_0_iREGpDD
	// 	/* LD B,SET 1,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xc8] = instrDDCB__LD_B_SET_1_iREGpDD
	// 	/* LD C,SET 1,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xc9] = instrDDCB__LD_C_SET_1_iREGpDD
	// 	/* LD D,SET 1,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xca] = instrDDCB__LD_D_SET_1_iREGpDD
	// 	/* LD E,SET 1,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xcb] = instrDDCB__LD_E_SET_1_iREGpDD
	// 	/* LD H,SET 1,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xcc] = instrDDCB__LD_H_SET_1_iREGpDD
	// 	/* LD L,SET 1,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xcd] = instrDDCB__LD_L_SET_1_iREGpDD
	// 	/* SET 1,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xce] = instrDDCB__SET_1_iREGpDD
	// 	/* LD A,SET 1,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xcf] = instrDDCB__LD_A_SET_1_iREGpDD
	// 	/* LD B,SET 2,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xd0] = instrDDCB__LD_B_SET_2_iREGpDD
	// 	/* LD C,SET 2,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xd1] = instrDDCB__LD_C_SET_2_iREGpDD
	// 	/* LD D,SET 2,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xd2] = instrDDCB__LD_D_SET_2_iREGpDD
	// 	/* LD E,SET 2,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xd3] = instrDDCB__LD_E_SET_2_iREGpDD
	// 	/* LD H,SET 2,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xd4] = instrDDCB__LD_H_SET_2_iREGpDD
	// 	/* LD L,SET 2,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xd5] = instrDDCB__LD_L_SET_2_iREGpDD
	// 	/* SET 2,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xd6] = instrDDCB__SET_2_iREGpDD
	// 	/* LD A,SET 2,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xd7] = instrDDCB__LD_A_SET_2_iREGpDD
	// 	/* LD B,SET 3,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xd8] = instrDDCB__LD_B_SET_3_iREGpDD
	// 	/* LD C,SET 3,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xd9] = instrDDCB__LD_C_SET_3_iREGpDD
	// 	/* LD D,SET 3,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xda] = instrDDCB__LD_D_SET_3_iREGpDD
	// 	/* LD E,SET 3,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xdb] = instrDDCB__LD_E_SET_3_iREGpDD
	// 	/* LD H,SET 3,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xdc] = instrDDCB__LD_H_SET_3_iREGpDD
	// 	/* LD L,SET 3,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xdd] = instrDDCB__LD_L_SET_3_iREGpDD
	// 	/* SET 3,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xde] = instrDDCB__SET_3_iREGpDD
	// 	/* LD A,SET 3,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xdf] = instrDDCB__LD_A_SET_3_iREGpDD
	// 	/* LD B,SET 4,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xe0] = instrDDCB__LD_B_SET_4_iREGpDD
	// 	/* LD C,SET 4,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xe1] = instrDDCB__LD_C_SET_4_iREGpDD
	// 	/* LD D,SET 4,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xe2] = instrDDCB__LD_D_SET_4_iREGpDD
	// 	/* LD E,SET 4,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xe3] = instrDDCB__LD_E_SET_4_iREGpDD
	// 	/* LD H,SET 4,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xe4] = instrDDCB__LD_H_SET_4_iREGpDD
	// 	/* LD L,SET 4,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xe5] = instrDDCB__LD_L_SET_4_iREGpDD
	// 	/* SET 4,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xe6] = instrDDCB__SET_4_iREGpDD
	// 	/* LD A,SET 4,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xe7] = instrDDCB__LD_A_SET_4_iREGpDD
	// 	/* LD B,SET 5,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xe8] = instrDDCB__LD_B_SET_5_iREGpDD
	// 	/* LD C,SET 5,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xe9] = instrDDCB__LD_C_SET_5_iREGpDD
	// 	/* LD D,SET 5,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xea] = instrDDCB__LD_D_SET_5_iREGpDD
	// 	/* LD E,SET 5,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xeb] = instrDDCB__LD_E_SET_5_iREGpDD
	// 	/* LD H,SET 5,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xec] = instrDDCB__LD_H_SET_5_iREGpDD
	// 	/* LD L,SET 5,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xed] = instrDDCB__LD_L_SET_5_iREGpDD
	// 	/* SET 5,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xee] = instrDDCB__SET_5_iREGpDD
	// 	/* LD A,SET 5,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xef] = instrDDCB__LD_A_SET_5_iREGpDD
	// 	/* LD B,SET 6,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xf0] = instrDDCB__LD_B_SET_6_iREGpDD
	// 	/* LD C,SET 6,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xf1] = instrDDCB__LD_C_SET_6_iREGpDD
	// 	/* LD D,SET 6,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xf2] = instrDDCB__LD_D_SET_6_iREGpDD
	// 	/* LD E,SET 6,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xf3] = instrDDCB__LD_E_SET_6_iREGpDD
	// 	/* LD H,SET 6,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xf4] = instrDDCB__LD_H_SET_6_iREGpDD
	// 	/* LD L,SET 6,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xf5] = instrDDCB__LD_L_SET_6_iREGpDD
	// 	/* SET 6,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xf6] = instrDDCB__SET_6_iREGpDD
	// 	/* LD A,SET 6,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xf7] = instrDDCB__LD_A_SET_6_iREGpDD
	// 	/* LD B,SET 7,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xf8] = instrDDCB__LD_B_SET_7_iREGpDD
	// 	/* LD C,SET 7,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xf9] = instrDDCB__LD_C_SET_7_iREGpDD
	// 	/* LD D,SET 7,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xfa] = instrDDCB__LD_D_SET_7_iREGpDD
	// 	/* LD E,SET 7,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xfb] = instrDDCB__LD_E_SET_7_iREGpDD
	// 	/* LD H,SET 7,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xfc] = instrDDCB__LD_H_SET_7_iREGpDD
	// 	/* LD L,SET 7,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xfd] = instrDDCB__LD_L_SET_7_iREGpDD
	// 	/* SET 7,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xfe] = instrDDCB__SET_7_iREGpDD
	// 	/* LD A,SET 7,(REGISTER+dd) */
	// 	OpcodesMap[SHIFT_0xDDCB+0xff] = instrDDCB__LD_A_SET_7_iREGpDD

}

// //--

// /* LD B,RLC (REGISTER+dd) */
// func instrDDCB__LD_B_RLC_iREGpDD(z80 *Z80) {
// 	z80.B = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.B = z80.rlc(z80.B)
// 	z80.memory.WriteByte(z80.tempaddr, z80.B)
// }

// /* LD C,RLC (REGISTER+dd) */
// func instrDDCB__LD_C_RLC_iREGpDD(z80 *Z80) {
// 	z80.C = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.C = z80.rlc(z80.C)
// 	z80.memory.WriteByte(z80.tempaddr, z80.C)
// }

// /* LD D,RLC (REGISTER+dd) */
// func instrDDCB__LD_D_RLC_iREGpDD(z80 *Z80) {
// 	z80.D = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.D = z80.rlc(z80.D)
// 	z80.memory.WriteByte(z80.tempaddr, z80.D)
// }

// /* LD E,RLC (REGISTER+dd) */
// func instrDDCB__LD_E_RLC_iREGpDD(z80 *Z80) {
// 	z80.E = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.E = z80.rlc(z80.E)
// 	z80.memory.WriteByte(z80.tempaddr, z80.E)
// }

// /* LD H,RLC (REGISTER+dd) */
// func instrDDCB__LD_H_RLC_iREGpDD(z80 *Z80) {
// 	z80.H = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.H = z80.rlc(z80.H)
// 	z80.memory.WriteByte(z80.tempaddr, z80.H)
// }

// /* LD L,RLC (REGISTER+dd) */
// func instrDDCB__LD_L_RLC_iREGpDD(z80 *Z80) {
// 	z80.L = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.L = z80.rlc(z80.L)
// 	z80.memory.WriteByte(z80.tempaddr, z80.L)
// }

// /* RLC (REGISTER+dd) */
// func instrDDCB__RLC_iREGpDD(z80 *Z80) {
// 	var bytetemp byte = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	bytetemp = z80.rlc(bytetemp)
// 	z80.memory.WriteByte(z80.tempaddr, bytetemp)
// }

// /* LD A,RLC (REGISTER+dd) */
// func instrDDCB__LD_A_RLC_iREGpDD(z80 *Z80) {
// 	z80.A = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.A = z80.rlc(z80.A)
// 	z80.memory.WriteByte(z80.tempaddr, z80.A)
// }

// /* LD B,RRC (REGISTER+dd) */
// func instrDDCB__LD_B_RRC_iREGpDD(z80 *Z80) {
// 	z80.B = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.B = z80.rrc(z80.B)
// 	z80.memory.WriteByte(z80.tempaddr, z80.B)
// }

// /* LD C,RRC (REGISTER+dd) */
// func instrDDCB__LD_C_RRC_iREGpDD(z80 *Z80) {
// 	z80.C = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.C = z80.rrc(z80.C)
// 	z80.memory.WriteByte(z80.tempaddr, z80.C)
// }

// /* LD D,RRC (REGISTER+dd) */
// func instrDDCB__LD_D_RRC_iREGpDD(z80 *Z80) {
// 	z80.D = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.D = z80.rrc(z80.D)
// 	z80.memory.WriteByte(z80.tempaddr, z80.D)
// }

// /* LD E,RRC (REGISTER+dd) */
// func instrDDCB__LD_E_RRC_iREGpDD(z80 *Z80) {
// 	z80.E = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.E = z80.rrc(z80.E)
// 	z80.memory.WriteByte(z80.tempaddr, z80.E)
// }

// /* LD H,RRC (REGISTER+dd) */
// func instrDDCB__LD_H_RRC_iREGpDD(z80 *Z80) {
// 	z80.H = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.H = z80.rrc(z80.H)
// 	z80.memory.WriteByte(z80.tempaddr, z80.H)
// }

// /* LD L,RRC (REGISTER+dd) */
// func instrDDCB__LD_L_RRC_iREGpDD(z80 *Z80) {
// 	z80.L = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.L = z80.rrc(z80.L)
// 	z80.memory.WriteByte(z80.tempaddr, z80.L)
// }

// /* RRC (REGISTER+dd) */
// func instrDDCB__RRC_iREGpDD(z80 *Z80) {
// 	var bytetemp byte = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	bytetemp = z80.rrc(bytetemp)
// 	z80.memory.WriteByte(z80.tempaddr, bytetemp)
// }

// /* LD A,RRC (REGISTER+dd) */
// func instrDDCB__LD_A_RRC_iREGpDD(z80 *Z80) {
// 	z80.A = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.A = z80.rrc(z80.A)
// 	z80.memory.WriteByte(z80.tempaddr, z80.A)
// }

// /* LD B,RL (REGISTER+dd) */
// func instrDDCB__LD_B_RL_iREGpDD(z80 *Z80) {
// 	z80.B = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.B = z80.rl(z80.B)
// 	z80.memory.WriteByte(z80.tempaddr, z80.B)
// }

// /* LD C,RL (REGISTER+dd) */
// func instrDDCB__LD_C_RL_iREGpDD(z80 *Z80) {
// 	z80.C = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.C = z80.rl(z80.C)
// 	z80.memory.WriteByte(z80.tempaddr, z80.C)
// }

// /* LD D,RL (REGISTER+dd) */
// func instrDDCB__LD_D_RL_iREGpDD(z80 *Z80) {
// 	z80.D = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.D = z80.rl(z80.D)
// 	z80.memory.WriteByte(z80.tempaddr, z80.D)
// }

// /* LD E,RL (REGISTER+dd) */
// func instrDDCB__LD_E_RL_iREGpDD(z80 *Z80) {
// 	z80.E = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.E = z80.rl(z80.E)
// 	z80.memory.WriteByte(z80.tempaddr, z80.E)
// }

// /* LD H,RL (REGISTER+dd) */
// func instrDDCB__LD_H_RL_iREGpDD(z80 *Z80) {
// 	z80.H = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.H = z80.rl(z80.H)
// 	z80.memory.WriteByte(z80.tempaddr, z80.H)
// }

// /* LD L,RL (REGISTER+dd) */
// func instrDDCB__LD_L_RL_iREGpDD(z80 *Z80) {
// 	z80.L = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.L = z80.rl(z80.L)
// 	z80.memory.WriteByte(z80.tempaddr, z80.L)
// }

// /* RL (REGISTER+dd) */
// func instrDDCB__RL_iREGpDD(z80 *Z80) {
// 	var bytetemp byte = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	bytetemp = z80.rl(bytetemp)
// 	z80.memory.WriteByte(z80.tempaddr, bytetemp)
// }

// /* LD A,RL (REGISTER+dd) */
// func instrDDCB__LD_A_RL_iREGpDD(z80 *Z80) {
// 	z80.A = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.A = z80.rl(z80.A)
// 	z80.memory.WriteByte(z80.tempaddr, z80.A)
// }

// /* LD B,RR (REGISTER+dd) */
// func instrDDCB__LD_B_RR_iREGpDD(z80 *Z80) {
// 	z80.B = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.B = z80.rr(z80.B)
// 	z80.memory.WriteByte(z80.tempaddr, z80.B)
// }

// /* LD C,RR (REGISTER+dd) */
// func instrDDCB__LD_C_RR_iREGpDD(z80 *Z80) {
// 	z80.C = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.C = z80.rr(z80.C)
// 	z80.memory.WriteByte(z80.tempaddr, z80.C)
// }

// /* LD D,RR (REGISTER+dd) */
// func instrDDCB__LD_D_RR_iREGpDD(z80 *Z80) {
// 	z80.D = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.D = z80.rr(z80.D)
// 	z80.memory.WriteByte(z80.tempaddr, z80.D)
// }

// /* LD E,RR (REGISTER+dd) */
// func instrDDCB__LD_E_RR_iREGpDD(z80 *Z80) {
// 	z80.E = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.E = z80.rr(z80.E)
// 	z80.memory.WriteByte(z80.tempaddr, z80.E)
// }

// /* LD H,RR (REGISTER+dd) */
// func instrDDCB__LD_H_RR_iREGpDD(z80 *Z80) {
// 	z80.H = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.H = z80.rr(z80.H)
// 	z80.memory.WriteByte(z80.tempaddr, z80.H)
// }

// /* LD L,RR (REGISTER+dd) */
// func instrDDCB__LD_L_RR_iREGpDD(z80 *Z80) {
// 	z80.L = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.L = z80.rr(z80.L)
// 	z80.memory.WriteByte(z80.tempaddr, z80.L)
// }

// /* RR (REGISTER+dd) */
// func instrDDCB__RR_iREGpDD(z80 *Z80) {
// 	var bytetemp byte = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	bytetemp = z80.rr(bytetemp)
// 	z80.memory.WriteByte(z80.tempaddr, bytetemp)
// }

// /* LD A,RR (REGISTER+dd) */
// func instrDDCB__LD_A_RR_iREGpDD(z80 *Z80) {
// 	z80.A = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.A = z80.rr(z80.A)
// 	z80.memory.WriteByte(z80.tempaddr, z80.A)
// }

// /* LD B,SLA (REGISTER+dd) */
// func instrDDCB__LD_B_SLA_iREGpDD(z80 *Z80) {
// 	z80.B = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.B = z80.sla(z80.B)
// 	z80.memory.WriteByte(z80.tempaddr, z80.B)
// }

// /* LD C,SLA (REGISTER+dd) */
// func instrDDCB__LD_C_SLA_iREGpDD(z80 *Z80) {
// 	z80.C = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.C = z80.sla(z80.C)
// 	z80.memory.WriteByte(z80.tempaddr, z80.C)
// }

// /* LD D,SLA (REGISTER+dd) */
// func instrDDCB__LD_D_SLA_iREGpDD(z80 *Z80) {
// 	z80.D = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.D = z80.sla(z80.D)
// 	z80.memory.WriteByte(z80.tempaddr, z80.D)
// }

// /* LD E,SLA (REGISTER+dd) */
// func instrDDCB__LD_E_SLA_iREGpDD(z80 *Z80) {
// 	z80.E = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.E = z80.sla(z80.E)
// 	z80.memory.WriteByte(z80.tempaddr, z80.E)
// }

// /* LD H,SLA (REGISTER+dd) */
// func instrDDCB__LD_H_SLA_iREGpDD(z80 *Z80) {
// 	z80.H = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.H = z80.sla(z80.H)
// 	z80.memory.WriteByte(z80.tempaddr, z80.H)
// }

// /* LD L,SLA (REGISTER+dd) */
// func instrDDCB__LD_L_SLA_iREGpDD(z80 *Z80) {
// 	z80.L = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.L = z80.sla(z80.L)
// 	z80.memory.WriteByte(z80.tempaddr, z80.L)
// }

// /* SLA (REGISTER+dd) */
// func instrDDCB__SLA_iREGpDD(z80 *Z80) {
// 	var bytetemp byte = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	bytetemp = z80.sla(bytetemp)
// 	z80.memory.WriteByte(z80.tempaddr, bytetemp)
// }

// /* LD A,SLA (REGISTER+dd) */
// func instrDDCB__LD_A_SLA_iREGpDD(z80 *Z80) {
// 	z80.A = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.A = z80.sla(z80.A)
// 	z80.memory.WriteByte(z80.tempaddr, z80.A)
// }

// /* LD B,SRA (REGISTER+dd) */
// func instrDDCB__LD_B_SRA_iREGpDD(z80 *Z80) {
// 	z80.B = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.B = z80.sra(z80.B)
// 	z80.memory.WriteByte(z80.tempaddr, z80.B)
// }

// /* LD C,SRA (REGISTER+dd) */
// func instrDDCB__LD_C_SRA_iREGpDD(z80 *Z80) {
// 	z80.C = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.C = z80.sra(z80.C)
// 	z80.memory.WriteByte(z80.tempaddr, z80.C)
// }

// /* LD D,SRA (REGISTER+dd) */
// func instrDDCB__LD_D_SRA_iREGpDD(z80 *Z80) {
// 	z80.D = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.D = z80.sra(z80.D)
// 	z80.memory.WriteByte(z80.tempaddr, z80.D)
// }

// /* LD E,SRA (REGISTER+dd) */
// func instrDDCB__LD_E_SRA_iREGpDD(z80 *Z80) {
// 	z80.E = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.E = z80.sra(z80.E)
// 	z80.memory.WriteByte(z80.tempaddr, z80.E)
// }

// /* LD H,SRA (REGISTER+dd) */
// func instrDDCB__LD_H_SRA_iREGpDD(z80 *Z80) {
// 	z80.H = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.H = z80.sra(z80.H)
// 	z80.memory.WriteByte(z80.tempaddr, z80.H)
// }

// /* LD L,SRA (REGISTER+dd) */
// func instrDDCB__LD_L_SRA_iREGpDD(z80 *Z80) {
// 	z80.L = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.L = z80.sra(z80.L)
// 	z80.memory.WriteByte(z80.tempaddr, z80.L)
// }

// /* SRA (REGISTER+dd) */
// func instrDDCB__SRA_iREGpDD(z80 *Z80) {
// 	var bytetemp byte = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	bytetemp = z80.sra(bytetemp)
// 	z80.memory.WriteByte(z80.tempaddr, bytetemp)
// }

// /* LD A,SRA (REGISTER+dd) */
// func instrDDCB__LD_A_SRA_iREGpDD(z80 *Z80) {
// 	z80.A = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.A = z80.sra(z80.A)
// 	z80.memory.WriteByte(z80.tempaddr, z80.A)
// }

// /* LD B,SLL (REGISTER+dd) */
// func instrDDCB__LD_B_SLL_iREGpDD(z80 *Z80) {
// 	z80.B = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.B = z80.sll(z80.B)
// 	z80.memory.WriteByte(z80.tempaddr, z80.B)
// }

// /* LD C,SLL (REGISTER+dd) */
// func instrDDCB__LD_C_SLL_iREGpDD(z80 *Z80) {
// 	z80.C = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.C = z80.sll(z80.C)
// 	z80.memory.WriteByte(z80.tempaddr, z80.C)
// }

// /* LD D,SLL (REGISTER+dd) */
// func instrDDCB__LD_D_SLL_iREGpDD(z80 *Z80) {
// 	z80.D = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.D = z80.sll(z80.D)
// 	z80.memory.WriteByte(z80.tempaddr, z80.D)
// }

// /* LD E,SLL (REGISTER+dd) */
// func instrDDCB__LD_E_SLL_iREGpDD(z80 *Z80) {
// 	z80.E = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.E = z80.sll(z80.E)
// 	z80.memory.WriteByte(z80.tempaddr, z80.E)
// }

// /* LD H,SLL (REGISTER+dd) */
// func instrDDCB__LD_H_SLL_iREGpDD(z80 *Z80) {
// 	z80.H = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.H = z80.sll(z80.H)
// 	z80.memory.WriteByte(z80.tempaddr, z80.H)
// }

// /* LD L,SLL (REGISTER+dd) */
// func instrDDCB__LD_L_SLL_iREGpDD(z80 *Z80) {
// 	z80.L = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.L = z80.sll(z80.L)
// 	z80.memory.WriteByte(z80.tempaddr, z80.L)
// }

// /* SLL (REGISTER+dd) */
// func instrDDCB__SLL_iREGpDD(z80 *Z80) {
// 	var bytetemp byte = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	bytetemp = z80.sll(bytetemp)
// 	z80.memory.WriteByte(z80.tempaddr, bytetemp)
// }

// /* LD A,SLL (REGISTER+dd) */
// func instrDDCB__LD_A_SLL_iREGpDD(z80 *Z80) {
// 	z80.A = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.A = z80.sll(z80.A)
// 	z80.memory.WriteByte(z80.tempaddr, z80.A)
// }

// /* LD B,SRL (REGISTER+dd) */
// func instrDDCB__LD_B_SRL_iREGpDD(z80 *Z80) {
// 	z80.B = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.B = z80.srl(z80.B)
// 	z80.memory.WriteByte(z80.tempaddr, z80.B)
// }

// /* LD C,SRL (REGISTER+dd) */
// func instrDDCB__LD_C_SRL_iREGpDD(z80 *Z80) {
// 	z80.C = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.C = z80.srl(z80.C)
// 	z80.memory.WriteByte(z80.tempaddr, z80.C)
// }

// /* LD D,SRL (REGISTER+dd) */
// func instrDDCB__LD_D_SRL_iREGpDD(z80 *Z80) {
// 	z80.D = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.D = z80.srl(z80.D)
// 	z80.memory.WriteByte(z80.tempaddr, z80.D)
// }

// /* LD E,SRL (REGISTER+dd) */
// func instrDDCB__LD_E_SRL_iREGpDD(z80 *Z80) {
// 	z80.E = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.E = z80.srl(z80.E)
// 	z80.memory.WriteByte(z80.tempaddr, z80.E)
// }

// /* LD H,SRL (REGISTER+dd) */
// func instrDDCB__LD_H_SRL_iREGpDD(z80 *Z80) {
// 	z80.H = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.H = z80.srl(z80.H)
// 	z80.memory.WriteByte(z80.tempaddr, z80.H)
// }

// /* LD L,SRL (REGISTER+dd) */
// func instrDDCB__LD_L_SRL_iREGpDD(z80 *Z80) {
// 	z80.L = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.L = z80.srl(z80.L)
// 	z80.memory.WriteByte(z80.tempaddr, z80.L)
// }

// /* SRL (REGISTER+dd) */
// func instrDDCB__SRL_iREGpDD(z80 *Z80) {
// 	var bytetemp byte = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	bytetemp = z80.srl(bytetemp)
// 	z80.memory.WriteByte(z80.tempaddr, bytetemp)
// }

// /* LD A,SRL (REGISTER+dd) */
// func instrDDCB__LD_A_SRL_iREGpDD(z80 *Z80) {
// 	z80.A = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.A = z80.srl(z80.A)
// 	z80.memory.WriteByte(z80.tempaddr, z80.A)
// }

// /* BIT 0,(REGISTER+dd) */
// func instrDDCB__BIT_0_iREGpDD(z80 *Z80) {
// 	bytetemp := z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.biti(0, bytetemp, z80.tempaddr)
// }

// /* BIT 1,(REGISTER+dd) */
// func instrDDCB__BIT_1_iREGpDD(z80 *Z80) {
// 	bytetemp := z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.biti(1, bytetemp, z80.tempaddr)
// }

// /* BIT 2,(REGISTER+dd) */
// func instrDDCB__BIT_2_iREGpDD(z80 *Z80) {
// 	bytetemp := z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.biti(2, bytetemp, z80.tempaddr)
// }

// /* BIT 3,(REGISTER+dd) */
// func instrDDCB__BIT_3_iREGpDD(z80 *Z80) {
// 	bytetemp := z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.biti(3, bytetemp, z80.tempaddr)
// }

// /* BIT 4,(REGISTER+dd) */
// func instrDDCB__BIT_4_iREGpDD(z80 *Z80) {
// 	bytetemp := z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.biti(4, bytetemp, z80.tempaddr)
// }

// /* BIT 5,(REGISTER+dd) */
// func instrDDCB__BIT_5_iREGpDD(z80 *Z80) {
// 	bytetemp := z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.biti(5, bytetemp, z80.tempaddr)
// }

// /* BIT 6,(REGISTER+dd) */
// func instrDDCB__BIT_6_iREGpDD(z80 *Z80) {
// 	bytetemp := z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.biti(6, bytetemp, z80.tempaddr)
// }

// /* BIT 7,(REGISTER+dd) */
// func instrDDCB__BIT_7_iREGpDD(z80 *Z80) {
// 	bytetemp := z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.biti(7, bytetemp, z80.tempaddr)
// }

// /* LD B,RES 0,(REGISTER+dd) */
// func instrDDCB__LD_B_RES_0_iREGpDD(z80 *Z80) {
// 	z80.B = z80.memory.ReadByte(z80.tempaddr) & 0xfe
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.B)
// }

// /* LD C,RES 0,(REGISTER+dd) */
// func instrDDCB__LD_C_RES_0_iREGpDD(z80 *Z80) {
// 	z80.C = z80.memory.ReadByte(z80.tempaddr) & 0xfe
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.C)
// }

// /* LD D,RES 0,(REGISTER+dd) */
// func instrDDCB__LD_D_RES_0_iREGpDD(z80 *Z80) {
// 	z80.D = z80.memory.ReadByte(z80.tempaddr) & 0xfe
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.D)
// }

// /* LD E,RES 0,(REGISTER+dd) */
// func instrDDCB__LD_E_RES_0_iREGpDD(z80 *Z80) {
// 	z80.E = z80.memory.ReadByte(z80.tempaddr) & 0xfe
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.E)
// }

// /* LD H,RES 0,(REGISTER+dd) */
// func instrDDCB__LD_H_RES_0_iREGpDD(z80 *Z80) {
// 	z80.H = z80.memory.ReadByte(z80.tempaddr) & 0xfe
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.H)
// }

// /* LD L,RES 0,(REGISTER+dd) */
// func instrDDCB__LD_L_RES_0_iREGpDD(z80 *Z80) {
// 	z80.L = z80.memory.ReadByte(z80.tempaddr) & 0xfe
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.L)
// }

// /* RES 0,(REGISTER+dd) */
// func instrDDCB__RES_0_iREGpDD(z80 *Z80) {
// 	var bytetemp byte = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, bytetemp&0xfe)
// }

// /* LD A,RES 0,(REGISTER+dd) */
// func instrDDCB__LD_A_RES_0_iREGpDD(z80 *Z80) {
// 	z80.A = z80.memory.ReadByte(z80.tempaddr) & 0xfe
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.A)
// }

// /* LD B,RES 1,(REGISTER+dd) */
// func instrDDCB__LD_B_RES_1_iREGpDD(z80 *Z80) {
// 	z80.B = z80.memory.ReadByte(z80.tempaddr) & 0xfd
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.B)
// }

// /* LD C,RES 1,(REGISTER+dd) */
// func instrDDCB__LD_C_RES_1_iREGpDD(z80 *Z80) {
// 	z80.C = z80.memory.ReadByte(z80.tempaddr) & 0xfd
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.C)
// }

// /* LD D,RES 1,(REGISTER+dd) */
// func instrDDCB__LD_D_RES_1_iREGpDD(z80 *Z80) {
// 	z80.D = z80.memory.ReadByte(z80.tempaddr) & 0xfd
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.D)
// }

// /* LD E,RES 1,(REGISTER+dd) */
// func instrDDCB__LD_E_RES_1_iREGpDD(z80 *Z80) {
// 	z80.E = z80.memory.ReadByte(z80.tempaddr) & 0xfd
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.E)
// }

// /* LD H,RES 1,(REGISTER+dd) */
// func instrDDCB__LD_H_RES_1_iREGpDD(z80 *Z80) {
// 	z80.H = z80.memory.ReadByte(z80.tempaddr) & 0xfd
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.H)
// }

// /* LD L,RES 1,(REGISTER+dd) */
// func instrDDCB__LD_L_RES_1_iREGpDD(z80 *Z80) {
// 	z80.L = z80.memory.ReadByte(z80.tempaddr) & 0xfd
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.L)
// }

// /* RES 1,(REGISTER+dd) */
// func instrDDCB__RES_1_iREGpDD(z80 *Z80) {
// 	var bytetemp byte = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, bytetemp&0xfd)
// }

// /* LD A,RES 1,(REGISTER+dd) */
// func instrDDCB__LD_A_RES_1_iREGpDD(z80 *Z80) {
// 	z80.A = z80.memory.ReadByte(z80.tempaddr) & 0xfd
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.A)
// }

// /* LD B,RES 2,(REGISTER+dd) */
// func instrDDCB__LD_B_RES_2_iREGpDD(z80 *Z80) {
// 	z80.B = z80.memory.ReadByte(z80.tempaddr) & 0xfb
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.B)
// }

// /* LD C,RES 2,(REGISTER+dd) */
// func instrDDCB__LD_C_RES_2_iREGpDD(z80 *Z80) {
// 	z80.C = z80.memory.ReadByte(z80.tempaddr) & 0xfb
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.C)
// }

// /* LD D,RES 2,(REGISTER+dd) */
// func instrDDCB__LD_D_RES_2_iREGpDD(z80 *Z80) {
// 	z80.D = z80.memory.ReadByte(z80.tempaddr) & 0xfb
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.D)
// }

// /* LD E,RES 2,(REGISTER+dd) */
// func instrDDCB__LD_E_RES_2_iREGpDD(z80 *Z80) {
// 	z80.E = z80.memory.ReadByte(z80.tempaddr) & 0xfb
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.E)
// }

// /* LD H,RES 2,(REGISTER+dd) */
// func instrDDCB__LD_H_RES_2_iREGpDD(z80 *Z80) {
// 	z80.H = z80.memory.ReadByte(z80.tempaddr) & 0xfb
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.H)
// }

// /* LD L,RES 2,(REGISTER+dd) */
// func instrDDCB__LD_L_RES_2_iREGpDD(z80 *Z80) {
// 	z80.L = z80.memory.ReadByte(z80.tempaddr) & 0xfb
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.L)
// }

// /* RES 2,(REGISTER+dd) */
// func instrDDCB__RES_2_iREGpDD(z80 *Z80) {
// 	var bytetemp byte = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, bytetemp&0xfb)
// }

// /* LD A,RES 2,(REGISTER+dd) */
// func instrDDCB__LD_A_RES_2_iREGpDD(z80 *Z80) {
// 	z80.A = z80.memory.ReadByte(z80.tempaddr) & 0xfb
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.A)
// }

// /* LD B,RES 3,(REGISTER+dd) */
// func instrDDCB__LD_B_RES_3_iREGpDD(z80 *Z80) {
// 	z80.B = z80.memory.ReadByte(z80.tempaddr) & 0xf7
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.B)
// }

// /* LD C,RES 3,(REGISTER+dd) */
// func instrDDCB__LD_C_RES_3_iREGpDD(z80 *Z80) {
// 	z80.C = z80.memory.ReadByte(z80.tempaddr) & 0xf7
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.C)
// }

// /* LD D,RES 3,(REGISTER+dd) */
// func instrDDCB__LD_D_RES_3_iREGpDD(z80 *Z80) {
// 	z80.D = z80.memory.ReadByte(z80.tempaddr) & 0xf7
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.D)
// }

// /* LD E,RES 3,(REGISTER+dd) */
// func instrDDCB__LD_E_RES_3_iREGpDD(z80 *Z80) {
// 	z80.E = z80.memory.ReadByte(z80.tempaddr) & 0xf7
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.E)
// }

// /* LD H,RES 3,(REGISTER+dd) */
// func instrDDCB__LD_H_RES_3_iREGpDD(z80 *Z80) {
// 	z80.H = z80.memory.ReadByte(z80.tempaddr) & 0xf7
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.H)
// }

// /* LD L,RES 3,(REGISTER+dd) */
// func instrDDCB__LD_L_RES_3_iREGpDD(z80 *Z80) {
// 	z80.L = z80.memory.ReadByte(z80.tempaddr) & 0xf7
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.L)
// }

// /* RES 3,(REGISTER+dd) */
// func instrDDCB__RES_3_iREGpDD(z80 *Z80) {
// 	var bytetemp byte = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, bytetemp&0xf7)
// }

// /* LD A,RES 3,(REGISTER+dd) */
// func instrDDCB__LD_A_RES_3_iREGpDD(z80 *Z80) {
// 	z80.A = z80.memory.ReadByte(z80.tempaddr) & 0xf7
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.A)
// }

// /* LD B,RES 4,(REGISTER+dd) */
// func instrDDCB__LD_B_RES_4_iREGpDD(z80 *Z80) {
// 	z80.B = z80.memory.ReadByte(z80.tempaddr) & 0xef
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.B)
// }

// /* LD C,RES 4,(REGISTER+dd) */
// func instrDDCB__LD_C_RES_4_iREGpDD(z80 *Z80) {
// 	z80.C = z80.memory.ReadByte(z80.tempaddr) & 0xef
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.C)
// }

// /* LD D,RES 4,(REGISTER+dd) */
// func instrDDCB__LD_D_RES_4_iREGpDD(z80 *Z80) {
// 	z80.D = z80.memory.ReadByte(z80.tempaddr) & 0xef
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.D)
// }

// /* LD E,RES 4,(REGISTER+dd) */
// func instrDDCB__LD_E_RES_4_iREGpDD(z80 *Z80) {
// 	z80.E = z80.memory.ReadByte(z80.tempaddr) & 0xef
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.E)
// }

// /* LD H,RES 4,(REGISTER+dd) */
// func instrDDCB__LD_H_RES_4_iREGpDD(z80 *Z80) {
// 	z80.H = z80.memory.ReadByte(z80.tempaddr) & 0xef
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.H)
// }

// /* LD L,RES 4,(REGISTER+dd) */
// func instrDDCB__LD_L_RES_4_iREGpDD(z80 *Z80) {
// 	z80.L = z80.memory.ReadByte(z80.tempaddr) & 0xef
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.L)
// }

// /* RES 4,(REGISTER+dd) */
// func instrDDCB__RES_4_iREGpDD(z80 *Z80) {
// 	var bytetemp byte = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, bytetemp&0xef)
// }

// /* LD A,RES 4,(REGISTER+dd) */
// func instrDDCB__LD_A_RES_4_iREGpDD(z80 *Z80) {
// 	z80.A = z80.memory.ReadByte(z80.tempaddr) & 0xef
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.A)
// }

// /* LD B,RES 5,(REGISTER+dd) */
// func instrDDCB__LD_B_RES_5_iREGpDD(z80 *Z80) {
// 	z80.B = z80.memory.ReadByte(z80.tempaddr) & 0xdf
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.B)
// }

// /* LD C,RES 5,(REGISTER+dd) */
// func instrDDCB__LD_C_RES_5_iREGpDD(z80 *Z80) {
// 	z80.C = z80.memory.ReadByte(z80.tempaddr) & 0xdf
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.C)
// }

// /* LD D,RES 5,(REGISTER+dd) */
// func instrDDCB__LD_D_RES_5_iREGpDD(z80 *Z80) {
// 	z80.D = z80.memory.ReadByte(z80.tempaddr) & 0xdf
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.D)
// }

// /* LD E,RES 5,(REGISTER+dd) */
// func instrDDCB__LD_E_RES_5_iREGpDD(z80 *Z80) {
// 	z80.E = z80.memory.ReadByte(z80.tempaddr) & 0xdf
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.E)
// }

// /* LD H,RES 5,(REGISTER+dd) */
// func instrDDCB__LD_H_RES_5_iREGpDD(z80 *Z80) {
// 	z80.H = z80.memory.ReadByte(z80.tempaddr) & 0xdf
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.H)
// }

// /* LD L,RES 5,(REGISTER+dd) */
// func instrDDCB__LD_L_RES_5_iREGpDD(z80 *Z80) {
// 	z80.L = z80.memory.ReadByte(z80.tempaddr) & 0xdf
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.L)
// }

// /* RES 5,(REGISTER+dd) */
// func instrDDCB__RES_5_iREGpDD(z80 *Z80) {
// 	var bytetemp byte = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, bytetemp&0xdf)
// }

// /* LD A,RES 5,(REGISTER+dd) */
// func instrDDCB__LD_A_RES_5_iREGpDD(z80 *Z80) {
// 	z80.A = z80.memory.ReadByte(z80.tempaddr) & 0xdf
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.A)
// }

// /* LD B,RES 6,(REGISTER+dd) */
// func instrDDCB__LD_B_RES_6_iREGpDD(z80 *Z80) {
// 	z80.B = z80.memory.ReadByte(z80.tempaddr) & 0xbf
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.B)
// }

// /* LD C,RES 6,(REGISTER+dd) */
// func instrDDCB__LD_C_RES_6_iREGpDD(z80 *Z80) {
// 	z80.C = z80.memory.ReadByte(z80.tempaddr) & 0xbf
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.C)
// }

// /* LD D,RES 6,(REGISTER+dd) */
// func instrDDCB__LD_D_RES_6_iREGpDD(z80 *Z80) {
// 	z80.D = z80.memory.ReadByte(z80.tempaddr) & 0xbf
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.D)
// }

// /* LD E,RES 6,(REGISTER+dd) */
// func instrDDCB__LD_E_RES_6_iREGpDD(z80 *Z80) {
// 	z80.E = z80.memory.ReadByte(z80.tempaddr) & 0xbf
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.E)
// }

// /* LD H,RES 6,(REGISTER+dd) */
// func instrDDCB__LD_H_RES_6_iREGpDD(z80 *Z80) {
// 	z80.H = z80.memory.ReadByte(z80.tempaddr) & 0xbf
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.H)
// }

// /* LD L,RES 6,(REGISTER+dd) */
// func instrDDCB__LD_L_RES_6_iREGpDD(z80 *Z80) {
// 	z80.L = z80.memory.ReadByte(z80.tempaddr) & 0xbf
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.L)
// }

// /* RES 6,(REGISTER+dd) */
// func instrDDCB__RES_6_iREGpDD(z80 *Z80) {
// 	var bytetemp byte = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, bytetemp&0xbf)
// }

// /* LD A,RES 6,(REGISTER+dd) */
// func instrDDCB__LD_A_RES_6_iREGpDD(z80 *Z80) {
// 	z80.A = z80.memory.ReadByte(z80.tempaddr) & 0xbf
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.A)
// }

// /* LD B,RES 7,(REGISTER+dd) */
// func instrDDCB__LD_B_RES_7_iREGpDD(z80 *Z80) {
// 	z80.B = z80.memory.ReadByte(z80.tempaddr) & 0x7f
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.B)
// }

// /* LD C,RES 7,(REGISTER+dd) */
// func instrDDCB__LD_C_RES_7_iREGpDD(z80 *Z80) {
// 	z80.C = z80.memory.ReadByte(z80.tempaddr) & 0x7f
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.C)
// }

// /* LD D,RES 7,(REGISTER+dd) */
// func instrDDCB__LD_D_RES_7_iREGpDD(z80 *Z80) {
// 	z80.D = z80.memory.ReadByte(z80.tempaddr) & 0x7f
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.D)
// }

// /* LD E,RES 7,(REGISTER+dd) */
// func instrDDCB__LD_E_RES_7_iREGpDD(z80 *Z80) {
// 	z80.E = z80.memory.ReadByte(z80.tempaddr) & 0x7f
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.E)
// }

// /* LD H,RES 7,(REGISTER+dd) */
// func instrDDCB__LD_H_RES_7_iREGpDD(z80 *Z80) {
// 	z80.H = z80.memory.ReadByte(z80.tempaddr) & 0x7f
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.H)
// }

// /* LD L,RES 7,(REGISTER+dd) */
// func instrDDCB__LD_L_RES_7_iREGpDD(z80 *Z80) {
// 	z80.L = z80.memory.ReadByte(z80.tempaddr) & 0x7f
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.L)
// }

// /* RES 7,(REGISTER+dd) */
// func instrDDCB__RES_7_iREGpDD(z80 *Z80) {
// 	var bytetemp byte = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, bytetemp&0x7f)
// }

// /* LD A,RES 7,(REGISTER+dd) */
// func instrDDCB__LD_A_RES_7_iREGpDD(z80 *Z80) {
// 	z80.A = z80.memory.ReadByte(z80.tempaddr) & 0x7f
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.A)
// }

// /* LD B,SET 0,(REGISTER+dd) */
// func instrDDCB__LD_B_SET_0_iREGpDD(z80 *Z80) {
// 	z80.B = z80.memory.ReadByte(z80.tempaddr) | 0x01
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.B)
// }

// /* LD C,SET 0,(REGISTER+dd) */
// func instrDDCB__LD_C_SET_0_iREGpDD(z80 *Z80) {
// 	z80.C = z80.memory.ReadByte(z80.tempaddr) | 0x01
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.C)
// }

// /* LD D,SET 0,(REGISTER+dd) */
// func instrDDCB__LD_D_SET_0_iREGpDD(z80 *Z80) {
// 	z80.D = z80.memory.ReadByte(z80.tempaddr) | 0x01
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.D)
// }

// /* LD E,SET 0,(REGISTER+dd) */
// func instrDDCB__LD_E_SET_0_iREGpDD(z80 *Z80) {
// 	z80.E = z80.memory.ReadByte(z80.tempaddr) | 0x01
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.E)
// }

// /* LD H,SET 0,(REGISTER+dd) */
// func instrDDCB__LD_H_SET_0_iREGpDD(z80 *Z80) {
// 	z80.H = z80.memory.ReadByte(z80.tempaddr) | 0x01
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.H)
// }

// /* LD L,SET 0,(REGISTER+dd) */
// func instrDDCB__LD_L_SET_0_iREGpDD(z80 *Z80) {
// 	z80.L = z80.memory.ReadByte(z80.tempaddr) | 0x01
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.L)
// }

// /* SET 0,(REGISTER+dd) */
// func instrDDCB__SET_0_iREGpDD(z80 *Z80) {
// 	var bytetemp byte = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, bytetemp|0x01)
// }

// /* LD A,SET 0,(REGISTER+dd) */
// func instrDDCB__LD_A_SET_0_iREGpDD(z80 *Z80) {
// 	z80.A = z80.memory.ReadByte(z80.tempaddr) | 0x01
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.A)
// }

// /* LD B,SET 1,(REGISTER+dd) */
// func instrDDCB__LD_B_SET_1_iREGpDD(z80 *Z80) {
// 	z80.B = z80.memory.ReadByte(z80.tempaddr) | 0x02
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.B)
// }

// /* LD C,SET 1,(REGISTER+dd) */
// func instrDDCB__LD_C_SET_1_iREGpDD(z80 *Z80) {
// 	z80.C = z80.memory.ReadByte(z80.tempaddr) | 0x02
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.C)
// }

// /* LD D,SET 1,(REGISTER+dd) */
// func instrDDCB__LD_D_SET_1_iREGpDD(z80 *Z80) {
// 	z80.D = z80.memory.ReadByte(z80.tempaddr) | 0x02
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.D)
// }

// /* LD E,SET 1,(REGISTER+dd) */
// func instrDDCB__LD_E_SET_1_iREGpDD(z80 *Z80) {
// 	z80.E = z80.memory.ReadByte(z80.tempaddr) | 0x02
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.E)
// }

// /* LD H,SET 1,(REGISTER+dd) */
// func instrDDCB__LD_H_SET_1_iREGpDD(z80 *Z80) {
// 	z80.H = z80.memory.ReadByte(z80.tempaddr) | 0x02
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.H)
// }

// /* LD L,SET 1,(REGISTER+dd) */
// func instrDDCB__LD_L_SET_1_iREGpDD(z80 *Z80) {
// 	z80.L = z80.memory.ReadByte(z80.tempaddr) | 0x02
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.L)
// }

// /* SET 1,(REGISTER+dd) */
// func instrDDCB__SET_1_iREGpDD(z80 *Z80) {
// 	var bytetemp byte = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, bytetemp|0x02)
// }

// /* LD A,SET 1,(REGISTER+dd) */
// func instrDDCB__LD_A_SET_1_iREGpDD(z80 *Z80) {
// 	z80.A = z80.memory.ReadByte(z80.tempaddr) | 0x02
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.A)
// }

// /* LD B,SET 2,(REGISTER+dd) */
// func instrDDCB__LD_B_SET_2_iREGpDD(z80 *Z80) {
// 	z80.B = z80.memory.ReadByte(z80.tempaddr) | 0x04
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.B)
// }

// /* LD C,SET 2,(REGISTER+dd) */
// func instrDDCB__LD_C_SET_2_iREGpDD(z80 *Z80) {
// 	z80.C = z80.memory.ReadByte(z80.tempaddr) | 0x04
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.C)
// }

// /* LD D,SET 2,(REGISTER+dd) */
// func instrDDCB__LD_D_SET_2_iREGpDD(z80 *Z80) {
// 	z80.D = z80.memory.ReadByte(z80.tempaddr) | 0x04
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.D)
// }

// /* LD E,SET 2,(REGISTER+dd) */
// func instrDDCB__LD_E_SET_2_iREGpDD(z80 *Z80) {
// 	z80.E = z80.memory.ReadByte(z80.tempaddr) | 0x04
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.E)
// }

// /* LD H,SET 2,(REGISTER+dd) */
// func instrDDCB__LD_H_SET_2_iREGpDD(z80 *Z80) {
// 	z80.H = z80.memory.ReadByte(z80.tempaddr) | 0x04
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.H)
// }

// /* LD L,SET 2,(REGISTER+dd) */
// func instrDDCB__LD_L_SET_2_iREGpDD(z80 *Z80) {
// 	z80.L = z80.memory.ReadByte(z80.tempaddr) | 0x04
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.L)
// }

// /* SET 2,(REGISTER+dd) */
// func instrDDCB__SET_2_iREGpDD(z80 *Z80) {
// 	var bytetemp byte = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, bytetemp|0x04)
// }

// /* LD A,SET 2,(REGISTER+dd) */
// func instrDDCB__LD_A_SET_2_iREGpDD(z80 *Z80) {
// 	z80.A = z80.memory.ReadByte(z80.tempaddr) | 0x04
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.A)
// }

// /* LD B,SET 3,(REGISTER+dd) */
// func instrDDCB__LD_B_SET_3_iREGpDD(z80 *Z80) {
// 	z80.B = z80.memory.ReadByte(z80.tempaddr) | 0x08
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.B)
// }

// /* LD C,SET 3,(REGISTER+dd) */
// func instrDDCB__LD_C_SET_3_iREGpDD(z80 *Z80) {
// 	z80.C = z80.memory.ReadByte(z80.tempaddr) | 0x08
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.C)
// }

// /* LD D,SET 3,(REGISTER+dd) */
// func instrDDCB__LD_D_SET_3_iREGpDD(z80 *Z80) {
// 	z80.D = z80.memory.ReadByte(z80.tempaddr) | 0x08
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.D)
// }

// /* LD E,SET 3,(REGISTER+dd) */
// func instrDDCB__LD_E_SET_3_iREGpDD(z80 *Z80) {
// 	z80.E = z80.memory.ReadByte(z80.tempaddr) | 0x08
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.E)
// }

// /* LD H,SET 3,(REGISTER+dd) */
// func instrDDCB__LD_H_SET_3_iREGpDD(z80 *Z80) {
// 	z80.H = z80.memory.ReadByte(z80.tempaddr) | 0x08
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.H)
// }

// /* LD L,SET 3,(REGISTER+dd) */
// func instrDDCB__LD_L_SET_3_iREGpDD(z80 *Z80) {
// 	z80.L = z80.memory.ReadByte(z80.tempaddr) | 0x08
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.L)
// }

// /* SET 3,(REGISTER+dd) */
// func instrDDCB__SET_3_iREGpDD(z80 *Z80) {
// 	var bytetemp byte = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, bytetemp|0x08)
// }

// /* LD A,SET 3,(REGISTER+dd) */
// func instrDDCB__LD_A_SET_3_iREGpDD(z80 *Z80) {
// 	z80.A = z80.memory.ReadByte(z80.tempaddr) | 0x08
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.A)
// }

// /* LD B,SET 4,(REGISTER+dd) */
// func instrDDCB__LD_B_SET_4_iREGpDD(z80 *Z80) {
// 	z80.B = z80.memory.ReadByte(z80.tempaddr) | 0x10
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.B)
// }

// /* LD C,SET 4,(REGISTER+dd) */
// func instrDDCB__LD_C_SET_4_iREGpDD(z80 *Z80) {
// 	z80.C = z80.memory.ReadByte(z80.tempaddr) | 0x10
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.C)
// }

// /* LD D,SET 4,(REGISTER+dd) */
// func instrDDCB__LD_D_SET_4_iREGpDD(z80 *Z80) {
// 	z80.D = z80.memory.ReadByte(z80.tempaddr) | 0x10
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.D)
// }

// /* LD E,SET 4,(REGISTER+dd) */
// func instrDDCB__LD_E_SET_4_iREGpDD(z80 *Z80) {
// 	z80.E = z80.memory.ReadByte(z80.tempaddr) | 0x10
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.E)
// }

// /* LD H,SET 4,(REGISTER+dd) */
// func instrDDCB__LD_H_SET_4_iREGpDD(z80 *Z80) {
// 	z80.H = z80.memory.ReadByte(z80.tempaddr) | 0x10
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.H)
// }

// /* LD L,SET 4,(REGISTER+dd) */
// func instrDDCB__LD_L_SET_4_iREGpDD(z80 *Z80) {
// 	z80.L = z80.memory.ReadByte(z80.tempaddr) | 0x10
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.L)
// }

// /* SET 4,(REGISTER+dd) */
// func instrDDCB__SET_4_iREGpDD(z80 *Z80) {
// 	var bytetemp byte = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, bytetemp|0x10)
// }

// /* LD A,SET 4,(REGISTER+dd) */
// func instrDDCB__LD_A_SET_4_iREGpDD(z80 *Z80) {
// 	z80.A = z80.memory.ReadByte(z80.tempaddr) | 0x10
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.A)
// }

// /* LD B,SET 5,(REGISTER+dd) */
// func instrDDCB__LD_B_SET_5_iREGpDD(z80 *Z80) {
// 	z80.B = z80.memory.ReadByte(z80.tempaddr) | 0x20
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.B)
// }

// /* LD C,SET 5,(REGISTER+dd) */
// func instrDDCB__LD_C_SET_5_iREGpDD(z80 *Z80) {
// 	z80.C = z80.memory.ReadByte(z80.tempaddr) | 0x20
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.C)
// }

// /* LD D,SET 5,(REGISTER+dd) */
// func instrDDCB__LD_D_SET_5_iREGpDD(z80 *Z80) {
// 	z80.D = z80.memory.ReadByte(z80.tempaddr) | 0x20
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.D)
// }

// /* LD E,SET 5,(REGISTER+dd) */
// func instrDDCB__LD_E_SET_5_iREGpDD(z80 *Z80) {
// 	z80.E = z80.memory.ReadByte(z80.tempaddr) | 0x20
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.E)
// }

// /* LD H,SET 5,(REGISTER+dd) */
// func instrDDCB__LD_H_SET_5_iREGpDD(z80 *Z80) {
// 	z80.H = z80.memory.ReadByte(z80.tempaddr) | 0x20
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.H)
// }

// /* LD L,SET 5,(REGISTER+dd) */
// func instrDDCB__LD_L_SET_5_iREGpDD(z80 *Z80) {
// 	z80.L = z80.memory.ReadByte(z80.tempaddr) | 0x20
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.L)
// }

// /* SET 5,(REGISTER+dd) */
// func instrDDCB__SET_5_iREGpDD(z80 *Z80) {
// 	var bytetemp byte = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, bytetemp|0x20)
// }

// /* LD A,SET 5,(REGISTER+dd) */
// func instrDDCB__LD_A_SET_5_iREGpDD(z80 *Z80) {
// 	z80.A = z80.memory.ReadByte(z80.tempaddr) | 0x20
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.A)
// }

// /* LD B,SET 6,(REGISTER+dd) */
// func instrDDCB__LD_B_SET_6_iREGpDD(z80 *Z80) {
// 	z80.B = z80.memory.ReadByte(z80.tempaddr) | 0x40
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.B)
// }

// /* LD C,SET 6,(REGISTER+dd) */
// func instrDDCB__LD_C_SET_6_iREGpDD(z80 *Z80) {
// 	z80.C = z80.memory.ReadByte(z80.tempaddr) | 0x40
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.C)
// }

// /* LD D,SET 6,(REGISTER+dd) */
// func instrDDCB__LD_D_SET_6_iREGpDD(z80 *Z80) {
// 	z80.D = z80.memory.ReadByte(z80.tempaddr) | 0x40
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.D)
// }

// /* LD E,SET 6,(REGISTER+dd) */
// func instrDDCB__LD_E_SET_6_iREGpDD(z80 *Z80) {
// 	z80.E = z80.memory.ReadByte(z80.tempaddr) | 0x40
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.E)
// }

// /* LD H,SET 6,(REGISTER+dd) */
// func instrDDCB__LD_H_SET_6_iREGpDD(z80 *Z80) {
// 	z80.H = z80.memory.ReadByte(z80.tempaddr) | 0x40
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.H)
// }

// /* LD L,SET 6,(REGISTER+dd) */
// func instrDDCB__LD_L_SET_6_iREGpDD(z80 *Z80) {
// 	z80.L = z80.memory.ReadByte(z80.tempaddr) | 0x40
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.L)
// }

// /* SET 6,(REGISTER+dd) */
// func instrDDCB__SET_6_iREGpDD(z80 *Z80) {
// 	var bytetemp byte = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, bytetemp|0x40)
// }

// /* LD A,SET 6,(REGISTER+dd) */
// func instrDDCB__LD_A_SET_6_iREGpDD(z80 *Z80) {
// 	z80.A = z80.memory.ReadByte(z80.tempaddr) | 0x40
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.A)
// }

// /* LD B,SET 7,(REGISTER+dd) */
// func instrDDCB__LD_B_SET_7_iREGpDD(z80 *Z80) {
// 	z80.B = z80.memory.ReadByte(z80.tempaddr) | 0x80
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.B)
// }

// /* LD C,SET 7,(REGISTER+dd) */
// func instrDDCB__LD_C_SET_7_iREGpDD(z80 *Z80) {
// 	z80.C = z80.memory.ReadByte(z80.tempaddr) | 0x80
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.C)
// }

// /* LD D,SET 7,(REGISTER+dd) */
// func instrDDCB__LD_D_SET_7_iREGpDD(z80 *Z80) {
// 	z80.D = z80.memory.ReadByte(z80.tempaddr) | 0x80
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.D)
// }

// /* LD E,SET 7,(REGISTER+dd) */
// func instrDDCB__LD_E_SET_7_iREGpDD(z80 *Z80) {
// 	z80.E = z80.memory.ReadByte(z80.tempaddr) | 0x80
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.E)
// }

// /* LD H,SET 7,(REGISTER+dd) */
// func instrDDCB__LD_H_SET_7_iREGpDD(z80 *Z80) {
// 	z80.H = z80.memory.ReadByte(z80.tempaddr) | 0x80
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.H)
// }

// /* LD L,SET 7,(REGISTER+dd) */
// func instrDDCB__LD_L_SET_7_iREGpDD(z80 *Z80) {
// 	z80.L = z80.memory.ReadByte(z80.tempaddr) | 0x80
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.L)
// }

// /* SET 7,(REGISTER+dd) */
// func instrDDCB__SET_7_iREGpDD(z80 *Z80) {
// 	var bytetemp byte = z80.memory.ReadByte(z80.tempaddr)
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, bytetemp|0x80)
// }

// /* LD A,SET 7,(REGISTER+dd) */
// func instrDDCB__LD_A_SET_7_iREGpDD(z80 *Z80) {
// 	z80.A = z80.memory.ReadByte(z80.tempaddr) | 0x80
// 	z80.memory.ContendReadNoMreq(z80.tempaddr, 1)
// 	z80.memory.WriteByte(z80.tempaddr, z80.A)
// }

//--
