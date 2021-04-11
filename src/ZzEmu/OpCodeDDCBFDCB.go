package ZzEmu

func initOpcodeDDCBFDCBMap() {

	// BEGIN of 0xddfdcb shifted opcodes
	/* LD B,RLC (REGISTER+dd) */
	OpcodeDDCBMap[0x00] = instrDDCB__LD_r_RLC_iREGpDD
	/* LD C,RLC (REGISTER+dd) */
	OpcodeDDCBMap[0x01] = instrDDCB__LD_r_RLC_iREGpDD
	/* LD D,RLC (REGISTER+dd) */
	OpcodeDDCBMap[0x02] = instrDDCB__LD_r_RLC_iREGpDD
	/* LD E,RLC (REGISTER+dd) */
	OpcodeDDCBMap[0x03] = instrDDCB__LD_r_RLC_iREGpDD
	/* LD H,RLC (REGISTER+dd) */
	OpcodeDDCBMap[0x04] = instrDDCB__LD_r_RLC_iREGpDD
	/* LD L,RLC (REGISTER+dd) */
	OpcodeDDCBMap[0x05] = instrDDCB__LD_r_RLC_iREGpDD
	/* RLC (REGISTER+dd) */
	OpcodeDDCBMap[0x06] = instrDDCB__RLC_iREGpDD
	/* LD A,RLC (REGISTER+dd) */
	OpcodeDDCBMap[0x07] = instrDDCB__LD_r_RLC_iREGpDD
	/* LD B,RRC (REGISTER+dd) */
	OpcodeDDCBMap[0x08] = instrDDCB__LD_r_RRC_iREGpDD
	/* LD C,RRC (REGISTER+dd) */
	OpcodeDDCBMap[0x09] = instrDDCB__LD_r_RRC_iREGpDD
	/* LD D,RRC (REGISTER+dd) */
	OpcodeDDCBMap[0x0a] = instrDDCB__LD_r_RRC_iREGpDD
	/* LD E,RRC (REGISTER+dd) */
	OpcodeDDCBMap[0x0b] = instrDDCB__LD_r_RRC_iREGpDD
	/* LD H,RRC (REGISTER+dd) */
	OpcodeDDCBMap[0x0c] = instrDDCB__LD_r_RRC_iREGpDD
	/* LD L,RRC (REGISTER+dd) */
	OpcodeDDCBMap[0x0d] = instrDDCB__LD_r_RRC_iREGpDD
	/* RRC (REGISTER+dd) */
	OpcodeDDCBMap[0x0e] = instrDDCB__RRC_iREGpDD
	/* LD A,RRC (REGISTER+dd) */
	OpcodeDDCBMap[0x0f] = instrDDCB__LD_r_RRC_iREGpDD
	/* LD B,RL (REGISTER+dd) */
	OpcodeDDCBMap[0x10] = instrDDCB__LD_r_RL_iREGpDD
	/* LD C,RL (REGISTER+dd) */
	OpcodeDDCBMap[0x11] = instrDDCB__LD_r_RL_iREGpDD
	/* LD D,RL (REGISTER+dd) */
	OpcodeDDCBMap[0x12] = instrDDCB__LD_r_RL_iREGpDD
	/* LD E,RL (REGISTER+dd) */
	OpcodeDDCBMap[0x13] = instrDDCB__LD_r_RL_iREGpDD
	/* LD H,RL (REGISTER+dd) */
	OpcodeDDCBMap[0x14] = instrDDCB__LD_r_RL_iREGpDD
	/* LD L,RL (REGISTER+dd) */
	OpcodeDDCBMap[0x15] = instrDDCB__LD_r_RL_iREGpDD
	/* RL (REGISTER+dd) */
	OpcodeDDCBMap[0x16] = instrDDCB__RL_iREGpDD
	/* LD A,RL (REGISTER+dd) */
	OpcodeDDCBMap[0x17] = instrDDCB__LD_r_RL_iREGpDD
	/* LD B,RR (REGISTER+dd) */
	OpcodeDDCBMap[0x18] = instrDDCB__LD_r_RR_iREGpDD
	/* LD C,RR (REGISTER+dd) */
	OpcodeDDCBMap[0x19] = instrDDCB__LD_r_RR_iREGpDD
	/* LD D,RR (REGISTER+dd) */
	OpcodeDDCBMap[0x1a] = instrDDCB__LD_r_RR_iREGpDD
	/* LD E,RR (REGISTER+dd) */
	OpcodeDDCBMap[0x1b] = instrDDCB__LD_r_RR_iREGpDD
	/* LD H,RR (REGISTER+dd) */
	OpcodeDDCBMap[0x1c] = instrDDCB__LD_r_RR_iREGpDD
	/* LD L,RR (REGISTER+dd) */
	OpcodeDDCBMap[0x1d] = instrDDCB__LD_r_RR_iREGpDD
	/* RR (REGISTER+dd) */
	OpcodeDDCBMap[0x1e] = instrDDCB__RR_iREGpDD
	/* LD A,RR (REGISTER+dd) */
	OpcodeDDCBMap[0x1f] = instrDDCB__LD_r_RR_iREGpDD
	/* LD B,SLA (REGISTER+dd) */
	OpcodeDDCBMap[0x20] = instrDDCB__LD_r_SLA_iREGpDD
	/* LD C,SLA (REGISTER+dd) */
	OpcodeDDCBMap[0x21] = instrDDCB__LD_r_SLA_iREGpDD
	/* LD D,SLA (REGISTER+dd) */
	OpcodeDDCBMap[0x22] = instrDDCB__LD_r_SLA_iREGpDD
	/* LD E,SLA (REGISTER+dd) */
	OpcodeDDCBMap[0x23] = instrDDCB__LD_r_SLA_iREGpDD
	/* LD H,SLA (REGISTER+dd) */
	OpcodeDDCBMap[0x24] = instrDDCB__LD_r_SLA_iREGpDD
	/* LD L,SLA (REGISTER+dd) */
	OpcodeDDCBMap[0x25] = instrDDCB__LD_r_SLA_iREGpDD
	/* SLA (REGISTER+dd) */
	OpcodeDDCBMap[0x26] = instrDDCB__SLA_iREGpDD
	/* LD A,SLA (REGISTER+dd) */
	OpcodeDDCBMap[0x27] = instrDDCB__LD_r_SLA_iREGpDD
	/* LD B,SRA (REGISTER+dd) */
	OpcodeDDCBMap[0x28] = instrDDCB__LD_r_SRA_iREGpDD
	/* LD C,SRA (REGISTER+dd) */
	OpcodeDDCBMap[0x29] = instrDDCB__LD_r_SRA_iREGpDD
	/* LD D,SRA (REGISTER+dd) */
	OpcodeDDCBMap[0x2a] = instrDDCB__LD_r_SRA_iREGpDD
	/* LD E,SRA (REGISTER+dd) */
	OpcodeDDCBMap[0x2b] = instrDDCB__LD_r_SRA_iREGpDD
	/* LD H,SRA (REGISTER+dd) */
	OpcodeDDCBMap[0x2c] = instrDDCB__LD_r_SRA_iREGpDD
	/* LD L,SRA (REGISTER+dd) */
	OpcodeDDCBMap[0x2d] = instrDDCB__LD_r_SRA_iREGpDD
	/* SRA (REGISTER+dd) */
	OpcodeDDCBMap[0x2e] = instrDDCB__SRA_iREGpDD
	/* LD A,SRA (REGISTER+dd) */
	OpcodeDDCBMap[0x2f] = instrDDCB__LD_r_SRA_iREGpDD
	/* LD B,SLL (REGISTER+dd) */
	OpcodeDDCBMap[0x30] = instrDDCB__LD_r_SLL_iREGpDD
	/* LD C,SLL (REGISTER+dd) */
	OpcodeDDCBMap[0x31] = instrDDCB__LD_r_SLL_iREGpDD
	/* LD D,SLL (REGISTER+dd) */
	OpcodeDDCBMap[0x32] = instrDDCB__LD_r_SLL_iREGpDD
	/* LD E,SLL (REGISTER+dd) */
	OpcodeDDCBMap[0x33] = instrDDCB__LD_r_SLL_iREGpDD
	/* LD H,SLL (REGISTER+dd) */
	OpcodeDDCBMap[0x34] = instrDDCB__LD_r_SLL_iREGpDD
	/* LD L,SLL (REGISTER+dd) */
	OpcodeDDCBMap[0x35] = instrDDCB__LD_r_SLL_iREGpDD
	/* SLL (REGISTER+dd) */
	OpcodeDDCBMap[0x36] = instrDDCB__SLL_iREGpDD
	/* LD A,SLL (REGISTER+dd) */
	OpcodeDDCBMap[0x37] = instrDDCB__LD_r_SLL_iREGpDD
	/* LD B,SRL (REGISTER+dd) */
	OpcodeDDCBMap[0x38] = instrDDCB__LD_r_SRL_iREGpDD
	/* LD C,SRL (REGISTER+dd) */
	OpcodeDDCBMap[0x39] = instrDDCB__LD_r_SRL_iREGpDD
	/* LD D,SRL (REGISTER+dd) */
	OpcodeDDCBMap[0x3a] = instrDDCB__LD_r_SRL_iREGpDD
	/* LD E,SRL (REGISTER+dd) */
	OpcodeDDCBMap[0x3b] = instrDDCB__LD_r_SRL_iREGpDD
	/* LD H,SRL (REGISTER+dd) */
	OpcodeDDCBMap[0x3c] = instrDDCB__LD_r_SRL_iREGpDD
	/* LD L,SRL (REGISTER+dd) */
	OpcodeDDCBMap[0x3d] = instrDDCB__LD_r_SRL_iREGpDD
	/* SRL (REGISTER+dd) */
	OpcodeDDCBMap[0x3e] = instrDDCB__SRL_iREGpDD
	/* LD A,SRL (REGISTER+dd) */
	OpcodeDDCBMap[0x3f] = instrDDCB__LD_r_SRL_iREGpDD
	// 	/* BIT 0,(REGISTER+dd) */
	// 	OpcodeDDCBMap[0x47] = instrDDCB__BIT_0_iREGpDD
	// 	// Fallthrough cases
	// 	OpcodeDDCBMap[0x40] = OpcodeDDCBMap[0x47]
	// 	OpcodeDDCBMap[0x41] = OpcodeDDCBMap[0x47]
	// 	OpcodeDDCBMap[0x42] = OpcodeDDCBMap[0x47]
	// 	OpcodeDDCBMap[0x43] = OpcodeDDCBMap[0x47]
	// 	OpcodeDDCBMap[0x44] = OpcodeDDCBMap[0x47]
	// 	OpcodeDDCBMap[0x45] = OpcodeDDCBMap[0x47]
	// 	OpcodeDDCBMap[0x46] = OpcodeDDCBMap[0x47]
	// 	/* BIT 1,(REGISTER+dd) */
	// 	OpcodeDDCBMap[0x4f] = instrDDCB__BIT_1_iREGpDD
	// 	// Fallthrough cases
	// 	OpcodeDDCBMap[0x48] = OpcodeDDCBMap[0x4f]
	// 	OpcodeDDCBMap[0x49] = OpcodeDDCBMap[0x4f]
	// 	OpcodeDDCBMap[0x4a] = OpcodeDDCBMap[0x4f]
	// 	OpcodeDDCBMap[0x4b] = OpcodeDDCBMap[0x4f]
	// 	OpcodeDDCBMap[0x4c] = OpcodeDDCBMap[0x4f]
	// 	OpcodeDDCBMap[0x4d] = OpcodeDDCBMap[0x4f]
	// 	OpcodeDDCBMap[0x4e] = OpcodeDDCBMap[0x4f]
	// 	/* BIT 2,(REGISTER+dd) */
	// 	OpcodeDDCBMap[0x57] = instrDDCB__BIT_2_iREGpDD
	// 	// Fallthrough cases
	// 	OpcodeDDCBMap[0x50] = OpcodeDDCBMap[0x57]
	// 	OpcodeDDCBMap[0x51] = OpcodeDDCBMap[0x57]
	// 	OpcodeDDCBMap[0x52] = OpcodeDDCBMap[0x57]
	// 	OpcodeDDCBMap[0x53] = OpcodeDDCBMap[0x57]
	// 	OpcodeDDCBMap[0x54] = OpcodeDDCBMap[0x57]
	// 	OpcodeDDCBMap[0x55] = OpcodeDDCBMap[0x57]
	// 	OpcodeDDCBMap[0x56] = OpcodeDDCBMap[0x57]
	// 	/* BIT 3,(REGISTER+dd) */
	// 	OpcodeDDCBMap[0x5f] = instrDDCB__BIT_3_iREGpDD
	// 	// Fallthrough cases
	// 	OpcodeDDCBMap[0x58] = OpcodeDDCBMap[0x5f]
	// 	OpcodeDDCBMap[0x59] = OpcodeDDCBMap[0x5f]
	// 	OpcodeDDCBMap[0x5a] = OpcodeDDCBMap[0x5f]
	// 	OpcodeDDCBMap[0x5b] = OpcodeDDCBMap[0x5f]
	// 	OpcodeDDCBMap[0x5c] = OpcodeDDCBMap[0x5f]
	// 	OpcodeDDCBMap[0x5d] = OpcodeDDCBMap[0x5f]
	// 	OpcodeDDCBMap[0x5e] = OpcodeDDCBMap[0x5f]
	// 	/* BIT 4,(REGISTER+dd) */
	// 	OpcodeDDCBMap[0x67] = instrDDCB__BIT_4_iREGpDD
	// 	// Fallthrough cases
	// 	OpcodeDDCBMap[0x60] = OpcodeDDCBMap[0x67]
	// 	OpcodeDDCBMap[0x61] = OpcodeDDCBMap[0x67]
	// 	OpcodeDDCBMap[0x62] = OpcodeDDCBMap[0x67]
	// 	OpcodeDDCBMap[0x63] = OpcodeDDCBMap[0x67]
	// 	OpcodeDDCBMap[0x64] = OpcodeDDCBMap[0x67]
	// 	OpcodeDDCBMap[0x65] = OpcodeDDCBMap[0x67]
	// 	OpcodeDDCBMap[0x66] = OpcodeDDCBMap[0x67]
	// 	/* BIT 5,(REGISTER+dd) */
	// 	OpcodeDDCBMap[0x6f] = instrDDCB__BIT_5_iREGpDD
	// 	// Fallthrough cases
	// 	OpcodeDDCBMap[0x68] = OpcodeDDCBMap[0x6f]
	// 	OpcodeDDCBMap[0x69] = OpcodeDDCBMap[0x6f]
	// 	OpcodeDDCBMap[0x6a] = OpcodeDDCBMap[0x6f]
	// 	OpcodeDDCBMap[0x6b] = OpcodeDDCBMap[0x6f]
	// 	OpcodeDDCBMap[0x6c] = OpcodeDDCBMap[0x6f]
	// 	OpcodeDDCBMap[0x6d] = OpcodeDDCBMap[0x6f]
	// 	OpcodeDDCBMap[0x6e] = OpcodeDDCBMap[0x6f]
	// 	/* BIT 6,(REGISTER+dd) */
	// 	OpcodeDDCBMap[0x77] = instrDDCB__BIT_6_iREGpDD
	// 	// Fallthrough cases
	// 	OpcodeDDCBMap[0x70] = OpcodeDDCBMap[0x77]
	// 	OpcodeDDCBMap[0x71] = OpcodeDDCBMap[0x77]
	// 	OpcodeDDCBMap[0x72] = OpcodeDDCBMap[0x77]
	// 	OpcodeDDCBMap[0x73] = OpcodeDDCBMap[0x77]
	// 	OpcodeDDCBMap[0x74] = OpcodeDDCBMap[0x77]
	// 	OpcodeDDCBMap[0x75] = OpcodeDDCBMap[0x77]
	// 	OpcodeDDCBMap[0x76] = OpcodeDDCBMap[0x77]
	// 	/* BIT 7,(REGISTER+dd) */
	// 	OpcodeDDCBMap[0x7f] = instrDDCB__BIT_7_iREGpDD
	// 	// Fallthrough cases
	// 	OpcodeDDCBMap[0x78] = OpcodeDDCBMap[0x7f]
	// 	OpcodeDDCBMap[0x79] = OpcodeDDCBMap[0x7f]
	// 	OpcodeDDCBMap[0x7a] = OpcodeDDCBMap[0x7f]
	// 	OpcodeDDCBMap[0x7b] = OpcodeDDCBMap[0x7f]
	// 	OpcodeDDCBMap[0x7c] = OpcodeDDCBMap[0x7f]
	// 	OpcodeDDCBMap[0x7d] = OpcodeDDCBMap[0x7f]
	// 	OpcodeDDCBMap[0x7e] = OpcodeDDCBMap[0x7f]
	/* LD B,RES 0,(REGISTER+dd) */
	OpcodeDDCBMap[0x80] = instrDDCB__LD_r_RES_b_iREGpDD
	/* LD C,RES 0,(REGISTER+dd) */
	OpcodeDDCBMap[0x81] = instrDDCB__LD_r_RES_b_iREGpDD
	/* LD D,RES 0,(REGISTER+dd) */
	OpcodeDDCBMap[0x82] = instrDDCB__LD_r_RES_b_iREGpDD
	/* LD E,RES 0,(REGISTER+dd) */
	OpcodeDDCBMap[0x83] = instrDDCB__LD_r_RES_b_iREGpDD
	/* LD H,RES 0,(REGISTER+dd) */
	OpcodeDDCBMap[0x84] = instrDDCB__LD_r_RES_b_iREGpDD
	/* LD L,RES 0,(REGISTER+dd) */
	OpcodeDDCBMap[0x85] = instrDDCB__LD_r_RES_b_iREGpDD
	/* RES 0,(REGISTER+dd) */
	OpcodeDDCBMap[0x86] = instrDDCB__RES_0_iREGpDD
	/* LD A,RES 0,(REGISTER+dd) */
	OpcodeDDCBMap[0x87] = instrDDCB__LD_r_RES_b_iREGpDD
	/* LD B,RES 1,(REGISTER+dd) */
	OpcodeDDCBMap[0x88] = instrDDCB__LD_r_RES_b_iREGpDD
	/* LD C,RES 1,(REGISTER+dd) */
	OpcodeDDCBMap[0x89] = instrDDCB__LD_r_RES_b_iREGpDD
	/* LD D,RES 1,(REGISTER+dd) */
	OpcodeDDCBMap[0x8a] = instrDDCB__LD_r_RES_b_iREGpDD
	/* LD E,RES 1,(REGISTER+dd) */
	OpcodeDDCBMap[0x8b] = instrDDCB__LD_r_RES_b_iREGpDD
	/* LD H,RES 1,(REGISTER+dd) */
	OpcodeDDCBMap[0x8c] = instrDDCB__LD_r_RES_b_iREGpDD
	/* LD L,RES 1,(REGISTER+dd) */
	OpcodeDDCBMap[0x8d] = instrDDCB__LD_r_RES_b_iREGpDD
	/* RES 1,(REGISTER+dd) */
	OpcodeDDCBMap[0x8e] = instrDDCB__RES_1_iREGpDD
	/* LD A,RES 1,(REGISTER+dd) */
	OpcodeDDCBMap[0x8f] = instrDDCB__LD_r_RES_b_iREGpDD
	/* LD B,RES 2,(REGISTER+dd) */
	OpcodeDDCBMap[0x90] = instrDDCB__LD_r_RES_b_iREGpDD
	/* LD C,RES 2,(REGISTER+dd) */
	OpcodeDDCBMap[0x91] = instrDDCB__LD_r_RES_b_iREGpDD
	/* LD D,RES 2,(REGISTER+dd) */
	OpcodeDDCBMap[0x92] = instrDDCB__LD_r_RES_b_iREGpDD
	/* LD E,RES 2,(REGISTER+dd) */
	OpcodeDDCBMap[0x93] = instrDDCB__LD_r_RES_b_iREGpDD
	/* LD H,RES 2,(REGISTER+dd) */
	OpcodeDDCBMap[0x94] = instrDDCB__LD_r_RES_b_iREGpDD
	/* LD L,RES 2,(REGISTER+dd) */
	OpcodeDDCBMap[0x95] = instrDDCB__LD_r_RES_b_iREGpDD
	/* RES 2,(REGISTER+dd) */
	OpcodeDDCBMap[0x96] = instrDDCB__RES_2_iREGpDD
	/* LD A,RES 2,(REGISTER+dd) */
	OpcodeDDCBMap[0x97] = instrDDCB__LD_r_RES_b_iREGpDD
	/* LD B,RES 3,(REGISTER+dd) */
	OpcodeDDCBMap[0x98] = instrDDCB__LD_r_RES_b_iREGpDD
	/* LD C,RES 3,(REGISTER+dd) */
	OpcodeDDCBMap[0x99] = instrDDCB__LD_r_RES_b_iREGpDD
	/* LD D,RES 3,(REGISTER+dd) */
	OpcodeDDCBMap[0x9a] = instrDDCB__LD_r_RES_b_iREGpDD
	/* LD E,RES 3,(REGISTER+dd) */
	OpcodeDDCBMap[0x9b] = instrDDCB__LD_r_RES_b_iREGpDD
	/* LD H,RES 3,(REGISTER+dd) */
	OpcodeDDCBMap[0x9c] = instrDDCB__LD_r_RES_b_iREGpDD
	/* LD L,RES 3,(REGISTER+dd) */
	OpcodeDDCBMap[0x9d] = instrDDCB__LD_r_RES_b_iREGpDD
	/* RES 3,(REGISTER+dd) */
	OpcodeDDCBMap[0x9e] = instrDDCB__RES_3_iREGpDD
	/* LD A,RES 3,(REGISTER+dd) */
	OpcodeDDCBMap[0x9f] = instrDDCB__LD_r_RES_b_iREGpDD
	/* LD B,RES 4,(REGISTER+dd) */
	OpcodeDDCBMap[0xa0] = instrDDCB__LD_r_RES_b_iREGpDD
	/* LD C,RES 4,(REGISTER+dd) */
	OpcodeDDCBMap[0xa1] = instrDDCB__LD_r_RES_b_iREGpDD
	/* LD D,RES 4,(REGISTER+dd) */
	OpcodeDDCBMap[0xa2] = instrDDCB__LD_r_RES_b_iREGpDD
	/* LD E,RES 4,(REGISTER+dd) */
	OpcodeDDCBMap[0xa3] = instrDDCB__LD_r_RES_b_iREGpDD
	/* LD H,RES 4,(REGISTER+dd) */
	OpcodeDDCBMap[0xa4] = instrDDCB__LD_r_RES_b_iREGpDD
	/* LD L,RES 4,(REGISTER+dd) */
	OpcodeDDCBMap[0xa5] = instrDDCB__LD_r_RES_b_iREGpDD
	/* RES 4,(REGISTER+dd) */
	OpcodeDDCBMap[0xa6] = instrDDCB__RES_4_iREGpDD
	/* LD A,RES 4,(REGISTER+dd) */
	OpcodeDDCBMap[0xa7] = instrDDCB__LD_r_RES_b_iREGpDD
	/* LD B,RES 5,(REGISTER+dd) */
	OpcodeDDCBMap[0xa8] = instrDDCB__LD_r_RES_b_iREGpDD
	/* LD C,RES 5,(REGISTER+dd) */
	OpcodeDDCBMap[0xa9] = instrDDCB__LD_r_RES_b_iREGpDD
	/* LD D,RES 5,(REGISTER+dd) */
	OpcodeDDCBMap[0xaa] = instrDDCB__LD_r_RES_b_iREGpDD
	/* LD E,RES 5,(REGISTER+dd) */
	OpcodeDDCBMap[0xab] = instrDDCB__LD_r_RES_b_iREGpDD
	/* LD H,RES 5,(REGISTER+dd) */
	OpcodeDDCBMap[0xac] = instrDDCB__LD_r_RES_b_iREGpDD
	/* LD L,RES 5,(REGISTER+dd) */
	OpcodeDDCBMap[0xad] = instrDDCB__LD_r_RES_b_iREGpDD
	/* RES 5,(REGISTER+dd) */
	OpcodeDDCBMap[0xae] = instrDDCB__RES_5_iREGpDD
	/* LD A,RES 5,(REGISTER+dd) */
	OpcodeDDCBMap[0xaf] = instrDDCB__LD_r_RES_b_iREGpDD
	/* LD B,RES 6,(REGISTER+dd) */
	OpcodeDDCBMap[0xb0] = instrDDCB__LD_r_RES_b_iREGpDD
	/* LD C,RES 6,(REGISTER+dd) */
	OpcodeDDCBMap[0xb1] = instrDDCB__LD_r_RES_b_iREGpDD
	/* LD D,RES 6,(REGISTER+dd) */
	OpcodeDDCBMap[0xb2] = instrDDCB__LD_r_RES_b_iREGpDD
	/* LD E,RES 6,(REGISTER+dd) */
	OpcodeDDCBMap[0xb3] = instrDDCB__LD_r_RES_b_iREGpDD
	/* LD H,RES 6,(REGISTER+dd) */
	OpcodeDDCBMap[0xb4] = instrDDCB__LD_r_RES_b_iREGpDD
	/* LD L,RES 6,(REGISTER+dd) */
	OpcodeDDCBMap[0xb5] = instrDDCB__LD_r_RES_b_iREGpDD
	/* RES 6,(REGISTER+dd) */
	OpcodeDDCBMap[0xb6] = instrDDCB__RES_6_iREGpDD
	/* LD A,RES 6,(REGISTER+dd) */
	OpcodeDDCBMap[0xb7] = instrDDCB__LD_r_RES_b_iREGpDD
	/* LD B,RES 7,(REGISTER+dd) */
	OpcodeDDCBMap[0xb8] = instrDDCB__LD_r_RES_b_iREGpDD
	/* LD C,RES 7,(REGISTER+dd) */
	OpcodeDDCBMap[0xb9] = instrDDCB__LD_r_RES_b_iREGpDD
	/* LD D,RES 7,(REGISTER+dd) */
	OpcodeDDCBMap[0xba] = instrDDCB__LD_r_RES_b_iREGpDD
	/* LD E,RES 7,(REGISTER+dd) */
	OpcodeDDCBMap[0xbb] = instrDDCB__LD_r_RES_b_iREGpDD
	/* LD H,RES 7,(REGISTER+dd) */
	OpcodeDDCBMap[0xbc] = instrDDCB__LD_r_RES_b_iREGpDD
	/* LD L,RES 7,(REGISTER+dd) */
	OpcodeDDCBMap[0xbd] = instrDDCB__LD_r_RES_b_iREGpDD
	/* RES 7,(REGISTER+dd) */
	OpcodeDDCBMap[0xbe] = instrDDCB__RES_7_iREGpDD
	/* LD A,RES 7,(REGISTER+dd) */
	OpcodeDDCBMap[0xbf] = instrDDCB__LD_r_RES_b_iREGpDD
	/* LD B,SET 0,(REGISTER+dd) */
	OpcodeDDCBMap[0xc0] = instrDDCB__LD_r_SET_b_iREGpDD
	/* LD C,SET 0,(REGISTER+dd) */
	OpcodeDDCBMap[0xc1] = instrDDCB__LD_r_SET_b_iREGpDD
	/* LD D,SET 0,(REGISTER+dd) */
	OpcodeDDCBMap[0xc2] = instrDDCB__LD_r_SET_b_iREGpDD
	/* LD E,SET 0,(REGISTER+dd) */
	OpcodeDDCBMap[0xc3] = instrDDCB__LD_r_SET_b_iREGpDD
	/* LD H,SET 0,(REGISTER+dd) */
	OpcodeDDCBMap[0xc4] = instrDDCB__LD_r_SET_b_iREGpDD
	/* LD L,SET 0,(REGISTER+dd) */
	OpcodeDDCBMap[0xc5] = instrDDCB__LD_r_SET_b_iREGpDD
	/* SET 0,(REGISTER+dd) */
	OpcodeDDCBMap[0xc6] = instrDDCB__SET_0_iREGpDD
	/* LD A,SET 0,(REGISTER+dd) */
	OpcodeDDCBMap[0xc7] = instrDDCB__LD_r_SET_b_iREGpDD
	/* LD B,SET 1,(REGISTER+dd) */
	OpcodeDDCBMap[0xc8] = instrDDCB__LD_r_SET_b_iREGpDD
	/* LD C,SET 1,(REGISTER+dd) */
	OpcodeDDCBMap[0xc9] = instrDDCB__LD_r_SET_b_iREGpDD
	/* LD D,SET 1,(REGISTER+dd) */
	OpcodeDDCBMap[0xca] = instrDDCB__LD_r_SET_b_iREGpDD
	/* LD E,SET 1,(REGISTER+dd) */
	OpcodeDDCBMap[0xcb] = instrDDCB__LD_r_SET_b_iREGpDD
	/* LD H,SET 1,(REGISTER+dd) */
	OpcodeDDCBMap[0xcc] = instrDDCB__LD_r_SET_b_iREGpDD
	/* LD L,SET 1,(REGISTER+dd) */
	OpcodeDDCBMap[0xcd] = instrDDCB__LD_r_SET_b_iREGpDD
	/* SET 1,(REGISTER+dd) */
	OpcodeDDCBMap[0xce] = instrDDCB__SET_1_iREGpDD
	/* LD A,SET 1,(REGISTER+dd) */
	OpcodeDDCBMap[0xcf] = instrDDCB__LD_r_SET_b_iREGpDD
	/* LD B,SET 2,(REGISTER+dd) */
	OpcodeDDCBMap[0xd0] = instrDDCB__LD_r_SET_b_iREGpDD
	/* LD C,SET 2,(REGISTER+dd) */
	OpcodeDDCBMap[0xd1] = instrDDCB__LD_r_SET_b_iREGpDD
	/* LD D,SET 2,(REGISTER+dd) */
	OpcodeDDCBMap[0xd2] = instrDDCB__LD_r_SET_b_iREGpDD
	/* LD E,SET 2,(REGISTER+dd) */
	OpcodeDDCBMap[0xd3] = instrDDCB__LD_r_SET_b_iREGpDD
	/* LD H,SET 2,(REGISTER+dd) */
	OpcodeDDCBMap[0xd4] = instrDDCB__LD_r_SET_b_iREGpDD
	/* LD L,SET 2,(REGISTER+dd) */
	OpcodeDDCBMap[0xd5] = instrDDCB__LD_r_SET_b_iREGpDD
	/* SET 2,(REGISTER+dd) */
	OpcodeDDCBMap[0xd6] = instrDDCB__SET_2_iREGpDD
	/* LD A,SET 2,(REGISTER+dd) */
	OpcodeDDCBMap[0xd7] = instrDDCB__LD_r_SET_b_iREGpDD
	/* LD B,SET 3,(REGISTER+dd) */
	OpcodeDDCBMap[0xd8] = instrDDCB__LD_r_SET_b_iREGpDD
	/* LD C,SET 3,(REGISTER+dd) */
	OpcodeDDCBMap[0xd9] = instrDDCB__LD_r_SET_b_iREGpDD
	/* LD D,SET 3,(REGISTER+dd) */
	OpcodeDDCBMap[0xda] = instrDDCB__LD_r_SET_b_iREGpDD
	/* LD E,SET 3,(REGISTER+dd) */
	OpcodeDDCBMap[0xdb] = instrDDCB__LD_r_SET_b_iREGpDD
	/* LD H,SET 3,(REGISTER+dd) */
	OpcodeDDCBMap[0xdc] = instrDDCB__LD_r_SET_b_iREGpDD
	/* LD L,SET 3,(REGISTER+dd) */
	OpcodeDDCBMap[0xdd] = instrDDCB__LD_r_SET_b_iREGpDD
	/* SET 3,(REGISTER+dd) */
	OpcodeDDCBMap[0xde] = instrDDCB__SET_3_iREGpDD
	/* LD A,SET 3,(REGISTER+dd) */
	OpcodeDDCBMap[0xdf] = instrDDCB__LD_r_SET_b_iREGpDD
	/* LD B,SET 4,(REGISTER+dd) */
	OpcodeDDCBMap[0xe0] = instrDDCB__LD_r_SET_b_iREGpDD
	/* LD C,SET 4,(REGISTER+dd) */
	OpcodeDDCBMap[0xe1] = instrDDCB__LD_r_SET_b_iREGpDD
	/* LD D,SET 4,(REGISTER+dd) */
	OpcodeDDCBMap[0xe2] = instrDDCB__LD_r_SET_b_iREGpDD
	/* LD E,SET 4,(REGISTER+dd) */
	OpcodeDDCBMap[0xe3] = instrDDCB__LD_r_SET_b_iREGpDD
	/* LD H,SET 4,(REGISTER+dd) */
	OpcodeDDCBMap[0xe4] = instrDDCB__LD_r_SET_b_iREGpDD
	/* LD L,SET 4,(REGISTER+dd) */
	OpcodeDDCBMap[0xe5] = instrDDCB__LD_r_SET_b_iREGpDD
	/* SET 4,(REGISTER+dd) */
	OpcodeDDCBMap[0xe6] = instrDDCB__SET_4_iREGpDD
	/* LD A,SET 4,(REGISTER+dd) */
	OpcodeDDCBMap[0xe7] = instrDDCB__LD_r_SET_b_iREGpDD
	/* LD B,SET 5,(REGISTER+dd) */
	OpcodeDDCBMap[0xe8] = instrDDCB__LD_r_SET_b_iREGpDD
	/* LD C,SET 5,(REGISTER+dd) */
	OpcodeDDCBMap[0xe9] = instrDDCB__LD_r_SET_b_iREGpDD
	/* LD D,SET 5,(REGISTER+dd) */
	OpcodeDDCBMap[0xea] = instrDDCB__LD_r_SET_b_iREGpDD
	/* LD E,SET 5,(REGISTER+dd) */
	OpcodeDDCBMap[0xeb] = instrDDCB__LD_r_SET_b_iREGpDD
	/* LD H,SET 5,(REGISTER+dd) */
	OpcodeDDCBMap[0xec] = instrDDCB__LD_r_SET_b_iREGpDD
	/* LD L,SET 5,(REGISTER+dd) */
	OpcodeDDCBMap[0xed] = instrDDCB__LD_r_SET_b_iREGpDD
	/* SET 5,(REGISTER+dd) */
	OpcodeDDCBMap[0xee] = instrDDCB__SET_5_iREGpDD
	/* LD A,SET 5,(REGISTER+dd) */
	OpcodeDDCBMap[0xef] = instrDDCB__LD_r_SET_b_iREGpDD
	/* LD B,SET 6,(REGISTER+dd) */
	OpcodeDDCBMap[0xf0] = instrDDCB__LD_r_SET_b_iREGpDD
	/* LD C,SET 6,(REGISTER+dd) */
	OpcodeDDCBMap[0xf1] = instrDDCB__LD_r_SET_b_iREGpDD
	/* LD D,SET 6,(REGISTER+dd) */
	OpcodeDDCBMap[0xf2] = instrDDCB__LD_r_SET_b_iREGpDD
	/* LD E,SET 6,(REGISTER+dd) */
	OpcodeDDCBMap[0xf3] = instrDDCB__LD_r_SET_b_iREGpDD
	/* LD H,SET 6,(REGISTER+dd) */
	OpcodeDDCBMap[0xf4] = instrDDCB__LD_r_SET_b_iREGpDD
	/* LD L,SET 6,(REGISTER+dd) */
	OpcodeDDCBMap[0xf5] = instrDDCB__LD_r_SET_b_iREGpDD
	/* SET 6,(REGISTER+dd) */
	OpcodeDDCBMap[0xf6] = instrDDCB__SET_6_iREGpDD
	/* LD A,SET 6,(REGISTER+dd) */
	OpcodeDDCBMap[0xf7] = instrDDCB__LD_r_SET_b_iREGpDD
	/* LD B,SET 7,(REGISTER+dd) */
	OpcodeDDCBMap[0xf8] = instrDDCB__LD_r_SET_b_iREGpDD
	/* LD C,SET 7,(REGISTER+dd) */
	OpcodeDDCBMap[0xf9] = instrDDCB__LD_r_SET_b_iREGpDD
	/* LD D,SET 7,(REGISTER+dd) */
	OpcodeDDCBMap[0xfa] = instrDDCB__LD_r_SET_b_iREGpDD
	/* LD E,SET 7,(REGISTER+dd) */
	OpcodeDDCBMap[0xfb] = instrDDCB__LD_r_SET_b_iREGpDD
	/* LD H,SET 7,(REGISTER+dd) */
	OpcodeDDCBMap[0xfc] = instrDDCB__LD_r_SET_b_iREGpDD
	/* LD L,SET 7,(REGISTER+dd) */
	OpcodeDDCBMap[0xfd] = instrDDCB__LD_r_SET_b_iREGpDD
	/* SET 7,(REGISTER+dd) */
	OpcodeDDCBMap[0xfe] = instrDDCB__SET_7_iREGpDD
	/* LD A,SET 7,(REGISTER+dd) */
	OpcodeDDCBMap[0xff] = instrDDCB__LD_r_SET_b_iREGpDD

}

