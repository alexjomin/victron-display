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
	minimunDelayBetweenClic = time.Millisecond * 300
	timeout                 = time.Second * 30
	water000                = machine.GP10
	water033                = machine.GP11
	water066                = machine.GP12
	water100                = machine.GP13

	button = machine.GP15
)

const (
	baudRate      = 19200
	numberOfpages = 5
)

func initButton() {
	button.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})

	refreshScreen := func(p machine.Pin) {
		delta := time.Now().Sub(lastClic)
		if delta > minimunDelayBetweenClic {
			incPage()
			displayPage()
			lastClic = time.Now()
		}
	}

	err := button.SetInterrupt(machine.PinRising, refreshScreen)
	if err != nil {
		println(err)
	}
}

func clearDisplayAfterTimeout() {
	for {
		time.Sleep(timeout)
		if time.Now().Sub(lastClic) >= timeout && currentPage != pageStandBy {
			currentPage = pageStandBy
			displayPage()
		}
	}
}

func main() {

	water000.Configure(machine.PinConfig{
		Mode: machine.PinInput,
	})

	go clearDisplayAfterTimeout()
	initButton()

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
