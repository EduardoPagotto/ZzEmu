package ZzEmu

import "log"

type Memory interface {
	Read(address uint16) byte
	Write(address uint16, value byte)

	contendMem(address, states, time uint16) uint16
	contendIO(address, states, time uint16) uint16
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

func (mem *cpuMemory) contendMem(address, states, time uint16) uint16 {
	return time
}

func (mem *cpuMemory) contendIO(address, states, time uint16) uint16 {
	return time
}

// func contendMemory(z80 *z80.Z80, address uint16, time int) {
// 	tstates_p := &z80.Tstates
// 	tstates := *tstates_p

// 	if (address & 0xc000) == 0x4000 {
// 		tstates += int(delay_table[tstates])
// 	}

// 	tstates += time

// 	*tstates_p = tstates
// }
