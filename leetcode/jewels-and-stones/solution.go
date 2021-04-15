package jewels_and_stones

func numJewelsInStones(jewels string, stones string) int {
	jmap := map[string]bool{}
	for _, j := range jewels {
		jmap[string(j)] = true
	}

	result := 0
	for _, s := range stones {
		if jmap[string(s)] {
			result += 1
		}
	}

	return result
}
