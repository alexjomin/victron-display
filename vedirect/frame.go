package vedirect

import (
	"fmt"
	"strconv"
)

type OperationState string

var operationStates = map[string]string{
	"0": "OFF",
	"1": "Low power",
	"2": "Fault",
	"3": "Bulk",
	"4": "Absorption",
	"5": "Float",
	"9": "Inverting",
}

var productList = map[string]string{
	"0x203":  "BMV-700",
	"0x204":  "BMV-702",
	"0x205":  "BMV-700H",
	"0xA04C": "BlueSolar MPPT 75/10",
	"0x300":  "BlueSolar MPPT 70/15",
	"0xA042": "BlueSolar MPPT 75/15",
	"0xA043": "BlueSolar MPPT 100/15",
	"0xA044": "BlueSolar MPPT 100/30 Rev1",
	"0xA04A": "BlueSolar MPPT 100/30 Rev2",
	"0xA041": "BlueSolar MPPT 150/35 Rev1",
	"0xA04B": "BlueSolar MPPT 150/35 Rev2",
	"0xA04D": "BlueSolar MPPT 150/45",
	"0xA040": "BlueSolar MPPT 75/50",
	"0xA045": "BlueSolar MPPT 100/50 Rev1",
	"0xA049": "BlueSolar MPPT 100/50 Rev2",
	"0xA04E": "BlueSolar MPPT 150/60",
	"0xA046": "BlueSolar MPPT 150/70",
	"0xA04F": "BlueSolar MPPT 150/85",
	"0xA047": "BlueSolar MPPT 150/100",
	"0xA051": "SmartSolar MPPT 150/100",
	"0xA060": "SmartSolar MPPT 100/20",
	"0xA050": "SmartSolar MPPT 250/100",
}

type Frame struct {
	BatteryCurrent     *int    `json:"battery_current,omitempty"`
	DaySequenceNumber  *int    `json:"day_sequence_number,omitempty"`
	ErrorCode          *string `json:"error_code,omitempty"`
	FirmwareVersion    *string `json:"firmware_version,omitempty"`
	InstantaneousPower *int    `json:"instantaneous_power,omitempty"`
	LoadCurrent        *int    `json:"load_current,omitempty"`
	LoadOutputState    *bool   `json:"load_output_state,omitempty"`
	MainBatteryVoltage *int    `json:"main_battery_voltage,omitempty"`
	MaxPowerToday      *int    `json:"max_power_today,omitempty"`
	MaxPowerYesterday  *int    `json:"max_power_yesterday,omitempty"`
	PanelPower         *int    `json:"panel_power,omitempty"`
	PanelVoltage       *int    `json:"panel_voltage,omitempty"`
	ProductName        *string `json:"product_id,omitempty"`
	RelayState         *bool   `json:"relay_state,omitempty"`
	SerialNumber       *string `json:"serial_number,omitempty"`
	StateOfOperation   *string `json:"state_of_operation,omitempty"`
	YieldToday         *int    `json:"yield_today,omitempty"`
	YieldTotal         *int    `json:"yield_total,omitempty"`
	YieldYesterday     *int    `json:"yield_yesterday,omitempty"`
}

func NewFrame(kv KeyValue) (*Frame, error) {
	f := &Frame{}
	for k, v := range kv {
		err := f.parseKV(k, v)
		if err != nil {
			return nil, err
		}
	}
	return f, nil
}

func toBoolPointer(b bool) *bool {
	return &b
}

func parseInt(v string) (*int, error) {
	ps, err := strconv.Atoi(v)
	if err != nil {
		return nil, err
	}
	return &ps, nil
}

func parseLoadOutputState(v string) (state *bool, err error) {

	switch v {
	case "ON":
		state = toBoolPointer(true)
	case "OFF":
		state = toBoolPointer(false)
	default:
		return nil, fmt.Errorf("can't parse state value `%s`", v)
	}

	return
}

