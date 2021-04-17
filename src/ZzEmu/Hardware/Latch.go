package ZzEmu

type DeviceLatch struct {
	address uint16
	value   byte
}

func (d *DeviceLatch) Read(address uint16) (byte, bool) {
	if d.address == address {
		return d.value, true
	}
	return 0xff, false
}

func (d *DeviceLatch) Write(address uint16, value byte) bool {
	if d.address == address {
		d.value = value
	}
	return false
}

func (d *DeviceLatch) Valid(address uint16) bool {
	return d.address == address
}

func NewDeviceLatch(address uint16) *DeviceLatch {
	dev := new(DeviceLatch)
	dev.address = address
	return dev
}
