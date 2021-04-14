package main

import (
	"ZzEmu"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func RetrieveROM(filename string) ([]byte, error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	stats, statsErr := file.Stat()
	if statsErr != nil {
		return nil, statsErr
	}

	var size int64 = stats.Size()
	bytes := make([]byte, size)

	bufr := bufio.NewReader(file)
	_, err = bufr.Read(bytes)

	return bytes, err
}

func main() {

	var buffer []byte
	var erro error

	pgm := "./examples/hello.bin"

	buffer, erro = RetrieveROM(pgm)
	if erro != nil {
		fmt.Println(erro)
		return
	}

	tot := len(buffer)
	fmt.Println("Carregado Programa: " + pgm + " Tamanho: " + strconv.FormatInt(int64(tot), 10))

	var console *ZzEmu.Console = ZzEmu.NewConsole()
	var cpu = console.CPU

	copy(console.ROM[:], buffer)

	for {
		cpu.DoOpcode()
		if cpu.Halted {
			break
		}

		if len(console.Input) > 0 {
			for k, v := range console.Input {
				fmt.Println("TState: " + strconv.FormatInt(int64(cpu.Tstates), 10) +
					" Addr: " + strconv.FormatInt(int64(k), 10) +
					" Val:" + strconv.FormatInt(int64(v), 10))

				newAddr := k & 0x0f
				console.Output[newAddr+1] = v

				delete(console.Input, k)
			}

		}
	}

	fmt.Println("Final: " + strconv.FormatInt(int64(cpu.Tstates), 10))
	fmt.Println("Dump " + fmt.Sprint(cpu))
}
