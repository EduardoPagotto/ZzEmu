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

	var console *ZzEmu.Console = new(ZzEmu.Console)
	var cpu *ZzEmu.Z80 = ZzEmu.CreateCPU(console)

	copy(console.ROM[:], buffer)

	for {
		cpu.DoOpcode()
		if cpu.Halted {
			break
		}
	}

	fmt.Println("Final: " + strconv.FormatInt(int64(cpu.Tstates), 10))
	fmt.Println("Dump " + fmt.Sprint(cpu))
}
