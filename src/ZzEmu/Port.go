package ZzEmu

type PortInterface interface {
	ReadPort(address uint16) byte
	WritePort(address uint16, b byte)
}

type CpuPort struct {
	Input  *BufferIO
	Output *BufferIO
}

func (port *CpuPort) ReadPort(address uint16) byte {
	return port.Output.Read(address)
}

func (port *CpuPort) WritePort(address uint16, b byte) {
	port.Input.Write(address, b)
}
