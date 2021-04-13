package ZzEmu

type PortInterface interface {
	ReadPort(address uint16) byte
	WritePort(address uint16, b byte)
	//ReadPortInternal(address uint16, contend bool) byte
	//WritePortInternal(address uint16, b byte, contend bool)
	//ContendPortPreio(address uint16)
	//ContendPortPostio(address uint16)
}

type CpuPort struct {
	io *[TotIO]byte
	//console *Console

}

// func (p *CpuPort) ReadPortInternal(address uint16, contend bool) byte {
// 	if contend {
// 		p.ContendPortPreio(address)
// 	}

// 	var r byte = byte(address >> 8)
// 	//events = append(events, fmt.Sprintf("%5d PR %04x %02x\n", p.console.CPU.Tstates, address, r))

// 	if contend {
// 		p.ContendPortPostio(address)
// 	}
// 	return r
// }

func (port *CpuPort) ReadPort(address uint16) byte {
	return port.io[address] //p.ReadPortInternal(port, true)
}

// func (p *CpuPort) WritePortInternal(address uint16, b byte, contend bool) {
// 	if contend {
// 		p.ContendPortPreio(address)
// 	}
// 	//events = append(events, fmt.Sprintf("%5d PW %04x %02x\n", p.console.CPU.Tstates, address, b))

// 	if contend {
// 		p.ContendPortPostio(address)
// 	}
// }

func (port *CpuPort) WritePort(address uint16, b byte) {
	port.io[address] = b
	//p.WritePortInternal(port, b, true)
}

// func (p *CpuPort) ContendPortPreio(port uint16) {
// 	if (port & 0xc000) == 0x4000 {
// 		//events = append(events, fmt.Sprintf("%5d PC %04x\n", p.console.CPU.Tstates, port))
// 	}
// 	//p.console.CPU.Tstates += 1
// }

// func (p *CpuPort) ContendPortPostio(port uint16) {
// 	if (port & 0x0001) == 1 {
// 		if (port & 0xc000) == 0x4000 {
// 			for i := 0; i < 3; i++ {
// 				//events = append(events, fmt.Sprintf("%5d PC %04x\n", p.console.CPU.Tstates, port))
// 				//contendPort(p.console.CPU, 1)
// 			}
// 		} else {
// 			//p.console.CPU.Tstates += 3
// 		}

// 	} else {
// 		//events = append(events, fmt.Sprintf("%5d PC %04x\n", p.console.CPU.Tstates, port))
// 		//contendPort(p.console.CPU, 3)
// 	}
// }
