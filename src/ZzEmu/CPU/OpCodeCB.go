package ZzEmu

func initOpcodeCBMap() {
	// BEGIN of 0xcb shifted opcodes
	/* RLC B */
	OpcodeCBMap[0x00] = instrCB__RLC_r
	/* RLC C */
	OpcodeCBMap[0x01] = instrCB__RLC_r
	/* RLC D */
	OpcodeCBMap[0x02] = instrCB__RLC_r
	/* RLC E */
	OpcodeCBMap[0x03] = instrCB__RLC_r
	/* RLC H */
	OpcodeCBMap[0x04] = instrCB__RLC_r
	/* RLC L */
	OpcodeCBMap[0x05] = instrCB__RLC_r
	/* RLC (HL) */
	OpcodeCBMap[0x06] = instrCB__RLC_iHL
	/* RLC A */
	OpcodeCBMap[0x07] = instrCB__RLC_r
	/* RRC B */
	OpcodeCBMap[0x08] = instrCB__RRC_r
	/* RRC C */
	OpcodeCBMap[0x09] = instrCB__RRC_r
	/* RRC D */
	OpcodeCBMap[0x0a] = instrCB__RRC_r
	/* RRC E */
	OpcodeCBMap[0x0b] = instrCB__RRC_r
	/* RRC H */
	OpcodeCBMap[0x0c] = instrCB__RRC_r
	/* RRC L */
	OpcodeCBMap[0x0d] = instrCB__RRC_r
	/* RRC (HL) */
	OpcodeCBMap[0x0e] = instrCB__RRC_iHL
	/* RRC A */
	OpcodeCBMap[0x0f] = instrCB__RRC_r
	/* RL B */
	OpcodeCBMap[0x10] = instrCB__RL_r
	/* RL C */
	OpcodeCBMap[0x11] = instrCB__RL_r
	/* RL D */
	OpcodeCBMap[0x12] = instrCB__RL_r
	/* RL E */
	OpcodeCBMap[0x13] = instrCB__RL_r
	/* RL H */
	OpcodeCBMap[0x14] = instrCB__RL_r
	/* RL L */
	OpcodeCBMap[0x15] = instrCB__RL_r
	/* RL (HL) */
	OpcodeCBMap[0x16] = instrCB__RL_iHL
	/* RL A */
	OpcodeCBMap[0x17] = instrCB__RL_r
	/* RR B */
	OpcodeCBMap[0x18] = instrCB__RR_r
	/* RR C */
	OpcodeCBMap[0x19] = instrCB__RR_r
	/* RR D */
	OpcodeCBMap[0x1a] = instrCB__RR_r
	/* RR E */
	OpcodeCBMap[0x1b] = instrCB__RR_r
	/* RR H */
	OpcodeCBMap[0x1c] = instrCB__RR_r
	/* RR L */
	OpcodeCBMap[0x1d] = instrCB__RR_r
	/* RR (HL) */
	OpcodeCBMap[0x1e] = instrCB__RR_iHL
	/* RR A */
	OpcodeCBMap[0x1f] = instrCB__RR_r
	/* SLA B */
	OpcodeCBMap[0x20] = instrCB__SLA_r
	/* SLA C */
	OpcodeCBMap[0x21] = instrCB__SLA_r
	/* SLA D */
	OpcodeCBMap[0x22] = instrCB__SLA_r
	/* SLA E */
	OpcodeCBMap[0x23] = instrCB__SLA_r
	/* SLA H */
	OpcodeCBMap[0x24] = instrCB__SLA_r
	/* SLA L */
	OpcodeCBMap[0x25] = instrCB__SLA_r
	/* SLA (HL) */
	OpcodeCBMap[0x26] = instrCB__SLA_iHL
	/* SLA A */
	OpcodeCBMap[0x27] = instrCB__SLA_r
	/* SRA B */
	OpcodeCBMap[0x28] = instrCB__SRA_r
	/* SRA C */
	OpcodeCBMap[0x29] = instrCB__SRA_r
	/* SRA D */
	OpcodeCBMap[0x2a] = instrCB__SRA_r
	/* SRA E */
	OpcodeCBMap[0x2b] = instrCB__SRA_r
	/* SRA H */
	OpcodeCBMap[0x2c] = instrCB__SRA_r
	/* SRA L */
	OpcodeCBMap[0x2d] = instrCB__SRA_r
	/* SRA (HL) */
	OpcodeCBMap[0x2e] = instrCB__SRA_iHL
	/* SRA A */
	OpcodeCBMap[0x2f] = instrCB__SRA_r
	/* SLL B */
	OpcodeCBMap[0x30] = instrCB__SLL_r
	/* SLL C */
	OpcodeCBMap[0x31] = instrCB__SLL_r
	/* SLL D */
	OpcodeCBMap[0x32] = instrCB__SLL_r
	/* SLL E */
	OpcodeCBMap[0x33] = instrCB__SLL_r
	/* SLL H */
	OpcodeCBMap[0x34] = instrCB__SLL_r
	/* SLL L */
	OpcodeCBMap[0x35] = instrCB__SLL_r
	/* SLL (HL) */
	OpcodeCBMap[0x36] = instrCB__SLL_iHL
	/* SLL A */
	OpcodeCBMap[0x37] = instrCB__SLL_r
	/* SRL B */
	OpcodeCBMap[0x38] = instrCB__SRL_r
	/* SRL C */
	OpcodeCBMap[0x39] = instrCB__SRL_r
	/* SRL D */
	OpcodeCBMap[0x3a] = instrCB__SRL_r
	/* SRL E */
	OpcodeCBMap[0x3b] = instrCB__SRL_r
	/* SRL H */
	OpcodeCBMap[0x3c] = instrCB__SRL_r
	/* SRL L */
	OpcodeCBMap[0x3d] = instrCB__SRL_r
	/* SRL (HL) */
	OpcodeCBMap[0x3e] = instrCB__SRL_iHL
	/* SRL A */
	OpcodeCBMap[0x3f] = instrCB__SRL_r
	/* BIT 0,B */
	OpcodeCBMap[0x40] = instrCB__BIT_b_r
	/* BIT 0,C */
	OpcodeCBMap[0x41] = instrCB__BIT_b_r
	/* BIT 0,D */
	OpcodeCBMap[0x42] = instrCB__BIT_b_r
	/* BIT 0,E */
	OpcodeCBMap[0x43] = instrCB__BIT_b_r
	/* BIT 0,H */
	OpcodeCBMap[0x44] = instrCB__BIT_b_r
	/* BIT 0,L */
	OpcodeCBMap[0x45] = instrCB__BIT_b_r
	/* BIT 0,(HL) */
	OpcodeCBMap[0x46] = instrCB__BIT_b_iHL
	/* BIT 0,A */
	OpcodeCBMap[0x47] = instrCB__BIT_b_r
	/* BIT 1,B */
	OpcodeCBMap[0x48] = instrCB__BIT_b_r
	/* BIT 1,C */
	OpcodeCBMap[0x49] = instrCB__BIT_b_r
	/* BIT 1,D */
	OpcodeCBMap[0x4a] = instrCB__BIT_b_r
	/* BIT 1,E */
	OpcodeCBMap[0x4b] = instrCB__BIT_b_r
	/* BIT 1,H */
	OpcodeCBMap[0x4c] = instrCB__BIT_b_r
	/* BIT 1,L */
	OpcodeCBMap[0x4d] = instrCB__BIT_b_r
	/* BIT 1,(HL) */
	OpcodeCBMap[0x4e] = instrCB__BIT_b_iHL
	/* BIT 1,A */
	OpcodeCBMap[0x4f] = instrCB__BIT_b_r
	/* BIT 2,B */
	OpcodeCBMap[0x50] = instrCB__BIT_b_r
	/* BIT 2,C */
	OpcodeCBMap[0x51] = instrCB__BIT_b_r
	/* BIT 2,D */
	OpcodeCBMap[0x52] = instrCB__BIT_b_r
	/* BIT 2,E */
	OpcodeCBMap[0x53] = instrCB__BIT_b_r
	/* BIT 2,H */
	OpcodeCBMap[0x54] = instrCB__BIT_b_r
	/* BIT 2,L */
	OpcodeCBMap[0x55] = instrCB__BIT_b_r
	/* BIT 2,(HL) */
	OpcodeCBMap[0x56] = instrCB__BIT_b_iHL
	/* BIT 2,A */
	OpcodeCBMap[0x57] = instrCB__BIT_b_r
	/* BIT 3,B */
	OpcodeCBMap[0x58] = instrCB__BIT_b_r
	/* BIT 3,C */
	OpcodeCBMap[0x59] = instrCB__BIT_b_r
	/* BIT 3,D */
	OpcodeCBMap[0x5a] = instrCB__BIT_b_r
	/* BIT 3,E */
	OpcodeCBMap[0x5b] = instrCB__BIT_b_r
	/* BIT 3,H */
	OpcodeCBMap[0x5c] = instrCB__BIT_b_r
	/* BIT 3,L */
	OpcodeCBMap[0x5d] = instrCB__BIT_b_r
	/* BIT 3,(HL) */
	OpcodeCBMap[0x5e] = instrCB__BIT_b_iHL
	/* BIT 3,A */
	OpcodeCBMap[0x5f] = instrCB__BIT_b_r
	/* BIT 4,B */
	OpcodeCBMap[0x60] = instrCB__BIT_b_r
	/* BIT 4,C */
	OpcodeCBMap[0x61] = instrCB__BIT_b_r
	/* BIT 4,D */
	OpcodeCBMap[0x62] = instrCB__BIT_b_r
	/* BIT 4,E */
	OpcodeCBMap[0x63] = instrCB__BIT_b_r
	/* BIT 4,H */
	OpcodeCBMap[0x64] = instrCB__BIT_b_r
	/* BIT 4,L */
	OpcodeCBMap[0x65] = instrCB__BIT_b_r
	/* BIT 4,(HL) */
	OpcodeCBMap[0x66] = instrCB__BIT_b_iHL
	/* BIT 4,A */
	OpcodeCBMap[0x67] = instrCB__BIT_b_r
	/* BIT 5,B */
	OpcodeCBMap[0x68] = instrCB__BIT_b_r
	/* BIT 5,C */
	OpcodeCBMap[0x69] = instrCB__BIT_b_r
	/* BIT 5,D */
	OpcodeCBMap[0x6a] = instrCB__BIT_b_r
	/* BIT 5,E */
	OpcodeCBMap[0x6b] = instrCB__BIT_b_r
	/* BIT 5,H */
	OpcodeCBMap[0x6c] = instrCB__BIT_b_r
	/* BIT 5,L */
	OpcodeCBMap[0x6d] = instrCB__BIT_b_r
	/* BIT 5,(HL) */
	OpcodeCBMap[0x6e] = instrCB__BIT_b_iHL
	/* BIT 5,A */
	OpcodeCBMap[0x6f] = instrCB__BIT_b_r
	/* BIT 6,B */
	OpcodeCBMap[0x70] = instrCB__BIT_b_r
	/* BIT 6,C */
	OpcodeCBMap[0x71] = instrCB__BIT_b_r
	/* BIT 6,D */
	OpcodeCBMap[0x72] = instrCB__BIT_b_r
	/* BIT 6,E */
	OpcodeCBMap[0x73] = instrCB__BIT_b_r
	/* BIT 6,H */
	OpcodeCBMap[0x74] = instrCB__BIT_b_r
	/* BIT 6,L */
	OpcodeCBMap[0x75] = instrCB__BIT_b_r
	/* BIT 6,(HL) */
	OpcodeCBMap[0x76] = instrCB__BIT_b_iHL
	/* BIT 6,A */
	OpcodeCBMap[0x77] = instrCB__BIT_b_r
	/* BIT 7,B */
	OpcodeCBMap[0x78] = instrCB__BIT_b_r
	/* BIT 7,C */
	OpcodeCBMap[0x79] = instrCB__BIT_b_r
	/* BIT 7,D */
	OpcodeCBMap[0x7a] = instrCB__BIT_b_r
	/* BIT 7,E */
	OpcodeCBMap[0x7b] = instrCB__BIT_b_r
	/* BIT 7,H */
	OpcodeCBMap[0x7c] = instrCB__BIT_b_r
	/* BIT 7,L */
	OpcodeCBMap[0x7d] = instrCB__BIT_b_r
	/* BIT 7,(HL) */
	OpcodeCBMap[0x7e] = instrCB__BIT_b_iHL
	/* BIT 7,A */
	OpcodeCBMap[0x7f] = instrCB__BIT_b_r
	/* RES 0,B */
	OpcodeCBMap[0x80] = instrCB__RES_b_r
	/* RES 0,C */
	OpcodeCBMap[0x81] = instrCB__RES_b_r
	/* RES 0,D */
	OpcodeCBMap[0x82] = instrCB__RES_b_r
	/* RES 0,E */
	OpcodeCBMap[0x83] = instrCB__RES_b_r
	/* RES 0,H */
	OpcodeCBMap[0x84] = instrCB__RES_b_r
	/* RES 0,L */
	OpcodeCBMap[0x85] = instrCB__RES_b_r
	/* RES 0,(HL) */
	OpcodeCBMap[0x86] = instrCB__RES_b_iHL
	/* RES 0,A */
	OpcodeCBMap[0x87] = instrCB__RES_b_r
	/* RES 1,B */
	OpcodeCBMap[0x88] = instrCB__RES_b_r
	/* RES 1,C */
	OpcodeCBMap[0x89] = instrCB__RES_b_r
	/* RES 1,D */
	OpcodeCBMap[0x8a] = instrCB__RES_b_r
	/* RES 1,E */
	OpcodeCBMap[0x8b] = instrCB__RES_b_r
	/* RES 1,H */
	OpcodeCBMap[0x8c] = instrCB__RES_b_r
	/* RES 1,L */
	OpcodeCBMap[0x8d] = instrCB__RES_b_r
	/* RES 1,(HL) */
	OpcodeCBMap[0x8e] = instrCB__RES_b_iHL
	/* RES 1,A */
	OpcodeCBMap[0x8f] = instrCB__RES_b_r
	/* RES 2,B */
	OpcodeCBMap[0x90] = instrCB__RES_b_r
	/* RES 2,C */
	OpcodeCBMap[0x91] = instrCB__RES_b_r
	/* RES 2,D */
	OpcodeCBMap[0x92] = instrCB__RES_b_r
	/* RES 2,E */
	OpcodeCBMap[0x93] = instrCB__RES_b_r
	/* RES 2,H */
	OpcodeCBMap[0x94] = instrCB__RES_b_r
	/* RES 2,L */
	OpcodeCBMap[0x95] = instrCB__RES_b_r
	/* RES 2,(HL) */
	OpcodeCBMap[0x96] = instrCB__RES_b_iHL
	/* RES 2,A */
	OpcodeCBMap[0x97] = instrCB__RES_b_r
	/* RES 3,B */
	OpcodeCBMap[0x98] = instrCB__RES_b_r
	/* RES 3,C */
	OpcodeCBMap[0x99] = instrCB__RES_b_r
	/* RES 3,D */
	OpcodeCBMap[0x9a] = instrCB__RES_b_r
	/* RES 3,E */
	OpcodeCBMap[0x9b] = instrCB__RES_b_r
	/* RES 3,H */
	OpcodeCBMap[0x9c] = instrCB__RES_b_r
	/* RES 3,L */
	OpcodeCBMap[0x9d] = instrCB__RES_b_r
	/* RES 3,(HL) */
	OpcodeCBMap[0x9e] = instrCB__RES_b_iHL
	/* RES 3,A */
	OpcodeCBMap[0x9f] = instrCB__RES_b_r
	/* RES 4,B */
	OpcodeCBMap[0xa0] = instrCB__RES_b_r
	/* RES 4,C */
	OpcodeCBMap[0xa1] = instrCB__RES_b_r
	/* RES 4,D */
	OpcodeCBMap[0xa2] = instrCB__RES_b_r
	/* RES 4,E */
	OpcodeCBMap[0xa3] = instrCB__RES_b_r
	/* RES 4,H */
	OpcodeCBMap[0xa4] = instrCB__RES_b_r
	/* RES 4,L */
	OpcodeCBMap[0xa5] = instrCB__RES_b_r
	/* RES 4,(HL) */
	OpcodeCBMap[0xa6] = instrCB__RES_b_iHL
	/* RES 4,A */
	OpcodeCBMap[0xa7] = instrCB__RES_b_r
	/* RES 5,B */
	OpcodeCBMap[0xa8] = instrCB__RES_b_r
	/* RES 5,C */
	OpcodeCBMap[0xa9] = instrCB__RES_b_r
	/* RES 5,D */
	OpcodeCBMap[0xaa] = instrCB__RES_b_r
	/* RES 5,E */
	OpcodeCBMap[0xab] = instrCB__RES_b_r
	/* RES 5,H */
	OpcodeCBMap[0xac] = instrCB__RES_b_r
	/* RES 5,L */
	OpcodeCBMap[0xad] = instrCB__RES_b_r
	/* RES 5,(HL) */
	OpcodeCBMap[0xae] = instrCB__RES_b_iHL
	/* RES 5,A */
	OpcodeCBMap[0xaf] = instrCB__RES_b_r
	/* RES 6,B */
	OpcodeCBMap[0xb0] = instrCB__RES_b_r
	/* RES 6,C */
	OpcodeCBMap[0xb1] = instrCB__RES_b_r
	/* RES 6,D */
	OpcodeCBMap[0xb2] = instrCB__RES_b_r
	/* RES 6,E */
	OpcodeCBMap[0xb3] = instrCB__RES_b_r
	/* RES 6,H */
	OpcodeCBMap[0xb4] = instrCB__RES_b_r
	/* RES 6,L */
	OpcodeCBMap[0xb5] = instrCB__RES_b_r
	/* RES 6,(HL) */
	OpcodeCBMap[0xb6] = instrCB__RES_b_iHL
	/* RES 6,A */
	OpcodeCBMap[0xb7] = instrCB__RES_b_r
	/* RES 7,B */
	OpcodeCBMap[0xb8] = instrCB__RES_b_r
	/* RES 7,C */
	OpcodeCBMap[0xb9] = instrCB__RES_b_r
	/* RES 7,D */
	OpcodeCBMap[0xba] = instrCB__RES_b_r
	/* RES 7,E */
	OpcodeCBMap[0xbb] = instrCB__RES_b_r
	/* RES 7,H */
	OpcodeCBMap[0xbc] = instrCB__RES_b_r
	/* RES 7,L */
	OpcodeCBMap[0xbd] = instrCB__RES_b_r
	/* RES 7,(HL) */
	OpcodeCBMap[0xbe] = instrCB__RES_b_iHL
	/* RES 7,A */
	OpcodeCBMap[0xbf] = instrCB__RES_b_r
	/* SET 0,B */
	OpcodeCBMap[0xc0] = instrCB__SET_b_r
	/* SET 0,C */
	OpcodeCBMap[0xc1] = instrCB__SET_b_r
	/* SET 0,D */
	OpcodeCBMap[0xc2] = instrCB__SET_b_r
	/* SET 0,E */
	OpcodeCBMap[0xc3] = instrCB__SET_b_r
	/* SET 0,H */
	OpcodeCBMap[0xc4] = instrCB__SET_b_r
	/* SET 0,L */
	OpcodeCBMap[0xc5] = instrCB__SET_b_r
	/* SET 0,(HL) */
	OpcodeCBMap[0xc6] = instrCB__SET_b_iHL
	/* SET 0,A */
	OpcodeCBMap[0xc7] = instrCB__SET_b_r
	/* SET 1,B */
	OpcodeCBMap[0xc8] = instrCB__SET_b_r
	/* SET 1,C */
	OpcodeCBMap[0xc9] = instrCB__SET_b_r
	/* SET 1,D */
	OpcodeCBMap[0xca] = instrCB__SET_b_r
	/* SET 1,E */
	OpcodeCBMap[0xcb] = instrCB__SET_b_r
	/* SET 1,H */
	OpcodeCBMap[0xcc] = instrCB__SET_b_r
	/* SET 1,L */
	OpcodeCBMap[0xcd] = instrCB__SET_b_r
	/* SET 1,(HL) */
	OpcodeCBMap[0xce] = instrCB__SET_b_iHL
	/* SET 1,A */
	OpcodeCBMap[0xcf] = instrCB__SET_b_r
	/* SET 2,B */
	OpcodeCBMap[0xd0] = instrCB__SET_b_r
	/* SET 2,C */
	OpcodeCBMap[0xd1] = instrCB__SET_b_r
	/* SET 2,D */
	OpcodeCBMap[0xd2] = instrCB__SET_b_r
	/* SET 2,E */
	OpcodeCBMap[0xd3] = instrCB__SET_b_r
	/* SET 2,H */
	OpcodeCBMap[0xd4] = instrCB__SET_b_r
	/* SET 2,L */
	OpcodeCBMap[0xd5] = instrCB__SET_b_r
	/* SET 2,(HL) */
	OpcodeCBMap[0xd6] = instrCB__SET_b_iHL
	/* SET 2,A */
	OpcodeCBMap[0xd7] = instrCB__SET_b_r
	/* SET 3,B */
	OpcodeCBMap[0xd8] = instrCB__SET_b_r
	/* SET 3,C */
	OpcodeCBMap[0xd9] = instrCB__SET_b_r
	/* SET 3,D */
	OpcodeCBMap[0xda] = instrCB__SET_b_r
	/* SET 3,E */
	OpcodeCBMap[0xdb] = instrCB__SET_b_r
	/* SET 3,H */
	OpcodeCBMap[0xdc] = instrCB__SET_b_r
	/* SET 3,L */
	OpcodeCBMap[0xdd] = instrCB__SET_b_r
	/* SET 3,(HL) */
	OpcodeCBMap[0xde] = instrCB__SET_b_iHL
	/* SET 3,A */
	OpcodeCBMap[0xdf] = instrCB__SET_b_r
	/* SET 4,B */
	OpcodeCBMap[0xe0] = instrCB__SET_b_r
	/* SET 4,C */
	OpcodeCBMap[0xe1] = instrCB__SET_b_r
	/* SET 4,D */
	OpcodeCBMap[0xe2] = instrCB__SET_b_r
	/* SET 4,E */
	OpcodeCBMap[0xe3] = instrCB__SET_b_r
	/* SET 4,H */
	OpcodeCBMap[0xe4] = instrCB__SET_b_r
	/* SET 4,L */
	OpcodeCBMap[0xe5] = instrCB__SET_b_r
	/* SET 4,(HL) */
	OpcodeCBMap[0xe6] = instrCB__SET_b_iHL
	/* SET 4,A */
	OpcodeCBMap[0xe7] = instrCB__SET_b_r
	/* SET 5,B */
	OpcodeCBMap[0xe8] = instrCB__SET_b_r
	/* SET 5,C */
	OpcodeCBMap[0xe9] = instrCB__SET_b_r
	/* SET 5,D */
	OpcodeCBMap[0xea] = instrCB__SET_b_r
	/* SET 5,E */
	OpcodeCBMap[0xeb] = instrCB__SET_b_r
	/* SET 5,H */
	OpcodeCBMap[0xec] = instrCB__SET_b_r
	/* SET 5,L */
	OpcodeCBMap[0xed] = instrCB__SET_b_r
	/* SET 5,(HL) */
	OpcodeCBMap[0xee] = instrCB__SET_b_iHL
	/* SET 5,A */
	OpcodeCBMap[0xef] = instrCB__SET_b_r
	/* SET 6,B */
	OpcodeCBMap[0xf0] = instrCB__SET_b_r
	/* SET 6,C */
	OpcodeCBMap[0xf1] = instrCB__SET_b_r
	/* SET 6,D */
	OpcodeCBMap[0xf2] = instrCB__SET_b_r
	/* SET 6,E */
	OpcodeCBMap[0xf3] = instrCB__SET_b_r
	/* SET 6,H */
	OpcodeCBMap[0xf4] = instrCB__SET_b_r
	/* SET 6,L */
	OpcodeCBMap[0xf5] = instrCB__SET_b_r
	/* SET 6,(HL) */
	OpcodeCBMap[0xf6] = instrCB__SET_b_iHL
	/* SET 6,A */
	OpcodeCBMap[0xf7] = instrCB__SET_b_r
	/* SET 7,B */
	OpcodeCBMap[0xf8] = instrCB__SET_b_r
	/* SET 7,C */
	OpcodeCBMap[0xf9] = instrCB__SET_b_r
	/* SET 7,D */
	OpcodeCBMap[0xfa] = instrCB__SET_b_r
	/* SET 7,E */
	OpcodeCBMap[0xfb] = instrCB__SET_b_r
	/* SET 7,H */
	OpcodeCBMap[0xfc] = instrCB__SET_b_r
	/* SET 7,L */
	OpcodeCBMap[0xfd] = instrCB__SET_b_r
	/* SET 7,(HL) */
	OpcodeCBMap[0xfe] = instrCB__SET_b_iHL
	/* SET 7,A */
	OpcodeCBMap[0xff] = instrCB__SET_b_r

	// END of 0xcb shifted opcodes
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
	//z.Memory.ContendReadNoMreq(z.HL.Get(), 1)
	bytetemp = z.rlc(bytetemp)
	z.Memory.Write(z.HL.Get(), bytetemp)
}

