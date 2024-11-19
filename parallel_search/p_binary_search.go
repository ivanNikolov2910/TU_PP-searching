package parallel_search

import (
	"sync"
)

// Global wait group to synchronize goroutines
var bsWaitGroup sync.WaitGroup

func ParallelBinarySearch(arr []int, key int, threadMax int) bool {
	result := make(chan int)
	partitionSize := len(arr) / threadMax

	for i := 0; i < threadMax; i++ {
		bsWaitGroup.Add(1)
		start := i * partitionSize
		end := start + partitionSize - 1
		if i == threadMax-1 {
			end = len(arr) - 1
		}
		go func(start, end int) {
			defer bsWaitGroup.Done()
			low, high := start, end
			for low <= high {
				mid := low + (high-low)/2
				if arr[mid] == key {
					result <- mid
					return
				} else if arr[mid] < key {
					low = mid + 1
				} else {
					high = mid - 1
				}
			}
			result <- -1
		}(start, end)
	}

	go func() {
		bsWaitGroup.Wait()
		close(result)
	}()

	keyFound := -1
	for res := range result {
		if res != -1 {
			keyFound = res
			break
		}
	}

	if keyFound != -1 {
		return true
	}
	return false
}