//--

/* LD r,RLC (REGISTER+dd) */
func instrDDCB__LD_r_RLC_iREGpDD(z *Z80, opcode byte, address uint16) {

	z.Tstates += 15 //FIXME ??

	var pReg *byte = z.GetPrtRegisterValByte(opcode)
	(*pReg) = z.Memory.Read(address)
	//z.Memory.ContendReadNoMreq(address, 1)
	(*pReg) = z.rlc((*pReg))
	z.Memory.Write(address, (*pReg))
}

/* RLC (REGISTER+dd) */
func instrDDCB__RLC_iREGpDD(z *Z80, opcode byte, address uint16) {

	z.Tstates += 15 //FIXME ??

	var bytetemp byte = z.Memory.Read(address)
	//z.Memory.ContendReadNoMreq(address, 1)
	bytetemp = z.rlc(bytetemp)
	z.Memory.Write(address, bytetemp)
}

/* LD r,RRC (REGISTER+dd) */
func instrDDCB__LD_r_RRC_iREGpDD(z *Z80, opcode byte, address uint16) {

	z.Tstates += 15 //FIXME ??

	var pReg *byte = z.GetPrtRegisterValByte(opcode)
	(*pReg) = z.Memory.Read(address)
	//z.Memory.ContendReadNoMreq(address, 1)
	(*pReg) = z.rrc((*pReg))
	z.Memory.Write(address, (*pReg))
}

