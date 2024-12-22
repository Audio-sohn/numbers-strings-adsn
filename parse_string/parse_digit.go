package parse_string

import (
	"slices"
)

// ParseDigit erwartet einen String von 0 bis 9 oder A bis F und liefert die zugehörige Zahl.
// Dabei gilt A=10, B=11, ..., F=15.
// Ist der String kein gültiger Wert, wird -1 zurückgegeben.
func ParseDigit(digit string) int {

	allowed_strings := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F"}

	//muss einstellig sein
	if len(digit) > 1 {

		return -1

	}
	//suche den string in der slice
	decoded_int, found := slices.BinarySearch(allowed_strings, digit)

	if found {

		//convenient: der index ist gleich der dekodierte integer
		return decoded_int

	}

	return -1
}
