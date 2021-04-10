package ZzEmu

const TotROM = 0x1000
const StartRAM = 0x2000
const SizeRAM = 0x1000
const TotRAM = StartRAM + SizeRAM - 1
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

func NewCPUMemory(console *Console) Memory {
	return &cpuMemory{console}
}

func NewCPUPort(console *Console) PortAccessor {
	return &cpuPort{console}
}

// FIXME: lugar esquisito
func contendPort(z80 *Z80, time uint) {
	// tstates_p := &z80.Tstates
	// *tstates_p += time
	z80.Tstates += uint16(time)
}
