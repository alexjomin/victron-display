package tools

func Cube(v float64) float64 {
	return v * v * v
}
func Square(v float64) float64 {
	return v * v
}

func WaterProbe(v uint16) float64 {
	f := float64(v)
	r := 0.0495*Cube(f) - 1.3180*Square(f) + 18.0382*f - 9.0416
	if r < 0 {
		return 0
	} else {
		return r
	}
}
