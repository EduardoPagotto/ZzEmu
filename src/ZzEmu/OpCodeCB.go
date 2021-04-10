package ZzEmu

func initOpcodeCBMap() {
	// 	// BEGIN of 0xcb shifted opcodes
	// 	/* RLC B */
	OpcodeCBMap[0x00] = instrCB__RLC_r
	// 	/* RLC C */
	OpcodeCBMap[0x01] = instrCB__RLC_r
	// 	/* RLC D */
	OpcodeCBMap[0x02] = instrCB__RLC_r
	// 	/* RLC E */
	OpcodeCBMap[0x03] = instrCB__RLC_r
	// 	/* RLC H */
	OpcodeCBMap[0x04] = instrCB__RLC_r
	// 	/* RLC L */
	OpcodeCBMap[0x05] = instrCB__RLC_r
	// 	/* RLC (HL) */
	OpcodeCBMap[0x06] = instrCB__RLC_iHL
	// 	/* RLC A */
	OpcodeCBMap[0x07] = instrCB__RLC_r
	// 	/* RRC B */
	// 	OpcodeCBMap[0x08] = instrCB__RRC_B
	// 	/* RRC C */
	// 	OpcodeCBMap[0x09] = instrCB__RRC_C
	// 	/* RRC D */
	// 	OpcodeCBMap[0x0a] = instrCB__RRC_D
	// 	/* RRC E */
	// 	OpcodeCBMap[0x0b] = instrCB__RRC_E
	// 	/* RRC H */
	// 	OpcodeCBMap[0x0c] = instrCB__RRC_H
	// 	/* RRC L */
	// 	OpcodeCBMap[0x0d] = instrCB__RRC_L
	// 	/* RRC (HL) */
	// 	OpcodeCBMap[0x0e] = instrCB__RRC_iHL
	// 	/* RRC A */
	// 	OpcodeCBMap[0x0f] = instrCB__RRC_A
	// 	/* RL B */
	// 	OpcodeCBMap[0x10] = instrCB__RL_B
	// 	/* RL C */
	// 	OpcodeCBMap[0x11] = instrCB__RL_C
	// 	/* RL D */
	// 	OpcodeCBMap[0x12] = instrCB__RL_D
	// 	/* RL E */
	// 	OpcodeCBMap[0x13] = instrCB__RL_E
	// 	/* RL H */
	// 	OpcodeCBMap[0x14] = instrCB__RL_H
	// 	/* RL L */
	// 	OpcodeCBMap[0x15] = instrCB__RL_L
	// 	/* RL (HL) */
	// 	OpcodeCBMap[0x16] = instrCB__RL_iHL
	// 	/* RL A */
	// 	OpcodeCBMap[0x17] = instrCB__RL_A
	// 	/* RR B */
	// 	OpcodeCBMap[0x18] = instrCB__RR_B
	// 	/* RR C */
	// 	OpcodeCBMap[0x19] = instrCB__RR_C
	// 	/* RR D */
	// 	OpcodeCBMap[0x1a] = instrCB__RR_D
	// 	/* RR E */
	// 	OpcodeCBMap[0x1b] = instrCB__RR_E
	// 	/* RR H */
	// 	OpcodeCBMap[0x1c] = instrCB__RR_H
	// 	/* RR L */
	// 	OpcodeCBMap[0x1d] = instrCB__RR_L
	// 	/* RR (HL) */
	// 	OpcodeCBMap[0x1e] = instrCB__RR_iHL
	// 	/* RR A */
	// 	OpcodeCBMap[0x1f] = instrCB__RR_A
	// 	/* SLA B */
	// 	OpcodeCBMap[0x20] = instrCB__SLA_B
	// 	/* SLA C */
	// 	OpcodeCBMap[0x21] = instrCB__SLA_C
	// 	/* SLA D */
	// 	OpcodeCBMap[0x22] = instrCB__SLA_D
	// 	/* SLA E */
	// 	OpcodeCBMap[0x23] = instrCB__SLA_E
	// 	/* SLA H */
	// 	OpcodeCBMap[0x24] = instrCB__SLA_H
	// 	/* SLA L */
	// 	OpcodeCBMap[0x25] = instrCB__SLA_L
	// 	/* SLA (HL) */
	// 	OpcodeCBMap[0x26] = instrCB__SLA_iHL
	// 	/* SLA A */
	// 	OpcodeCBMap[0x27] = instrCB__SLA_A
	// 	/* SRA B */
	// 	OpcodeCBMap[0x28] = instrCB__SRA_B
	// 	/* SRA C */
	// 	OpcodeCBMap[0x29] = instrCB__SRA_C
	// 	/* SRA D */
	// 	OpcodeCBMap[0x2a] = instrCB__SRA_D
	// 	/* SRA E */
	// 	OpcodeCBMap[0x2b] = instrCB__SRA_E
	// 	/* SRA H */
	// 	OpcodeCBMap[0x2c] = instrCB__SRA_H
	// 	/* SRA L */
	// 	OpcodeCBMap[0x2d] = instrCB__SRA_L
	// 	/* SRA (HL) */
	// 	OpcodeCBMap[0x2e] = instrCB__SRA_iHL
	// 	/* SRA A */
	// 	OpcodeCBMap[0x2f] = instrCB__SRA_A
	// 	/* SLL B */
	// 	OpcodeCBMap[0x30] = instrCB__SLL_B
	// 	/* SLL C */
	// 	OpcodeCBMap[0x31] = instrCB__SLL_C
	// 	/* SLL D */
	// 	OpcodeCBMap[0x32] = instrCB__SLL_D
	// 	/* SLL E */
	// 	OpcodeCBMap[0x33] = instrCB__SLL_E
	// 	/* SLL H */
	// 	OpcodeCBMap[0x34] = instrCB__SLL_H
	// 	/* SLL L */
	// 	OpcodeCBMap[0x35] = instrCB__SLL_L
	// 	/* SLL (HL) */
	// 	OpcodeCBMap[0x36] = instrCB__SLL_iHL
	// 	/* SLL A */
	// 	OpcodeCBMap[0x37] = instrCB__SLL_A
	// 	/* SRL B */
	// 	OpcodeCBMap[0x38] = instrCB__SRL_B
	// 	/* SRL C */
	// 	OpcodeCBMap[0x39] = instrCB__SRL_C
	// 	/* SRL D */
	// 	OpcodeCBMap[0x3a] = instrCB__SRL_D
	// 	/* SRL E */
	// 	OpcodeCBMap[0x3b] = instrCB__SRL_E
	// 	/* SRL H */
	// 	OpcodeCBMap[0x3c] = instrCB__SRL_H
	// 	/* SRL L */
	// 	OpcodeCBMap[0x3d] = instrCB__SRL_L
	// 	/* SRL (HL) */
	// 	OpcodeCBMap[0x3e] = instrCB__SRL_iHL
	// 	/* SRL A */
	// 	OpcodeCBMap[0x3f] = instrCB__SRL_A
	// 	/* BIT 0,B */
	// 	OpcodeCBMap[0x40] = instrCB__BIT_0_B
	// 	/* BIT 0,C */
	// 	OpcodeCBMap[0x41] = instrCB__BIT_0_C
	// 	/* BIT 0,D */
	// 	OpcodeCBMap[0x42] = instrCB__BIT_0_D
	// 	/* BIT 0,E */
	// 	OpcodeCBMap[0x43] = instrCB__BIT_0_E
	// 	/* BIT 0,H */
	// 	OpcodeCBMap[0x44] = instrCB__BIT_0_H
	// 	/* BIT 0,L */
	// 	OpcodeCBMap[0x45] = instrCB__BIT_0_L
	// 	/* BIT 0,(HL) */
	// 	OpcodeCBMap[0x46] = instrCB__BIT_0_iHL
	// 	/* BIT 0,A */
	// 	OpcodeCBMap[0x47] = instrCB__BIT_0_A
	// 	/* BIT 1,B */
	// 	OpcodeCBMap[0x48] = instrCB__BIT_1_B
	// 	/* BIT 1,C */
	// 	OpcodeCBMap[0x49] = instrCB__BIT_1_C
	// 	/* BIT 1,D */
	// 	OpcodeCBMap[0x4a] = instrCB__BIT_1_D
	// 	/* BIT 1,E */
	// 	OpcodeCBMap[0x4b] = instrCB__BIT_1_E
	// 	/* BIT 1,H */
	// 	OpcodeCBMap[0x4c] = instrCB__BIT_1_H
	// 	/* BIT 1,L */
	// 	OpcodeCBMap[0x4d] = instrCB__BIT_1_L
	// 	/* BIT 1,(HL) */
	// 	OpcodeCBMap[0x4e] = instrCB__BIT_1_iHL
	// 	/* BIT 1,A */
	// 	OpcodeCBMap[0x4f] = instrCB__BIT_1_A
	// 	/* BIT 2,B */
	// 	OpcodeCBMap[0x50] = instrCB__BIT_2_B
	// 	/* BIT 2,C */
	// 	OpcodeCBMap[0x51] = instrCB__BIT_2_C
	// 	/* BIT 2,D */
	// 	OpcodeCBMap[0x52] = instrCB__BIT_2_D
	// 	/* BIT 2,E */
	// 	OpcodeCBMap[0x53] = instrCB__BIT_2_E
	// 	/* BIT 2,H */
	// 	OpcodeCBMap[0x54] = instrCB__BIT_2_H
	// 	/* BIT 2,L */
	// 	OpcodeCBMap[0x55] = instrCB__BIT_2_L
	// 	/* BIT 2,(HL) */
	// 	OpcodeCBMap[0x56] = instrCB__BIT_2_iHL
	// 	/* BIT 2,A */
	// 	OpcodeCBMap[0x57] = instrCB__BIT_2_A
	// 	/* BIT 3,B */
	// 	OpcodeCBMap[0x58] = instrCB__BIT_3_B
	// 	/* BIT 3,C */
	// 	OpcodeCBMap[0x59] = instrCB__BIT_3_C
	// 	/* BIT 3,D */
	// 	OpcodeCBMap[0x5a] = instrCB__BIT_3_D
	// 	/* BIT 3,E */
	// 	OpcodeCBMap[0x5b] = instrCB__BIT_3_E
	// 	/* BIT 3,H */
	// 	OpcodeCBMap[0x5c] = instrCB__BIT_3_H
	// 	/* BIT 3,L */
	// 	OpcodeCBMap[0x5d] = instrCB__BIT_3_L
	// 	/* BIT 3,(HL) */
	// 	OpcodeCBMap[0x5e] = instrCB__BIT_3_iHL
	// 	/* BIT 3,A */
	// 	OpcodeCBMap[0x5f] = instrCB__BIT_3_A
	// 	/* BIT 4,B */
	// 	OpcodeCBMap[0x60] = instrCB__BIT_4_B
	// 	/* BIT 4,C */
	// 	OpcodeCBMap[0x61] = instrCB__BIT_4_C
	// 	/* BIT 4,D */
	// 	OpcodeCBMap[0x62] = instrCB__BIT_4_D
	// 	/* BIT 4,E */
	// 	OpcodeCBMap[0x63] = instrCB__BIT_4_E
	// 	/* BIT 4,H */
	// 	OpcodeCBMap[0x64] = instrCB__BIT_4_H
	// 	/* BIT 4,L */
	// 	OpcodeCBMap[0x65] = instrCB__BIT_4_L
	// 	/* BIT 4,(HL) */
	// 	OpcodeCBMap[0x66] = instrCB__BIT_4_iHL
	// 	/* BIT 4,A */
	// 	OpcodeCBMap[0x67] = instrCB__BIT_4_A
	// 	/* BIT 5,B */
	// 	OpcodeCBMap[0x68] = instrCB__BIT_5_B
	// 	/* BIT 5,C */
	// 	OpcodeCBMap[0x69] = instrCB__BIT_5_C
	// 	/* BIT 5,D */
	// 	OpcodeCBMap[0x6a] = instrCB__BIT_5_D
	// 	/* BIT 5,E */
	// 	OpcodeCBMap[0x6b] = instrCB__BIT_5_E
	// 	/* BIT 5,H */
	// 	OpcodeCBMap[0x6c] = instrCB__BIT_5_H
	// 	/* BIT 5,L */
	// 	OpcodeCBMap[0x6d] = instrCB__BIT_5_L
	// 	/* BIT 5,(HL) */
	// 	OpcodeCBMap[0x6e] = instrCB__BIT_5_iHL
	// 	/* BIT 5,A */
	// 	OpcodeCBMap[0x6f] = instrCB__BIT_5_A
	// 	/* BIT 6,B */
	// 	OpcodeCBMap[0x70] = instrCB__BIT_6_B
	// 	/* BIT 6,C */
	// 	OpcodeCBMap[0x71] = instrCB__BIT_6_C
	// 	/* BIT 6,D */
	// 	OpcodeCBMap[0x72] = instrCB__BIT_6_D
	// 	/* BIT 6,E */
	// 	OpcodeCBMap[0x73] = instrCB__BIT_6_E
	// 	/* BIT 6,H */
	// 	OpcodeCBMap[0x74] = instrCB__BIT_6_H
	// 	/* BIT 6,L */
	// 	OpcodeCBMap[0x75] = instrCB__BIT_6_L
	// 	/* BIT 6,(HL) */
	// 	OpcodeCBMap[0x76] = instrCB__BIT_6_iHL
	// 	/* BIT 6,A */
	// 	OpcodeCBMap[0x77] = instrCB__BIT_6_A
	// 	/* BIT 7,B */
	// 	OpcodeCBMap[0x78] = instrCB__BIT_7_B
	// 	/* BIT 7,C */
	// 	OpcodeCBMap[0x79] = instrCB__BIT_7_C
	// 	/* BIT 7,D */
	// 	OpcodeCBMap[0x7a] = instrCB__BIT_7_D
	// 	/* BIT 7,E */
	// 	OpcodeCBMap[0x7b] = instrCB__BIT_7_E
	// 	/* BIT 7,H */
	// 	OpcodeCBMap[0x7c] = instrCB__BIT_7_H
	// 	/* BIT 7,L */
	// 	OpcodeCBMap[0x7d] = instrCB__BIT_7_L
	// 	/* BIT 7,(HL) */
	// 	OpcodeCBMap[0x7e] = instrCB__BIT_7_iHL
	// 	/* BIT 7,A */
	// 	OpcodeCBMap[0x7f] = instrCB__BIT_7_A
	// 	/* RES 0,B */
	// 	OpcodeCBMap[0x80] = instrCB__RES_0_B
	// 	/* RES 0,C */
	// 	OpcodeCBMap[0x81] = instrCB__RES_0_C
	// 	/* RES 0,D */
	// 	OpcodeCBMap[0x82] = instrCB__RES_0_D
	// 	/* RES 0,E */
	// 	OpcodeCBMap[0x83] = instrCB__RES_0_E
	// 	/* RES 0,H */
	// 	OpcodeCBMap[0x84] = instrCB__RES_0_H
	// 	/* RES 0,L */
	// 	OpcodeCBMap[0x85] = instrCB__RES_0_L
	// 	/* RES 0,(HL) */
	// 	OpcodeCBMap[0x86] = instrCB__RES_0_iHL
	// 	/* RES 0,A */
	// 	OpcodeCBMap[0x87] = instrCB__RES_0_A
	// 	/* RES 1,B */
	// 	OpcodeCBMap[0x88] = instrCB__RES_1_B
	// 	/* RES 1,C */
	// 	OpcodeCBMap[0x89] = instrCB__RES_1_C
	// 	/* RES 1,D */
	// 	OpcodeCBMap[0x8a] = instrCB__RES_1_D
	// 	/* RES 1,E */
	// 	OpcodeCBMap[0x8b] = instrCB__RES_1_E
	// 	/* RES 1,H */
	// 	OpcodeCBMap[0x8c] = instrCB__RES_1_H
	// 	/* RES 1,L */
	// 	OpcodeCBMap[0x8d] = instrCB__RES_1_L
	// 	/* RES 1,(HL) */
	// 	OpcodeCBMap[0x8e] = instrCB__RES_1_iHL
	// 	/* RES 1,A */
	// 	OpcodeCBMap[0x8f] = instrCB__RES_1_A
	// 	/* RES 2,B */
	// 	OpcodeCBMap[0x90] = instrCB__RES_2_B
	// 	/* RES 2,C */
	// 	OpcodeCBMap[0x91] = instrCB__RES_2_C
	// 	/* RES 2,D */
	// 	OpcodeCBMap[0x92] = instrCB__RES_2_D
	// 	/* RES 2,E */
	// 	OpcodeCBMap[0x93] = instrCB__RES_2_E
	// 	/* RES 2,H */
	// 	OpcodeCBMap[0x94] = instrCB__RES_2_H
	// 	/* RES 2,L */
	// 	OpcodeCBMap[0x95] = instrCB__RES_2_L
	// 	/* RES 2,(HL) */
	// 	OpcodeCBMap[0x96] = instrCB__RES_2_iHL
	// 	/* RES 2,A */
	// 	OpcodeCBMap[0x97] = instrCB__RES_2_A
	// 	/* RES 3,B */
	// 	OpcodeCBMap[0x98] = instrCB__RES_3_B
	// 	/* RES 3,C */
	// 	OpcodeCBMap[0x99] = instrCB__RES_3_C
	// 	/* RES 3,D */
	// 	OpcodeCBMap[0x9a] = instrCB__RES_3_D
	// 	/* RES 3,E */
	// 	OpcodeCBMap[0x9b] = instrCB__RES_3_E
	// 	/* RES 3,H */
	// 	OpcodeCBMap[0x9c] = instrCB__RES_3_H
	// 	/* RES 3,L */
	// 	OpcodeCBMap[0x9d] = instrCB__RES_3_L
	// 	/* RES 3,(HL) */
	// 	OpcodeCBMap[0x9e] = instrCB__RES_3_iHL
	// 	/* RES 3,A */
	// 	OpcodeCBMap[0x9f] = instrCB__RES_3_A
	// 	/* RES 4,B */
	// 	OpcodeCBMap[0xa0] = instrCB__RES_4_B
	// 	/* RES 4,C */
	// 	OpcodeCBMap[0xa1] = instrCB__RES_4_C
	// 	/* RES 4,D */
	// 	OpcodeCBMap[0xa2] = instrCB__RES_4_D
	// 	/* RES 4,E */
	// 	OpcodeCBMap[0xa3] = instrCB__RES_4_E
	// 	/* RES 4,H */
	// 	OpcodeCBMap[0xa4] = instrCB__RES_4_H
	// 	/* RES 4,L */
	// 	OpcodeCBMap[0xa5] = instrCB__RES_4_L
	// 	/* RES 4,(HL) */
	// 	OpcodeCBMap[0xa6] = instrCB__RES_4_iHL
	// 	/* RES 4,A */
	// 	OpcodeCBMap[0xa7] = instrCB__RES_4_A
	// 	/* RES 5,B */
	// 	OpcodeCBMap[0xa8] = instrCB__RES_5_B
	// 	/* RES 5,C */
	// 	OpcodeCBMap[0xa9] = instrCB__RES_5_C
	// 	/* RES 5,D */
	// 	OpcodeCBMap[0xaa] = instrCB__RES_5_D
	// 	/* RES 5,E */
	// 	OpcodeCBMap[0xab] = instrCB__RES_5_E
	// 	/* RES 5,H */
	// 	OpcodeCBMap[0xac] = instrCB__RES_5_H
	// 	/* RES 5,L */
	// 	OpcodeCBMap[0xad] = instrCB__RES_5_L
	// 	/* RES 5,(HL) */
	// 	OpcodeCBMap[0xae] = instrCB__RES_5_iHL
	// 	/* RES 5,A */
	// 	OpcodeCBMap[0xaf] = instrCB__RES_5_A
	// 	/* RES 6,B */
	// 	OpcodeCBMap[0xb0] = instrCB__RES_6_B
	// 	/* RES 6,C */
	// 	OpcodeCBMap[0xb1] = instrCB__RES_6_C
	// 	/* RES 6,D */
	// 	OpcodeCBMap[0xb2] = instrCB__RES_6_D
	// 	/* RES 6,E */
	// 	OpcodeCBMap[0xb3] = instrCB__RES_6_E
	// 	/* RES 6,H */
	// 	OpcodeCBMap[0xb4] = instrCB__RES_6_H
	// 	/* RES 6,L */
	// 	OpcodeCBMap[0xb5] = instrCB__RES_6_L
	// 	/* RES 6,(HL) */
	// 	OpcodeCBMap[0xb6] = instrCB__RES_6_iHL
	// 	/* RES 6,A */
	// 	OpcodeCBMap[0xb7] = instrCB__RES_6_A
	// 	/* RES 7,B */
	// 	OpcodeCBMap[0xb8] = instrCB__RES_7_B
	// 	/* RES 7,C */
	// 	OpcodeCBMap[0xb9] = instrCB__RES_7_C
	// 	/* RES 7,D */
	// 	OpcodeCBMap[0xba] = instrCB__RES_7_D
	// 	/* RES 7,E */
	// 	OpcodeCBMap[0xbb] = instrCB__RES_7_E
	// 	/* RES 7,H */
	// 	OpcodeCBMap[0xbc] = instrCB__RES_7_H
	// 	/* RES 7,L */
	// 	OpcodeCBMap[0xbd] = instrCB__RES_7_L
	// 	/* RES 7,(HL) */
	// 	OpcodeCBMap[0xbe] = instrCB__RES_7_iHL
	// 	/* RES 7,A */
	// 	OpcodeCBMap[0xbf] = instrCB__RES_7_A
	// 	/* SET 0,B */
	// 	OpcodeCBMap[0xc0] = instrCB__SET_0_B
	// 	/* SET 0,C */
	// 	OpcodeCBMap[0xc1] = instrCB__SET_0_C
	// 	/* SET 0,D */
	// 	OpcodeCBMap[0xc2] = instrCB__SET_0_D
	// 	/* SET 0,E */
	// 	OpcodeCBMap[0xc3] = instrCB__SET_0_E
	// 	/* SET 0,H */
	// 	OpcodeCBMap[0xc4] = instrCB__SET_0_H
	// 	/* SET 0,L */
	// 	OpcodeCBMap[0xc5] = instrCB__SET_0_L
	// 	/* SET 0,(HL) */
	// 	OpcodeCBMap[0xc6] = instrCB__SET_0_iHL
	// 	/* SET 0,A */
	// 	OpcodeCBMap[0xc7] = instrCB__SET_0_A
	// 	/* SET 1,B */
	// 	OpcodeCBMap[0xc8] = instrCB__SET_1_B
	// 	/* SET 1,C */
	// 	OpcodeCBMap[0xc9] = instrCB__SET_1_C
	// 	/* SET 1,D */
	// 	OpcodeCBMap[0xca] = instrCB__SET_1_D
	// 	/* SET 1,E */
	// 	OpcodeCBMap[0xcb] = instrCB__SET_1_E
	// 	/* SET 1,H */
	// 	OpcodeCBMap[0xcc] = instrCB__SET_1_H
	// 	/* SET 1,L */
	// 	OpcodeCBMap[0xcd] = instrCB__SET_1_L
	// 	/* SET 1,(HL) */
	// 	OpcodeCBMap[0xce] = instrCB__SET_1_iHL
	// 	/* SET 1,A */
	// 	OpcodeCBMap[0xcf] = instrCB__SET_1_A
	// 	/* SET 2,B */
	// 	OpcodeCBMap[0xd0] = instrCB__SET_2_B
	// 	/* SET 2,C */
	// 	OpcodeCBMap[0xd1] = instrCB__SET_2_C
	// 	/* SET 2,D */
	// 	OpcodeCBMap[0xd2] = instrCB__SET_2_D
	// 	/* SET 2,E */
	// 	OpcodeCBMap[0xd3] = instrCB__SET_2_E
	// 	/* SET 2,H */
	// 	OpcodeCBMap[0xd4] = instrCB__SET_2_H
	// 	/* SET 2,L */
	// 	OpcodeCBMap[0xd5] = instrCB__SET_2_L
	// 	/* SET 2,(HL) */
	// 	OpcodeCBMap[0xd6] = instrCB__SET_2_iHL
	// 	/* SET 2,A */
	// 	OpcodeCBMap[0xd7] = instrCB__SET_2_A
	// 	/* SET 3,B */
	// 	OpcodeCBMap[0xd8] = instrCB__SET_3_B
	// 	/* SET 3,C */
	// 	OpcodeCBMap[0xd9] = instrCB__SET_3_C
	// 	/* SET 3,D */
	// 	OpcodeCBMap[0xda] = instrCB__SET_3_D
	// 	/* SET 3,E */
	// 	OpcodeCBMap[0xdb] = instrCB__SET_3_E
	// 	/* SET 3,H */
	// 	OpcodeCBMap[0xdc] = instrCB__SET_3_H
	// 	/* SET 3,L */
	// 	OpcodeCBMap[0xdd] = instrCB__SET_3_L
	// 	/* SET 3,(HL) */
	// 	OpcodeCBMap[0xde] = instrCB__SET_3_iHL
	// 	/* SET 3,A */
	// 	OpcodeCBMap[0xdf] = instrCB__SET_3_A
	// 	/* SET 4,B */
	// 	OpcodeCBMap[0xe0] = instrCB__SET_4_B
	// 	/* SET 4,C */
	// 	OpcodeCBMap[0xe1] = instrCB__SET_4_C
	// 	/* SET 4,D */
	// 	OpcodeCBMap[0xe2] = instrCB__SET_4_D
	// 	/* SET 4,E */
	// 	OpcodeCBMap[0xe3] = instrCB__SET_4_E
	// 	/* SET 4,H */
	// 	OpcodeCBMap[0xe4] = instrCB__SET_4_H
	// 	/* SET 4,L */
	// 	OpcodeCBMap[0xe5] = instrCB__SET_4_L
	// 	/* SET 4,(HL) */
	// 	OpcodeCBMap[0xe6] = instrCB__SET_4_iHL
	// 	/* SET 4,A */
	// 	OpcodeCBMap[0xe7] = instrCB__SET_4_A
	// 	/* SET 5,B */
	// 	OpcodeCBMap[0xe8] = instrCB__SET_5_B
	// 	/* SET 5,C */
	// 	OpcodeCBMap[0xe9] = instrCB__SET_5_C
	// 	/* SET 5,D */
	// 	OpcodeCBMap[0xea] = instrCB__SET_5_D
	// 	/* SET 5,E */
	// 	OpcodeCBMap[0xeb] = instrCB__SET_5_E
	// 	/* SET 5,H */
	// 	OpcodeCBMap[0xec] = instrCB__SET_5_H
	// 	/* SET 5,L */
	// 	OpcodeCBMap[0xed] = instrCB__SET_5_L
	// 	/* SET 5,(HL) */
	// 	OpcodeCBMap[0xee] = instrCB__SET_5_iHL
	// 	/* SET 5,A */
	// 	OpcodeCBMap[0xef] = instrCB__SET_5_A
	// 	/* SET 6,B */
	// 	OpcodeCBMap[0xf0] = instrCB__SET_6_B
	// 	/* SET 6,C */
	// 	OpcodeCBMap[0xf1] = instrCB__SET_6_C
	// 	/* SET 6,D */
	// 	OpcodeCBMap[0xf2] = instrCB__SET_6_D
	// 	/* SET 6,E */
	// 	OpcodeCBMap[0xf3] = instrCB__SET_6_E
	// 	/* SET 6,H */
	// 	OpcodeCBMap[0xf4] = instrCB__SET_6_H
	// 	/* SET 6,L */
	// 	OpcodeCBMap[0xf5] = instrCB__SET_6_L
	// 	/* SET 6,(HL) */
	// 	OpcodeCBMap[0xf6] = instrCB__SET_6_iHL
	// 	/* SET 6,A */
	// 	OpcodeCBMap[0xf7] = instrCB__SET_6_A
	// 	/* SET 7,B */
	// 	OpcodeCBMap[0xf8] = instrCB__SET_7_B
	// 	/* SET 7,C */
	// 	OpcodeCBMap[0xf9] = instrCB__SET_7_C
	// 	/* SET 7,D */
	// 	OpcodeCBMap[0xfa] = instrCB__SET_7_D
	// 	/* SET 7,E */
	// 	OpcodeCBMap[0xfb] = instrCB__SET_7_E
	// 	/* SET 7,H */
	// 	OpcodeCBMap[0xfc] = instrCB__SET_7_H
	// 	/* SET 7,L */
	// 	OpcodeCBMap[0xfd] = instrCB__SET_7_L
	// 	/* SET 7,(HL) */
	// 	OpcodeCBMap[0xfe] = instrCB__SET_7_iHL
	// 	/* SET 7,A */
	// 	OpcodeCBMap[0xff] = instrCB__SET_7_A

	// 	// END of 0xcb shifted opcodes
}

