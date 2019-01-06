package main

import (
  "fmt"
  "log"
  "net/http"
  
  _ "github.com/go-sql-driver/mysql"
  "github.com/gorilla/mux"
)

const PORT = ":8080"

func handleIndex(w http.ResponseWriter, r *http.Request){
  response := "This is web service for Stall POS"
  fmt.Fprintln(w, response)
}

func main(){
  router := mux.NewRouter()

  router.HandleFunc("/", handleIndex).Methods("GET", "POST")

  http.Handle("/", router)

  fmt.Printf("Connected to port : %v", PORT)
  log.Fatal(http.ListenAndServe(PORT, router))
}