package ZzEmu

type DeviceInterface interface {
	Read(address uint16) (byte, bool)
	Write(address uint16, value byte) bool
	Valid(address uint16) bool
}

//--

type Bus struct {
	memory map[string]DeviceInterface
	io     map[string]DeviceInterface
}

func (b *Bus) GetMemory(name string) DeviceInterface {
	return b.memory[name]
}

func (b *Bus) GetIO(name string) DeviceInterface {
	return b.io[name]
}

func (b *Bus) AddMemory(name string, device DeviceInterface) {
	b.memory[name] = device
}

func (b *Bus) AddIO(name string, device DeviceInterface) {
	b.io[name] = device
}

func (b *Bus) ReadMemory(address uint16) byte {
	for _, dev := range b.memory {
		value, ok := dev.Read(address)
		if ok {
			return value
		}
	}
	return 0xff
}

func (b *Bus) WriteMemory(address uint16, value byte) {
	for _, dev := range b.memory {
		ok := dev.Write(address, value)
		if ok {
			return
		}
	}
}

func (b *Bus) ReadIO(address uint16) byte {
	for _, dev := range b.io {
		value, ok := dev.Read(address)
		if ok {
			return value
		}
	}
	return 0xff
}

func (b *Bus) WriteIO(address uint16, value byte) {
	for _, dev := range b.io {
		ok := dev.Write(address, value)
		if ok {
			return
		}
	}
}

func NewBuz() *Bus {
	bus := new(Bus)
	bus.memory = make(map[string]DeviceInterface)
	bus.io = make(map[string]DeviceInterface)

	return bus
}
