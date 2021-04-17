package ZzEmu

import (
	"container/list"
)

type BufferIO struct {
	mapIO map[uint16]*list.List
	start uint16
	top   uint16
}

func NewBufferIO(start, size uint16) *BufferIO {

	dev := new(BufferIO)
	dev.start = start
	dev.top = start + size
	dev.mapIO = make(map[uint16]*list.List)
	return dev

}

func (b *BufferIO) Write(address uint16, value byte) bool {

	if (address >= b.start) && (address < b.top) {
		addrFinal := address - b.start

		queue, ok := b.mapIO[addrFinal]
		if ok {
			queue.PushBack(value)
		} else {
			newQueue := list.New()
			newQueue.PushBack(value)
			b.mapIO[address] = newQueue
		}
	}

	return false

}

func (b *BufferIO) Read(address uint16) (byte, bool) {

	if (address >= b.start) && (address < b.top) {
		addrFinal := address - b.start

		queue, ok := b.mapIO[addrFinal]
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

func (b *BufferIO) Valid(address uint16) bool {
	if (address >= b.start) && (address < b.top) {
		return true
	}

	return false
}

func (b *BufferIO) Len() int {
	var total int = 0
	for _, queue := range b.mapIO {
		total += queue.Len()
	}
	return total
}

func (b *BufferIO) LenAddress(address uint16) (int, bool) {

	if (address >= b.start) && (address < b.top) {
		addrFinal := address - b.start
		var total int = 0
		for addq, queue := range b.mapIO {
			if addq == addrFinal {
				total += queue.Len()
			}
		}

		return total, true
	}

	return 0, false
}

func (b *BufferIO) ReadAll() (uint16, byte, bool) {

	for address, queue := range b.mapIO {
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
