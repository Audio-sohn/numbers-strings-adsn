package natural

// NumberStringGreater20 erwartet eine Zahl >= 20 und liefert den zugehörigen String.
func NumberStringGreater20(n int) string {

	//wenn zahl zu groß, OVER zurückgeben
	if n > 999 {

		return "OVER"

	} else if n < 20 { // wenn unter 20, spezielle routine verwenden

		return NumberStringBelow20(n)

	}

	//slice für digits initialisieren
	number_digits := make([]int, 3)

	//buffer für zwischenergebnis
	buf := n

	//nummer in digits zerlegen
	for i := 0; i < len(number_digits); i++ {

		number_digits[i] = buf % 10
		buf = buf / 10

	}

	//Sonderfall: wenn die letzten zwei stellen unter 20 ergeben
	if number_digits[1] < 2 {

		if (number_digits[1] + number_digits[0]) == 0 {

			return DigitString100(number_digits[2])
		}

		return DigitString100(number_digits[2]) + NumberStringBelow20(n-number_digits[2]*100)

	}

	//zahl mit drei stellen, letze 2 stellen über 20
	return DigitString100(number_digits[2]) + DigitString1(number_digits[0]) + DigitString10(number_digits[1])
}
