package number_of_islands

import "testing"

func Test_numIslands(t *testing.T) {
	tests := []struct {
		grid   [][]byte
		expect int
	}{
		{
			grid: [][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			expect: 1,
		},
		{
			grid: [][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			expect: 3,
		},
	}

	for _, test := range tests {
		if numIslands(test.grid) != test.expect {
			t.Fatal("failed")
		}
	}
}
