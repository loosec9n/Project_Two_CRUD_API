package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/loosec9n/Project_Two_CRUD_API/routes"
)

func main() {
	log.Println("CRUD API Examples")

	// If the file doesn't exist, create it or append to the file
	file, err := os.OpenFile("CRUD_Project_Two.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)

	r := mux.NewRouter()

	r.HandleFunc("/movies", routes.GetMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", routes.GetMovie).Methods("GET")
	r.HandleFunc("/movies", routes.CreateMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", routes.UpdateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", routes.DeleteMovie).Methods("DELETE")

	log.Println("Starting server at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
