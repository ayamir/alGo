package main

import (
	"algo/sort"
)

func main() {
	sort.SortBenchmark(sort.QuickSort, sort.BubbleSort, 1, 10000)
}