/* RRC (REGISTER+dd) */
func instrDDCB__RRC_iREGpDD(z *Z80, opcode byte, address uint16) {

	z.Tstates += 15 //FIXME ??

	var bytetemp byte = z.Memory.Read(address)
	//z.Memory.ContendReadNoMreq(address, 1)
	bytetemp = z.rrc(bytetemp)
	z.Memory.Write(address, bytetemp)
}

/* LD B,RL (REGISTER+dd) */
func instrDDCB__LD_r_RL_iREGpDD(z *Z80, opcode byte, address uint16) {

	z.Tstates += 15 //FIXME ??
	var pReg *byte = z.GetPrtRegisterValByte(opcode)

	(*pReg) = z.Memory.Read(address)
	//z.Memory.ContendReadNoMreq(address, 1)
	(*pReg) = z.rl((*pReg))
	z.Memory.Write(address, (*pReg))
}

/* RL (REGISTER+dd) */
func instrDDCB__RL_iREGpDD(z *Z80, opcode byte, address uint16) {

	z.Tstates += 15 //FIXME ??

	var bytetemp byte = z.Memory.Read(address)
	//z.Memory.ContendReadNoMreq(address, 1)
	bytetemp = z.rl(bytetemp)
	z.Memory.Write(address, bytetemp)
}

