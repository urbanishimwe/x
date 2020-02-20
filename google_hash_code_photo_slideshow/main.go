package main

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
	"fmt"
)

var input Input
var slides []Slide
var marks int

// Algorithm you're looking for this
func Algorithm(file string) {
	ReadInputIntoVariable(file)
	for index, value := range input.Photo {
		slide := Slide{
			point: value.point,
			tags:  value.tags,
			pic:   append([]int{}, index),
		}
		if value.point == 'H' {
			slides = append(slides, slide)
		} else {
			GroupVerticalPhoto(slide, index)
		}
	}
	slides = RemoveSlideWithSingleVerticalPic()
	// Putting slides in an interesting order
	for count := 1; count < len(slides); count++ {
		to := count + 499 // Minimize slideshow's complexity by making group of 500 slides
		winner := RateMostInterestingSlide(slides[count -1], count, to)
		temp := slides[count]
		slides[count] = slides[winner]
		slides[winner] = temp
		fmt.Println("winner", count)
	}
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
	fmt.Println(marks)
	Error(WriteAnswerToOuputFile(file + ".out"))
}
