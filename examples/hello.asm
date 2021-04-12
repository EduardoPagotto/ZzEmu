; Primeiro Teste
org 0x0000
main:   LD SP, 0x0020    ; Top Memoria
        LD HL, 0x1234    ; 
        PUSH HL
        LD DE, 0x5678
        PUSH de
        NOP
        POP HL
        POP DE
        HALT

