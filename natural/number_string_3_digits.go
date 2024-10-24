package natural

// NumberString3Digits erwartet eine Zahl 0 <= n <= 999 und liefert den zugehörigen String.
func NumberString3Digits(n int) string {
	if n <= 19 {
		return NumberStringBelow20(n)
	}
	return NumberStringGreater20(n)
}

// HINT
// Verwenden Sie die Funktionen NumberStringBelow20 und NumberStringGreater20.
