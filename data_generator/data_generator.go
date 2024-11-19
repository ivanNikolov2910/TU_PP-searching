package data_generator

import (
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"
)

func GenerateOrderedData(size int, maxElement int) []int {
	return generateRandomOrderedData(size, maxElement)
}

func GenerateReverseOrderedData(size int, maxElement int) []int {
	data := generateRandomOrderedData(size, maxElement)
	reverseSlice(data)
	return data
}

func GenerateRandomData(size int, maxElement int) []int {
	return generateRandomData(size, maxElement)
}

func generateRandomOrderedData(size int, maxElement int) []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	data := make([]int, size)

	for i := 0; i < size; i++ {
		data[i] = r.Intn(maxElement + 1)
	}
	sort.Ints(data)
	return data
}

func generateRandomData(size int, maxElement int) []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	data := make([]int, size)

	for i := range data {
		data[i] = r.Intn(maxElement)
	}
	return data
}

func reverseSlice(data []int) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}

func intSliceToString(data []int) string {
	strData := make([]string, len(data))
	for i, v := range data {
		strData[i] = strconv.Itoa(v)
	}
	return strings.Join(strData, ",")
}
