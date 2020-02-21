package main

import (
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// Atoi convert string to integer
func Atoi(s string) int {
	number, err := strconv.Atoi(s)
	Error(err)
	return number
}

// Error handlers
func Error(e error) {
	if e != nil {
		panic(e)
	}
}

// ReadInputIntoVariable read variables from file
func ReadInputIntoVariable(file string) {
	var (
		err   error
		data  []byte
		lines []string
		line  []string
	)
	data, err = ioutil.ReadFile(file)
	Error(err)
	lines = strings.Split(string(data), "\n")
	line = strings.Split(lines[0], " ")
	input.B = Atoi(line[0])
	input.L = Atoi(line[1])
	input.D = Atoi(line[2])
	// Parse books scores
	line = strings.Split(lines[1], " ")
	input.Scores = make([]int, input.B)
	for count := 0; count < input.B; count++ {
		input.Scores[count] = Atoi(line[count])
	}
	input.Libraries = make([]Library, input.L)
	index := 0
	for count := 2; count <= (input.L * 2); count += 2 {
		line = strings.Split(lines[count], " ")
		input.Libraries[index].N = Atoi(line[0])
		input.Libraries[index].T = Atoi(line[1])
		input.Libraries[index].M = Atoi(line[2])
		input.Libraries[index].index = index
		line = strings.Split(lines[count+1], " ")
		input.Libraries[index].Books = make([]int, input.Libraries[index].N)
		for i := 0; i < input.Libraries[index].N; i++ {
			input.Libraries[index].Books[i] = Atoi(line[i])
		}
		index++
	}
}

// SortLibrary by their scores
func SortLibrary() {
	swapped := true
	prev := 0
	current := 0
	for swapped {
		swapped = false
		for i := 1; i < input.L; i++ {
			prev = LibraryScore(input.Libraries[i-1], input.Libraries[i-1].T)
			current = LibraryScore(input.Libraries[i], input.Libraries[i].T)
			if prev < current {
				temp := input.Libraries[i-1]
				input.Libraries[i-1] = input.Libraries[i]
				input.Libraries[i] = temp
				swapped = true
			}
		}
	}
}

// SortLibraryBook sort the books of a library
func SortLibraryBook(library int) {
	length := input.Libraries[library].N
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < length; i++ {
			prev := input.Scores[input.Libraries[library].Books[i-1]]
			current := input.Scores[input.Libraries[library].Books[i]]
			if prev < current {
				temp := input.Libraries[library].Books[i-1]
				input.Libraries[library].Books[i-1] = input.Libraries[library].Books[i]
				input.Libraries[library].Books[i] = temp
				swapped = true
			}
		}
	}
}

// LibraryScore rate library by its scores
func LibraryScore(library Library, days int) (score int) {
	score = 0
	count := 0
	for count < library.N && days < input.D {
		for i := 0; i < library.M && count < library.N; i++ {
			score += input.Scores[library.Books[count]]
			count++
		}
		days++
	}
	return score
}

// RemoveDuplicateBook remove duplicate book wherever necessary
func RemoveDuplicateBook(book []int, inLib []int) (answer []int) {
	for _, i := range inLib {
		found := false
		for _, j := range book {
			if j == i {
				found = true
			}
		}
		if !found {
			answer = append(answer, i)
		}
	}
	return answer
}

// JoinBooks merge two tags
func JoinBooks(a []int, b []int) []int {
	return append(a, b...)
}

//RemoveEmptyLibrary remove library with empty books
func RemoveEmptyLibrary() (answer []Library) {
	for _, library := range input.Libraries {
		if library.N >= 1 {
			answer = append(answer, library)
		}
	}
	return answer
}

// WriteAnswerToOuputFile flush answer to ouput file
func WriteAnswerToOuputFile(filePath string) (err error) {
	var (
		file *os.File
		data string
	)
	file, err = os.Create(filePath)
	Error(err)
	defer file.Close()
	data = strconv.Itoa(len(input.Libraries)) + "\n"
	for _, value := range input.Libraries {
		data += strconv.Itoa(value.index) + " " + strconv.Itoa(len(value.Books)) + "\n"
		for count := 0; count < value.N; count++ {
			data += (strconv.Itoa(value.Books[count]) + " ")
		}
		data = strings.Trim(data, " ")
		data += "\n"
	}
	_, err = io.WriteString(file, data)
	Error(err)
	return file.Sync()
} 