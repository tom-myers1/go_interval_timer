package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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

// readJSON reads configs from config.json
func getTimers() []Timer {

	fmt.Println("getting saved configs")
	timers := make([]Timer, 10)
	file, err := ioutil.ReadFile(FILE)
	check(err)
	json.Unmarshal(file, &timers)

	return timers

}

func writeJSON(w []Timer) {
	// do a write
	file, err := json.MarshalIndent(w, "", " ")
	check(err)
	_ = ioutil.WriteFile("end.json", file, 0644)
}

func main() {
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

	// loads timers from config file
	timers := getTimers()
	fmt.Printf("there are currently %d saved timers\n", len(timers))
	for _, ti := range timers {
		// can use ti.Names to compare names etc
		fmt.Println(ti.toString())
	}
	fmt.Println(" ")
	fmt.Println(timers)
	// delete timers - need to allow this to use input to search - currently searching for "timer22"
	for i, t := range timers {
		i++
		if t.Name == "timer23" {
			fmt.Printf("timer23 found - removing timer from position:%d", i)

			timers = append(timers[:i-1], timers[i:]...)

		}
	}

	// save timer
	// temp values
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
	if len(timers) < 10 {
		fmt.Println()
		// check for unique name
		for _, tt := range timers {
			if tt.Name == n {
				fmt.Println("there is already a timer named: ", n)
				fmt.Println("please pick a new name")
			} else {
				timers = append(timers, newConfig)
			}
		}

	} else {
		fmt.Println("you already have 10 saved timers, you need to delete one before saving")
	}

	// load timer - actuall think we can keep in mems and just select from timers

	// last thing to do is write config back to file
	writeJSON(timers)
}
