package main

import (
	"machine"
	"runtime"

	"github.com/alexjomin/victron/vedirect"
)

const (
	baudRate = 19200
)

func main() {
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	uart := machine.UART0
	c := machine.UARTConfig{
		BaudRate: baudRate,
		TX:       machine.UART0_TX_PIN,
		RX:       machine.UART0_RX_PIN,
	}
	err := uart.Configure(c)
	if err != nil {
		led.High()
	}

	parser, err := vedirect.NewParser()
	if err != nil {
		println(err)
	}

	state, err := vedirect.NewState()
	if err != nil {
		println(err)
	}

	ms := runtime.MemStats{}

	for {
		if uart.Buffered() > 0 {
			led.High()

			data, err := uart.ReadByte()
			if err != nil {
				println(err)
				continue
			}
			// print(string(data))
			err = parser.ParseByte(data)
			if err != nil {
				println(err)
				continue
			}

			if parser.Ready {
				data, _ := parser.GetFrame()
				if data == nil {
					continue
				}
				f, err := vedirect.NewFrame(*data)
				if err != nil {
					println(err)
					continue
				}
				state.Update(*f)
				println(state.BatteryVoltage, state.OperationState)

				runtime.ReadMemStats(&ms)
				println("Heap before GC. Used: ", ms.HeapInuse, " Free: ", ms.HeapIdle, " Meta: ", ms.GCSys)

				// Force Garbage Collection
				runtime.GC()

			}

			led.Low()

		}
	}
}
