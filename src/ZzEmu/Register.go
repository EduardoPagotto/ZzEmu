package ZzEmu

// Registrador
// type Register interface {
// 	get() byte
// 	set(value byte)
// 	setR(value Register)
// 	inc()
// 	dec()
// 	and(value byte)
// 	or(value byte)
// 	add(value byte)
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
