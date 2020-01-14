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
const FILE = "./config.json"

// EOL is used for end of line in bufio string reader
// const EOL = '\r' // win
const EOL = '\n' // nix

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

// writeJSON writes to config file
func writeJSON(w []Timer) {
	// do a write
	file, err := json.MarshalIndent(w, "", " ")
	check(err)
	_ = ioutil.WriteFile(FILE, file, 0644)
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
func deleteTimer(current Timer, timers []Timer) {

	fmt.Println("press 1 to select config or press 9 to delete all")
	reader := bufio.NewReader(os.Stdin)
	input, _, err := reader.ReadRune()
	check(err)
	if input == '9' {
		fmt.Println("WARNING - THIS WILL DELETE ALL SAVED CONFIGS")
		fmt.Println("press y to comtinue")
		reader2 := bufio.NewReader(os.Stdin)
		input2, _, err := reader2.ReadRune()
		check(err)
		if input2 == 'y' {
			// delete all configs, return to main
			empty := make([]Timer, 9)
			writeJSON(empty)
			main()
		}
	}

	selected := selectTimer(loadTimer())
	sName := selected.Name
	deleted := 0
	// delete timers
	for i, t := range timers {
		i++
		if t.Name == sName {
			fmt.Printf("%s found - removing timer from position:%d", sName, i)

			timers = append(timers[:i-1], timers[i:]...)
			deleted++
		}
	}

	if deleted == 0 {

		fmt.Println("%s not found", sName)
	}

	fmt.Println("press 1 to delete more or any other key to return to menu")
	reader3 := bufio.NewReader(os.Stdin)
	input3, _, err := reader3.ReadRune()
	check(err)
	if input3 == '1' {
		deleteTimer(current, timers)
	}
	// save updated slice of timers
	fmt.Println("saving current configs to file and returning to menu...")
	writeJSON(timers)
	menu(current)

}

// selectTimer creates console reader to allow user to select a timer from the list
func selectTimer(timers []Timer) Timer {

	current := Timer{
		Name: "znxxzn7xx",
		Work: 0,
		Rest: 0,
		Sets: 0,
	}

	x := 0
	fmt.Printf("there are currently %d saved timers\n", len(timers))
	for _, ti := range timers {
		x++
		fmt.Printf("%d) %s - ", x, ti.Name)
		fmt.Printf("work: %d ", ti.Work)
		fmt.Printf("	rest: %d ", ti.Rest)
		fmt.Printf("	sets: %d\n", ti.Sets)
	}

	fmt.Printf("\nselect config to load from 1 - %d or press m to return to menu\n", x)

	reader := bufio.NewReader(os.Stdin)
	input, _, err := reader.ReadRune()
	check(err)

	if input == 'm' {
		menu(current)
	}

	i := input - 48
	fmt.Print(i)
	fmt.Printf("\nx = %d\n", x)
	fmt.Printf("i = %d\n", i)
	// validate input is between 1 and x
	if input < '1' || i > rune(x) {
		fmt.Printf("\nplease input a number between 1 and %d\n", x)
		selectTimer(timers)
	}
	t := timers[i-1] // TODO: currently doing something strange if you input an out of range number then select an actual config
	fmt.Printf("you have selected timer: %s\n", t.Name)

	return t

}

// loadTimers is call from menu to read config file into memory and display any timers
func loadTimer() []Timer {
	// loads timers from config file
	current := Timer{
		Name: "znxxzn7xx",
		Work: 0,
		Rest: 0,
		Sets: 0,
	}

	fmt.Println("getting saved configs")
	t := make([]Timer, 9)
	file, err := ioutil.ReadFile(FILE)
	check(err)
	json.Unmarshal(file, &t) // TODO: something here isnt populating the the file contentss
	if !bytes.ContainsAny(file, "Name") {
		fmt.Println("there are no saved configs")
		menu(current)
	}
	fmt.Println("in laod...")
	x := 0
	for _, ti := range t {
		x++
		fmt.Printf("%d) %s - ", x, ti.Name)
		fmt.Printf("work: %d ", ti.Work)
		fmt.Printf("	rest: %d ", ti.Rest)
		fmt.Printf("	sets: %d\n", ti.Sets)
	}
	return t

}

// saveTimer adds timer to in memory list if less than 9 exist
func saveTimer(t []Timer, current Timer) {

	// go back to menu if full
	if len(t) > 8 {
		fmt.Println("there is already 9 configs saved, please delete some to make room...")
		menu(current)
	}

	fmt.Println("press 1 to add current timer, 2 to input new timer or m to return to menu")
	r := bufio.NewReader(os.Stdin)
	i, _, err := r.ReadRune()

	switch i {
	case 'm':

		menu(current)

	case '2':

		current = userInput()

	default:
		fmt.Println("unkonw input...")
		saveTimer(t, current)

	}
getName:
	fmt.Println("please enter a name for the timer and press enter")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString(EOL)
	check(err)
	n := strings.TrimRight(input, "\r\n")

	// check for unique name
	for _, tt := range t {
		fmt.Println("checking for name match - ", tt.Name)
		if tt.Name == n {
			fmt.Println("there is already a timer named: ", n)
			fmt.Println("please pick a new name")
			goto getName
		}
	}

	current.Name = n
	t = append(t, current)
	writeJSON(t)

	menu(current)
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
		userInput()

	}
	println("valid")
	return z
}

// userInput collects user input config from command line
func userInput() Timer {
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
		Name: "activezxzxzxzxzxzxzx",
		Work: working,
		Rest: rest,
		Sets: sets,
	}
	return current
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
	fmt.Println("well done, workout complete...")
	fmt.Println("would you like to save the current config?")
	fmt.Println("press 1 to save or any other key to return to menu")
	reader := bufio.NewReader(os.Stdin)
	input, _, err := reader.ReadRune()
	check(err)
	if input == '1' {
		saveTimer(loadTimer(), current)
	}
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

	fmt.Println(" ") // adding spacing for menu

	reader := bufio.NewReader(os.Stdin)
	input, _, err := reader.ReadRune()
	check(err)

	switch input {

	case '1':

		selected := selectTimer(loadTimer())
		runTimer(selected)

	case '2':

		runTimer(userInput())

	case '3':

		saveTimer(loadTimer(), current)

	case '4':

		deleteTimer(current, loadTimer())

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

}
