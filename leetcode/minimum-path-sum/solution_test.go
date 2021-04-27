package minimum_path_sum

import "testing"

func Test_minPathSum(t *testing.T) {
	tests := []struct {
		grid   [][]int
		expect int
	}{
		{
			grid:   [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}},
			expect: 7,
		},
	}

	for _, test := range tests {
		if minPathSum(test.grid) != test.expect {
			t.Fatal("failed")
		}
	}
}
