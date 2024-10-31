package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var task string

func POSTHandler(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)
	var t struct {
		Task string `json:"task"`
	}
	if err := decoder.Decode(&t); err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	task = t.Task
	w.WriteHeader(http.StatusOK)
}




func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello", task)
}


func main() {
	InitDB()

	DB.AutoMigrate(&Message{})

	router := mux.NewRouter()
	router.HandleFunc("/api/hello", HelloHandler).Methods("GET")
	router.HandleFunc("/api/hello", POSTHandler).Methods("POST")
	http.ListenAndServe(":8080", router)
}