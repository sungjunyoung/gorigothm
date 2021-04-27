package possible_bipartition1

import "testing"

func Test_possibleBipartition(t *testing.T) {
	tests := []struct {
		N        int
		dislikes [][]int
		expect   bool
	}{
		{
			N:        4,
			dislikes: [][]int{{1, 2}, {1, 3}, {2, 4}},
			expect:   true,
		},
		{
			N:        3,
			dislikes: [][]int{{1, 2}, {1, 3}, {2, 3}},
			expect:   false,
		},
	}

	for _, test := range tests {
		if possibleBipartition(test.N, test.dislikes) != test.expect {
			t.Fatal("failed")
		}
	}
}
