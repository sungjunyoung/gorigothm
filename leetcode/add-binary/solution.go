package add_binary

func addBinary(a string, b string) string {
	additional := 0
	aIndex := len(a) - 1
	bIndex := len(b) - 1
	result := ""
	for aIndex >= 0 || bIndex >= 0 {
		aSub := 0
		bSub := 0
		if aIndex >= 0 {
			if a[aIndex] == '1' {
				aSub = 1
			}
		}
		if bIndex >= 0 {
			if b[bIndex] == '1' {
				bSub = 1
			}
		}

		subResult := additional + aSub + bSub
		if subResult > 2 {
			additional = 1
			subResult = 1
		} else if subResult > 1 {
			additional = 1
			subResult = 0
		} else {
			additional = 0
		}

		if subResult == 0 {
			result = "0" + result
		} else {
			result = "1" + result
		}

		aIndex--
		bIndex--
	}

	if additional > 0 {
		result = "1" + result
	}

	return result
}
