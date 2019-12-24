package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

// FILE is used for config file
const FILE = "config.json"

// Timer is json structure for timer configs
type Timer struct {
	Name string `json:"name"`
	Work int    `json:"work"`
	Rest int    `json:"rest"`
	Sets int    `json:"sets"`
}

// check is a basic error checker
func check(e error) {
	if e != nil {
		fmt.Println("an error occured... panic!!!")
		panic(e)
	}
}

func (t Timer) toString() string {
	bytes, err := json.Marshal(t)
	check(err)
	return string(bytes)
}

// getTimers reads configs from config.json
func getTimers() []Timer {

	fmt.Println("getting saved configs")
	timers := make([]Timer, 10)
	file, err := ioutil.ReadFile(FILE)
	check(err)
	json.Unmarshal(file, &timers)
	return timers

}

// writeJSON writes to config file
func writeJSON(w []Timer) {
	// do a write
	file, err := json.MarshalIndent(w, "", " ")
	check(err)
	_ = ioutil.WriteFile("end.json", file, 0644)
}

// checkConfig looks for config file and creates if missing
func checkConfig() {
	// checking for config file
	fmt.Println("checking for config file")
	if _, err := os.Stat(FILE); os.IsNotExist(err) {
		fmt.Println("file does not exist, creating file", FILE)
		_, err := os.Create(FILE)
		check(err)
	} else {
		fmt.Println("file exists")
		fmt.Println("")
	}
}

// deleteTimer removes timer based on name
func deleteTimer(name string, timers []Timer) {
	// delete timers - need to allow this to use input to search - currently searching for "timer22"
	for i, t := range timers {
		i++
		if t.Name == "timer23" {
			fmt.Printf("timer23 found - removing timer from position:%d", i)

			timers = append(timers[:i-1], timers[i:]...)

		}
	}

}

// saveTimer adds timer to in memory list if less than 10 exist
func saveTimer(t []Timer) {

	// save timer
	// temp values - will need ot be able to accept values
	fmt.Println("\nadding new config to slice")
	n := "hitt23"
	x := 1
	y := 2
	z := 3
	newConfig := Timer{
		Name: n,
		Work: x,
		Rest: y,
		Sets: z,
	}
	// slice is set to length 10 to save memory - check that slice is less than 10 before adding
	if len(t) < 10 {
		fmt.Println()
		// check for unique name
		for _, tt := range t {
			if tt.Name == n {
				fmt.Println("there is already a timer named: ", n)
				fmt.Println("please pick a new name")

			} else {
				t = append(t, newConfig)

			}
		}

	} else {
		fmt.Println("you already have 10 saved timers, you need to delete one before saving")

	}

}

// runTimer runs desired timer
func runTimer(w, r, s int) {

	fmt.Printf("Number of sets = %d\n", s)
	fmt.Printf("Working seconds = %d\n", w)
	fmt.Printf("Resting seconds = %d\n", r)

	fmt.Println("press any key to start")
	_, err := fmt.Scanln()
	check(err)

	fmt.Printf("\n*********Starting timer*********\n")
	for x := 5; x > 0; x-- {
		fmt.Printf("%d\n", x)
		time.Sleep(time.Second)
	}

	for s > 0 {
		fmt.Printf("*****starting set %d*****", s)

		fmt.Printf("\n***Working***\n")
		w1 := w
		for w1 > 0 {
			fmt.Printf("%d\n", w1)
			time.Sleep(time.Second)
			w1--
		}
		r1 := r
		fmt.Printf("\n***Resting***\n")
		for r1 > 0 {
			fmt.Printf("%d\n", r1)
			time.Sleep(time.Second)
			r1--
		}
		s--
	}
	fmt.Printf("\n\n*********DONE!!!*********\n\n")
}

func main() {

	// look for config file + make if not found
	checkConfig()

	// loads timers from config file
	timers := getTimers()
	fmt.Printf("there are currently %d saved timers\n", len(timers))
	for _, ti := range timers {
		// can use ti.Names to compare names etc
		fmt.Println(ti.toString())
	}

	// if click save
	//saveTimer(timers)

	// last thing to do is write config back to file
	//writeJSON(timers)

	runTimer()
}
