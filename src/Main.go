package main

import (
	"ZzEmu"
	"fmt"
	"strconv"
)

func main() {

	pgm := "./examples/hello.bin"

	var console *ZzEmu.Console = ZzEmu.NewConsole()
	var cpu = console.CPU

	size, erro := console.LoadRom(pgm)
	if erro != nil {
		fmt.Println(erro)
		return
	}

	fmt.Println("Carregado Programa: " + pgm + " Tamanho: " + strconv.FormatInt(int64(size), 10))

	for {
		cpu.DoOpcode()
		if cpu.Halted {
			break
		}

		for {
			address, value, ok := console.Input.ReadAll()
			if ok {
				fmt.Println("TState: " + strconv.FormatInt(int64(cpu.Tstates), 10) +
					" Addr: " + strconv.FormatInt(int64(address), 10) +
					" Val:" + strconv.FormatInt(int64(value), 10))

				//newAddr := address & 0x0f
				console.Output.Write(address+1, value)

			} else {
				break
			}

		}
	}

	fmt.Println("Final: " + strconv.FormatInt(int64(cpu.Tstates), 10))
	fmt.Println("Dump " + fmt.Sprint(cpu))
}
