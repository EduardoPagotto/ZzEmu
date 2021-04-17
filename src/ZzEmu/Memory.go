package ZzEmu

import (
	"bufio"
	"os"
)

type DeviceMemory struct {
	mem   []byte
	start uint16
	top   uint16
}

func (m *DeviceMemory) Read(address uint16) (byte, bool) {

	if (address >= m.start) && (address < m.top) {
		addrFinal := address - m.start
		return m.mem[addrFinal], true
	}

	return 0xff, false
}

func (m *DeviceMemory) Write(address uint16, value byte) bool {

	if (address >= m.start) && (address < m.top) {
		addrFinal := address - m.start
		m.mem[addrFinal] = value
		return true
	}
	return false

}

func (m *DeviceMemory) Valid(address uint16) bool {

	if (address >= m.start) && (address < m.top) {
		return true
	}

	return false
}

func NewDeviceMemory(start, size uint16) *DeviceMemory {

	device := new(DeviceMemory)
	device.start = start
	device.top = start + size
	device.mem = make([]byte, size)
	return device

}

func (m *DeviceMemory) LoadRom(filename string) (int, error) {

	buffer, erro := readBinary(filename)
	if erro != nil {
		return -1, erro
	}

	tot := copy(m.mem, buffer)
	//tot := len(buffer)
	return tot, nil
}

func readBinary(filename string) ([]byte, error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	stats, statsErr := file.Stat()
	if statsErr != nil {
		return nil, statsErr
	}

	var size int64 = stats.Size()
	bytes := make([]byte, size)

	bufr := bufio.NewReader(file)
	_, err = bufr.Read(bytes)

	return bytes, err
}
