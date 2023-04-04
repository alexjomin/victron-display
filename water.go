package main

type WaterLevel struct {
	Level000 bool
	Level033 bool
	Level066 bool
	Level100 bool
}

func GetWaterLevel(w WaterLevel) string {
	switch {
	case w.Level100:
		return "100%"
	case w.Level066:
		return "66%"
	case w.Level033:
		return "33%"
	default:
		return "0%"
	}
}
