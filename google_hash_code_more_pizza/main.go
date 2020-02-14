package google_hash_code_more_pizza

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"sort"
	"strconv"
	"strings"
)

/**
	Error
	Handle Errors
**/
func Error(e error) {
	if e != nil {
		panic(e)
	}
}

/**
 arraySum
 get sum of S's element with index in P,
**/
func arraySum(S []int, P []int) int {
	sum := 0
	for _, value := range P {
		sum = sum + S[value]
	}
	return sum
}

/**
	ReadFileVariable
	Read input from input.in
**/
func ReadFileVariable(file string) (M int, N int, S []int) {
	var (
		err   error
		data  []byte
		lines []string
		MandN []string
	)
	Error(err)
	data, err = ioutil.ReadFile(file)
	Error(err)
	lines = strings.Split(string(data), "\n")[0:2] // read all lines
	MandN = strings.Split(lines[0], " ")           // split first line of M and N
	M, err = strconv.Atoi(MandN[0])                // converting string to integer
	Error(err)
	N, err = strconv.Atoi(MandN[1])
	Error(err)
	S = make([]int, N)
	for index, value := range strings.Split(lines[1], " ") {
		S[index], err = strconv.Atoi(value)
		Error(err)
	}
	return M, N, S
}

/**
	Algorithm
	this where the real stuff happen
**/
func Algorithm(file string) (L int, P []int) { // L is number of Pizzas, P stores the index from S
	M, N, S := ReadFileVariable(file) // M is max number of slices, N len of Pizza, S is Pizzas
	sum := 0
	for count := 0; count < N; count++ {
		if (sum + S[count]) > M {
			for index, value := range P {
				if (sum + S[count] - S[value]) <= M {
					P[index] = count
					break
				}
			}
		} else {
			P = append(P, count)
		}
		sum = arraySum(S, P)
	}
	rate := ((float32(sum) / float32(M)) * 100)
	fmt.Printf("sum of ordered slices = %d, expected = %d, rate = %.2f\n", sum, M, rate)
	L = len(P)
	return L, P
}

/**
	WriteOutputVariable
	write the answer to the ouput file
**/
func WriteOutputVariable(filePath string, L int, P []int) (err error) {
	var (
		file  *os.File
		data  string
		array []string
	)
	file, err = os.Create(filePath)
	Error(err)
	defer file.Close()
	data = strconv.Itoa(L) + "\n"
	sort.Ints(P)
	for _, value := range P {
		array = append(array, strconv.Itoa(value))
	}
	data += (strings.Join(array, " ") + "\n")
	_, err = io.WriteString(file, data)
	Error(err)
	return file.Sync()
}

/**
  the start
**/
func main() {
	var (
		dir  string
		err  error
		file string
		L    int
		P    []int
	)
	_, file, _, _ = runtime.Caller(0)
	dir, err = filepath.Abs(filepath.Dir(file))
	Error(err)
	file = path.Join(dir, os.Args[1])
	L, P = Algorithm(file) // simply the final answer is here
	file = path.Join(dir, os.Args[1]+".out")
	Error(WriteOutputVariable(file, L, P))
	fmt.Println("done")
}
