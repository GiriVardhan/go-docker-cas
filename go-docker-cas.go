package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
)

var Session *gocql.Session

type Emp struct {
	Id        string `json:"empid"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}

type AllEmpsResponse struct {
	Emps []Emp `json:"emps"`
}

func init() {
	var err error
	//Set IP address and Keyspace values accordingly.
	cluster := gocql.NewCluster("172.18.0.4")
	cluster.Keyspace = "clusterdb"
	Session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	fmt.Println("cassandra init done")
}

func getEmps(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting all Employees")
	var empList []Emp
	m := map[string]interface{}{}

	iter := Session.Query("SELECT empid,first_name,last_name,age FROM emps").Iter()
	for iter.MapScan(m) {
		empList = append(empList, Emp{
			Id:        m["empid"].(string),
			FirstName: m["first_name"].(string),
			LastName:  m["last_name"].(string),
			Age:       m["age"].(int),
		})
		m = map[string]interface{}{}

	}

	json.NewEncoder(w).Encode(AllEmpsResponse{Emps: empList})
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", getEmps)

	log.Fatal(http.ListenAndServe(":8080", router))

}
