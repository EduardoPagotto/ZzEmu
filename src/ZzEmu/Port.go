package ZzEmu

type PortInterface interface {
	ReadPort(address uint16) byte
	WritePort(address uint16, b byte)
}

type CpuPort struct {
	io *map[uint16]byte
}

func (port *CpuPort) ReadPort(address uint16) byte {
	var pMapIO *map[uint16]byte = (*port).io
	val := (*pMapIO)[address]
	return val

}

func (port *CpuPort) WritePort(address uint16, b byte) {
	var pMapIO *map[uint16]byte = (*port).io
	(*pMapIO)[address] = b
}