/* RLC r */
func instrCB__RLC_r(z *Z80, opcode byte) {
	z.Tstates += 8
	var reg *byte = z.GetPrtRegisterValByte(opcode)
	*reg = z.rlc(*reg)
}

/* RLC (HL) */
func instrCB__RLC_iHL(z *Z80, opcode byte) {
	z.Tstates += 16
	var bytetemp byte = z.Memory.Read(z.HL.Get())
	//z.memory.ContendReadNoMreq(z.HL(), 1)
	bytetemp = z.rlc(bytetemp)
	z.Memory.Write(z.HL.Get(), bytetemp)
}

// /* RRC B */
// func instrCB__RRC_B(z *Z80, opcode byte) {
// 	z.B = z.rrc(z.B)
// }

// /* RRC C */
// func instrCB__RRC_C(z *Z80, opcode byte) {
// 	z.C = z.rrc(z.C)
// }

// /* RRC D */
// func instrCB__RRC_D(z *Z80, opcode byte) {
// 	z.D = z.rrc(z.D)
// }

// /* RRC E */
// func instrCB__RRC_E(z *Z80, opcode byte) {
// 	z.E = z.rrc(z.E)
// }

// /* RRC H */
// func instrCB__RRC_H(z *Z80, opcode byte) {
// 	z.H = z.rrc(z.H)
// }