/* LD r,RR (REGISTER+dd) */
func instrDDCB__LD_r_RR_iREGpDD(z *Z80, opcode byte, address uint16) {

	z.Tstates += 15 //FIXME ??
	var pReg *byte = z.GetPrtRegisterValByte(opcode)

	(*pReg) = z.Memory.Read(address)
	//z.Memory.ContendReadNoMreq(address, 1)
	(*pReg) = z.rr((*pReg))
	z.Memory.Write(address, (*pReg))
}

/* RR (REGISTER+dd) */
func instrDDCB__RR_iREGpDD(z *Z80, opcode byte, address uint16) {

	z.Tstates += 15 //FIXME ??

	var bytetemp byte = z.Memory.Read(address)
	//z.Memory.ContendReadNoMreq(address, 1)
	bytetemp = z.rr(bytetemp)
	z.Memory.Write(address, bytetemp)
}

/* LD r,SLA (REGISTER+dd) */
func instrDDCB__LD_r_SLA_iREGpDD(z *Z80, opcode byte, address uint16) {

	z.Tstates += 15 //FIXME ??
	var pReg *byte = z.GetPrtRegisterValByte(opcode)

	(*pReg) = z.Memory.Read(address)
	//z.Memory.ContendReadNoMreq(address, 1)
	(*pReg) = z.sla((*pReg))
	z.Memory.Write(address, (*pReg))
}

