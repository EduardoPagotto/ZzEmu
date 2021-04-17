package ZzEmu

type Console struct {
	CPU *Z80
	Bus *Bus
}

func NewConsole() Console {
	console := Console{Bus: NewBuz()}
	console.CPU = NewZ80(console.Bus)
	return console
}

func (console *Console) Exec() {

	for {
		console.CPU.DoOpcode()

		if (console.CPU.Tstates > 50) && (console.CPU.Tstates < 500) {
			console.CPU.Interrupt()
		}

		// for {
		// 	// 1667 de ts para 1/60 de segundos
		// 	address, value, ok := console.Input.ReadAll()
		// 	if ok {
		// 		fmt.Println("TState: " + strconv.FormatInt(int64(console.CPU.Tstates), 10) +
		// 			" Addr: " + strconv.FormatInt(int64(address), 10) +
		// 			" Val:" + strconv.FormatInt(int64(value), 10))

		// 		//newAddr := address & 0x0f
		// 		console.Output.Write(address+1, value)
		// 	} else {
		// 		break
		// 	}
		// }
	}
}
