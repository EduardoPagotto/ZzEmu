package ZzEmu

// type RegisterPairInterface interface {
// 	Hi() byte
// 	Lo() byte
// 	Get() uint16
// 	SetHi(h byte)
// 	SetLo(l byte)
// 	Set(value uint16)
// 	Inc()
// 	Dec()
// 	Add(value byte)
// 	GetRH() *Register
// 	GetRL() *Register
// }

type RegisterPair struct {
	pair uint16
}

func (rp *RegisterPair) Hi() byte {
	return byte(rp.pair >> 8)
}

func (rp *RegisterPair) Lo() byte {
	return byte(rp.pair & 0x00FF)
}

func (rp *RegisterPair) Get() uint16 {
	return rp.pair
}

func (rp *RegisterPair) SetHi(h byte) {
	rp.pair = (rp.pair & 0x00ff) | uint16(h)<<8
}

func (rp *RegisterPair) SetLo(l byte) {
	rp.pair = (rp.pair & uint16(0xFF00)) | uint16(l)
}

func (rp *RegisterPair) Set(value uint16) {
	rp.pair = value
}

func (rp *RegisterPair) Inc() {
	rp.pair++
}

func (rp *RegisterPair) Dec() {
	rp.pair--
}

func (rp *RegisterPair) Add(value byte) {
	rp.pair += uint16(value)
}

func (rp *RegisterPair) GetRH() *Register {
	var reg *Register = new(Register)
	reg.single = rp.Hi()
	return reg
}

func (rp *RegisterPair) GetRL() *Register {
	var reg *Register = new(Register)
	reg.single = rp.Lo()
	return reg
}
