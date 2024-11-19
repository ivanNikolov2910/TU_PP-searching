package search

import (
	"fmt"
	generator "parallel.searching/v2/data_generator"
)

type SearchType string
type DataType string

const (
	LINEAR  SearchType = "linear"
	BINARY  SearchType = "binary"
	HASH    SearchType = "hash"
	TREE    SearchType = "tree"
	MANUAL  DataType   = "manual"
	ORDERED DataType   = "ordered"
	REVERSE DataType   = "reverse"
	RANDOM  DataType   = "random"
)

type SearchResult struct {
	Result      bool       `json:"result"`
	SearchType  SearchType `json:"search_type"`
	SearchedKey int        `json:"searched_key"`
	ListSize    int        `json:"list_size"`
	Data        string     `json:"data"`
	Range       string     `json:"range"`
	SearchTime  string     `json:"search_time"`
}

type Context struct {
	SearchType SearchType
	Data       []int
	DataRange  int
	Size       int
	Key        int
}

// GetDataByType generates data based on the type provided.
func GetDataByType(dataType DataType, size, maxElement int) ([]int, error) {
	switch dataType {
	case ORDERED:
		return generator.GenerateOrderedData(size, maxElement), nil
	case REVERSE:
		return generator.GenerateReverseOrderedData(size, maxElement), nil
	case RANDOM:
		return generator.GenerateRandomData(size, maxElement), nil
	default:
		return nil, fmt.Errorf("invalid data type")
	}
}
