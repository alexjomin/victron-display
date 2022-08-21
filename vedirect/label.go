package vedirect

import (
	"errors"
)

var (
	ErrLabelIsUnknown = errors.New("label is unknown")
)

// only MPTT is available for now

const (
	BatteryCurrent     = "I"
	Checksum           = "Checksum"
	DaySequenceNumber  = "HSDS"
	ErrorCode          = "ERR"
	FirmwareVersion    = "FW"
	InstantaneousPower = "P"
	LoadCurrent        = "IL"
	LoadOutputState    = "LOAD"
	MainBatteryVoltage = "V"
	MaxPowerToday      = "H21"
	MaxPowerYesterday  = "H23"
	PanelPower         = "PPV"
	PanelVoltage       = "VPV"
	ProductID          = "PID"
	RelayState         = "Relay"
	SerialNumber       = "SER#"
	StateOfOperation   = "CS"
	YieldToday         = "H20"
	YieldTotal         = "H19"
	YieldYesterday     = "H22"
)

var AvailableLabels = map[string]string{
	BatteryCurrent:     "battery_current",
	Checksum:           "checksum",
	DaySequenceNumber:  "day_sequence_number",
	ErrorCode:          "error_code",
	FirmwareVersion:    "firmware_version",
	InstantaneousPower: "instantaneous_power",
	LoadCurrent:        "load_current",
	LoadOutputState:    "load_output_state",
	MainBatteryVoltage: "main_battery_voltage",
	MaxPowerToday:      "max_power_today",
	MaxPowerYesterday:  "max_power_testerday",
	PanelPower:         "panel_power",
	PanelVoltage:       "panel_voltage",
	ProductID:          "product_id",
	RelayState:         "relay_state",
	SerialNumber:       "serial_number",
	StateOfOperation:   "state_operation",
	YieldToday:         "yield_today",
	YieldTotal:         "yield_total",
	YieldYesterday:     "yield_yesterday",
}

func IsKnownLabel(key string) error {
	for k := range AvailableLabels {
		if key == k {
			return nil
		}
	}
	return ErrLabelIsUnknown
}
