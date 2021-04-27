package possible_bipartition1

func possibleBipartition(N int, dislikes [][]int) bool {
	graph := make([][]int, N+1)
	visits := make([]bool, N+1)
	colors := make([]bool, N+1)

	for _, dislike := range dislikes {
		a := dislike[0]
		b := dislike[1]
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	for i := 1; i < N+1; i++ {
		if !visits[i] {
			visits[i] = true
			res := traverse(i, graph, visits, colors)
			if !res {
				return false
			}
		}
	}

	return true
}

func traverse(current int, graph [][]int, visits []bool, colors []bool) bool {
	for _, next := range graph[current] {
		if !visits[next] {
			visits[current] = true
			colors[next] = !colors[current]
			res := traverse(next, graph, visits, colors)

			if !res {
				return false
			}
		} else if colors[current] == colors[next] {
			return false
		}
	}

	return true
}