// /* RRC L */
// func instrCB__RRC_L(z *Z80, opcode byte) {
// 	z.L = z.rrc(z.L)
// }

// /* RRC (HL) */
// func instrCB__RRC_iHL(z *Z80, opcode byte) {
// 	var bytetemp byte = z.memory.ReadByte(z.HL())
// 	z.memory.ContendReadNoMreq(z.HL(), 1)
// 	bytetemp = z.rrc(bytetemp)
// 	z.memory.WriteByte(z.HL(), bytetemp)
// }

// /* RRC A */
// func instrCB__RRC_A(z *Z80, opcode byte) {
// 	z.A = z.rrc(z.A)
// }

// /* RL B */
// func instrCB__RL_B(z *Z80, opcode byte) {
// 	z.B = z.rl(z.B)
// }

// /* RL C */
// func instrCB__RL_C(z *Z80, opcode byte) {
// 	z.C = z.rl(z.C)
// }

// /* RL D */
// func instrCB__RL_D(z *Z80, opcode byte) {
// 	z.D = z.rl(z.D)
// }

// /* RL E */
// func instrCB__RL_E(z *Z80, opcode byte) {
// 	z.E = z.rl(z.E)
// }

// /* RL H */
// func instrCB__RL_H(z *Z80, opcode byte) {
// 	z.H = z.rl(z.H)
// }

