; Primeiro Teste
org 0x0000
main:   ld sp,0x1fff
        ld a,1
        ld b,2
        ld c,3
        ld h,b
        ld l,c
        rst main