/* SLA (REGISTER+dd) */
func instrDDCB__SLA_iREGpDD(z *Z80, opcode byte, address uint16) {

	z.Tstates += 15 //FIXME ??

	var bytetemp byte = z.Memory.Read(address)
	//z.Memory.ContendReadNoMreq(address, 1)
	bytetemp = z.sla(bytetemp)
	z.Memory.Write(address, bytetemp)
}

/* LD B,SRA (REGISTER+dd) */
func instrDDCB__LD_r_SRA_iREGpDD(z *Z80, opcode byte, address uint16) {

	z.Tstates += 15 //FIXME ??
	var pReg *byte = z.GetPrtRegisterValByte(opcode)

	(*pReg) = z.Memory.Read(address)
	//z.Memory.ContendReadNoMreq(address, 1)
	(*pReg) = z.sra((*pReg))
	z.Memory.Write(address, (*pReg))
}

/* SRA (REGISTER+dd) */
func instrDDCB__SRA_iREGpDD(z *Z80, opcode byte, address uint16) {

	z.Tstates += 15 //FIXME ??

	var bytetemp byte = z.Memory.Read(address)
	//z.Memory.ContendReadNoMreq(address, 1)
	bytetemp = z.sra(bytetemp)
	z.Memory.Write(address, bytetemp)
}

