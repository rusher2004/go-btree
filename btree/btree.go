package btree

import (
	"errors"
	"fmt"
)

// Data metadate about a file
type Data struct {
	FileName string
	FilePath string
}

type node struct {
	Data  Data
	left  *node
	right *node
}

type Tree struct {
	root *node
}

func (n *node) find(s string) (Data, error) {
	fmt.Println(s, " == ", n.Data.FileName)
	if n.Data.FileName == s {
		return n.Data, nil
	} else if n.Data.FileName > s && n.left != nil {
		n.left.find(s)
	} else if n.Data.FileName < s && n.right != nil {
		n.right.find(s)
	}

	return Data{}, errors.New("not found")
}

func (n *node) insert(d Data) {
	if n == nil {
		return
	} else if n.Data.FileName >= d.FileName {
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

func (n *node) print() {
	fmt.Printf("%+v\n", n.Data)

	if n.left != nil {
		n.left.print()
	}

	if n.right != nil {
		n.right.print()
	}
}

func (t *Tree) Find(s string) (Data, error) {
	if t.root == nil {
		return Data{}, errors.New("tree is empty")
	}

	found, err := t.root.find(s)

	// if t.root.Data.FileName == s {
	// 	return t.root.Data, nil
	// } else if t.root.Data.FileName > s {
	// 	return t.root.left.find(s)
	// } else if t.root.right.Data.FileName < s {
	// 	return t.root.right.find(s)
	// }

	return found, fmt.Errorf("%s not in tree", err.Error())
}

func (t *Tree) Insert(d Data) *Tree {
	if t.root == nil {
		t.root = &node{Data: d, left: nil, right: nil}
	} else {
		t.root.insert(d)
	}

	return t
}

func (t *Tree) Print() {
	if t.root == nil {
		fmt.Println("nothing here.")
		return
	}

	t.root.print()
}