/* RRC r */
func instrCB__RRC_r(z *Z80, opcode byte) {
	z.Tstates += 8
	var reg *byte = z.GetPrtRegisterValByte(opcode)
	*reg = z.rrc(*reg)
}

/* RRC (HL) */
func instrCB__RRC_iHL(z *Z80, opcode byte) {
	z.Tstates += 15
	var bytetemp byte = z.Memory.Read(z.HL.Get())
	//z.Memory.ContendReadNoMreq(z.HL.Get(), 1)
	bytetemp = z.rrc(bytetemp)
	z.Memory.Write(z.HL.Get(), bytetemp)
}

/* RL r */
func instrCB__RL_r(z *Z80, opcode byte) {
	z.Tstates += 8
	var reg *byte = z.GetPrtRegisterValByte(opcode)
	*reg = z.rl(*reg)
}

/* RL (HL) */
func instrCB__RL_iHL(z *Z80, opcode byte) {
	z.Tstates += 15
	var bytetemp byte = z.Memory.Read(z.HL.Get())
	//z.Memory.ContendReadNoMreq(z.HL.Get(), 1)
	bytetemp = z.rl(bytetemp)
	z.Memory.Write(z.HL.Get(), bytetemp)
}

/* RR r */
func instrCB__RR_r(z *Z80, opcode byte) {
	z.Tstates += 8
	var reg *byte = z.GetPrtRegisterValByte(opcode)
	*reg = z.rr(*reg)
}

