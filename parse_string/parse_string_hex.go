package parse_string

import (
	"math"
	"slices"
)

// ParseStringDecimal erwartet einen String, der eine Hexadezimalzahl repräsentiert, und liefert die zugehörige Zahl.
// Ist der String kein gültiger Wert, wird -1 zurückgegeben.
func ParseStringHexadecimal(s string) int {

	result := 0

	s_slice := []byte(s)
	slices.Reverse((s_slice))

	for i, speci := range s_slice {

		//stellenwert pro ziffer berechnen
		digitvalue := ParseDigit(string(speci)) * int(math.Pow(16, float64(i)))

		//error catchen, ergebnis muss negativ sein wenn einer der chars illegal ist
		if digitvalue < 0 {

			return -1

		}

		//stellenwerte aufaddieren um endgültigen integer zu errechnen
		result += digitvalue

	}

	return result
}