/* LD r,SLL (REGISTER+dd) */
func instrDDCB__LD_r_SLL_iREGpDD(z *Z80, opcode byte, address uint16) {

	z.Tstates += 15 //FIXME ??
	var pReg *byte = z.GetPrtRegisterValByte(opcode)

	(*pReg) = z.Memory.Read(address)
	//z.Memory.ContendReadNoMreq(address, 1)
	(*pReg) = z.sll((*pReg))
	z.Memory.Write(address, (*pReg))
}

/* SLL (REGISTER+dd) */
func instrDDCB__SLL_iREGpDD(z *Z80, opcode byte, address uint16) {

	z.Tstates += 15 //FIXME ??

	var bytetemp byte = z.Memory.Read(address)
	//z.Memory.ContendReadNoMreq(address, 1)
	bytetemp = z.sll(bytetemp)
	z.Memory.Write(address, bytetemp)
}

/* LD r,SRL (REGISTER+dd) */
func instrDDCB__LD_r_SRL_iREGpDD(z *Z80, opcode byte, address uint16) {

	z.Tstates += 15 //FIXME ??
	var pReg *byte = z.GetPrtRegisterValByte(opcode)

	(*pReg) = z.Memory.Read(address)
	//z.Memory.ContendReadNoMreq(address, 1)
	(*pReg) = z.srl((*pReg))
	z.Memory.Write(address, (*pReg))
}

