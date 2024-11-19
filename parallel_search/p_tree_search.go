package parallel_search

import (
	"fmt"
	"sync"

	. "parallel.searching/v2/model"
)

// Global wait group to synchronize goroutines
var tsWaitGroup sync.WaitGroup

func ParallelTreeSearch(root *Node, key int) bool {
	found := make(chan bool)
	tsWaitGroup.Add(1)
	go func() {
		defer tsWaitGroup.Done()
		fmt.Println("Calling `parellelSearch` for a thread")
		parallelSearch(root, key, found)
	}()

	go func() {
		tsWaitGroup.Wait()
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

func parallelSearch(node *Node, key int, found chan bool) {
	if node == nil {
		return
	}

	if node.Value == key {
		found <- true
		return
	}

	if node.Left != nil {
		tsWaitGroup.Add(1)
		go func() {
			defer tsWaitGroup.Done()
			parallelSearch(node.Left, key, found)
		}()
	}

	if node.Right != nil {
		tsWaitGroup.Add(1)
		go func() {
			defer tsWaitGroup.Done()
			parallelSearch(node.Right, key, found)
		}()
	}
}
