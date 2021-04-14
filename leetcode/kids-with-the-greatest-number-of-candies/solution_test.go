package kids_with_the_greatest_number_of_candies

import "testing"

func Test_kidsWithCandies(t *testing.T) {
	tests := []struct {
		candies      []int
		extraCandies int
		expect       []bool
	}{
		{
			candies:      []int{2, 3, 5, 1, 3},
			extraCandies: 3,
			expect:       []bool{true, true, true, false, true},
		},
		{
			candies:      []int{4, 2, 1, 1, 2},
			extraCandies: 1,
			expect:       []bool{true, false, false, false, false},
		},
	}

	for _, test := range tests {
		result := kidsWithCandies(test.candies, test.extraCandies)
		if len(result) != len(test.expect) {
			t.Fatal("result length should be equal")
		}

		for i := range result {
			if result[i] != test.expect[i] {
				t.Fatal("failed")
			}
		}
	}
}
