package ZzEmu

const TotROM = 0x100 //0x0100
const StartRAM = TotROM
const SizeRAM = 0x100 //0x0100

const TopAddr = StartRAM + SizeRAM // usado pelo stackpointer
const TotRAM = TopAddr - StartRAM

const TotIO = 0x16

type Console struct {
	CPU *Z80
	ROM [TotROM]byte
	RAM [TotRAM]byte
	IO  [TotIO]byte
}

func NewCPUMemory(console *Console) MemoryInterface {
	return &CpuMemory{rom: &console.ROM, ram: &console.RAM}
}

func NewCPUPort(console *Console) PortInterface {
	return &CpuPort{io: &console.IO}
}

func NewConsole() *Console {

	console := Console{}
	var mem MemoryInterface = NewCPUMemory(&console) //&CpuMemory{rom: &console.ROM, ram: &console.RAM}
	var port PortInterface = NewCPUPort(&console)    //&CpuPort{io: &console.IO}
	console.CPU = NewZ80(mem, port)

	return &console
}
