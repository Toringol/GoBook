/*
* Напишите метод String для типа *tree из gopl.io/ch4/treesort
* (раздел 4.4), который показывает последовательность значений в дереве.
 */

package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type tree struct {
	value       int
	left, right *tree
}

// Sort sorts values in place.
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func (tr *tree) String() string {
	treeInOrder := make([]int, 0)
	treeInOrder = appendValues(treeInOrder, tr)

	if len(treeInOrder) == 0 {
		return "[]"
	}

	resultString := "[ "
	for _, value := range treeInOrder {
		resultString += strconv.Itoa(value) + " "
	}
	resultString += "]"

	return resultString
}

func main() {
	data := make([]int, 50)
	for i := range data {
		data[i] = rand.Int() % 50
	}

	// Sort(data)

	fmt.Println(data)
}
