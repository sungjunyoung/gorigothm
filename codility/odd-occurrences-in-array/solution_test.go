package odd_occurrences_in_array

import "testing"

func Test_Solution(t *testing.T) {
	tests := []struct {
		A      []int
		expect int
	}{
		{
			A:      []int{9, 3, 9, 3, 9, 7, 9},
			expect: 7,
		},
	}

	for _, test := range tests {
		if Solution(test.A) != test.expect {
			t.Fatal("failed")
		}
	}
}
