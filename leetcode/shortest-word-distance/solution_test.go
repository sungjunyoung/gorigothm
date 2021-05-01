package shortest_word_distance

import "testing"

func Test_shortestDistance(t *testing.T) {
	tests := []struct {
		wordsDict []string
		word1     string
		word2     string
		expect    int
	}{
		{
			wordsDict: []string{"practice", "makes", "perfect", "coding", "makes"},
			word1:     "coding",
			word2:     "practice",
			expect:    3,
		},
		{
			wordsDict: []string{"practice", "makes", "perfect", "coding", "makes"},
			word1:     "makes",
			word2:     "coding",
			expect:    1,
		},
	}

	for i, test := range tests {
		if shortestDistance(test.wordsDict, test.word1, test.word2) != test.expect {
			t.Errorf("test %d failed", i)
		}
	}
}
