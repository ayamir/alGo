package sort

import "sort"

func SelectionSort(nums []int) {
	n := len(nums)
	for i := 0; i < n; i++ {
		min_idx := i
		for j := i + 1; j < n; j++ {
			if nums[j] < nums[min_idx] {
				min_idx = j
			}
		}
		nums[i], nums[min_idx] = nums[min_idx], nums[i]
	}
}

func BubbleSort(nums []int) {
	n := len(nums)
	for i := n - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

func InsertionSort(nums []int) {
	n := len(nums)
	for i := 1; i < n; i++ {
		base := nums[i]
		j := i - 1
		for j > 0 && nums[j] > base {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = base
	}
}

func partition(nums []int) int {
	pivot := nums[0]
	left, right := 1, len(nums)-1
	for left < right {
		for left < right && nums[left] <= pivot {
			left++
		}
		for left < right && nums[right] >= pivot {
			right--
		}
		if left < right {
			nums[left], nums[right] = nums[right], nums[left]
		}
	}
	nums[0], nums[right] = nums[right], nums[0]
	return right
}

func QuickSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	pivotIndex := partition(nums)
	QuickSort(nums[:pivotIndex])
	QuickSort(nums[pivotIndex+1:])
}

func MergeSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	mid := len(nums) / 2
	left := nums[:mid]
	right := nums[mid:]
	MergeSort(left)
	MergeSort(right)
	merge(nums, left, right)
}

func merge(nums, left, right []int) {
	i, j, k := 0, 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			nums[k] = left[i]
			i++
		} else {
			nums[k] = right[j]
			j++
		}
		k++
	}
	for i < len(left) {
		nums[k] = left[i]
		i++
		k++
	}
	for j < len(right) {
		nums[k] = right[j]
		j++
		k++
	}
}

func HeapSort(nums []int) {
	n := len(nums)
	for i := n/2 - 1; i >= 0; i-- {
		heapify(nums, n, i)
	}
	for i := n - 1; i > 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapify(nums, i, 0)
	}
}

func heapify(nums []int, n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2
	if left < n && nums[left] > nums[largest] {
		largest = left
	}
	if right < n && nums[right] > nums[largest] {
		largest = right
	}
	if largest != i {
		nums[i], nums[largest] = nums[largest], nums[i]
		heapify(nums, n, largest)
	}
}

func getMaxValue(nums []int) int {
	ret := nums[0]
	for _, num := range nums {
		if num > ret {
			ret = num
		}
	}
	return ret
}

func BucketSort(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}
	maxValue := getMaxValue(nums)
	buckets := make([][]int, n)
	for _, num := range nums {
		index := (num * (n - 1)) / (maxValue + 1)
		buckets[index] = append(buckets[index], num)
	}
	pos := 0
	for _, bucket := range buckets {
		if len(bucket) > 0 {
			sort.Ints(bucket)
			copy(nums[pos:], bucket)
			pos += len(bucket)
		}
	}
}
