package vedirect

const (
	numberOfVoltage = 10
)

type State struct {
	currentIndex         int
	batteryVoltageBuffer [numberOfVoltage]int
	batteryVoltageReady  bool

	PanelPower   int
	PanelVoltage int

	LoadCurrent     int
	LoadOutputState bool

	MinBatteryVoltage int
	MaxBatteryVoltage int
	BatteryVoltage    int
	BatteryCurrent    int

	MaxPower       int
	OperationState string
}

func (s *State) Update(f Frame) {
	if f.MainBatteryVoltage != nil {
		if s.currentIndex < numberOfVoltage {
			s.batteryVoltageBuffer[s.currentIndex] = *f.MainBatteryVoltage
			s.currentIndex += 1
		} else {
			s.currentIndex = 0
			s.batteryVoltageReady = true
		}
	}

	if s.batteryVoltageReady {
		buffer := 0
		for _, v := range s.batteryVoltageBuffer {
			buffer += v
		}

		s.BatteryVoltage = buffer / numberOfVoltage

		if s.MinBatteryVoltage == 0 {
			s.MinBatteryVoltage = s.BatteryVoltage
		}

		if s.MaxBatteryVoltage == 0 {
			s.MaxBatteryVoltage = s.BatteryVoltage
		}

		if s.BatteryVoltage < s.MinBatteryVoltage {
			s.MinBatteryVoltage = s.BatteryVoltage
		}

		if s.BatteryVoltage > s.MaxBatteryVoltage {
			s.MaxBatteryVoltage = s.BatteryVoltage
		}
	}

	if f.PanelPower != nil {
		s.PanelPower = *f.PanelPower
	}

	if f.PanelVoltage != nil {
		s.PanelVoltage = *f.PanelVoltage
	}

	if f.BatteryCurrent != nil {
		s.BatteryCurrent = *f.BatteryCurrent
	}

	if f.StateOfOperation != nil {
		s.OperationState = *f.StateOfOperation
	}

	if f.LoadCurrent != nil {
		s.LoadCurrent = *f.LoadCurrent
	}

	if f.LoadOutputState != nil {
		s.LoadOutputState = *f.LoadOutputState
	}

	if f.MaxPowerToday != nil {
		s.MaxPower = *f.MaxPowerToday
	}

}

func NewState() (*State, error) {
	s := State{}
	return &s, nil
}
