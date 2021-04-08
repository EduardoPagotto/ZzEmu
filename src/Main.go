package main

import (
	"ZzEmu"
	"fmt"
)

func main() {

	var console *ZzEmu.Console = new(ZzEmu.Console)
	var cpu *ZzEmu.Z80 = ZzEmu.CreateCPU(console)

	console.ROM[0] = 0x00 // nop
	console.ROM[1] = 0x03 // inc bc
	console.ROM[2] = 0x04 // inc b
	console.ROM[3] = 0x05 // dec b
	console.ROM[4] = 0x06 // ld b, 13
	console.ROM[5] = 0x0d

	cpu.DoOpcode()
	cpu.DoOpcode()
	cpu.DoOpcode()
	cpu.DoOpcode()
	cpu.DoOpcode()
	cpu.DoOpcode()
	cpu.DoOpcode()

	fmt.Println("Valor " + fmt.Sprint(cpu))

}