// /* RL L */
// func instrCB__RL_L(z *Z80, opcode byte) {
// 	z.L = z.rl(z.L)
// }

// /* RL (HL) */
// func instrCB__RL_iHL(z *Z80, opcode byte) {
// 	var bytetemp byte = z.memory.ReadByte(z.HL())
// 	z.memory.ContendReadNoMreq(z.HL(), 1)
// 	bytetemp = z.rl(bytetemp)
// 	z.memory.WriteByte(z.HL(), bytetemp)
// }

// /* RL A */
// func instrCB__RL_A(z *Z80, opcode byte) {
// 	z.A = z.rl(z.A)
// }

// /* RR B */
// func instrCB__RR_B(z *Z80, opcode byte) {
// 	z.B = z.rr(z.B)
// }

// /* RR C */
// func instrCB__RR_C(z *Z80, opcode byte) {
// 	z.C = z.rr(z.C)
// }

// /* RR D */
// func instrCB__RR_D(z *Z80, opcode byte) {
// 	z.D = z.rr(z.D)
// }

// /* RR E */
// func instrCB__RR_E(z *Z80, opcode byte) {
// 	z.E = z.rr(z.E)
// }

// /* RR H */
// func instrCB__RR_H(z *Z80, opcode byte) {
// 	z.H = z.rr(z.H)
// }

// /* RR L */
// func instrCB__RR_L(z *Z80, opcode byte) {
// 	z.L = z.rr(z.L)
// }

// /* RR (HL) */
// func instrCB__RR_iHL(z *Z80, opcode byte) {
// 	var bytetemp byte = z.memory.ReadByte(z.HL())
// 	z.memory.ContendReadNoMreq(z.HL(), 1)
// 	bytetemp = z.rr(bytetemp)
// 	z.memory.WriteByte(z.HL(), bytetemp)
// }

// /* RR A */
// func instrCB__RR_A(z *Z80, opcode byte) {
// 	z.A = z.rr(z.A)
// }

// /* SLA B */
// func instrCB__SLA_B(z *Z80, opcode byte) {
// 	z.B = z.sla(z.B)
// }

// /* SLA C */
// func instrCB__SLA_C(z *Z80, opcode byte) {
// 	z.C = z.sla(z.C)
// }

// /* SLA D */
// func instrCB__SLA_D(z *Z80, opcode byte) {
// 	z.D = z.sla(z.D)
// }

// /* SLA E */
// func instrCB__SLA_E(z *Z80, opcode byte) {
// 	z.E = z.sla(z.E)
// }

// /* SLA H */
// func instrCB__SLA_H(z *Z80, opcode byte) {
// 	z.H = z.sla(z.H)
// }

// /* SLA L */
// func instrCB__SLA_L(z *Z80, opcode byte) {
// 	z.L = z.sla(z.L)
// }

// /* SLA (HL) */
// func instrCB__SLA_iHL(z *Z80, opcode byte) {
// 	var bytetemp byte = z.memory.ReadByte(z.HL())
// 	z.memory.ContendReadNoMreq(z.HL(), 1)
// 	bytetemp = z.sla(bytetemp)
// 	z.memory.WriteByte(z.HL(), bytetemp)
// }

// /* SLA A */
// func instrCB__SLA_A(z *Z80, opcode byte) {
// 	z.A = z.sla(z.A)
// }

// /* SRA B */
// func instrCB__SRA_B(z *Z80, opcode byte) {
// 	z.B = z.sra(z.B)
// }

