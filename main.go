package main

import (
	"fmt"
)

type Node struct {
	value byte
	left  *Node
	right *Node
}

func (n *Node) add(value byte) bool {
	if n.value == value {
		return true
	}
	if n.value > value {
		if n.left == nil {
			n.left = &Node{value: value}
		} else {
			if ok := n.left.add(value); ok {
				return true
			}
		}
	} else if n.value < value {
		if n.right == nil {
			n.right = &Node{value: value}
		} else {
			if ok := n.right.add(value); ok {
				return true
			}
		}
	}
	return false
}

func findLongestUniqueSubstring(line string) int {
	count := 0
	for i := 0; i < len(line); i++ {
		tempCount := 1
		tree := &Node{value: line[i]}
		for _, char := range line[i+1:] {
			if ok := tree.add(byte(char)); ok {
				break
			}
			tempCount++
		}
		if count < tempCount {
			count = tempCount
		}
		if count >= len(line[i:]) {
			return count
		}
	}
	return count
}

func main() {
	lines := []string{"abcbada", "axbxcxd", "aaaaaaaa", "abcdefg"}

	for _, line := range lines {
		result := findLongestUniqueSubstring(line)
		fmt.Println(result)
	}
}
