package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/loosec9n/Project_Two_CRUD_API/models"
	"github.com/loosec9n/Project_Two_CRUD_API/routes"

)

var movies []models.Movie

func main() {
	log.Println("CRUD API Examples")

	 // If the file doesn't exist, create it or append to the file
	 file, err := os.OpenFile("CRUD_Project_Two.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	 if err != nil {
		 log.Fatal(err)
	 }
 
	 log.SetOutput(file)

	r := mux.NewRouter()


	 //Seeding movies data into the struct:
		movies = append(movies, Movie{
			ID:"1",
			Isbn:"438227",
			Title:"Movie One",
			Director: &Director{
				FirstName:"John",
				LastName:"Doe"
			}
		})

		movies = append(movies, Movie{
			ID:"2",
			Isbn:"459237",
			Title:"Movie Two",
			Director: &Director{
				FirstName:"Steven",
				LastName:"Smith"
			}
		})

	r.HandleFunc("/movies",routes.GetMovies).Method("GET")
	r.HandleFunc("/movies/{id}",routes.GetMovie).Method("GET")
	r.HandleFunc("/movies",routes.CreateMovie).Method("POST")
	r.HandleFunc("/movies/{id}",routes.UpdateMovie).Method("PUT")
	r.HandleFunc("/movies/{id}",routes.DeleteMovie).Method("DELETE")

	log.Println("Starting server at port 8080")
	log.Fatal(http.ListenAndServe(":8080",r))


}