// /* SRA C */
// func instrCB__SRA_C(z *Z80, opcode byte) {
// 	z.C = z.sra(z.C)
// }

// /* SRA D */
// func instrCB__SRA_D(z *Z80, opcode byte) {
// 	z.D = z.sra(z.D)
// }

// /* SRA E */
// func instrCB__SRA_E(z *Z80, opcode byte) {
// 	z.E = z.sra(z.E)
// }

// /* SRA H */
// func instrCB__SRA_H(z *Z80, opcode byte) {
// 	z.H = z.sra(z.H)
// }

// /* SRA L */
// func instrCB__SRA_L(z *Z80, opcode byte) {
// 	z.L = z.sra(z.L)
// }

// /* SRA (HL) */
// func instrCB__SRA_iHL(z *Z80, opcode byte) {
// 	var bytetemp byte = z.memory.ReadByte(z.HL())
// 	z.memory.ContendReadNoMreq(z.HL(), 1)
// 	bytetemp = z.sra(bytetemp)
// 	z.memory.WriteByte(z.HL(), bytetemp)
// }

// /* SRA A */
// func instrCB__SRA_A(z *Z80, opcode byte) {
// 	z.A = z.sra(z.A)
// }

// /* SLL B */
// func instrCB__SLL_B(z *Z80, opcode byte) {
// 	z.B = z.sll(z.B)
// }

// /* SLL C */
// func instrCB__SLL_C(z *Z80, opcode byte) {
// 	z.C = z.sll(z.C)
// }

// /* SLL D */
// func instrCB__SLL_D(z *Z80, opcode byte) {
// 	z.D = z.sll(z.D)
// }

// /* SLL E */
// func instrCB__SLL_E(z *Z80, opcode byte) {
// 	z.E = z.sll(z.E)
// }

// /* SLL H */
// func instrCB__SLL_H(z *Z80, opcode byte) {
// 	z.H = z.sll(z.H)
// }

// /* SLL L */
// func instrCB__SLL_L(z *Z80, opcode byte) {
// 	z.L = z.sll(z.L)
// }

// /* SLL (HL) */
// func instrCB__SLL_iHL(z *Z80, opcode byte) {
// 	var bytetemp byte = z.memory.ReadByte(z.HL())
// 	z.memory.ContendReadNoMreq(z.HL(), 1)
// 	bytetemp = z.sll(bytetemp)
// 	z.memory.WriteByte(z.HL(), bytetemp)
// }

// /* SLL A */
// func instrCB__SLL_A(z *Z80, opcode byte) {
// 	z.A = z.sll(z.A)
// }

// /* SRL B */
// func instrCB__SRL_B(z *Z80, opcode byte) {
// 	z.B = z.srl(z.B)
// }

// /* SRL C */
// func instrCB__SRL_C(z *Z80, opcode byte) {
// 	z.C = z.srl(z.C)
// }

// /* SRL D */
// func instrCB__SRL_D(z *Z80, opcode byte) {
// 	z.D = z.srl(z.D)
// }

// /* SRL E */
// func instrCB__SRL_E(z *Z80, opcode byte) {
// 	z.E = z.srl(z.E)
// }

// /* SRL H */
// func instrCB__SRL_H(z *Z80, opcode byte) {
// 	z.H = z.srl(z.H)
// }

// /* SRL L */
// func instrCB__SRL_L(z *Z80, opcode byte) {
// 	z.L = z.srl(z.L)
// }

// /* SRL (HL) */
// func instrCB__SRL_iHL(z *Z80, opcode byte) {
// 	var bytetemp byte = z.memory.ReadByte(z.HL())
// 	z.memory.ContendReadNoMreq(z.HL(), 1)
// 	bytetemp = z.srl(bytetemp)
// 	z.memory.WriteByte(z.HL(), bytetemp)
// }

// /* SRL A */
// func instrCB__SRL_A(z *Z80, opcode byte) {
// 	z.A = z.srl(z.A)
// }

// /* BIT 0,B */
// func instrCB__BIT_0_B(z *Z80, opcode byte) {
// 	z.bit(0, z.B)
// }

// /* BIT 0,C */
// func instrCB__BIT_0_C(z *Z80, opcode byte) {
// 	z.bit(0, z.C)
// }

// /* BIT 0,D */
// func instrCB__BIT_0_D(z *Z80, opcode byte) {
// 	z.bit(0, z.D)
// }

// /* BIT 0,E */
// func instrCB__BIT_0_E(z *Z80, opcode byte) {
// 	z.bit(0, z.E)
// }

// /* BIT 0,H */
// func instrCB__BIT_0_H(z *Z80, opcode byte) {
// 	z.bit(0, z.H)
// }

// /* BIT 0,L */
// func instrCB__BIT_0_L(z *Z80, opcode byte) {
// 	z.bit(0, z.L)
// }

// /* BIT 0,(HL) */
// func instrCB__BIT_0_iHL(z *Z80, opcode byte) {
// 	bytetemp := z.memory.ReadByte(z.HL())
// 	z.memory.ContendReadNoMreq(z.HL(), 1)
// 	z.bit(0, bytetemp)
// }

// /* BIT 0,A */
// func instrCB__BIT_0_A(z *Z80, opcode byte) {
// 	z.bit(0, z.A)
// }

// /* BIT 1,B */
// func instrCB__BIT_1_B(z *Z80, opcode byte) {
// 	z.bit(1, z.B)
// }

// /* BIT 1,C */
// func instrCB__BIT_1_C(z *Z80, opcode byte) {
// 	z.bit(1, z.C)
// }

// /* BIT 1,D */
// func instrCB__BIT_1_D(z *Z80, opcode byte) {
// 	z.bit(1, z.D)
// }

// /* BIT 1,E */
// func instrCB__BIT_1_E(z *Z80, opcode byte) {
// 	z.bit(1, z.E)
// }

// /* BIT 1,H */
// func instrCB__BIT_1_H(z *Z80, opcode byte) {
// 	z.bit(1, z.H)
// }

// /* BIT 1,L */
// func instrCB__BIT_1_L(z *Z80, opcode byte) {
// 	z.bit(1, z.L)
// }

// /* BIT 1,(HL) */
// func instrCB__BIT_1_iHL(z *Z80, opcode byte) {
// 	bytetemp := z.memory.ReadByte(z.HL())
// 	z.memory.ContendReadNoMreq(z.HL(), 1)
// 	z.bit(1, bytetemp)
// }

// /* BIT 1,A */
// func instrCB__BIT_1_A(z *Z80, opcode byte) {
// 	z.bit(1, z.A)
// }

// /* BIT 2,B */
// func instrCB__BIT_2_B(z *Z80, opcode byte) {
// 	z.bit(2, z.B)
// }

// /* BIT 2,C */
// func instrCB__BIT_2_C(z *Z80, opcode byte) {
// 	z.bit(2, z.C)
// }

// /* BIT 2,D */
// func instrCB__BIT_2_D(z *Z80, opcode byte) {
// 	z.bit(2, z.D)
// }

// /* BIT 2,E */
// func instrCB__BIT_2_E(z *Z80, opcode byte) {
// 	z.bit(2, z.E)
// }

// /* BIT 2,H */
// func instrCB__BIT_2_H(z *Z80, opcode byte) {
// 	z.bit(2, z.H)
// }

// /* BIT 2,L */
// func instrCB__BIT_2_L(z *Z80, opcode byte) {
// 	z.bit(2, z.L)
// }

// /* BIT 2,(HL) */
// func instrCB__BIT_2_iHL(z *Z80, opcode byte) {
// 	bytetemp := z.memory.ReadByte(z.HL())
// 	z.memory.ContendReadNoMreq(z.HL(), 1)
// 	z.bit(2, bytetemp)
// }

// /* BIT 2,A */
// func instrCB__BIT_2_A(z *Z80, opcode byte) {
// 	z.bit(2, z.A)
// }

// /* BIT 3,B */
// func instrCB__BIT_3_B(z *Z80, opcode byte) {
// 	z.bit(3, z.B)
// }

// /* BIT 3,C */
// func instrCB__BIT_3_C(z *Z80, opcode byte) {
// 	z.bit(3, z.C)
// }

// /* BIT 3,D */
// func instrCB__BIT_3_D(z *Z80, opcode byte) {
// 	z.bit(3, z.D)
// }

// /* BIT 3,E */
// func instrCB__BIT_3_E(z *Z80, opcode byte) {
// 	z.bit(3, z.E)
// }

// /* BIT 3,H */
// func instrCB__BIT_3_H(z *Z80, opcode byte) {
// 	z.bit(3, z.H)
// }

// /* BIT 3,L */
// func instrCB__BIT_3_L(z *Z80, opcode byte) {
// 	z.bit(3, z.L)
// }

