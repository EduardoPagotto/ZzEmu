package ZzEmu

type CpuPort struct {
	Input  *BufferIO
	Output *BufferIO
}

func (port *CpuPort) Read(address uint16) byte {
	return port.Output.Read(address)
}

func (port *CpuPort) Write(address uint16, b byte) {
	port.Input.Write(address, b)
}
