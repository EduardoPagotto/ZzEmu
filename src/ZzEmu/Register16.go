package ZzEmu

// type Register16Interface interface {
// 	Hi() byte
// 	Lo() byte
// 	Get() uint16
// 	SetHi(h byte)
// 	SetLo(l byte)
// 	Set(value uint16)
// 	Inc()
// 	Dec()
// 	Add(value byte)
// }

type Register16 struct {
	high *byte
	low  *byte
}

func (r16 *Register16) Hi() byte {
	return *r16.high
}

func (r16 *Register16) Lo() byte {
	return *r16.low
}

func (r16 Register16) Get() uint16 {
	return joinBytes(*r16.high, *r16.low)
}

func (r16 Register16) Set(value uint16) {
	*r16.high, *r16.low = splitWord(value)
}

func (r16 *Register16) SetHi(h byte) {
	*r16.high = h
}

func (r16 *Register16) SetLo(l byte) {
	*r16.low = l
}

func (r16 Register16) Inc() {
	*r16.high, *r16.low = splitWord(r16.Get() + 1)
}

func (r16 *Register16) Dec() {
	*r16.high, *r16.low = splitWord(r16.Get() - 1)
}

func (r16 *Register16) Add(value byte) {
	*r16.high, *r16.low = splitWord(r16.Get() + uint16(value))
}

// Auxiliares

func splitWord(word uint16) (byte, byte) {
	return byte(word >> 8), byte(word & 0xff)
}

func joinBytes(h, l byte) uint16 {
	return uint16(l) | (uint16(h) << 8)
}
