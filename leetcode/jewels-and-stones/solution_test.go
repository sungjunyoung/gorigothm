package jewels_and_stones

import "testing"

func Test_numJewelsInStones(t *testing.T) {
	tests := []struct {
		jewels string
		stones string
		expect int
	}{
		{
			jewels: "aA",
			stones: "aAAbbbb",
			expect: 3,
		},
	}

	for _, test := range tests {
		if numJewelsInStones(test.jewels, test.stones) != test.expect {
			t.Fatal("failed")
		}
	}
}
