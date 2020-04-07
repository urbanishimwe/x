package main

import (
	"io"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

// Input attributes
type Input struct {
	N     int
	Photo []Pic
}

// Pic photo attributes
type Pic struct {
	point rune
	tags  []string
}

// Slide slide structure
type Slide struct {
	point rune
	tags  []string
	pic   []int
}

// Error handlers
func Error(e error) {
	if e != nil {
		panic(e)
	}
}

// Atoi convert string to integer
func Atoi(s string) int {
	number, err := strconv.Atoi(s)
	Error(err)
	return number
}

// Abs do math.absolute
func Abs(number int) int {
	return int(math.Abs(float64(number)))
}

// ReadInputIntoVariable read variables from file
func ReadInputIntoVariable(file string) {
	var (
		err   error
		data  []byte
		lines []string
	)
	data, err = ioutil.ReadFile(file)
	Error(err)
	lines = strings.Split(string(data), "\n")
	input.N = Atoi(lines[0])
	input.Photo = make([]Pic, input.N)
	for count := 0; count < input.N; count++ {
		line := strings.Split(lines[count+1], " ")
		input.Photo[count].point = rune(line[0][0])
		input.Photo[count].tags = line[2:]
	}
}

// GroupVerticalPhoto we form new slide or we add to existing slide
func GroupVerticalPhoto(photo Slide, at int) {
	var (
		itsIndex    = -1
		min         = 0
		found       = false
		commonTags  []string
		tagsInAOnly []string
		tagsInBOnly []string
	)
	for index, value := range slides {
		if value.point == 'V' && len(value.pic) < 2 {
			if itsIndex == -1 {
				itsIndex = index
				found = true
			}
			commonTags = CommonTags(value.tags, photo.tags)
			tagsInAOnly = TagsInAOnly(value.tags, photo.tags)
			tagsInBOnly = TagsInAOnly(photo.tags, value.tags)
			minimum := Min([]int{len(commonTags), len(tagsInAOnly), len(tagsInBOnly)})
			if minimum > min {
				min = minimum
				itsIndex = index
				found = true
			}
		}
	}
	if found {
		commonTags = CommonTags(slides[itsIndex].tags, photo.tags)
		tagsInAOnly = TagsInAOnly(slides[itsIndex].tags, photo.tags)
		tagsInBOnly = TagsInAOnly(photo.tags, slides[itsIndex].tags)
		slides[itsIndex].tags = JoinTags(JoinTags(commonTags, tagsInAOnly), tagsInBOnly)
		slides[itsIndex].pic = append(slides[itsIndex].pic, at)
	} else {
		slides = append(slides, photo)
	}
}

// CommonTags between two slide or picture
func CommonTags(a []string, b []string) (answer []string) {
	for _, i := range a {
		for _, j := range b {
			if i == j {
				answer = append(answer, i)
			}
		}
	}
	return answer
}

// TagsInAOnly tags which can be found in first but not in second
func TagsInAOnly(a []string, b []string) (tags []string) {
	for _, i := range a {
		notFound := true
		for _, j := range b {
			if j == i {
				notFound = false
			}
		}
		if notFound {
			tags = append(tags, i)
		}
	}
	return tags
}

// JoinTags merge two tags
func JoinTags(a []string, b []string) []string {
	return append(a, b...)
}

// Min in an array
func Min(a []int) (min int) {
	min = a[0]
	for _, value := range a {
		if value < min {
			min = value
		}
	}
	return min
}

// RemoveSlideWithSingleVerticalPic since it is not allowed
func RemoveSlideWithSingleVerticalPic() (answer []Slide) {
	for _, value := range slides {
		if value.point == 'V' && len(value.pic) == 2 {
			answer = append(answer, value)
		} else if value.point == 'H' {
			answer = append(answer, value)
		}
	}
	return answer
}

// RateMostInterestingSlide before putting slide in the next queue we need to check how it is interesting with regarding to itsprev slide
func RateMostInterestingSlide(prev Slide, from int, to int) (index int) {
	var (
		commonTags int
		inPrevOnly int
		inNextOnly int
		max        int
	)
	for i := from; i < to && to < len(slides); i++ {
		commonTags = len(CommonTags(prev.tags, slides[i].tags))
		inPrevOnly = len(TagsInAOnly(prev.tags, slides[i].tags))
		inNextOnly = len(TagsInAOnly(slides[i].tags, prev.tags))
		minimum := Min([]int{commonTags, inPrevOnly, inNextOnly})
		if minimum >= max {
			index = i
			max = minimum
		}
	}
	marks += max
	return index
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
	data = strconv.Itoa(len(slides)) + "\n"
	for _, value := range slides {
		if value.point == 'H' {
			data += (strconv.Itoa(value.pic[0]) + "\n")
		} else {
			data += (strconv.Itoa(value.pic[0]) + " " + strconv.Itoa(value.pic[1]) + "\n")
		}
	}
	_, err = io.WriteString(file, data)
	Error(err)
	return file.Sync()
}
