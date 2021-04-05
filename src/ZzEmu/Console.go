package ZzEmu

const TotROM = 0x8000
const StartRAM = 0x8000
const SizeRAM = 0x8000
const TotRAM = StartRAM + SizeRAM - 1

type Console struct {
	CPU *CPU
	ROM [TotROM]byte
	RAM [TotRAM]byte
}

func NewCPU(console *Console) *CPU {
	cpu := CPU{Memory: NewCPUMemory(console)}
	cpu.createTable()
	cpu.Reset()
	return &cpu
}

func NewCPUMemory(console *Console) Memory {
	return &cpuMemory{console}
}
