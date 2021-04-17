package group_anagrams

import (
	"testing"
)

func Test_groupAnagrams(t *testing.T) {
	tests := []struct {
		strs   []string
		expect [][]string
	}{
		{
			strs:   []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			expect: [][]string{{"tan", "nat"}, {"eat", "tea", "ate"}, {"bat"}},
		},
		{
			strs:   []string{""},
			expect: [][]string{{""}},
		},
		{
			strs:   []string{"a"},
			expect: [][]string{{"a"}},
		},
	}

	for _, test := range tests {
		if len(groupAnagrams(test.strs)) != len(test.expect) {
			t.Fatal("failed")
		}
	}
}
