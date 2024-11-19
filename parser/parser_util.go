package parser

import (
	"strconv"
	"strings"
	"sync"

	"parallel.searching/v2/model"
)

// ParseStringToArray Function to parse a string and return an array of integers
func ParseStringToArray(data string) []int {
	data = strings.TrimSpace(data)

	n := 1
	for i := 0; i < len(data); i++ {
		if data[i] == ',' {
			n++
		}
	}

	arr := make([]int, 0, n)
	start := 0

	for i := 0; i < len(data); i++ {
		if data[i] == ',' {
			if val, err := strconv.Atoi(strings.TrimSpace(data[start:i])); err == nil {
				arr = append(arr, val)
			}
			start = i + 1
		}
	}

	if val, err := strconv.Atoi(strings.TrimSpace(data[start:])); err == nil {
		arr = append(arr, val)
	}

	return arr
}

// ParseArrayToString Function to parse an array of integers and return a string
func ParseArrayToString(arr []int) string {
	var builder strings.Builder
	builder.Grow(len(arr))

	for i, val := range arr {
		if i > 0 {
			builder.WriteString(", ")
		}
		builder.WriteString(strconv.Itoa(val))
	}

	return builder.String()
}

// ParseArrayToSet Function to parse a string and return a set (hash map)
func ParseArrayToSet(data []int) map[int]bool {
	var wg sync.WaitGroup
	set := sync.Map{}

	numWorkers := 8
	chunkSize := len(data) / numWorkers

	for i := 0; i < numWorkers; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == numWorkers-1 {
			end = len(data)
		}

		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			for _, d := range data[start:end] {
				set.Store(d, true)
			}
		}(start, end)
	}

	wg.Wait()

	result := make(map[int]bool, len(data))
	set.Range(func(key, value interface{}) bool {
		result[key.(int)] = true
		return true
	})

	return result
}

// ParseArrayToNode converts a slice of integers into a binary tree.
func ParseArrayToNode(data []int) *model.Node {
	return model.BuildNode(data, 0, len(data)-1)
}
