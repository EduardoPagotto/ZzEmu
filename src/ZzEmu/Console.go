package ZzEmu

const TotROM = 0x100 //0x0100
const StartRAM = TotROM
const SizeRAM = 0x100 //0x0100

const TopAddr = StartRAM + SizeRAM // usado pelo stackpointer
const TotRAM = TopAddr - StartRAM

const TotIO = 0x16

var (
	events []string
)

type Console struct {
	CPU *Z80
	ROM [TotROM]byte
	RAM [TotRAM]byte
	IO  [TotIO]byte
}

func CreateCPU(console *Console) *Z80 {
	cpu := NewZ80(NewCPUMemory(console), NewCPUPort(console))
	return cpu
}

func NewCPUMemory(console *Console) MemoryInterface {
	return &CpuMemory{console}
}

func NewCPUPort(console *Console) PortInterface {
	return &CpuPort{console}
}

// FIXME: lugar esquisito
func contendPort(z80 *Z80, time uint) {
	// tstates_p := &z80.Tstates
	// *tstates_p += time
	z80.Tstates += uint16(time)
}
