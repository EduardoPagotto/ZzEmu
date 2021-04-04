package ZzEmu

import (
	"log"

	"ZzEmu"
)

type Console struct {
	CPU *ZzEmu.CPU
	//Mapper Mapper
	RAM []byte
}

type cpuMemory struct {
	console *Console
}

type Memory interface {
	Read(address uint16) byte
	Write(address uint16, value byte)
}

func NewCPU(console *Console) *ZzEmu.CPU {
	cpu := ZzEmu.CPU{Memory: NewCPUMemory(console)}
	cpu.createTable()
	cpu.Reset()
}

func NewCPUMemory(console *Console) Memory {
	return &cpuMemory{console}

func (mem *cpuMemory) Read(address uint16) byte {
	switch {
	case address < 0x2000:
		return mem.console.RAM[address%0x0800]
	// case address >= 0x2000:
	// 	return mem.console.Mapper.Read(address)
	default:
		log.Fatal("Memoria fora de range", address)
	}
	return 0
}

// type Mapper interface {
// 	Read(address uint16) byte
// 	Write(address uint16, value byte)
// }
