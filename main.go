package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

// FILE is used for config file
const FILE = "config.json"

// EOL is used for end of line in bufio string reader
const EOL = '\r'

// Timer is json structure for timer configs
type Timer struct {
	Name string `json:"name"`
	Work int64  `json:"work"`
	Rest int64  `json:"rest"`
	Sets int64  `json:"sets"`
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

	current := Timer{
		Name: "znxxzn7xx",
		Work: 0,
		Rest: 0,
		Sets: 0,
	}

	fmt.Println("getting saved configs")
	timers := make([]Timer, 9)
	file, err := ioutil.ReadFile(FILE)
	check(err)
	json.Unmarshal(file, &timers)
	if !bytes.ContainsAny(file, "Name") {
		fmt.Println("there are no saved configs")
		menu(current)
	}
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

// deleteTimer removes timer based on name or all
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

// selectTimer creates console reader to allow user to select a timer from the list
func selectTimer(x int, y int) rune {
	// if sent from deleteTimers y == 1
	current := Timer{
		Name: "znxxzn7xx",
		Work: 0,
		Rest: 0,
		Sets: 0,
	}
	reader := bufio.NewReader(os.Stdin)
	input, _, err := reader.ReadRune()
	check(err)

	i := input - 48

	if input == 'm' {
		menu(current)
	}

	if input == 'a' && y == 1 {
		return input
	}

	// validate inpiut is between 1 and x
	if input < '1' || i > rune(x) {
		fmt.Printf("\nplease input a number between 1 and %d\n", x)
		selectTimer(x, y)
	}

	return input

}

// loadTimers is call from menu to read config file into memory and display any timers
func loadTimer() {
	// loads timers from config file
	timers := getTimers()
	x := 0
	fmt.Printf("there are currently %d saved timers\n", len(timers))
	for _, ti := range timers {
		x++
		// can use ti.Names to compare names etc
		//fmt.Println(ti.toString())
		fmt.Printf("%d) %s - ", x, ti.Name)
		fmt.Printf("work: %d ", ti.Work)
		fmt.Printf("	rest: %d ", ti.Rest)
		fmt.Printf("	sets: %d\n", ti.Sets)
	}

	fmt.Printf("\nselect config to load from 1 - %d or press m to return to menu\n", x)
	t := selectTimer(x, 2)
	t = t -49 // -49 to compensate for zero index
	fmt.Printf("you have selected timer: %s\n", timers[t].Name) 
	runTimer(timers[t])
}

// saveTimer adds timer to in memory list if less than 10 exist
func saveTimer(t []Timer, current Timer) {

	// save timer
	// temp values - will need ot be able to accept values
	fmt.Println("\nadding new config to slice")
	n := "newConfig"

	var w, r, s int64
	w = 1
	r = 1
	s = 1

	newConfig := Timer{
		Name: n,
		Work: w,
		Rest: r,
		Sets: s,
	}

	// expecting time to be passed

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

func validate(x string) int64 {

	x = strings.TrimRight(x, "\r\n")
	fmt.Println("validating -", x)
	z, err := strconv.ParseInt(x, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print("checking ", z, " is within accepted rance of 1 - 180\n")

	if z < 1 || z >= 180 {

		fmt.Println("cannot be less than 1 or greater than 180!")
		manualTimer()

	}
	println("valid")
	return z
}

// manualTimer allows user to add config manually
func manualTimer() {
	readerS := bufio.NewReader(os.Stdin)
	readerW := bufio.NewReader(os.Stdin)
	readerR := bufio.NewReader(os.Stdin)
	fmt.Println("\nplease enter number of sets and hit enter: ")
	s, err := readerS.ReadString(EOL)
	check(err)
	sets := validate(s)
	fmt.Println("\nplease enter number of seconds for working and hit enter: ")
	w, err := readerW.ReadString(EOL)
	check(err)
	working := validate(w)
	fmt.Println("\nplease enter number of seconds for resting and hit enter: ")
	r, err := readerR.ReadString(EOL)
	check(err)
	rest := validate(r)
	current := Timer{
		Name: "active",
		Work: working,
		Rest: rest,
		Sets: sets,
	}
	runTimer(current)

}

// runTimer runs desired timer
func runTimer(current Timer) {

	s := current.Sets
	w := current.Work
	r := current.Rest

	fmt.Printf("Number of sets = %d\n", s)
	fmt.Printf("Working seconds = %d\n", w)
	fmt.Printf("Resting seconds = %d\n", r)

	fmt.Println("press enter key to start")
	_, err := fmt.Scanln()
	check(err)

	fmt.Printf("\n*********Starting timer*********\n")
	for x := 5; x > 0; x-- {
		fmt.Printf("%d\n", x)
		time.Sleep(time.Second)
	}

	// sets loop
	for s > 0 {
		fmt.Printf("*****starting set %d*****", s)

		fmt.Printf("\n***Working***\n")
		w1 := w
		// working loop
		for w1 > 0 {
			fmt.Printf("%d\n", w1)
			time.Sleep(time.Second)
			w1--
		}
		r1 := r
		fmt.Printf("\n***Resting***\n")
		// resting loop
		for r1 > 0 {
			fmt.Printf("%d\n", r1)
			time.Sleep(time.Second)
			r1--
		}
		s--
	}
	fmt.Printf("\n\n*********DONE!!!*********\n\n")
	menu(current)
}

// menu is main menu for command line
func menu(current Timer) {
	// ask to load or manually set timers
	fmt.Print(`
	****************
	***** MENU *****
	****************

	* press 1 to load from config file
	* press 2 to input settings 
	* press 3 to save to config file 
	* press 4 to delete a saved config 
	* press q to quit`)

	reader := bufio.NewReader(os.Stdin)
	input, _, err := reader.ReadRune()
	check(err)

	switch input {

	case '1':
		loadTimer()

	case '2':
		manualTimer()

	case '3':
		//		saveTimer()

	case '4':
		//		deleteTimer()

	case 'q':
		fmt.Println("exiting...")
		os.Exit(1)

	default:
		fmt.Println("you seem to have missed 1, 2, 3, 4 or q... please try again")
		menu(current)
		return
	}

}

func main() {
// initialising a temp timer as to allow for passing timer to menu 
tempConfig := Timer{
	Name: "znxxzn7xx",
	Work: 0,
	Rest: 0,
	Sets: 0,
}

	// look for config file + make if not found
	checkConfig()

	menu(tempConfig)

	//runTimer(10, 3, 1)

	// if click save
	//saveTimer(timers)

	// last thing to do is write config back to file
	//writeJSON(timers)
}
