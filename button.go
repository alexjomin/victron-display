package main

import "machine"

func initButton() error {
	button.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
	return nil
}
