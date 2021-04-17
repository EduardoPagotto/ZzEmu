package ZzEmu

// type CpuMemory struct {
// 	rom *[TotROM]byte
// 	ram *[TotRAM]byte
// }

// func (mem *CpuMemory) Read(address uint16) byte {

// 	if address < TotROM {
// 		return mem.rom[address]
// 	} else if (address >= StartRAM) && (address < TopAddr) {
// 		addrFinal := address % StartRAM
// 		return mem.ram[addrFinal]
// 	}
// 	// else {
// 	// 	log.Fatal("Memoria fora de range", address)
// 	// }

// 	return 0xff
// }

// func (mem *CpuMemory) Write(address uint16, value byte) {
// 	if (address >= StartRAM) && (address < TopAddr) {
// 		addrFinal := address % StartRAM
// 		mem.ram[addrFinal] = value
// 	} else {
// 		//log.Fatal("Escrita em memoria invalida")
// 	}
// }

type DeviceMemory struct {
	mem   []byte
	start uint16
	size  uint16
}

func (m *DeviceMemory) Read(address uint16) (byte, bool) {

	var addrFinal uint16 = address - m.start
	if (addrFinal > 0) && (addrFinal < m.size) {
		return m.mem[addrFinal], true
	}

	return 0xff, false
}

func (m *DeviceMemory) Write(address uint16, value byte) bool {

	var addrFinal uint16 = address - m.start
	if (addrFinal > 0) && (addrFinal < m.size) {
		m.mem[addrFinal] = value
		return true
	}
	return false

}

func (m *DeviceMemory) Valid(address uint16) bool {

	val := address - m.start
	if (val > 0) && (val < m.size) {
		return true
	}

	return false
}

func NewDeviceMemory(start, size uint16) *DeviceMemory {

	device := new(DeviceMemory)
	device.start = start
	device.size = size
	device.mem = make([]byte, 0, size)
	return device

}
