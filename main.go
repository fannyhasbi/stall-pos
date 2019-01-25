package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/fannyhasbi/stall-pos/employee"
	"github.com/fannyhasbi/stall-pos/order"
	"github.com/fannyhasbi/stall-pos/product"
)

const PORT = ":8080"

func init(){
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	response := "This is web service for Stall POS"
	fmt.Fprintln(w, response)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", handleIndex).Methods("GET", "POST")

	router.HandleFunc("/api/product", product.GetProducts).Methods("GET")

	router.HandleFunc("/api/employee", employee.HandleEmployee).Methods("POST")

	router.HandleFunc("/api/order", order.HandleOrder).Methods("POST")

	http.Handle("/", router)

	fmt.Printf("Server running on localhost%v\n", PORT)
	log.Fatal(http.ListenAndServe(PORT, router))
}
