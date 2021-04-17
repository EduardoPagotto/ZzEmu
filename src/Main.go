package main

import (
	ZzEmu "ZzEmu/Hardware"
	"fmt"
	"strconv"
)

func main() {

	//pgm := "./examples/hello.bin"
	//pgm := "./examples/pilha.bin"
	//pgm := "./examples/indices.bin"
	pgm := "./examples/zx81.bin" //"./examples/interrup1.bin"

	var rom ZzEmu.DeviceInterface = ZzEmu.NewDeviceMemory(0x0000, 0x0100)
	ram = ZzEmu.NewDeviceMemory(0x0100, 0x0100)

	var console *ZzEmu.Console = ZzEmu.NewConsole()
	console.Bus.AddMemory("RAM", ram)

	size, erro := console.LoadRom(pgm)
	if erro != nil {
		fmt.Println(erro)
		return
	}

	fmt.Println("Carregado Programa: " + pgm + " Tamanho: " + strconv.FormatInt(int64(size), 10))

	console.Exec()

	fmt.Println("Final: " + strconv.FormatInt(int64(console.CPU.Tstates), 10))
	fmt.Println("Dump " + fmt.Sprint(console.CPU))
}
