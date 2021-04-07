package ZzEmu

import "log"

type Memory interface {
	Read(address uint16) byte
	Write(address uint16, value byte)
}

type cpuMemory struct {
	console *Console
}

func (mem *cpuMemory) Read(address uint16) byte {

	if address < TotROM {
		return mem.console.ROM[address]
	} else if (address >= StartRAM) && (address < TotRAM) {
		return mem.console.RAM[address%StartRAM]
	} else {
		log.Fatal("Memoria fora de range", address)
	}

	return 0
}

func (mem *cpuMemory) Write(address uint16, value byte) {
	if (address >= StartRAM) && (address < TotRAM) {
		mem.console.RAM[address%StartRAM] = value
		return
	}

	log.Fatal("Escrita em memoria invalida")
}