// /* BIT 3,(HL) */
// func instrCB__BIT_3_iHL(z *Z80, opcode byte) {
// 	bytetemp := z.memory.ReadByte(z.HL())
// 	z.memory.ContendReadNoMreq(z.HL(), 1)
// 	z.bit(3, bytetemp)
// }

// /* BIT 3,A */
// func instrCB__BIT_3_A(z *Z80, opcode byte) {
// 	z.bit(3, z.A)
// }

// /* BIT 4,B */
// func instrCB__BIT_4_B(z *Z80, opcode byte) {
// 	z.bit(4, z.B)
// }

// /* BIT 4,C */
// func instrCB__BIT_4_C(z *Z80, opcode byte) {
// 	z.bit(4, z.C)
// }

// /* BIT 4,D */
// func instrCB__BIT_4_D(z *Z80, opcode byte) {
// 	z.bit(4, z.D)
// }

// /* BIT 4,E */
// func instrCB__BIT_4_E(z *Z80, opcode byte) {
// 	z.bit(4, z.E)
// }

// /* BIT 4,H */
// func instrCB__BIT_4_H(z *Z80, opcode byte) {
// 	z.bit(4, z.H)
// }

// /* BIT 4,L */
// func instrCB__BIT_4_L(z *Z80, opcode byte) {
// 	z.bit(4, z.L)
// }

// /* BIT 4,(HL) */
// func instrCB__BIT_4_iHL(z *Z80, opcode byte) {
// 	bytetemp := z.memory.ReadByte(z.HL())
// 	z.memory.ContendReadNoMreq(z.HL(), 1)
// 	z.bit(4, bytetemp)
// }

// /* BIT 4,A */
// func instrCB__BIT_4_A(z *Z80, opcode byte) {
// 	z.bit(4, z.A)
// }

// /* BIT 5,B */
// func instrCB__BIT_5_B(z *Z80, opcode byte) {
// 	z.bit(5, z.B)
// }

// /* BIT 5,C */
// func instrCB__BIT_5_C(z *Z80, opcode byte) {
// 	z.bit(5, z.C)
// }

// /* BIT 5,D */
// func instrCB__BIT_5_D(z *Z80, opcode byte) {
// 	z.bit(5, z.D)
// }

// /* BIT 5,E */
// func instrCB__BIT_5_E(z *Z80, opcode byte) {
// 	z.bit(5, z.E)
// }

// /* BIT 5,H */
// func instrCB__BIT_5_H(z *Z80, opcode byte) {
// 	z.bit(5, z.H)
// }

// /* BIT 5,L */
// func instrCB__BIT_5_L(z *Z80, opcode byte) {
// 	z.bit(5, z.L)
// }

// /* BIT 5,(HL) */
// func instrCB__BIT_5_iHL(z *Z80, opcode byte) {
// 	bytetemp := z.memory.ReadByte(z.HL())
// 	z.memory.ContendReadNoMreq(z.HL(), 1)
// 	z.bit(5, bytetemp)
// }

// /* BIT 5,A */
// func instrCB__BIT_5_A(z *Z80, opcode byte) {
// 	z.bit(5, z.A)
// }

// /* BIT 6,B */
// func instrCB__BIT_6_B(z *Z80, opcode byte) {
// 	z.bit(6, z.B)
// }

// /* BIT 6,C */
// func instrCB__BIT_6_C(z *Z80, opcode byte) {
// 	z.bit(6, z.C)
// }

// /* BIT 6,D */
// func instrCB__BIT_6_D(z *Z80, opcode byte) {
// 	z.bit(6, z.D)
// }

// /* BIT 6,E */
// func instrCB__BIT_6_E(z *Z80, opcode byte) {
// 	z.bit(6, z.E)
// }

// /* BIT 6,H */
// func instrCB__BIT_6_H(z *Z80, opcode byte) {
// 	z.bit(6, z.H)
// }

// /* BIT 6,L */
// func instrCB__BIT_6_L(z *Z80, opcode byte) {
// 	z.bit(6, z.L)
// }

// /* BIT 6,(HL) */
// func instrCB__BIT_6_iHL(z *Z80, opcode byte) {
// 	bytetemp := z.memory.ReadByte(z.HL())
// 	z.memory.ContendReadNoMreq(z.HL(), 1)
// 	z.bit(6, bytetemp)
// }

// /* BIT 6,A */
// func instrCB__BIT_6_A(z *Z80, opcode byte) {
// 	z.bit(6, z.A)
// }

// /* BIT 7,B */
// func instrCB__BIT_7_B(z *Z80, opcode byte) {
// 	z.bit(7, z.B)
// }

// /* BIT 7,C */
// func instrCB__BIT_7_C(z *Z80, opcode byte) {
// 	z.bit(7, z.C)
// }

// /* BIT 7,D */
// func instrCB__BIT_7_D(z *Z80, opcode byte) {
// 	z.bit(7, z.D)
// }

// /* BIT 7,E */
// func instrCB__BIT_7_E(z *Z80, opcode byte) {
// 	z.bit(7, z.E)
// }

// /* BIT 7,H */
// func instrCB__BIT_7_H(z *Z80, opcode byte) {
// 	z.bit(7, z.H)
// }

// /* BIT 7,L */
// func instrCB__BIT_7_L(z *Z80, opcode byte) {
// 	z.bit(7, z.L)
// }

// /* BIT 7,(HL) */
// func instrCB__BIT_7_iHL(z *Z80, opcode byte) {
// 	bytetemp := z.memory.ReadByte(z.HL())
// 	z.memory.ContendReadNoMreq(z.HL(), 1)
// 	z.bit(7, bytetemp)
// }

// /* BIT 7,A */
// func instrCB__BIT_7_A(z *Z80, opcode byte) {
// 	z.bit(7, z.A)
// }

// /* RES 0,B */
// func instrCB__RES_0_B(z *Z80, opcode byte) {
// 	z.B &= 0xfe
// }

// /* RES 0,C */
// func instrCB__RES_0_C(z *Z80, opcode byte) {
// 	z.C &= 0xfe
// }

// /* RES 0,D */
// func instrCB__RES_0_D(z *Z80, opcode byte) {
// 	z.D &= 0xfe
// }

// /* RES 0,E */
// func instrCB__RES_0_E(z *Z80, opcode byte) {
// 	z.E &= 0xfe
// }

// /* RES 0,H */
// func instrCB__RES_0_H(z *Z80, opcode byte) {
// 	z.H &= 0xfe
// }

// /* RES 0,L */
// func instrCB__RES_0_L(z *Z80, opcode byte) {
// 	z.L &= 0xfe
// }

// /* RES 0,(HL) */
// func instrCB__RES_0_iHL(z *Z80, opcode byte) {
// 	var bytetemp byte = z.memory.ReadByte(z.HL())
// 	z.memory.ContendReadNoMreq(z.HL(), 1)
// 	z.memory.WriteByte(z.HL(), bytetemp&0xfe)
// }

// /* RES 0,A */
// func instrCB__RES_0_A(z *Z80, opcode byte) {
// 	z.A &= 0xfe
// }

// /* RES 1,B */
// func instrCB__RES_1_B(z *Z80, opcode byte) {
// 	z.B &= 0xfd
// }

// /* RES 1,C */
// func instrCB__RES_1_C(z *Z80, opcode byte) {
// 	z.C &= 0xfd
// }

// /* RES 1,D */
// func instrCB__RES_1_D(z *Z80, opcode byte) {
// 	z.D &= 0xfd
// }

// /* RES 1,E */
// func instrCB__RES_1_E(z *Z80, opcode byte) {
// 	z.E &= 0xfd
// }

// /* RES 1,H */
// func instrCB__RES_1_H(z *Z80, opcode byte) {
// 	z.H &= 0xfd
// }

// /* RES 1,L */
// func instrCB__RES_1_L(z *Z80, opcode byte) {
// 	z.L &= 0xfd
// }

// /* RES 1,(HL) */
// func instrCB__RES_1_iHL(z *Z80, opcode byte) {
// 	var bytetemp byte = z.memory.ReadByte(z.HL())
// 	z.memory.ContendReadNoMreq(z.HL(), 1)
// 	z.memory.WriteByte(z.HL(), bytetemp&0xfd)
// }

// /* RES 1,A */
// func instrCB__RES_1_A(z *Z80, opcode byte) {
// 	z.A &= 0xfd
// }

// /* RES 2,B */
// func instrCB__RES_2_B(z *Z80, opcode byte) {
// 	z.B &= 0xfb
// }

// /* RES 2,C */
// func instrCB__RES_2_C(z *Z80, opcode byte) {
// 	z.C &= 0xfb
// }

// /* RES 2,D */
// func instrCB__RES_2_D(z *Z80, opcode byte) {
// 	z.D &= 0xfb
// }

// /* RES 2,E */
// func instrCB__RES_2_E(z *Z80, opcode byte) {
// 	z.E &= 0xfb
// }

