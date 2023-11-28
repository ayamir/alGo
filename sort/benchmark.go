package sort

import (
	"fmt"
	"math/rand"
	"reflect"
	"runtime"
	"time"
)

func randomNums(scale int) []int {
	rand.New(rand.NewSource(time.Now().Unix()))
	return rand.Perm(scale)
}

func deepCopySlice(source []int) []int {
	copied := make([]int, len(source))
	for i, v := range source {
		copied[i] = v
	}
	return copied
}

func calculateTotalTime(times []time.Duration) time.Duration {
	var totalTime time.Duration
	for _, t := range times {
		totalTime += t
	}
	return totalTime
}

func SortBenchmark(algo1 func(nums []int), algo2 func(nums []int), times int, scale int) {
	elapsedTimes1 := make([]time.Duration, times)
	elapsedTimes2 := make([]time.Duration, times)

	for i := 0; i < times; i++ {
		nums1 := randomNums(scale)
		nums2 := make([]int, scale)
		copy(nums2, nums1)

		start1 := time.Now()
		algo1(nums1)
		elapsed1 := time.Since(start1)
		elapsedTimes1[i] = elapsed1

		start2 := time.Now()
		algo2(nums2)
		elapsed2 := time.Since(start2)
		elapsedTimes2[i] = elapsed2
	}
	totalTime1 := calculateTotalTime(elapsedTimes1)
	totalTime2 := calculateTotalTime(elapsedTimes2)
	algoName1 := runtime.FuncForPC(reflect.ValueOf(algo1).Pointer()).Name()
	algoName2 := runtime.FuncForPC(reflect.ValueOf(algo2).Pointer()).Name()
	fmt.Printf("Total time for %s: %v\n", algoName1, totalTime1)
	fmt.Printf("Total time for %s: %v\n", algoName2, totalTime2)

	// Calculate the number of times Algorithm 1 is faster than Algorithm 2
	if totalTime1 < totalTime2 {
		fmt.Printf("%s is faster than %s by %.2f times\n", algoName1, algoName2, float64(totalTime2)/float64(totalTime1))
	} else if totalTime2 < totalTime1 {
		fmt.Printf("%s is faster than %s by %.2f times\n", algoName2, algoName1, float64(totalTime1)/float64(totalTime2))
	} else {
		fmt.Println("Both algorithms have the same total time")
	}
}
