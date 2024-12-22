package parse_string

import "strconv"

// ParseString erwartet einen String, der entweder eine Dezimal- oder eine Hexadezimalzahl repräsentiert, und liefert die zugehörige Zahl.
// Ist der String kein gültiger Wert, wird -1 zurückgegeben.
func ParseString(s string) int {

	intFromIntString, err := strconv.Atoi(s)

	if err == nil {

		return intFromIntString

	}

	return ParseStringHexadecimal(s)

}