/* SRL (REGISTER+dd) */
func instrDDCB__SRL_iREGpDD(z *Z80, opcode byte, address uint16) {

	z.Tstates += 15 //FIXME ??

	var bytetemp byte = z.Memory.Read(address)
	// z.Memory.ContendReadNoMreq(address, 1)
	bytetemp = z.srl(bytetemp)
	z.Memory.Write(address, bytetemp)
}

// /* BIT 0,(REGISTER+dd) */
// func instrDDCB__BIT_0_iREGpDD(z *Z80, opcode byte, address uint16) {
// 	bytetemp := z.Memory.Read(address)
// 	z.Memory.ContendReadNoMreq(address, 1)
// 	z.biti(0, bytetemp, address)
// }

// /* BIT 1,(REGISTER+dd) */
// func instrDDCB__BIT_1_iREGpDD(z *Z80, opcode byte, address uint16) {
// 	bytetemp := z.Memory.Read(address)
// 	z.Memory.ContendReadNoMreq(address, 1)
// 	z.biti(1, bytetemp, address)
// }

// /* BIT 2,(REGISTER+dd) */
// func instrDDCB__BIT_2_iREGpDD(z *Z80, opcode byte, address uint16) {
// 	bytetemp := z.Memory.Read(address)
// 	z.Memory.ContendReadNoMreq(address, 1)
// 	z.biti(2, bytetemp, address)
// }

// /* BIT 3,(REGISTER+dd) */
// func instrDDCB__BIT_3_iREGpDD(z *Z80, opcode byte, address uint16) {
// 	bytetemp := z.Memory.Read(address)
// 	z.Memory.ContendReadNoMreq(address, 1)
// 	z.biti(3, bytetemp, address)
// }

// /* BIT 4,(REGISTER+dd) */
// func instrDDCB__BIT_4_iREGpDD(z *Z80, opcode byte, address uint16) {
// 	bytetemp := z.Memory.Read(address)
// 	z.Memory.ContendReadNoMreq(address, 1)
// 	z.biti(4, bytetemp, address)
// }

// /* BIT 5,(REGISTER+dd) */
// func instrDDCB__BIT_5_iREGpDD(z *Z80, opcode byte, address uint16) {
// 	bytetemp := z.Memory.Read(address)
// 	z.Memory.ContendReadNoMreq(address, 1)
// 	z.biti(5, bytetemp, address)
// }

// /* BIT 6,(REGISTER+dd) */
// func instrDDCB__BIT_6_iREGpDD(z *Z80, opcode byte, address uint16) {
// 	bytetemp := z.Memory.Read(address)
// 	z.Memory.ContendReadNoMreq(address, 1)
// 	z.biti(6, bytetemp, address)
// }

// /* BIT 7,(REGISTER+dd) */
// func instrDDCB__BIT_7_iREGpDD(z *Z80, opcode byte, address uint16) {
// 	bytetemp := z.Memory.Read(address)
// 	z.Memory.ContendReadNoMreq(address, 1)
// 	z.biti(7, bytetemp, address)
// }