func parseErrorCode(v string) (*string, error) {
	var errorMsg string

	switch v {
	case "0":
		return nil, nil
	case "2":
		errorMsg = "Battery voltage too high"
	case "17":
		errorMsg = "Charger temperature too high"
	case "18":
		errorMsg = "Charger over current"
	case "19":
		errorMsg = "Charger current reversed"
	case "20":
		errorMsg = "Bulk time limit exceeded"
	case "21":
		errorMsg = "Current sensor issue (sensor bias/sensor broken)"
	case "26":
		errorMsg = "Terminals overheated"
	case "33":
		errorMsg = "Input voltage too high (solar panel)"
	case "34":
		errorMsg = "Input current too high (solar panel)"
	case "38":
		errorMsg = "Input shutdown (due to excessive battery voltage)"
	case "116":
		errorMsg = "Factory calibration data lost"
	case "117":
		errorMsg = "Invalid/incompatible firmware"
	case "119":
		errorMsg = "User settings invalid"
	default:
		return nil, fmt.Errorf("error code `%s` is unknown", v)
	}

	return &errorMsg, nil
}

func parseDevice(v string) (*string, error) {
	if d, ok := productList[v]; ok {
		devicename := string(d)
		return &devicename, nil
	}
	return nil, fmt.Errorf("can't find specified device %s", v)
}

func parseOperationState(v string) (*string, error) {
	if state, ok := operationStates[v]; ok {
		return &state, nil
	}
	return nil, fmt.Errorf("can't find specified operation state %s", v)
}

func (f *Frame) parseKV(k, v string) error {
	switch k {
	case BatteryCurrent:
		i, err := parseInt(v)
		if err != nil {
			return fmt.Errorf("can't parse battery value %s - %w", v, err)
		}
		f.BatteryCurrent = i

	case DaySequenceNumber:
		d, err := parseInt(v)
		if err != nil {
			return fmt.Errorf("can't parse day sequence number value %s - %w", v, err)
		}
		f.DaySequenceNumber = d

	case ErrorCode:
		e, err := parseErrorCode(v)
		if err != nil {
			return err
		}
		f.ErrorCode = e

	case FirmwareVersion:
		f.FirmwareVersion = &v

	case InstantaneousPower:
		p, err := parseInt(v)
		if err != nil {
			return fmt.Errorf("can't parse instantaneous power %s - %w", v, err)
		}
		f.InstantaneousPower = p

	case LoadCurrent:
		c, err := parseInt(v)
		if err != nil {
			return fmt.Errorf("can't parse load current %s - %w", v, err)
		}
		f.LoadCurrent = c

	case LoadOutputState:
		s, err := parseLoadOutputState(v)
		if err != nil {
			return err
		}
		f.LoadOutputState = s

	case MainBatteryVoltage:
		s, err := parseInt(v)
		if err != nil {
			return fmt.Errorf("can't parse main battery voltage %s - %w", v, err)
		}
		f.MainBatteryVoltage = s

	case MaxPowerToday:
		s, err := parseInt(v)
		if err != nil {
			return fmt.Errorf("can't parse max power today %s - %w", v, err)
		}
		f.MaxPowerToday = s

	case MaxPowerYesterday:
		s, err := parseInt(v)
		if err != nil {
			return fmt.Errorf("can't parse max power yesterday %s - %w", v, err)
		}
		f.MaxPowerYesterday = s

	case PanelPower:
		s, err := parseInt(v)
		if err != nil {
			return fmt.Errorf("can't parse panel power %s - %w", v, err)
		}
		f.PanelPower = s

	case PanelVoltage:
		s, err := parseInt(v)
		if err != nil {
			return fmt.Errorf("can't parse panel voltage %s - %w", v, err)
		}
		f.PanelVoltage = s

	case ProductID:
		s, err := parseDevice(v)
		if err != nil {
			return fmt.Errorf("can't parse device %s - %w", v, err)
		}
		f.ProductName = s

	case RelayState:
		s, err := parseLoadOutputState(v)
		if err != nil {
			return fmt.Errorf("can't parse device %s - %w", v, err)
		}
		f.RelayState = s

	case SerialNumber:
		f.SerialNumber = &v

	case StateOfOperation:
		s, err := parseOperationState(v)
		if err != nil {
			return err
		}
		f.StateOfOperation = s

	case YieldToday:
		s, err := parseInt(v)
		if err != nil {
			return fmt.Errorf("can't parse yield today %s - %w", v, err)
		}
		f.YieldToday = s

	case YieldTotal:
		s, err := parseInt(v)
		if err != nil {
			return fmt.Errorf("can't parse yield total %s - %w", v, err)
		}
		f.YieldTotal = s

	case YieldYesterday:
		s, err := parseInt(v)
		if err != nil {
			return fmt.Errorf("can't parse yield yesterday %s - %w", v, err)
		}
		f.YieldYesterday = s

	}

	return nil
}
