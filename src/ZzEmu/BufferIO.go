package ZzEmu

import (
	"container/list"
)

type BufferIO struct {
	mapIO map[uint16]*list.List
}

func NewBufferIO() *BufferIO {
	return &BufferIO{mapIO: make(map[uint16]*list.List)}
}

func (buffer *BufferIO) Write(address uint16, value byte) {
	queue, ok := buffer.mapIO[address]
	if ok {
		queue.PushBack(value)
	} else {
		newQueue := list.New()
		newQueue.PushBack(value)
		buffer.mapIO[address] = newQueue
	}
}

func (buffer *BufferIO) Read(address uint16) byte {

	queue, ok := buffer.mapIO[address]
	if ok {
		if queue.Len() > 0 {
			iterator := queue.Front()
			interf := iterator.Value
			value, _ := interf.(byte)
			queue.Remove(iterator)
			return value
		}
	}

	return 0x00
}

func (buffer *BufferIO) Len() int {
	var total int = 0
	for _, queue := range buffer.mapIO {
		total += queue.Len()
	}
	return total
}

func (buffer *BufferIO) LenAddress(address uint16) int {
	var total int = 0
	for addq, queue := range buffer.mapIO {
		if addq == address {
			total += queue.Len()
		}
	}
	return total
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
	return 0, 0, false
}