/* LD r,RES b,(REGISTER+dd) */
func instrDDCB__LD_r_RES_b_iREGpDD(z *Z80, opcode byte, address uint16) {

	z.Tstates += 15 //FIXME ??
	var pReg *byte = z.GetPrtRegisterValByte(opcode)

	offset := (opcode >> 4) & 0x07
	mask := getMaskBitReset(offset)

	(*pReg) = z.Memory.Read(address) & mask //  0xfe, 0xfd, 0xfb, 0xf7, 0xef, 0xdf, 0xbf, 0x7f

	//z.Memory.ContendReadNoMreq(address, 1)
	z.Memory.Write(address, (*pReg))
}

/* RES 0,(REGISTER+dd) */
func instrDDCB__RES_0_iREGpDD(z *Z80, opcode byte, address uint16) {
	var bytetemp byte = z.Memory.Read(address)
	// z.Memory.ContendReadNoMreq(address, 1)
	z.Memory.Write(address, bytetemp&0xfe)
}

/* RES 1,(REGISTER+dd) */
func instrDDCB__RES_1_iREGpDD(z *Z80, opcode byte, address uint16) {
	var bytetemp byte = z.Memory.Read(address)
	// z.Memory.ContendReadNoMreq(address, 1)
	z.Memory.Write(address, bytetemp&0xfd)
}

/* RES 2,(REGISTER+dd) */
func instrDDCB__RES_2_iREGpDD(z *Z80, opcode byte, address uint16) {
	var bytetemp byte = z.Memory.Read(address)
	// z.Memory.ContendReadNoMreq(address, 1)
	z.Memory.Write(address, bytetemp&0xfb)
}

/* RES 3,(REGISTER+dd) */
func instrDDCB__RES_3_iREGpDD(z *Z80, opcode byte, address uint16) {
	var bytetemp byte = z.Memory.Read(address)
	// z.Memory.ContendReadNoMreq(address, 1)
	z.Memory.Write(address, bytetemp&0xf7)
}

/* RES 4,(REGISTER+dd) */
func instrDDCB__RES_4_iREGpDD(z *Z80, opcode byte, address uint16) {
	var bytetemp byte = z.Memory.Read(address)
	// z.Memory.ContendReadNoMreq(address, 1)
	z.Memory.Write(address, bytetemp&0xef)
}

/* RES 5,(REGISTER+dd) */
func instrDDCB__RES_5_iREGpDD(z *Z80, opcode byte, address uint16) {
	var bytetemp byte = z.Memory.Read(address)
	//z.Memory.ContendReadNoMreq(address, 1)
	z.Memory.Write(address, bytetemp&0xdf)
}

/* RES 6,(REGISTER+dd) */
func instrDDCB__RES_6_iREGpDD(z *Z80, opcode byte, address uint16) {
	var bytetemp byte = z.Memory.Read(address)
	//z.Memory.ContendReadNoMreq(address, 1)
	z.Memory.Write(address, bytetemp&0xbf)
}

/* RES 7,(REGISTER+dd) */
func instrDDCB__RES_7_iREGpDD(z *Z80, opcode byte, address uint16) {
	var bytetemp byte = z.Memory.Read(address)
	//z.Memory.ContendReadNoMreq(address, 1)
	z.Memory.Write(address, bytetemp&0x7f)
}

/* LD r,SET b,(REGISTER+dd) */
func instrDDCB__LD_r_SET_b_iREGpDD(z *Z80, opcode byte, address uint16) {

	z.Tstates += 15 //FIXME ??
	var pReg *byte = z.GetPrtRegisterValByte(opcode)

	offset := (opcode >> 4) & 0x07
	mask := ^getMaskBitReset(offset)

	(*pReg) = z.Memory.Read(address) | mask
	//z.Memory.ContendReadNoMreq(address, 1)
	z.Memory.Write(address, (*pReg))
}

/* SET 0,(REGISTER+dd) */
func instrDDCB__SET_0_iREGpDD(z *Z80, opcode byte, address uint16) {

	z.Tstates += 15 //FIXME ??

	var bytetemp byte = z.Memory.Read(address)
	//z.Memory.ContendReadNoMreq(address, 1)
	z.Memory.Write(address, bytetemp|0x01)
}

/* SET 1,(REGISTER+dd) */
func instrDDCB__SET_1_iREGpDD(z *Z80, opcode byte, address uint16) {

	z.Tstates += 15 //FIXME ??

	var bytetemp byte = z.Memory.Read(address)
	//z.Memory.ContendReadNoMreq(address, 1)
	z.Memory.Write(address, bytetemp|0x02)
}

/* SET 2,(REGISTER+dd) */
func instrDDCB__SET_2_iREGpDD(z *Z80, opcode byte, address uint16) {

	z.Tstates += 15 //FIXME ??

	var bytetemp byte = z.Memory.Read(address)
	//z.Memory.ContendReadNoMreq(address, 1)
	z.Memory.Write(address, bytetemp|0x04)
}

/* SET 3,(REGISTER+dd) */
func instrDDCB__SET_3_iREGpDD(z *Z80, opcode byte, address uint16) {

	z.Tstates += 15 //FIXME ??

	var bytetemp byte = z.Memory.Read(address)
	//z.Memory.ContendReadNoMreq(address, 1)
	z.Memory.Write(address, bytetemp|0x08)
}

/* SET 4,(REGISTER+dd) */
func instrDDCB__SET_4_iREGpDD(z *Z80, opcode byte, address uint16) {

	z.Tstates += 15 //FIXME ??

	var bytetemp byte = z.Memory.Read(address)
	//z.Memory.ContendReadNoMreq(address, 1)
	z.Memory.Write(address, bytetemp|0x10)
}

/* SET 5,(REGISTER+dd) */
func instrDDCB__SET_5_iREGpDD(z *Z80, opcode byte, address uint16) {

	z.Tstates += 15 //FIXME ??

	var bytetemp byte = z.Memory.Read(address)
	//z.Memory.ContendReadNoMreq(address, 1)
	z.Memory.Write(address, bytetemp|0x20)
}

/* SET 6,(REGISTER+dd) */
func instrDDCB__SET_6_iREGpDD(z *Z80, opcode byte, address uint16) {

	z.Tstates += 15 //FIXME ??

	var bytetemp byte = z.Memory.Read(address)
	//z.Memory.ContendReadNoMreq(address, 1)
	z.Memory.Write(address, bytetemp|0x40)
}

/* SET 7,(REGISTER+dd) */
func instrDDCB__SET_7_iREGpDD(z *Z80, opcode byte, address uint16) {

	z.Tstates += 15 //FIXME ??

	var bytetemp byte = z.Memory.Read(address)
	//z.Memory.ContendReadNoMreq(address, 1)
	z.Memory.Write(address, bytetemp|0x80)
}
