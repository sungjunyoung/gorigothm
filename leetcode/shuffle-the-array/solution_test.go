package shuffle_the_array

import "testing"

func Test_shuffle(t *testing.T) {
	tests := []struct {
		nums   []int
		n      int
		expect []int
	}{
		{
			nums:   []int{2, 5, 1, 3, 4, 7},
			n:      3,
			expect: []int{2, 3, 5, 4, 1, 7},
		},
	}

	for _, test := range tests {
		result := shuffle(test.nums, test.n)
		if len(result) != len(test.expect) {
			t.Fatal("result length should be equal with expected")
		}

		for i := range result {
			if result[i] != test.expect[i] {
				t.Fatal("failed")
			}
		}
	}
}
