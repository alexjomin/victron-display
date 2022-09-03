package main

import "machine"

func initUART() (*machine.UART, error) {
	uart := machine.UART0
	c := machine.UARTConfig{
		BaudRate: baudRate,
		TX:       machine.UART0_TX_PIN,
		RX:       machine.UART0_RX_PIN,
	}
	err := uart.Configure(c)
	if err != nil {
		return nil, err
	}

	return uart, nil
}
