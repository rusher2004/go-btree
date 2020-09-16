package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	charSet    = "abcdefghijklmnopqrstuvwxyz"
	fileCount  = 1000000
	fileLength = 10
	findLength = 10000
)

var knownFiles []Data

func init() {
	knownFiles = getFiles(1000)
}

func main() {
	rand.Seed(time.Now().Unix())

	// get a slice of mock random metadata
	files := getFiles(fileCount)

	// take the slice and put it into a binary search tree
	tree := buildTree(files)

	// run a check to see if we can actually find files
	findFiles(tree)
}

// buildTree will take in f and return a pointer to the first node
// of a tree, sorted by Data.fileName
func buildTree(d []Data) *Tree {
	t := &Tree{}

	for _, v := range d {
		t.Insert(v)
	}

	return t
}

// getFiles creates the slice of file metadata that will be inserted into the tree as nodes
func getFiles(l int) []Data {
	res := []Data{}

	for i := 0; i < l; i++ {
		res = append(res, fileName())
	}

	for _, d := range knownFiles {
		res = append(res, d)
	}

	return res
}

// fileName gives you a random filename ending in `.csv`
func fileName() Data {
	var output strings.Builder

	for i := 0; i < fileLength; i++ {
		char := charSet[rand.Intn(len(charSet))]
		output.WriteString(string(char))
	}

	return Data{
		FileName: output.String() + ".csv",
		FilePath: "path/to/this/file",
	}
}

// findFiles will find our known files to validate our Tree.Find()
func findFiles(t *Tree) {
	// find the known files
	for _, d := range knownFiles {
		found, err := t.Find(d.FileName)

		if err != nil {
			fmt.Printf("%s not found: %s\n", d.FileName, err.Error())
		} else {
			fmt.Printf("found %s at %s\n", found.FileName, found.FilePath)
		}
	}
}
