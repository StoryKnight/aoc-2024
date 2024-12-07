package day07

import (
	"aoc-2024/utils"
	"fmt"
	"strconv"
	"strings"
)

// TreeNode represents a node in the tree.
type TreeNode struct {
	Value  int
	Symbol rune
	Left   *TreeNode
	Right  *TreeNode
}

// Tree represents the entire binary tree.
type Tree struct {
	Root *TreeNode
}

// Add fills the next layer of the tree with nodes containing the given value.
// Left child gets '+' symbol, and the right child gets '*'.

func (t *Tree) SetRoot(value int) {

	newNode := func(value int) *TreeNode {
		return &TreeNode{Value: value}
	}
	t.Root = newNode(value)
}

func (t *Tree) Add(value int) {
	newNode := func(symbol rune) *TreeNode {
		return &TreeNode{Value: value, Symbol: symbol}
	}

	if t.Root == nil {
		// Create the root node if the tree is empty
		t.Root = newNode(' ')
		return
	}

	// Use a queue to perform a level-order traversal
	queue := []*TreeNode{t.Root}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		// Check if the current node is missing children
		if current.Left == nil {
			current.Left = newNode('+')
			current.Right = newNode('*')
			return
		}

		// Add the children of the current node to the queue for further processing
		queue = append(queue, current.Left, current.Right)
	}
}

// PrintLevelOrder prints the tree in a level-order format for debugging.
func (t *Tree) PrintLevelOrder() {
	if t.Root == nil {
		fmt.Println("Tree is empty.")
		return
	}

	queue := []*TreeNode{t.Root}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		fmt.Printf("%d(%c) ", current.Value, current.Symbol)

		if current.Left != nil {
			queue = append(queue, current.Left)
		}
		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}
	fmt.Println()
}

func (t *Tree) isValid(node *TreeNode) bool {

	return false
}

func Part1() {
	lines := utils.ReadLines("problems/day07/test.txt")

	for _, line := range lines {

		parts := strings.Split(line, ":")
		total, _ := strconv.Atoi(strings.TrimSpace(parts[0]))

		tree := &Tree{}
		tree.SetRoot(total)

		operands := make([]int, len(strings.Fields(parts[1])))

		// add operands to array
		for i, s := range strings.Fields(parts[1]) {
			num, _ := strconv.Atoi(s)
			operands[i] = num

			tree.Add(num)
			tree.Add(num)
		}

		fmt.Println("tree")
		tree.PrintLevelOrder()
		break
	}
}
