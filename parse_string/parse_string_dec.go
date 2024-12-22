package parse_string

import (
	"strconv"
)

// ParseStringDecimal erwartet einen String, der eine Dezimalzahl repräsentiert, und liefert die zugehörige Zahl.
// Ist der String kein gültiger Wert, wird -1 zurückgegeben.
func ParseStringDecimal(s string) int {

	//convenient: go standard library hat eine atoi funktion
	result, err := strconv.Atoi(s)

	//error handling
	if err != nil {

		return -1

	}

	return result
}