/* RR (HL) */
func instrCB__RR_iHL(z *Z80, opcode byte) {
	z.Tstates += 15
	var bytetemp byte = z.Memory.Read(z.HL.Get())
	//z.Memory.ContendReadNoMreq(z.HL.Get(), 1)
	bytetemp = z.rr(bytetemp)
	z.Memory.Write(z.HL.Get(), bytetemp)
}

/* SLA r */
func instrCB__SLA_r(z *Z80, opcode byte) {
	z.Tstates += 8
	var reg *byte = z.GetPrtRegisterValByte(opcode)
	*reg = z.sla(*reg)
}

/* SLA (HL) */
func instrCB__SLA_iHL(z *Z80, opcode byte) {
	z.Tstates += 15
	var bytetemp byte = z.Memory.Read(z.HL.Get())
	//z.Memory.ContendReadNoMreq(z.HL.Get(), 1)
	bytetemp = z.sla(bytetemp)
	z.Memory.Write(z.HL.Get(), bytetemp)
}

/* SRA r */
func instrCB__SRA_r(z *Z80, opcode byte) {
	z.Tstates += 8
	var reg *byte = z.GetPrtRegisterValByte(opcode)
	*reg = z.sra(*reg)
}

/* SRA (HL) */
func instrCB__SRA_iHL(z *Z80, opcode byte) {
	z.Tstates += 15
	var bytetemp byte = z.Memory.Read(z.HL.Get())
	//z.Memory.ContendReadNoMreq(z.HL.Get(), 1)
	bytetemp = z.sra(bytetemp)
	z.Memory.Write(z.HL.Get(), bytetemp)
}

