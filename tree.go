package main

import (
	"errors"
	"fmt"
)

// Data metadata about a file
type Data struct {
	FileName string
	FilePath string
}

// node holds the data values and points to child nodes
type node struct {
	Data  Data
	left  *node
	right *node
}

// Tree just holds the pointer to the root node of the tree
type Tree struct {
	root *node
}

// find scans the tree for Data.FileName in nodes, from left to right
func (n *node) find(s string) (Data, error) {
	if n.Data.FileName == s {
		return n.Data, nil
	} else if n.Data.FileName > s && n.left != nil {
		return n.left.find(s)
	} else if n.Data.FileName < s && n.right != nil {
		return n.right.find(s)
	}

	return Data{}, errors.New("not found")
}

// insert d into the tree, from left to right
func (n *node) insert(d Data) {
	if n.Data.FileName >= d.FileName {
		if n.left == nil {
			n.left = &node{Data: d, left: nil, right: nil}
		} else {
			n.left.insert(d)
		}
	} else {
		if n.right == nil {
			n.right = &node{Data: d, left: nil, right: nil}
		} else {
			n.right.insert(d)
		}
	}
}

// print the value in this node, and any child nodes, from left to right
func (n *node) print() {
	fmt.Printf("%+v\n", n.Data)

	if n.left != nil {
		n.left.print()
	}

	if n.right != nil {
		n.right.print()
	}
}

// Find a value s in the tree
func (t *Tree) Find(s string) (Data, error) {
	if t.root == nil {
		return Data{}, errors.New("tree is empty")
	}

	found, err := t.root.find(s)

	return found, err
}

// Insert d into the tree as root if t is empty, otherwise in the tree as necessary
func (t *Tree) Insert(d Data) {
	if t.root == nil {
		t.root = &node{Data: d, left: nil, right: nil}
	} else {
		t.root.insert(d)
	}
}

// Print the whole tree, scanning from left to right
func (t *Tree) Print() {
	if t.root == nil {
		fmt.Println("nothing here.")
		return
	}

	t.root.print()
}
