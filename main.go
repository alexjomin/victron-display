package main

import (
	"fmt"
	"machine"
	"time"

	"github.com/alexjomin/victron/tools"
	"github.com/alexjomin/victron/vedirect"
	"tinygo.org/x/drivers/ssd1306"
)

var (
	currentPage = 0
	display     ssd1306.Device
	state       vedirect.State
	lastClic    = time.Now()
)

const (
	baudRate                = 19200
	numberOfpages           = 4
	minimunDelayBetweenClic = time.Millisecond * 300
	timeout                 = time.Second * 30
)

func button() {
	button := machine.GP13
	button.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})

	callback := func(p machine.Pin) {
		delta := time.Now().Sub(lastClic)
		if delta > minimunDelayBetweenClic {
			incPage()
			displayPage()
			lastClic = time.Now()
		}
	}

	err := button.SetInterrupt(machine.PinRising, callback)
	if err != nil {
		println(err)
	}

}

func clearDisplayAfterTimeout() {
	for {
		time.Sleep(timeout)
		if time.Now().Sub(lastClic) >= timeout && currentPage != 4 {
			currentPage = 4
			displayPage()
		}
	}
}

func main() {

	// go clearDisplayAfterTimeout()
	// button()

	// uart, err := initUART()
	// if err != nil {
	// 	println(err)
	// }

	// parser, err := vedirect.NewParser()
	// if err != nil {
	// 	println(err)
	// }

	// state, err = vedirect.NewState()
	// if err != nil {
	// 	println(err)
	// }

	// display, err := initDisplay()
	// if err != nil {
	// 	println(err)
	// }

	// welcomePage(&display)

	for {
		machine.InitADC()
		// led := machine.LED
		// led.Configure(machine.PinConfig{Mode: machine.PinOutput})

		sensor := machine.ADC{Pin: machine.ADC2}
		sensor.Configure(machine.ADCConfig{})

		// cf := 3.3 / 65536.0
		// v := float64(sensor.Get()) * cf
		// value := fmt.Sprintf("%.2f", v)
		adc := sensor.Get()
		v := tools.WaterProbe(adc)
		fmt.Println(adc, v)
		time.Sleep(time.Millisecond * 400)
		// 1700 - 52000

	}

	// for {
	// 	if uart.Buffered() > 0 {
	// 		data, err := uart.ReadByte()
	// 		if err != nil {
	// 			println(err)
	// 			continue
	// 		}

	// 		parser, err = parser.ParseByte(data)
	// 		if err != nil {
	// 			println(err)
	// 			continue
	// 		}

	// 		if parser.Ready {
	// 			data, _ := parser.GetKV()
	// 			if data == nil {
	// 				continue
	// 			}
	// 			parser.Ready = false
	// 			f, err := vedirect.NewFrame(data)
	// 			if err != nil {
	// 				println(err)
	// 				continue
	// 			}
	// 			state = state.Update(f)
	// 		}
	// 	}
	// 	time.Sleep(time.Microsecond * 100)
	// }
}
