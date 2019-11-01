// interval timer

package main

import (
	"fmt"
	"time"
)

var sets int
var work int
var rest int
var getSetsCounter = 0
var getWorkCounter = 0
var getRestCounter = 0

func validatorExit() {
	fmt.Println()
	fmt.Println("you have tried three times now without getting it right...exiting")
	fmt.Println()
	main()

}

func working(w int) {

	for w > 0 {
		fmt.Println("working :", w)
		time.Sleep(1 * time.Second)
		w--
	}
}

func resting(r int) {
	for r > 0 {
		fmt.Println("resting :", r)
		time.Sleep(1 * time.Second)
		r--
	}
}

func getSets() {
	getSetsCounter++
	if getSetsCounter > 3 {
		getSetsCounter = 0
		validatorExit()
	}
	fmt.Println("input the number of sets:")
	fmt.Scan(&sets)
	validate(sets, "sets")
}

func getWork() {
	getWorkCounter++
	if getWorkCounter > 3 {
		getWorkCounter = 0
		validatorExit()
	}
	fmt.Println("input the number of seconds per set:")
	fmt.Scan(&work)
	validate(work, "work")
}

func getRest() {
	getRestCounter++
	if getRestCounter > 3 {
		getRestCounter = 0
		validatorExit()
	}
	fmt.Println("input the number of seconds rest:")
	fmt.Scan(&rest)
	validate(rest, "rest")
}

func validate(x int, y string) bool {
	fmt.Println("validating", y, "value passed:", x)

	if x < 1 || x > 600 {
		fmt.Println(y, "please enter a NUMBER greater than 0 and less than 600 (you are doing it wrong!!!)")
		switch y {
		case "rest":
			getRest()
		case "sets":
			getSets()
		case "work":
			getWork()
		default:
			return false
		}

	}

	if x > 30 && y == "sets" {
		var response string
		fmt.Println("are you sure you want to do", x, "sets? (y/n)")
		fmt.Scan(&response)
		if response != "y" {
			getSets()
		}
		return false
	}

	return true

}

func main() {
	fmt.Println()
	fmt.Println("### Interval Timer ###")
	// get sets, working time and resting time
	// check that they are are int and not somethig ridiculous
	getSets()
	getWork()
	getRest()

	fmt.Println("number of sets:", sets)
	fmt.Println("working:", work)
	fmt.Println("rest:", rest)
	fmt.Println()
	fmt.Println("starting in 5")
	time.Sleep(1 * time.Second)
	fmt.Println("starting in 4")
	time.Sleep(1 * time.Second)
	fmt.Println("starting in 3")
	time.Sleep(1 * time.Second)
	fmt.Println("starting in 2")
	time.Sleep(1 * time.Second)
	fmt.Println("starting in 1")
	time.Sleep(1 * time.Second)
	fmt.Println()

	for s := 1; s < sets+1; s++ {
		fmt.Println("starting set :", s)
		fmt.Println()
		working(work)
		fmt.Println()
		resting(rest)
		fmt.Println()

	}
	fmt.Println("*** Done! ***")
	fmt.Println()

}
