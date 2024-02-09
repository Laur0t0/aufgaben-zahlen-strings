package numberstrings

// DigitString1 erwartet eine Ziffer als int.
// Die Funktion gibt den zugehörigen String zurück, wie er üblicherweise
// an der Einer-Stelle einer Zahl >= 21 vorkommen würde.
// Außerdem wird bei Ziffern != 0 das Wort "und" angehängt.
//
// Beispiel: Für 3 soll der String "dreiund" geliefert werden.
//
// Anmerkung:
// Dies ist eine Hilfsfunktion, die genutzt werden soll,
// um den Gesamt-String einer Zahl zusammenzusetzen.
// Diese Funktion muss nur für den Normalfall (Zahlen >= 21) funktionieren.
func DigitString1(digit int) string {
	/* Hinweis:
	   Sie können eine Reihe von If-Anweisungen, eine Switch-Anweisung oder
	   auch eine Liste mit den Strings für die Ziffern verwenden.
	*/
	return []string{"", "einund", "zweiund", "dreiund", "vierund", "fünfund", "sechsund", "siebenund", "achtund", "neunund"}[digit]
}

// DigitString10 erwartet eine Ziffer als int.
// Die Funktion gibt den zugehörigen String zurück, wie er üblicherweise
// an der Zehner-Stelle einer Zahl >= 21 vorkommen würde.
//
// Beispiel: Für 3 soll der String "dreißig" geliefert werden.
//
// Anmerkung:
// Dies ist eine Hilfsfunktion, die genutzt werden soll,
// um den Gesamt-String einer Zahl zusammenzusetzen.
func DigitString10(digit int) string {
	/* Hinweis:
	   Sie können eine Reihe von If-Anweisungen, eine Switch-Anweisung oder
	   auch eine Liste mit den Strings für die Ziffern verwenden.
	*/
	return []string{"", "", "zwanzig", "dreißig", "vierzig", "fünfzig", "sechzig", "siebzig", "achtzig", "neunzig"}[digit]
}

// DigitString100 erwartet eine Ziffer als int.
// Die Funktion gibt den zugehörigen String zurück, wie er üblicherweise
// an der Hunderter-Stelle einer Zahl >= 21 vorkommen würde.
//
// Beispiel: Für 3 soll der String "dreihundert" geliefert werden.
//
// Anmerkung:
// Dies ist eine Hilfsfunktion, die genutzt werden soll,
// um den Gesamt-String einer Zahl zusammenzusetzen.
func DigitString100(digit int) string {
	/* Hinweis:
	   Sie können eine Reihe von If-Anweisungen, eine Switch-Anweisung oder
	   auch eine Liste mit den Strings für die Ziffern verwenden.
	*/
	return []string{"", "einhundert", "zweihundert", "dreihundert", "vierhundert", "fünfhundert", "sechshundert", "siebenhundert", "achthundert", "neunhundert"}[digit]
}

// NumberStringGreater20 erwartet eine Zahl >= 20 und liefert den zugehörigen String.
func NumberStringGreater20(n int) string {
	/* Hinweis:
	   Verwenden Sie Modulo- und Divisions-Operationen, um die Ziffern der Zahl zu bestimmen.
	   Verwenden Sie die Funktionen DigitString1, DigitString10 und DigitString100,
	   um die Ziffern in Strings umzuwandeln.
	*/
	if n%100 < 20 && n%100 > 0 {
		return DigitString100(n/100) + NumberStringBelow20(n%100)
	}
	return DigitString100(n/100) + DigitString1(n%10) + DigitString10((n%100)/10)
}

// NumberStringBelow20 erwartet eine Zahl >= 20 und liefert den zugehörigen String.
func NumberStringBelow20(n int) string {
	/* Hinweis: Verwenden Sie eine Switch-Anweisung oder eine Liste. */
	switch n {
	case 0:
		return "null"
	case 1:
		return "eins"
	case 2:
		return "zwei"
	case 3:
		return "drei"
	case 4:
		return "vier"
	case 5:
		return "fünf"
	case 6:
		return "sechs"
	case 7:
		return "sieben"
	case 8:
		return "acht"
	case 9:
		return "neun"
	case 10:
		return "zehn"
	case 11:
		return "elf"
	case 12:
		return "zwölf"
	case 13:
		return "dreizehn"
	case 14:
		return "vierzehn"
	case 15:
		return "fünfzehn"
	case 16:
		return "sechzehn"
	case 17:
		return "siebzehn"
	case 18:
		return "achtzehn"
	case 19:
		return "neunzehn"
	}
	return ""
}

// NumberString3Digits erwartet eine Zahl 0 <= n <= 999 und liefert den zugehörigen String.
func NumberString3Digits(n int) string {
	/* Hinweis:
	   Verwenden Sie die Funktionen NumberStringBelow20 und NumberStringGreater20.
	*/
	if n <= 19 {
		return NumberStringBelow20(n)
	}
	return NumberStringGreater20(n)
}

// NumberString6Digits erwartet eine Zahl 0 <= n <= 999999 und liefert den zugehörigen String.
func NumberString6Digits(n int) string {
	/* Hinweis:
	   Verwenden Sie Modulo- und Divisions-Operationen, um die Zahl in zwei Dreierblöcke zu zerlegen.
	   Verwenden Sie die Funktion NumberString3Digits, um die Ziffern in Strings umzuwandeln.
	*/
	result := ""
	if n == 0 {
		return "null"
	}
	low := n % 1000
	high := n / 1000
	if high != 0 {
		result += NumberString3Digits(high) + "tausend"
	}
	if low != 0 {
		result += NumberString3Digits(low)
	}
	return result
}
