package main

import (
	"ZzEmu"
	"fmt"
)

func incremet(addr *uint16) uint16 {
	(*addr)++
	return *addr
}

func main() {

	var console *ZzEmu.Console = new(ZzEmu.Console)
	var cpu *ZzEmu.Z80 = ZzEmu.CreateCPU(console)

	var addr uint16 = 0

	console.ROM[0] = 0x00               // NOP
	console.ROM[incremet(&addr)] = 0x01 // LD BC, 0x2000
	console.ROM[incremet(&addr)] = 0x00 //
	console.ROM[incremet(&addr)] = 0x20 //

	console.ROM[incremet(&addr)] = 0x03 // inc bc
	console.ROM[incremet(&addr)] = 0x04 // inc b
	console.ROM[incremet(&addr)] = 0x05 // dec b
	console.ROM[incremet(&addr)] = 0x06 // ld b, 13
	console.ROM[incremet(&addr)] = 0x0d

	console.ROM[incremet(&addr)] = 0xed // LD SP,(0x1fff)
	console.ROM[incremet(&addr)] = 0x7b
	console.ROM[incremet(&addr)] = 0xff
	console.ROM[incremet(&addr)] = 0x1f

	console.ROM[incremet(&addr)] = 0x31 // LD SP, 0x1FFF
	console.ROM[incremet(&addr)] = 0xff
	console.ROM[incremet(&addr)] = 0x1f

	cpu.DoOpcode()
	cpu.DoOpcode()
	cpu.DoOpcode()
	cpu.DoOpcode()
	cpu.DoOpcode()
	cpu.DoOpcode()
	cpu.DoOpcode()

	fmt.Println("Valor " + fmt.Sprint(cpu))

}
