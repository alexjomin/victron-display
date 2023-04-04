package vedirect

import (
	"errors"
)

func Ftoa(f float64) string {
	if f == 0 {
		return "0"
	}
	// Définir une valeur d'exposant
	exp := 1
	for i := 0; i < 6; i++ {
		if f >= float64(exp)*10 || f == float64(int(f)) {
			break
		}
		exp /= 10
	}
	// Appliquer l'exposant pour obtenir le nombre entier
	i := int(f*float64(exp) + 0.5)
	// Convertir le nombre entier en string
	s := ""
	for i > 0 {
		digit := i % 10
		s = string(rune(digit+48)) + s
		i /= 10
	}
	// Ajouter le signe "-" pour les nombres négatifs
	if f < 0 {
		s = "-" + s
	}
	// Ajouter les décimales
	if exp > 1 {
		s = s[:len(s)-exp+1] + "." + s[len(s)-exp+1:]
	}
	return s

}

func Atoi(s string) (int, error) {
	result := 0
	sign := 1
	for _, c := range s {
		if c == '-' {
			sign = -1
		} else if c >= '0' && c <= '9' {
			result = result*10 + int(c-'0')
		} else {
			return 0, errors.New("invalid characters")
		}
	}
	return sign * result, nil
}

const smallsString = "00010203040506070809" +
	"10111213141516171819" +
	"20212223242526272829" +
	"30313233343536373839" +
	"40414243444546474849" +
	"50515253545556575859" +
	"60616263646566676869" +
	"70717273747576777879" +
	"80818283848586878889" +
	"90919293949596979899"

const digits = "0123456789abcdefghijklmnopqrstuvwxyz"

func small(i int) string {
	if i < 10 {
		return digits[i : i+1]
	}
	return smallsString[i*2 : i*2+2]
}

func IntToString(i int) string {
	return small(int(i))
}

// FormatVoltage transform a int value to something like 12.50 V
// Here we tried to avoid head allocation
func FormatVoltage(n int) string {
	ns := IntToString(n)

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
	return IntToString(n)
	// return fmt.Sprintf("%d W", n)
}

func intToFloat(i int) float64 {
	sign := 1.0
	if i < 0 {
		sign = -1.0
		i = -i
	}
	f := 0.0
	for i > 0 {
		f = f*10 + float64(i%10)
		i /= 10
	}
	return sign * f
}
