package ZzEmu

type PortInterface interface {
	ReadPort(address uint16) byte
	WritePort(address uint16, b byte)
}

type CpuPort struct {
	Input  *map[uint16]byte
	Output *map[uint16]byte
}

func (port *CpuPort) ReadPort(address uint16) byte {
	var pMapIO *map[uint16]byte = (*port).Output

	//endAddr := address & 0x0f
	val, found := (*pMapIO)[address]
	if found {
		delete(*pMapIO, address)
		return val
	}

	return 0x00
}

func (port *CpuPort) WritePort(address uint16, b byte) {
	var pMapIO *map[uint16]byte = (*port).Input
	(*pMapIO)[address] = b
}
