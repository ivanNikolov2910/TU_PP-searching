package parallel_search

import (
	"sync"
)

// Global wait group to synchronize goroutines
var lsWaitGroup sync.WaitGroup

func ParallelLinearSearch(arr []int, key int, threadMax int) bool {
	found := make(chan bool)
	partitionSize := len(arr) / threadMax

	for i := 0; i < threadMax; i++ {
		lsWaitGroup.Add(1)
		start := i * partitionSize
		end := start + partitionSize
		go func(start, end int) {
			defer lsWaitGroup.Done()
			for j := start; j < end; j++ {
				if arr[j] == key {
					found <- true
					return
				}
			}
		}(start, end)
	}

	go func() {
		lsWaitGroup.Wait()
		close(found)
	}()

	keyFound := false
	for v := range found {
		if v {
			keyFound = true
			break
		}
	}

	return keyFound
}
