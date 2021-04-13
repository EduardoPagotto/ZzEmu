package ZzEmu

const TotROM = 0x100 //0x0100
const StartRAM = TotROM
const SizeRAM = 0x100 //0x0100

const TopAddr = StartRAM + SizeRAM // usado pelo stackpointer
const TotRAM = TopAddr - StartRAM

type Console struct {
	CPU *Z80
	ROM [TotROM]byte
	RAM [TotRAM]byte
	IO  map[uint16]byte
}

func NewCPUMemory(console *Console) MemoryInterface {
	return &CpuMemory{rom: &console.ROM, ram: &console.RAM}
}

func NewCPUPort(console *Console) PortInterface {
	return &CpuPort{io: &console.IO}
}

func NewConsole() *Console {

	console := Console{}
	var mem MemoryInterface = NewCPUMemory(&console)
	var port PortInterface = NewCPUPort(&console)
	console.CPU = NewZ80(mem, port)

	return &console
}
