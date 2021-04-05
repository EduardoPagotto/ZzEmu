package main

import (
	"ZzEmu"
	"fmt"
)

func main() {

	var con *ZzEmu.Console = new(ZzEmu.Console)
	var cpu *ZzEmu.Z80 = ZzEmu.CreateCPU(con)

	fmt.Println("Valor " + fmt.Sprint(cpu))

}
