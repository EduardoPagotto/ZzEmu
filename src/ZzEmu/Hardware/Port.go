package ZzEmu

type PortInterface interface {
	Read(address uint16) (byte, bool)
	Write(address uint16, value byte)
}

type CpuPort struct {
	Input  *BufferIO
	Output *BufferIO
}

func (port *CpuPort) Read(address uint16) (byte, bool) {
	return port.Output.Read(address)
}

func (port *CpuPort) Write(address uint16, b byte) {
	port.Input.Write(address, b)
}