/* SLL B */
func instrCB__SLL_r(z *Z80, opcode byte) {
	z.Tstates += 8
	var reg *byte = z.GetPrtRegisterValByte(opcode)
	*reg = z.sll(*reg)
}

/* SLL (HL) */
func instrCB__SLL_iHL(z *Z80, opcode byte) {
	z.Tstates += 15
	var bytetemp byte = z.Memory.Read(z.HL.Get())
	//z.Memory.ContendReadNoMreq(z.HL.Get(), 1)
	bytetemp = z.sll(bytetemp)
	z.Memory.Write(z.HL.Get(), bytetemp)
}

/* SRL r */
func instrCB__SRL_r(z *Z80, opcode byte) {
	z.Tstates += 8
	var reg *byte = z.GetPrtRegisterValByte(opcode)
	*reg = z.srl(*reg)
}

/* SRL (HL) */
func instrCB__SRL_iHL(z *Z80, opcode byte) {
	z.Tstates += 15
	var bytetemp byte = z.Memory.Read(z.HL.Get())
	//z.Memory.ContendReadNoMreq(z.HL.Get(), 1)
	bytetemp = z.srl(bytetemp)
	z.Memory.Write(z.HL.Get(), bytetemp)
}

/* BIT b,r */
func instrCB__BIT_b_r(z *Z80, opcode byte) {
	var reg *byte = z.GetPrtRegisterValByte(opcode)
	var bitSel byte = (opcode & 0x38) >> 3
	z.Tstates += 8
	z.bit(bitSel, *reg)
}

