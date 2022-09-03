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
	YieldToday     int
	OperationState string

	InstantaneousPower int
}

func (s State) Update(f Frame) State {
	if f.MainBatteryVoltage != 0 {
		if s.currentIndex < numberOfVoltage {
			s.batteryVoltageBuffer[s.currentIndex] = f.MainBatteryVoltage
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

	s.PanelPower = f.PanelPower
	s.PanelVoltage = f.PanelVoltage
	s.BatteryCurrent = f.BatteryCurrent
	s.OperationState = f.StateOfOperation
	s.LoadCurrent = f.LoadCurrent
	s.LoadOutputState = f.LoadOutputState
	s.MaxPower = f.MaxPowerToday
	s.YieldToday = f.YieldToday
	s.InstantaneousPower = f.InstantaneousPower
	return s

}

func NewState() (State, error) {
	s := State{
		OperationState: StateAbsorption,
	}
	return s, nil
}
