package main

import (
	"fmt"
	"strconv"

	"github.com/EduardoPagotto/ZzEmu/src/ZzEmu"
)

func main() {

	//pgm := "./examples/hello.bin"
	pgm := "./examples/pilha.bin"
	//pgm := "./examples/indices.bin"
	//pgm := "./examples/zx81.bin" //"./examples/interrup1.bin"

	console := ZzEmu.NewConsole()

	rom := ZzEmu.NewDeviceMemory(0x0000, 0x0100)
	size, erro := rom.LoadRom(pgm)
	if erro != nil {
		fmt.Println(erro)
		return
	}

	fmt.Println("Carregado Programa: " + pgm + " Tamanho: " + strconv.FormatInt(int64(size), 10))

	console.Bus.AddMemory("ROM", rom)
	console.Bus.AddMemory("RAM", ZzEmu.NewDeviceMemory(0x0100, 0x0100))
	console.Bus.AddIO("porta", ZzEmu.NewDeviceLatch(0x0010))
	console.Bus.AddIO("portb", ZzEmu.NewDeviceLatch(0x0011))

	console.Exec()

	fmt.Println("Final: " + strconv.FormatInt(int64(console.CPU.Tstates), 10))
	fmt.Println("Dump " + fmt.Sprint(console.CPU))
}