/* BIT b,(HL) */
func instrCB__BIT_b_iHL(z *Z80, opcode byte) {
	var bitSel byte = (opcode & 0x38) >> 3
	z.Tstates += 12
	bytetemp := z.Memory.Read(z.HL.Get())
	//z.Memory.ContendReadNoMreq(z.HL.Get(), 1)
	z.bit(bitSel, bytetemp)
}

/* RES b,n */
func instrCB__RES_b_r(z *Z80, opcode byte) {
	z.Tstates += 8
	var reg *byte = z.GetPrtRegisterValByte(opcode)
	var bitSel byte = (opcode & 0x38) >> 3
	*reg &= getMaskBitReset(bitSel)

}

/* RES b,(HL) */
func instrCB__RES_b_iHL(z *Z80, opcode byte) {
	z.Tstates += 15
	var bytetemp byte = z.Memory.Read(z.HL.Get())
	var bitSel byte = (opcode & 0x38) >> 3
	mask := getMaskBitReset(bitSel)
	//z.Memory.ContendReadNoMreq(z.HL.Get(), 1)
	z.Memory.Write(z.HL.Get(), bytetemp&mask)
}

/* SET b,r */
func instrCB__SET_b_r(z *Z80, opcode byte) {
	z.Tstates += 8
	var reg *byte = z.GetPrtRegisterValByte(opcode)
	var bitSel byte = (opcode & 0x38) >> 3
	*reg |= ^getMaskBitReset(bitSel)
}

/* SET b,(HL) */
func instrCB__SET_b_iHL(z *Z80, opcode byte) {

	z.Tstates += 15
	var bytetemp byte = z.Memory.Read(z.HL.Get())
	var bitSel byte = (opcode & 0x38) >> 3
	mask := ^getMaskBitReset(bitSel)
	//z.Memory.ContendReadNoMreq(z.HL.Get(), 1)
	z.Memory.Write(z.HL.Get(), bytetemp|mask)
}
