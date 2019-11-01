// Interval timer
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/couchbase/gocb"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
)

// Config is the top level json structure for storing in database.
type Config struct {
	ID     string `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Config Data   `json:"config,omitempty"`
}

// Data is the second level json structure.
type Data struct {
	Work int `json:"work,omitempty"`
	Rest int `json:"rest,omitempty"`
	Sets int `json:"sets,omitempty"`
}

var bucket *gocb.Bucket
var bucketName string

// ListEndpoint returns all saved configs from database.
func ListEndpoint(w http.ResponseWriter, req *http.Request) {
	var configs []Config
	query := gocb.NewN1qlQuery("SELECT `" + bucketName + "`.* FROM `" + bucketName + "`")
	query.Consistency(gocb.RequestPlus)
	rows, _ := bucket.ExecuteN1qlQuery(query, nil)
	var row Config
	for rows.Next(&row) {
		configs = append(configs, row)
		row = Config{}
	}
	if configs == nil {
		configs = make([]Config, 0)
	}
	json.NewEncoder(w).Encode(configs)
}

// SearchEndpoint returns searched config from database.
func SearchEndpoint(w http.ResponseWriter, req *http.Request) {
	var configs []Config
	params := mux.Vars(req)
	var n1qlParams []interface{}
	n1qlParams = append(n1qlParams, strings.ToLower(params["title"]))
	query := gocb.NewN1qlQuery("SELECT `" + bucketName + "`.* FROM `" + bucketName + "` WHERE LOWER(name) LIKE '%' || $1 || '%'")
	query.Consistency(gocb.RequestPlus)
	rows, _ := bucket.ExecuteN1qlQuery(query, n1qlParams)
	var row Config
	for rows.Next(&row) {
		configs = append(configs, row)
		row = Config{}
	}
	if configs == nil {
		configs = make([]Config, 0)
	}
	json.NewEncoder(w).Encode(configs)
}

// CreateEndpoint saves configs to database.
func CreateEndpoint(w http.ResponseWriter, req *http.Request) {
	var config Config
	_ = json.NewDecoder(req.Body).Decode(&config)
	bucket.Insert(uuid.NewV4().String(), config, 0)
	json.NewEncoder(w).Encode(config)
}

func main() {
	fmt.Println("Starting server at http://localhost:12345...")
	cluster, _ := gocb.Connect("couchbase://localhost")
	bucketName = "interval"
	bucket, _ = cluster.OpenBucket(bucketName, "")
	router := mux.NewRouter()
	router.HandleFunc("/configs", ListEndpoint).Methods("GET")
	router.HandleFunc("/configs", CreateEndpoint).Methods("POST")
	router.HandleFunc("/search/{title}", SearchEndpoint).Methods("GET")
	log.Fatal(http.ListenAndServe(":12345", handlers.CORS(handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD"}), handlers.AllowedOrigins([]string{"*"}))(router)))
}
