package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"tree/btree"
)

const (
	charSet    = "abcdefghijklmnopqrstuvwxyz"
	fileCount  = 100
	fileLength = 10
	findLength = 10
)

var knownFiles []btree.Data

func init() {
	knownFiles = getFiles(10)
}

func main() {
	rand.Seed(time.Now().Unix())

	files := getFiles(fileLength)

	tree := buildTree(files)

	tree.Print()

	findFiles(tree)

}

// buildTree will take in f and return a pointer to the first node
// of a tree, sorted by data.fileName
func buildTree(d []btree.Data) *btree.Tree {
	t := &btree.Tree{}

	for _, v := range d {
		t.Insert(v)
	}

	return t
}

// getFiles creates the slice of file metadata that will be inserted into the tree as nodes
func getFiles(l int) []btree.Data {
	res := []btree.Data{}

	for i := 0; i < l; i++ {
		res = append(res, fileName())
	}

	for _, d := range knownFiles {
		res = append(res, d)
	}

	return res
}

// fileName gives you a random filename ending in `.csv`
func fileName() btree.Data {
	var output strings.Builder

	for i := 0; i < fileLength; i++ {
		char := charSet[rand.Intn(len(charSet))]
		output.WriteString(string(char))
	}

	return btree.Data{
		FileName: output.String() + ".csv",
		FilePath: "path/to/this/file",
	}
}

func findFiles(t *btree.Tree) {
	for i := 0; i < findLength; i++ {
		tempName := fileName().FileName
		found, err := t.Find(tempName)

		if err != nil {
			fmt.Printf("%s not found: %s\n", tempName, err.Error())
		} else {
			fmt.Printf("found %s at %s\n", found.FileName, found.FilePath)
		}
	}

	for _, d := range knownFiles {
		found, err := t.Find(d.FileName)

		if err != nil {
			fmt.Printf("%s not found: %s\n", d.FileName, err.Error())
		} else {
			fmt.Printf("found %s at %s\n", found.FileName, found.FilePath)
		}
	}
}
