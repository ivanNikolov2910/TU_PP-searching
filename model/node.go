package model

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// BuildNode builds a binary tree from a slice.
func BuildNode(data []int, start, end int) *Node {
	if start > end {
		return nil
	}

	mid := start + (end-start)/2
	root := &Node{Value: data[mid]}

	root.Left = BuildNode(data, start, mid-1)
	root.Right = BuildNode(data, mid+1, end)

	return root
}
