; Primeiro Teste
org 0x0000

ram:        equ 0100h   ; Inicio da ram, cfg em Console.TotROM
top:        equ 0200h   ; Fim da ram, cfg em Console.TotROM + Connsole.SizeRAM

start:  LD SP, top      ; fixa o SP no fima da memoria ram
        LD A, 0x10
        LD (ram), A
        LD A, 0x0
        LD A,(ram)
        CP 0x10
        JP NZ,fim
        DI
        ; rst     0x0
fim:    halt
; Data
; ram_top:        defw 0x0030
ds              10, 0xee
defb            0xff,0xff,0xff
;line1:    defb ' primeira,',13,'$'
;line2_3:  defb ' segunda,',13,'terceita,',13,'$'
;line4:    defb ' quarta.',13,13,'$'
end