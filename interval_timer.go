// interval timer

package main

import (
	"fmt"
	"time"
)

var sets int
var work int
var rest int

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
	fmt.Println("input the number of sets:")
	fmt.Scan(&sets)
	validate(sets, "sets")

}
func getWork() {
	fmt.Println("input the number of seconds per set:")
	fmt.Scan(&work)
	validate(work, "work")

}
func getRest() {
	fmt.Println("input the number of seconds rest:")
	fmt.Scan(&rest)
	validate(rest, "rest")

}

func validate(x int, y string) {
	if x <= 0 && y != "rest" {
		fmt.Println(y, "must be greater than 0")

	} else if x > 30 && y == "sets" {
		fmt.Println("are you sure you want to do", x, "sets? (y/n)")
		fmt.Fscanln(r)
		// go back to getSets
	} else if x > 500 {
		fmt.Println(x, "seems like a lot of", y, "are you sure? (y/n)")
		fmt.Fscan(r)
		// go back to getXXX
	} else {
		return true
	}
}

func main() {

	fmt.Println("### Interval Timer ###")
	// get sets, working time and resting time
	// check that they are are int and not somethig ridiculous
	getSets()
	getWork()
	getRest()

	fmt.Println("number of sets:", sets)

	return
	fmt.Println("working:", work)
	fmt.Println("rest:", rest)
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

	for s := 1; s < sets+1; s++ {
		fmt.Println("starting set :", s)
		working(work)

	}
	fmt.Println()

}