// /* RES 2,H */
// func instrCB__RES_2_H(z *Z80, opcode byte) {
// 	z.H &= 0xfb
// }

// /* RES 2,L */
// func instrCB__RES_2_L(z *Z80, opcode byte) {
// 	z.L &= 0xfb
// }

// /* RES 2,(HL) */
// func instrCB__RES_2_iHL(z *Z80, opcode byte) {
// 	var bytetemp byte = z.memory.ReadByte(z.HL())
// 	z.memory.ContendReadNoMreq(z.HL(), 1)
// 	z.memory.WriteByte(z.HL(), bytetemp&0xfb)
// }

// /* RES 2,A */
// func instrCB__RES_2_A(z *Z80, opcode byte) {
// 	z.A &= 0xfb
// }

// /* RES 3,B */
// func instrCB__RES_3_B(z *Z80, opcode byte) {
// 	z.B &= 0xf7
// }

// /* RES 3,C */
// func instrCB__RES_3_C(z *Z80, opcode byte) {
// 	z.C &= 0xf7
// }

// /* RES 3,D */
// func instrCB__RES_3_D(z *Z80, opcode byte) {
// 	z.D &= 0xf7
// }

// /* RES 3,E */
// func instrCB__RES_3_E(z *Z80, opcode byte) {
// 	z.E &= 0xf7
// }

// /* RES 3,H */
// func instrCB__RES_3_H(z *Z80, opcode byte) {
// 	z.H &= 0xf7
// }

// /* RES 3,L */
// func instrCB__RES_3_L(z *Z80, opcode byte) {
// 	z.L &= 0xf7
// }

// /* RES 3,(HL) */
// func instrCB__RES_3_iHL(z *Z80, opcode byte) {
// 	var bytetemp byte = z.memory.ReadByte(z.HL())
// 	z.memory.ContendReadNoMreq(z.HL(), 1)
// 	z.memory.WriteByte(z.HL(), bytetemp&0xf7)
// }

// /* RES 3,A */
// func instrCB__RES_3_A(z *Z80, opcode byte) {
// 	z.A &= 0xf7
// }

// /* RES 4,B */
// func instrCB__RES_4_B(z *Z80, opcode byte) {
// 	z.B &= 0xef
// }

// /* RES 4,C */
// func instrCB__RES_4_C(z *Z80, opcode byte) {
// 	z.C &= 0xef
// }

// /* RES 4,D */
// func instrCB__RES_4_D(z *Z80, opcode byte) {
// 	z.D &= 0xef
// }

// /* RES 4,E */
// func instrCB__RES_4_E(z *Z80, opcode byte) {
// 	z.E &= 0xef
// }

// /* RES 4,H */
// func instrCB__RES_4_H(z *Z80, opcode byte) {
// 	z.H &= 0xef
// }

// /* RES 4,L */
// func instrCB__RES_4_L(z *Z80, opcode byte) {
// 	z.L &= 0xef
// }

// /* RES 4,(HL) */
// func instrCB__RES_4_iHL(z *Z80, opcode byte) {
// 	var bytetemp byte = z.memory.ReadByte(z.HL())
// 	z.memory.ContendReadNoMreq(z.HL(), 1)
// 	z.memory.WriteByte(z.HL(), bytetemp&0xef)
// }

// /* RES 4,A */
// func instrCB__RES_4_A(z *Z80, opcode byte) {
// 	z.A &= 0xef
// }

// /* RES 5,B */
// func instrCB__RES_5_B(z *Z80, opcode byte) {
// 	z.B &= 0xdf
// }

// /* RES 5,C */
// func instrCB__RES_5_C(z *Z80, opcode byte) {
// 	z.C &= 0xdf
// }

// /* RES 5,D */
// func instrCB__RES_5_D(z *Z80, opcode byte) {
// 	z.D &= 0xdf
// }

// /* RES 5,E */
// func instrCB__RES_5_E(z *Z80, opcode byte) {
// 	z.E &= 0xdf
// }

// /* RES 5,H */
// func instrCB__RES_5_H(z *Z80, opcode byte) {
// 	z.H &= 0xdf
// }

// /* RES 5,L */
// func instrCB__RES_5_L(z *Z80, opcode byte) {
// 	z.L &= 0xdf
// }

// /* RES 5,(HL) */
// func instrCB__RES_5_iHL(z *Z80, opcode byte) {
// 	var bytetemp byte = z.memory.ReadByte(z.HL())
// 	z.memory.ContendReadNoMreq(z.HL(), 1)
// 	z.memory.WriteByte(z.HL(), bytetemp&0xdf)
// }

// /* RES 5,A */
// func instrCB__RES_5_A(z *Z80, opcode byte) {
// 	z.A &= 0xdf
// }

// /* RES 6,B */
// func instrCB__RES_6_B(z *Z80, opcode byte) {
// 	z.B &= 0xbf
// }

// /* RES 6,C */
// func instrCB__RES_6_C(z *Z80, opcode byte) {
// 	z.C &= 0xbf
// }

// /* RES 6,D */
// func instrCB__RES_6_D(z *Z80, opcode byte) {
// 	z.D &= 0xbf
// }

// /* RES 6,E */
// func instrCB__RES_6_E(z *Z80, opcode byte) {
// 	z.E &= 0xbf
// }

// /* RES 6,H */
// func instrCB__RES_6_H(z *Z80, opcode byte) {
// 	z.H &= 0xbf
// }

// /* RES 6,L */
// func instrCB__RES_6_L(z *Z80, opcode byte) {
// 	z.L &= 0xbf
// }

// /* RES 6,(HL) */
// func instrCB__RES_6_iHL(z *Z80, opcode byte) {
// 	var bytetemp byte = z.memory.ReadByte(z.HL())
// 	z.memory.ContendReadNoMreq(z.HL(), 1)
// 	z.memory.WriteByte(z.HL(), bytetemp&0xbf)
// }

// /* RES 6,A */
// func instrCB__RES_6_A(z *Z80, opcode byte) {
// 	z.A &= 0xbf
// }

// /* RES 7,B */
// func instrCB__RES_7_B(z *Z80, opcode byte) {
// 	z.B &= 0x7f
// }

// /* RES 7,C */
// func instrCB__RES_7_C(z *Z80, opcode byte) {
// 	z.C &= 0x7f
// }

// /* RES 7,D */
// func instrCB__RES_7_D(z *Z80, opcode byte) {
// 	z.D &= 0x7f
// }

// /* RES 7,E */
// func instrCB__RES_7_E(z *Z80, opcode byte) {
// 	z.E &= 0x7f
// }

// /* RES 7,H */
// func instrCB__RES_7_H(z *Z80, opcode byte) {
// 	z.H &= 0x7f
// }

// /* RES 7,L */
// func instrCB__RES_7_L(z *Z80, opcode byte) {
// 	z.L &= 0x7f
// }

// /* RES 7,(HL) */
// func instrCB__RES_7_iHL(z *Z80, opcode byte) {
// 	var bytetemp byte = z.memory.ReadByte(z.HL())
// 	z.memory.ContendReadNoMreq(z.HL(), 1)
// 	z.memory.WriteByte(z.HL(), bytetemp&0x7f)
// }

// /* RES 7,A */
// func instrCB__RES_7_A(z *Z80, opcode byte) {
// 	z.A &= 0x7f
// }

// /* SET 0,B */
// func instrCB__SET_0_B(z *Z80, opcode byte) {
// 	z.B |= 0x01
// }

// /* SET 0,C */
// func instrCB__SET_0_C(z *Z80, opcode byte) {
// 	z.C |= 0x01
// }

// /* SET 0,D */
// func instrCB__SET_0_D(z *Z80, opcode byte) {
// 	z.D |= 0x01
// }

// /* SET 0,E */
// func instrCB__SET_0_E(z *Z80, opcode byte) {
// 	z.E |= 0x01
// }

// /* SET 0,H */
// func instrCB__SET_0_H(z *Z80, opcode byte) {
// 	z.H |= 0x01
// }

// /* SET 0,L */
// func instrCB__SET_0_L(z *Z80, opcode byte) {
// 	z.L |= 0x01
// }

// /* SET 0,(HL) */
// func instrCB__SET_0_iHL(z *Z80, opcode byte) {
// 	var bytetemp byte = z.memory.ReadByte(z.HL())
// 	z.memory.ContendReadNoMreq(z.HL(), 1)
// 	z.memory.WriteByte(z.HL(), bytetemp|0x01)
// }

// /* SET 0,A */
// func instrCB__SET_0_A(z *Z80, opcode byte) {
// 	z.A |= 0x01
// }

// /* SET 1,B */
// func instrCB__SET_1_B(z *Z80, opcode byte) {
// 	z.B |= 0x02
// }

// /* SET 1,C */
// func instrCB__SET_1_C(z *Z80, opcode byte) {
// 	z.C |= 0x02
// }

