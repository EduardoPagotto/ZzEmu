package ZzEmu

const TotROM = 0x100 //0x0100
const StartRAM = TotROM
const SizeRAM = 0x100 //0x0100

const TopAddr = StartRAM + SizeRAM // usado pelo stackpointer
const TotRAM = TopAddr - StartRAM

type Console struct {
	CPU    *Z80
	ROM    [TotROM]byte
	RAM    [TotRAM]byte
	Input  map[uint16]byte
	Output map[uint16]byte
}

func NewCPUMemory(console *Console) MemoryInterface {
	return &CpuMemory{rom: &console.ROM, ram: &console.RAM}
}

func NewCPUPort(console *Console) PortInterface {
	return &CpuPort{Input: &console.Input, Output: &console.Output}
}

func NewConsole() *Console {

	console := Console{}
	console.Input = make(map[uint16]byte, 1)
	console.Output = make(map[uint16]byte, 1)

	var mem MemoryInterface = NewCPUMemory(&console)
	var port PortInterface = NewCPUPort(&console)
	console.CPU = NewZ80(mem, port)

	return &console
}
