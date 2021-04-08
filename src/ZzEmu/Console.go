package ZzEmu

const TotROM = 0x1000
const StartRAM = 0x2000
const SizeRAM = 0x1000
const TotRAM = StartRAM + SizeRAM - 1

type Console struct {
	CPU *Z80
	ROM [TotROM]byte
	RAM [TotRAM]byte
}

func CreateCPU(console *Console) *Z80 {
	cpu := NewZ80(NewCPUMemory(console))
	return cpu
}

func NewCPUMemory(console *Console) Memory {
	return &cpuMemory{console}
}
