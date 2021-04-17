package ZzEmu

import (
	"container/list"
)

type BufferIO struct {
	mapIO map[uint16]*list.List
	start uint16
	size  uint16
}

func NewBufferIO(start, size uint16) *BufferIO {

	dev := new(BufferIO)
	dev.start = start
	dev.size = size
	dev.mapIO = make(map[uint16]*list.List)
	return dev

}

func (buffer *BufferIO) Write(address uint16, value byte) bool {

	var addrFinal uint16 = address - buffer.start
	if (addrFinal > 0) && (addrFinal < buffer.size) {

		queue, ok := buffer.mapIO[address]
		if ok {
			queue.PushBack(value)
		} else {
			newQueue := list.New()
			newQueue.PushBack(value)
			buffer.mapIO[address] = newQueue
		}
	}

	return false

}

func (buffer *BufferIO) Read(address uint16) (byte, bool) {

	var addrFinal uint16 = address - buffer.start
	if (addrFinal > 0) && (addrFinal < buffer.size) {

		queue, ok := buffer.mapIO[address]
		if ok {
			if queue.Len() > 0 {
				iterator := queue.Front()
				interf := iterator.Value
				value, _ := interf.(byte)
				queue.Remove(iterator)
				return value, true
			}
		}
	}

	return 0xff, false

}

func (buffer *BufferIO) Len() int {
	var total int = 0
	for _, queue := range buffer.mapIO {
		total += queue.Len()
	}
	return total
}

func (buffer *BufferIO) LenAddress(address uint16) (int, bool) {

	var addrFinal uint16 = address - buffer.start
	if (addrFinal > 0) && (addrFinal < buffer.size) {
		var total int = 0
		for addq, queue := range buffer.mapIO {
			if addq == address {
				total += queue.Len()
			}
		}

		return total, true
	}

	return 0, false
}

func (buffer *BufferIO) ReadAll() (uint16, byte, bool) {

	for address, queue := range buffer.mapIO {
		if queue.Len() > 0 {
			iterator := queue.Front()
			interf := iterator.Value
			value, _ := interf.(byte)
			queue.Remove(iterator)
			return address, value, true
		}
	}

	return 0, 0xff, false

}
