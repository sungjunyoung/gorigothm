package zigzag_conversion

func convert(s string, numRows int) string {
	if len(s) == 1 {
		return s
	}
	if numRows == 1 {
		return s
	}

	length := len(s)

	var metrics [][]string
	for i := 0; i < length; i++ {
		metrics = append(metrics, []string{})
		for j := 0; j < length; j++ {
			metrics[i] = append(metrics[i], string(s[j]))
		}
	}

	result := ""
	for i := 0; i < numRows; i++ {
		mustPoint := i
		additionalPoint := numRows*2 - 2 - i
		for {
			if mustPoint > length-1 && additionalPoint > length-1 {
				break
			}

			if length > mustPoint {
				result += metrics[i][mustPoint]
			}

			if length > additionalPoint && i != 0 && i != numRows-1 {
				result += metrics[i][additionalPoint]
			}

			mustPoint += (2 * numRows) - 2
			additionalPoint += (2 * numRows) - 2
		}
	}
	return result
}
