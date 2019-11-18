// interval timer - using json file for saving and loading settings
package main

import (
	//"bufio"
	//"encoding/json"
	"fmt"
	"io/ioutil"
	//"log"
	//"net/http"
	//"strings"
	//"github.com/gofrs/uuid"
	//"github.com/gorilla/handlers"
	//"github.com/gorilla/mux"
)

// Config is the top level json structure for storing in database.
type Config struct {
	Name   string    `json:"name,omitempty"`
	Config TimerData `json:"config,omitempty"`
}

// TimerData is the second level json structure.
type TimerData struct {
	Work int `json:"work,omitempty"`
	Rest int `json:"rest,omitempty"`
	Sets int `json:"sets,omitempty"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readConfig(configFile) {
	conf, err := ioutil.ReadFile(configFile)
	check(err)
	fmt.Print(string(conf))
}

func writeConfig(timerConfigs, configFile) {
	err
}

func main() {
	configFile := "/config.json"
	f, err := os.Create(configFile)
	check(err)
	defer f.Close()
	c, err := readConfig(configFile)
	check(err)

	writeConfig(Config, configFile)

}
