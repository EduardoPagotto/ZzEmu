package ZzEmu

// type RegisterInterface interface {
// 	Get() byte
// 	Set(value byte)
// 	Inc()
// 	Dec()
// 	And(value byte)
// 	Or(value byte)
// 	Add(value byte)
// }

type Register struct {
	single byte
}

func (r *Register) Get() byte {
	return r.single
}

func (r *Register) Set(value byte) {
	r.single = value
}

func (r *Register) Inc() {
	r.single++
}

func (r *Register) Dec() {
	r.single--
}

func (r *Register) And(value byte) {
	r.single &= value
}

func (r *Register) Or(value byte) {
	r.single |= value
}

func (r *Register) Add(value byte) {
	r.single += value
}
