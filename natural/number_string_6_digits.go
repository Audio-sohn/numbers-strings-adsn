package natural

// NumberString6Digits erwartet eine Zahl 0 <= n <= 999999 und liefert den zugehörigen String.
func NumberString6Digits(n int) string {

	//error catchen
	if n > 999999 {

		return "OVER"

	} else if n < 20 {

		return NumberStringBelow20(n)

	}

	//slice für digits initialisieren
	number_digits := make([]int, 6)

	//buffer für zwischenergebnis
	buf := n

	//n in slice aus digits zerlegen
	for i := 0; i < len(number_digits); i++ {

		number_digits[i] = buf % 10
		buf = buf / 10

	}

	//hunderter sowie tausender stellen in 2 x dreistellige integer zerlegen
	hunderter_int := number_digits[0] + number_digits[1]*10 + number_digits[2]*100
	tausender_int := number_digits[3] + number_digits[4]*10 + number_digits[5]*100

	//beide integer in strings konvertieren
	hunderter_string := NumberStringGreater20(hunderter_int)
	tausender_string := NumberStringGreater20(tausender_int)

	//Fallunterscheidungen bei null
	if tausender_int == 0 {

		return NumberStringGreater20(hunderter_int)

	} else if hunderter_int == 0 {

		return NumberStringGreater20(tausender_int) + "tausend"

	}

	//default return
	return tausender_string + "tausend" + hunderter_string
}
