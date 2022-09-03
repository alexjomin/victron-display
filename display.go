package main

import (
	"fmt"
	"image/color"
	"machine"

	"github.com/alexjomin/victron/vedirect"
	"tinygo.org/x/drivers/ssd1306"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freesans"
	"tinygo.org/x/tinyfont/proggy"
)

var (
	white = color.RGBA{255, 255, 255, 255}
)

func incPage() {
	if currentPage > numberOfpages-1 {
		currentPage = 0
	} else {
		currentPage += 1
	}
}

func initDisplay() (display ssd1306.Device, err error) {
	err = machine.I2C0.Configure(machine.I2CConfig{
		Frequency: machine.TWI_FREQ_400KHZ,
	})

	if err != nil {
		return
	}

	display = ssd1306.NewI2C(machine.I2C0)

	display.Configure(ssd1306.Config{
		Address: 0x3C,
		Width:   128,
		Height:  64,
	})

	display.ClearDisplay()

	return display, err

}

func formatVoltage(v int) string {
	return fmt.Sprintf("%.2f V", float64(v)/1000.0)
}

func welcomePage(display *ssd1306.Device) {
	display.SetBuffer(loading)
	tinyfont.WriteLine(display, &freesans.Regular9pt7b, 32, 16, "Loading", color.RGBA{255, 255, 255, 255})
	tinyfont.WriteLine(display, &proggy.TinySZ8pt7b, 10, 55, "Captain Gantu v1.0", white)
	display.Display()
}

func displayPage() {

	display.ClearBuffer()

	if currentPage == 0 {
		display.Command(ssd1306.DISPLAYON)
		if state.OperationState == vedirect.StateFloat {
			display.SetBuffer(charged)
			tinyfont.WriteLine(&display, &freesans.Regular9pt7b, 50, 36, state.OperationState, white)
		} else {
			display.SetBuffer(inCharge)
			tinyfont.WriteLine(&display, &freesans.Regular9pt7b, 42, 28, state.OperationState, white)
			bc := fmt.Sprintf("%d A", state.BatteryCurrent)
			tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 43, 42, bc, white)
		}

	} else if currentPage == 1 {
		display.SetBuffer(sun)
		pv := formatVoltage(state.PanelVoltage)
		pw := fmt.Sprintf("%d W", state.PanelPower)
		tinyfont.WriteLine(&display, &freesans.Regular9pt7b, 50, 25, pv, white)
		tinyfont.WriteLine(&display, &freesans.Regular9pt7b, 50, 50, pw, white)

	} else if currentPage == 2 {
		display.SetBuffer(battery)
		bv := formatVoltage(state.BatteryVoltage)
		bvmin := "Min: " + formatVoltage(state.MinBatteryVoltage)
		bvmax := "Max: " + formatVoltage(state.MaxBatteryVoltage)

		tinyfont.WriteLine(&display, &freesans.Regular9pt7b, 50, 25, bv, white)
		tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 50, 40, bvmin, white)
		tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 50, 50, bvmax, white)
	} else if currentPage == 3 {

		display.SetBuffer(plug)
		lc := fmt.Sprintf("%d A - %d W", state.LoadCurrent, state.InstantaneousPower)
		yp := fmt.Sprintf("Today: %d W", state.YieldToday)
		mp := fmt.Sprintf("Max: %d W", state.MaxPower)
		tinyfont.WriteLine(&display, &freesans.Regular9pt7b, 50, 25, lc, white)
		tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 50, 40, yp, white)
		tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 50, 50, mp, white)

	} else if currentPage == 4 {
		display.Command(ssd1306.DISPLAYOFF)
	}

	display.Display()

}
