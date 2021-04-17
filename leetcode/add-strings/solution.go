package add_strings

import (
	"fmt"
	"strconv"
)

func addStrings(num1 string, num2 string) string {
	moreNum := 0
	result := ""
	for len(num1) != 0 || len(num2) != 0 {
		n1Str := "0"
		if len(num1) > 0 {
			n1Str, num1 = string(num1[len(num1)-1]), num1[:len(num1)-1]
		}

		n2Str := "0"
		if len(num2) > 0 {
			n2Str, num2 = string(num2[len(num2)-1]), num2[:len(num2)-1]
		}

		n1, _ := strconv.Atoi(n1Str)
		n2, _ := strconv.Atoi(n2Str)

		res := n1 + n2 + moreNum
		if res/10 > 0 {
			moreNum = 1
		} else {
			moreNum = 0
		}
		res = res % 10
		result = fmt.Sprintf("%d", res) + result
	}

	if moreNum == 1 {
		result = "1" + result
	}

	return result
}
