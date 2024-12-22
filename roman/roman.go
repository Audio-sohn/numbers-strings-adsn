package roman

// NToI erwartet eine Zahl und liefert die entsprechende Anzahl an I als String.
func NToI(n int) string {

	result := ""

	for i := 0; i < n; i++ {

		result += "I"

	}

	return result
}

// NToX erwartet eine Zahl und liefert die entsprechende Anzahl an X als String.
func NToX(n int) string {
	// TODO
	return ""
}

// RomanBelow10 erwartet eine Zahl 1 <= n <= 10 und liefert die römische Schreibweise für n als String.
func RomanBelow10(n int) string {

	switch n {

	case 0:

		return ""

	case 1:

		return "I"

	case 2:

		return "II"

	case 3:

		return "III"

	case 4:

		return "IV"

	case 5:

		return "V"

	case 6:

		return "VI"

	case 7:

		return "VII"

	case 8:

		return "VIII"

	case 9:

		return "IX"

	case 10:

		return "X"

	default:

		return "ERROR"

	}

}

// RomanBelow100 erwartet eine Zahl 1 <= n <= 100 und liefert die römische Schreibweise für n als String.
func RomanBelow100(n int) string {

	//wenn unter 11 dann andere routine benutzen
	if n < 11 {

		return RomanBelow10(n)

	}

	//buffer für return wert aufmachen
	result := ""

	//rest bei division durch 10 ermitteln
	rest := n % 10

	//ganzzahlige vielfache von 10 ermitteln
	tens := n / 10

	//10er stelle aufbauen
	switch tens {

	case 0:

		result += ""

	case 1:

		result += "X"

	case 2:

		result += "XX"

	case 3:

		result += "XXX"

	case 4:

		result += "XL"

	case 5:

		result += "L"

	case 6:

		result += "LX"

	case 7:

		result += "LXX"

	case 8:

		result += "LXXX"

	case 9:

		result += "XC"

	case 10:

		result += "C"

	}

	//1er stelle anhängen
	result += RomanBelow10(rest)

	return result

}
