package quicksort

func quickSort(arr []int) {
	start := 0
	end := len(arr) - 1
	divide(arr, start, end)
}

func divide(arr []int, start int, end int) {
	if start < end {
		p := partition(arr, start, end)

		divide(arr, start, p-1)
		divide(arr, p+1, end)
	}
}

func partition(arr []int, start int, end int) int {
	pivot := arr[end]
	i := start - 1

	for j := start; j <= end-1; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[end] = arr[end], arr[i+1]
	return i + 1
}
