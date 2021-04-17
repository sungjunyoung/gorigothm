package group_anagrams

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	resultMap := map[string][]string{}

	for _, str := range strs {
		key := parseAnagramKey(str)
		resultMap[key] = append(resultMap[key], str)
	}

	var result [][]string
	for _, vals := range resultMap {
		result = append(result, vals)
	}
	return result
}

func parseAnagramKey(str string) string {
	mapper := map[string]int{}
	for _, s := range str {
		mapper[string(s)] += 1
	}

	keys := make([]string, 0, len(mapper))
	for k := range mapper {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	key := ""
	for _, k := range keys {
		key += fmt.Sprintf("%s%d", k, mapper[k])
	}
	return key
}
