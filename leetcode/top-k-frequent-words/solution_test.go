package top_k_frequent_words

import (
	"reflect"
	"testing"
)

func Test_topKFrequent(t *testing.T) {
	tests := []struct {
		words   []string
		k       int
		expects []string
	}{
		{
			words:   []string{"i", "love", "leetcode", "i", "love", "coding"},
			k:       3,
			expects: []string{"i", "love", "coding"},
		},
		{
			words:   []string{"i", "love", "leetcode", "i", "love", "coding"},
			k:       2,
			expects: []string{"i", "love"},
		},
		{
			words:   []string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"},
			k:       4,
			expects: []string{"the", "is", "sunny", "day"},
		},
	}

	for _, test := range tests {
		if !reflect.DeepEqual(topKFrequent(test.words, test.k), test.expects) {
			t.Fatal("failed")
		}
	}
}
