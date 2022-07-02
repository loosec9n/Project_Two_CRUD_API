package routes

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/loosec9n/Project_Two_CRUD_API/models"
)

var movie []models.Movie

func GetMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movie)
	log.Println("I was here")
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movie {
		if item.ID == params["id"] {
			movie = append(movie[:index], movie[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movie)
}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movie {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movies models.Movie
	_ = json.NewDecoder(r.Body).Decode(&movies)
	movies.ID = strconv.Itoa(rand.Intn(1000000))
	movie = append(movie, movies)
	json.NewEncoder(w).Encode(movies)
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movie {
		if item.ID == params["ID"] {
			movie = append(movie[:index], movie[index+1:]...)
			var movies models.Movie
			_ = json.NewDecoder(r.Body).Decode(&movies)
			movies.ID = params["id"]
			movie = append(movie, movies)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}
