package main

import (
	"io"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

// InputVariable this the structure of the input as defined in project description
type InputVariable struct {
	R, C, F, N, B, T int
	Rides            []Ride
	Fleet            [][]int //rides mapped to their vehicle
}

// Ride should have the following structure as specified in challenge description
type Ride struct {
	start, end             []int
	eStart, lFinish, index int
}

// Error handler
func Error(e error) {
	if e != nil {
		panic(e)
	}
}

// ReadFileVariable Read input from input.in
func ReadFileVariable(file string) (input InputVariable) {
	var (
		data      []byte
		err       error
		lines     []string
		stringRow []string
		intRow    [6]int
	)
	Error(err)
	data, err = ioutil.ReadFile(file)
	Error(err)
	lines = strings.Split(string(data), "\n")
	stringRow = strings.Split(lines[0], " ")
	input.R = atoi(stringRow[0])
	input.C = atoi(stringRow[1])
	input.F = atoi(stringRow[2])
	input.N = atoi(stringRow[3])
	input.B = atoi(stringRow[4])
	input.T = atoi(stringRow[5])
	input.Rides = make([]Ride, input.N)
	for i := 0; i < input.N; i++ {
		stringRow = strings.Split(lines[i+1], " ")
		for j := 0; j < 6; j++ {
			intRow[j] = atoi(stringRow[j])
		}
		input.Rides[i].start = []int{intRow[0], intRow[1]}
		input.Rides[i].end = []int{intRow[2], intRow[3]}
		input.Rides[i].eStart = intRow[4]
		input.Rides[i].lFinish = intRow[5]
		input.Rides[i].index = i
	}
	return input
}

// Abs do math.absolute
func abs(number int) int {
	return int(math.Abs(float64(number)))
}

// find the distance between two point
func distance(a []int, b []int) int {
	return abs(a[0]-b[0]) + abs(a[1]-b[1])
}

// atoi convert string to integer
func atoi(s string) int {
	number, err := strconv.Atoi(s)
	Error(err)
	return number
}

func remove(ride int, tempRides []Ride) (temp []Ride) {
	for _, value := range tempRides {
		if value.index != ride {
			temp = append(temp, value)
		}
	}
	return temp
}

func prevRide(fleet int) (prev Ride) {
	prev = Ride{
		start:   []int{0, 0},
		end:     []int{0, 0},
		lFinish: 0,
		eStart:  0,
	}
	for _, value := range input.Fleet[fleet] {
		prev = input.Rides[value]
	}
	return prev
}

// WriteOutputVariable the function to generate output file
func WriteOutputVariable(filePath string, fleet [][]int) (err error) {
	var (
		file *os.File
		data = ""
	)
	file, err = os.Create(filePath)
	Error(err)
	defer file.Close()
	for _, value := range fleet {
		data += strconv.Itoa(len(value))
		for _, ride := range value {
			data += (" " + strconv.Itoa(ride))
		}
		data += "\n"
	}
	_, err = io.WriteString(file, data)
	Error(err)
	return file.Sync()
}
