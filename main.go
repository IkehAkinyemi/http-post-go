package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
  Id string `json:"Id"`
  Title string `json:"Title"`
  Content string `json:"Content"`
  Summary string `json:"Summary"`
}

// var post Article


func createNewArticle(w http.ResponseWriter, r *http.Request) {
  reqBody, _ := ioutil.ReadAll(r.Body)
  var post Article 
  json.Unmarshal(reqBody, &post)

  json.NewEncoder(w).Encode(post)

  newData, err := json.Marshal(post)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(string(newData))
  }
}

func handleReqs() {
  r := mux.NewRouter().StrictSlash(true)
  r.HandleFunc("/post", createNewArticle).Methods("POST")

  log.Fatal(http.ListenAndServe(":8000", r))
}

func main() {
	handleReqs();
}