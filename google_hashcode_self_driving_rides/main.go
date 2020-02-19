package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

var input InputVariable

func rateRide(ride Ride, f int, fleetDistance int) (canBeAdded bool, isBonus bool, totalDistance int) {
	var prev = prevRide(f)
	firstDistance := distance(prev.end, ride.start)
	if firstDistance < ride.eStart {
		firstDistance += ride.eStart
	}
	totalDistance = firstDistance + distance(ride.start, ride.end)
	canBeAdded = (fleetDistance + totalDistance) <= ride.lFinish
	isBonus = fleetDistance+distance(prev.end, ride.start) <= ride.eStart
	return canBeAdded, isBonus, totalDistance
}

func rankRide(fleet int, tempRides []Ride, fD int) (answer int, fDistance int) {
	type Sort struct {
		isBonus              bool
		totalDistance, index int
	}
	var temp []Sort
	for _, value := range tempRides {
		canBeAdded, isBonus, totalDistance := rateRide(value, fleet, fD)
		if canBeAdded {
			var obj Sort
			obj.index = value.index
			obj.isBonus = isBonus
			obj.totalDistance = totalDistance
			temp = append(temp, obj)
		}
	}
	if len(temp) == 0 {
		return -1, 0
	}
	fmt.Println("temp", temp)
	answer = 0
	fDistance = temp[answer].totalDistance
	for index, value := range temp {
		if temp[answer].isBonus {
			if value.isBonus {
				if value.totalDistance < temp[answer].totalDistance {
					answer = index
				}
			}
		} else if value.isBonus {
			answer = index
			fDistance = value.totalDistance
		} else {
			if value.totalDistance < temp[answer].totalDistance {
				answer = index
				fDistance = value.totalDistance
			}
		}
	}
	return temp[answer].index, fDistance
}

// Algorithm the real shit happens here.
func Algorithm(file string) {
	input = ReadFileVariable(file)
	input.Fleet = make([][]int, input.F)
	tempRides := input.Rides
	for count := 0; count < input.F; count++ {
		fleetDistance := 0
		answer, temp := rankRide(count, tempRides, 0)
		tempRides = remove(answer, tempRides)
		fleetDistance += temp
		for answer != -1 {
			fmt.Println("fleet distance", fleetDistance)
			fmt.Println("answer", input.Rides[answer])
			input.Fleet[count] = append(input.Fleet[count], answer)
			tempRides = remove(answer, tempRides)
			answer, temp = rankRide(count, tempRides, fleetDistance)
			fleetDistance += temp
		}
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
	Algorithm(file) // simply the final answer is here
	Error(WriteOutputVariable(file+".out", input.Fleet))
}
