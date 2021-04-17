package ZzEmu

import (
	ZzEmu "ZzEmu/CPU"
	"bufio"
	"os"
)

// const TotROM = 8192 //0x100 //0x0100
// const StartRAM = TotROM
// const SizeRAM = 2048 //0x020 //0x0100

// const TopAddr = StartRAM + SizeRAM // usado pelo stackpointer
// const TotRAM = TopAddr - StartRAM

type Console struct {
	CPU *ZzEmu.Z80
	Bus *ZzEmu.Bus
	// ROM    [TotROM]byte
	// RAM    [TotRAM]byte
	// Input  *BufferIO
	// Output *BufferIO
}

func NewConsole() *Console {

	console := new(Console)
	console.Bus = ZzEmu.NewBuz()
	console.CPU = ZzEmu.NewZ80(console.Bus)

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

func (console *Console) LoadRom(filename string) (int, error) {

	buffer, erro := readBinary(filename)
	if erro != nil {
		return -1, erro
	}

	//copy(console.ROM[:], buffer)
	tot := len(buffer)
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
