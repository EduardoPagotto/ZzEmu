; Primeiro Teste
org 0x8000
start:  
        LD SP, 0x0030    ; Top Memoria
        LD HL, 0x1234    ; 
        LD DE, 0x5678
        PUSH HL
        PUSH DE
        CALL valret
        Jr fim
        nop
        nop
valret:
        POP HL
        POP DE
        RET     
fim:
        nop
        halt
        
; Data
defb    0xff,0xff
;line1:    defb ' primeira,',13,'$'
;line2_3:  defb ' segunda,',13,'terceita,',13,'$'
;line4:    defb ' quarta.',13,13,'$'