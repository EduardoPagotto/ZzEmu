package main

import (
	"ZzEmu"
	"fmt"
)

func main() {

	var con *ZzEmu.Console = new(ZzEmu.Console)
	var cpu *ZzEmu.CPU = ZzEmu.CreateCPU(con)

	fmt.Println("Valor " + fmt.Sprint(cpu))

}