// /* SET 1,D */
// func instrCB__SET_1_D(z *Z80, opcode byte) {
// 	z.D |= 0x02
// }

// /* SET 1,E */
// func instrCB__SET_1_E(z *Z80, opcode byte) {
// 	z.E |= 0x02
// }

// /* SET 1,H */
// func instrCB__SET_1_H(z *Z80, opcode byte) {
// 	z.H |= 0x02
// }

// /* SET 1,L */
// func instrCB__SET_1_L(z *Z80, opcode byte) {
// 	z.L |= 0x02
// }

// /* SET 1,(HL) */
// func instrCB__SET_1_iHL(z *Z80, opcode byte) {
// 	var bytetemp byte = z.memory.ReadByte(z.HL())
// 	z.memory.ContendReadNoMreq(z.HL(), 1)
// 	z.memory.WriteByte(z.HL(), bytetemp|0x02)
// }

// /* SET 1,A */
// func instrCB__SET_1_A(z *Z80, opcode byte) {
// 	z.A |= 0x02
// }

// /* SET 2,B */
// func instrCB__SET_2_B(z *Z80, opcode byte) {
// 	z.B |= 0x04
// }

// /* SET 2,C */
// func instrCB__SET_2_C(z *Z80, opcode byte) {
// 	z.C |= 0x04
// }

// /* SET 2,D */
// func instrCB__SET_2_D(z *Z80, opcode byte) {
// 	z.D |= 0x04
// }

// /* SET 2,E */
// func instrCB__SET_2_E(z *Z80, opcode byte) {
// 	z.E |= 0x04
// }

// /* SET 2,H */
// func instrCB__SET_2_H(z *Z80, opcode byte) {
// 	z.H |= 0x04
// }

// /* SET 2,L */
// func instrCB__SET_2_L(z *Z80, opcode byte) {
// 	z.L |= 0x04
// }

// /* SET 2,(HL) */
// func instrCB__SET_2_iHL(z *Z80, opcode byte) {
// 	var bytetemp byte = z.memory.ReadByte(z.HL())
// 	z.memory.ContendReadNoMreq(z.HL(), 1)
// 	z.memory.WriteByte(z.HL(), bytetemp|0x04)
// }

// /* SET 2,A */
// func instrCB__SET_2_A(z *Z80, opcode byte) {
// 	z.A |= 0x04
// }

// /* SET 3,B */
// func instrCB__SET_3_B(z *Z80, opcode byte) {
// 	z.B |= 0x08
// }

// /* SET 3,C */
// func instrCB__SET_3_C(z *Z80, opcode byte) {
// 	z.C |= 0x08
// }

// /* SET 3,D */
// func instrCB__SET_3_D(z *Z80, opcode byte) {
// 	z.D |= 0x08
// }

// /* SET 3,E */
// func instrCB__SET_3_E(z *Z80, opcode byte) {
// 	z.E |= 0x08
// }

// /* SET 3,H */
// func instrCB__SET_3_H(z *Z80, opcode byte) {
// 	z.H |= 0x08
// }

// /* SET 3,L */
// func instrCB__SET_3_L(z *Z80, opcode byte) {
// 	z.L |= 0x08
// }

// /* SET 3,(HL) */
// func instrCB__SET_3_iHL(z *Z80, opcode byte) {
// 	var bytetemp byte = z.memory.ReadByte(z.HL())
// 	z.memory.ContendReadNoMreq(z.HL(), 1)
// 	z.memory.WriteByte(z.HL(), bytetemp|0x08)
// }

// /* SET 3,A */
// func instrCB__SET_3_A(z *Z80, opcode byte) {
// 	z.A |= 0x08
// }

// /* SET 4,B */
// func instrCB__SET_4_B(z *Z80, opcode byte) {
// 	z.B |= 0x10
// }

// /* SET 4,C */
// func instrCB__SET_4_C(z *Z80, opcode byte) {
// 	z.C |= 0x10
// }

// /* SET 4,D */
// func instrCB__SET_4_D(z *Z80, opcode byte) {
// 	z.D |= 0x10
// }

// /* SET 4,E */
// func instrCB__SET_4_E(z *Z80, opcode byte) {
// 	z.E |= 0x10
// }

// /* SET 4,H */
// func instrCB__SET_4_H(z *Z80, opcode byte) {
// 	z.H |= 0x10
// }

// /* SET 4,L */
// func instrCB__SET_4_L(z *Z80, opcode byte) {
// 	z.L |= 0x10
// }

// /* SET 4,(HL) */
// func instrCB__SET_4_iHL(z *Z80, opcode byte) {
// 	var bytetemp byte = z.memory.ReadByte(z.HL())
// 	z.memory.ContendReadNoMreq(z.HL(), 1)
// 	z.memory.WriteByte(z.HL(), bytetemp|0x10)
// }

// /* SET 4,A */
// func instrCB__SET_4_A(z *Z80, opcode byte) {
// 	z.A |= 0x10
// }

// /* SET 5,B */
// func instrCB__SET_5_B(z *Z80, opcode byte) {
// 	z.B |= 0x20
// }

// /* SET 5,C */
// func instrCB__SET_5_C(z *Z80, opcode byte) {
// 	z.C |= 0x20
// }

// /* SET 5,D */
// func instrCB__SET_5_D(z *Z80, opcode byte) {
// 	z.D |= 0x20
// }

// /* SET 5,E */
// func instrCB__SET_5_E(z *Z80, opcode byte) {
// 	z.E |= 0x20
// }

// /* SET 5,H */
// func instrCB__SET_5_H(z *Z80, opcode byte) {
// 	z.H |= 0x20
// }

// /* SET 5,L */
// func instrCB__SET_5_L(z *Z80, opcode byte) {
// 	z.L |= 0x20
// }

// /* SET 5,(HL) */
// func instrCB__SET_5_iHL(z *Z80, opcode byte) {
// 	var bytetemp byte = z.memory.ReadByte(z.HL())
// 	z.memory.ContendReadNoMreq(z.HL(), 1)
// 	z.memory.WriteByte(z.HL(), bytetemp|0x20)
// }

// /* SET 5,A */
// func instrCB__SET_5_A(z *Z80, opcode byte) {
// 	z.A |= 0x20
// }

// /* SET 6,B */
// func instrCB__SET_6_B(z *Z80, opcode byte) {
// 	z.B |= 0x40
// }

// /* SET 6,C */
// func instrCB__SET_6_C(z *Z80, opcode byte) {
// 	z.C |= 0x40
// }

// /* SET 6,D */
// func instrCB__SET_6_D(z *Z80, opcode byte) {
// 	z.D |= 0x40
// }

// /* SET 6,E */
// func instrCB__SET_6_E(z *Z80, opcode byte) {
// 	z.E |= 0x40
// }

// /* SET 6,H */
// func instrCB__SET_6_H(z *Z80, opcode byte) {
// 	z.H |= 0x40
// }

// /* SET 6,L */
// func instrCB__SET_6_L(z *Z80, opcode byte) {
// 	z.L |= 0x40
// }

// /* SET 6,(HL) */
// func instrCB__SET_6_iHL(z *Z80, opcode byte) {
// 	var bytetemp byte = z.memory.ReadByte(z.HL())
// 	z.memory.ContendReadNoMreq(z.HL(), 1)
// 	z.memory.WriteByte(z.HL(), bytetemp|0x40)
// }

// /* SET 6,A */
// func instrCB__SET_6_A(z *Z80, opcode byte) {
// 	z.A |= 0x40
// }

// /* SET 7,B */
// func instrCB__SET_7_B(z *Z80, opcode byte) {
// 	z.B |= 0x80
// }

// /* SET 7,C */
// func instrCB__SET_7_C(z *Z80, opcode byte) {
// 	z.C |= 0x80
// }

// /* SET 7,D */
// func instrCB__SET_7_D(z *Z80, opcode byte) {
// 	z.D |= 0x80
// }

// /* SET 7,E */
// func instrCB__SET_7_E(z *Z80, opcode byte) {
// 	z.E |= 0x80
// }

// /* SET 7,H */
// func instrCB__SET_7_H(z *Z80, opcode byte) {
// 	z.H |= 0x80
// }

// /* SET 7,L */
// func instrCB__SET_7_L(z *Z80, opcode byte) {
// 	z.L |= 0x80
// }

// /* SET 7,(HL) */
// func instrCB__SET_7_iHL(z *Z80, opcode byte) {
// 	var bytetemp byte = z.memory.ReadByte(z.HL())
// 	z.memory.ContendReadNoMreq(z.HL(), 1)
// 	z.memory.WriteByte(z.HL(), bytetemp|0x80)
// }

// /* SET 7,A */
// func instrCB__SET_7_A(z *Z80, opcode byte) {
// 	z.A |= 0x80
// }

// //--
