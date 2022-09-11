package vedirect

import "fmt"

// FormatVoltage transform a int value to something like 12.50 V
// Here we tried to avoid head allocation
func FormatVoltage(n int) string {
	ns := fmt.Sprint(n)

	start := 0
	end := len(ns) - 1

	switch {
	case n < 10:
		return "0.00" + " V"
	case n < 100:
		return "0.0" + ns[start:end] + " V"
	case n < 1000:
		return "0." + ns[start:end] + " V"
	case n < 10000:
		return ns[0:1] + "." + ns[1:end] + " V"
	case n < 100000:
		return ns[0:2] + "." + ns[2:end] + " V"
	}

	return ns

}

func FormatPower(n int) string {
	return fmt.Sprintf("%d W", n)
}
