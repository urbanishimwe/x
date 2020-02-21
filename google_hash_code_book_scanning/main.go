package main

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
)

var input Input

// Input structure as defined in project description
type Input struct {
	B, L, D   int
	Scores    []int
	Libraries []Library
}

// Library library should have this
type Library struct {
	N, T, M, index int
	Books          []int
}

// Algorithm this what you're looking for
func Algorithm(file string) {
	ReadInputIntoVariable(file)
	for count := 0; count < input.L; count++ {
		SortLibraryBook(count)
	}
	SortLibrary()

	// Removing duplicate books by increasing scores
	books := input.Libraries[0].Books
	for count := 1; count < input.L; count++ {
		answer := RemoveDuplicateBook(books, input.Libraries[count].Books)
		books = JoinBooks(books, answer)
		input.Libraries[count].Books = answer
		input.Libraries[count].N = len(answer)
	}

	//Sorting remaining data
	for count := 0; count < input.L; count++ {
		SortLibraryBook(count)
	}
	SortLibrary()
	input.Libraries = RemoveEmptyLibrary()
}

func main() {
	var (
		dir  string
		err  error
		file string
	)
	_, file, _, _ = runtime.Caller(0)
	dir, err = filepath.Abs(filepath.Dir(file))
	Error(err)
	file = path.Join(dir, os.Args[1])
	Algorithm(file)
	Error(WriteAnswerToOuputFile(path.Join(dir, "answer.out")))
}
