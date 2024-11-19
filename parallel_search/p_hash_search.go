package parallel_search

import (
	"context"
	"sync"
)

// Global wait group to synchronize goroutines
var hsWaitGroup sync.WaitGroup

func ParallelHashSearch(hashTable map[int]bool, key int, threadMax int) bool {
	found := make(chan bool, 1)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	itemCount := len(hashTable)
	partitionSize := itemCount / threadMax

	i := 0
	for k := range hashTable {
		if i%partitionSize == 0 && i/partitionSize < threadMax {
			hsWaitGroup.Add(1)
			go func(startKey int) {
				defer hsWaitGroup.Done()
				for checkKey, _ := range hashTable {
					select {
					case <-ctx.Done():
						return
					default:
						if checkKey == key {
							found <- true
							cancel()
							return
						}
					}
				}
			}(k)
		}
		i++
	}

	go func() {
		hsWaitGroup.Wait()
		close(found)
	}()

	keyFound := false
	for res := range found {
		if res {
			keyFound = true
			break
		}
	}

	return keyFound
}
