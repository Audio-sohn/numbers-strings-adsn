package natural

// NumberStringBelow20 erwartet eine Zahl >= 20 und liefert den zugehörigen String.
func NumberStringBelow20(n int) string {

	unterZehn := []string{"null", "eins", "zwei", "drei", "vier", "fünf", "sechs", "sieben", "acht", "neun"}

	ZehnbisZwanzig := []string{"zehn", "elf", "zwölf", "dreizehn", "vierzehn", "fünfzehn", "sechzehn", "siebzehn", "achtzehn", "neunzehn"}

	if n < 10 {

		return unterZehn[n]

	} else if n < 20 {

		return ZehnbisZwanzig[n-10]

	}

	return "OVER"
}
