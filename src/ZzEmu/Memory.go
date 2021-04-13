package ZzEmu

import "log"

type MemoryInterface interface {
	Read(address uint16) byte
	Write(address uint16, value byte)
}

type CpuMemory struct {
	ram *[TotROM]byte
	rom *[TotRAM]byte
}

func (mem *CpuMemory) Read(address uint16) byte {

	if address < TotROM {
		return mem.rom[address]
	} else if (address >= StartRAM) && (address < TopAddr) {
		addrFinal := address % StartRAM
		return mem.ram[addrFinal]
	} else {
		log.Fatal("Memoria fora de range", address)
	}

	return 0
}

func (mem *CpuMemory) Write(address uint16, value byte) {
	if (address >= StartRAM) && (address < TopAddr) {
		addrFinal := address % StartRAM
		mem.ram[addrFinal] = value
		return
	}

	log.Fatal("Escrita em memoria invalida")
}
