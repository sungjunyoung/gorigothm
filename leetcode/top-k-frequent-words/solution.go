package top_k_frequent_words

import (
	"container/heap"
)

type wordCounts []wordCount
type wordCount struct {
	word  string
	count int
}

func (wcs wordCounts) Len() int {
	return len(wcs)
}

func (wcs wordCounts) Less(i, j int) bool {
	if wcs[i].count != wcs[j].count {
		return wcs[i].count < wcs[j].count
	} else {
		return wcs[i].word > wcs[j].word
	}
}

func (wcs wordCounts) Swap(i, j int) {
	wcs[i], wcs[j] = wcs[j], wcs[i]
}

func (wcs *wordCounts) Push(x interface{}) {
	*wcs = append(*wcs, x.(wordCount))
}

func (wcs *wordCounts) Pop() interface{} {
	target := (*wcs)[len(*wcs)-1]
	*wcs = (*wcs)[:len(*wcs)-1]
	return target
}

func topKFrequent(words []string, k int) []string {
	counter := map[string]int{}
	for _, word := range words {
		counter[word] += 1
	}

	var wcs wordCounts
	heap.Init(&wcs)
	for word, count := range counter {
		heap.Push(&wcs, wordCount{
			word:  word,
			count: count,
		})

		if wcs.Len() > k {
			heap.Pop(&wcs)
		}
	}

	result := make([]string, k)
	for i := k - 1; i >= 0; i-- {
		result[i] = heap.Pop(&wcs).(wordCount).word
	}

	return result
}
