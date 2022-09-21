package main

import (
	"machine"
	"time"

	"github.com/alexjomin/victron/vedirect"
	"tinygo.org/x/drivers/ssd1306"
)

var (
	currentPage             = 0
	display                 ssd1306.Device
	state                   vedirect.State
	lastClic                = time.Now()
	minimunDelayBetweenClic = time.Millisecond * 500
	timeout                 = time.Second * 30
)

const (
	baudRate      = 19200
	numberOfpages = 4
)

func bg() {
	for {
		button := machine.GP13
		button.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
		now := time.Now()
		delta := now.Sub(lastClic)
		if button.Get() && delta > minimunDelayBetweenClic {
			incPage()
			displayPage()
			lastClic = now
		}
		if delta > timeout && currentPage != 4 {
			lastClic = now
			currentPage = 4
			displayPage()
		}
		time.Sleep(time.Millisecond * 200)
	}
}

func main() {

	go bg()

	uart, err := initUART()
	if err != nil {
		println(err)
	}

	parser, err := vedirect.NewParser()
	if err != nil {
		println(err)
	}

	state, err = vedirect.NewState()
	if err != nil {
		println(err)
	}

	display, err = initDisplay()
	if err != nil {
		println(err)
	}

	welcomePage(&display)

	for {
		if uart.Buffered() > 0 {
			data, err := uart.ReadByte()
			if err != nil {
				println(err)
				continue
			}

			parser, err = parser.ParseByte(data)
			if err != nil {
				println(err)
				continue
			}

			if parser.Ready {
				data, _ := parser.GetKV()
				if data == nil {
					continue
				}
				parser.Ready = false
				f, err := vedirect.NewFrame(data)
				if err != nil {
					println(err)
					continue
				}
				state = state.Update(f)
			}
		}
		time.Sleep(time.Microsecond * 100)

	}
}
