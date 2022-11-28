package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var movies []Movie

func init() {
	johnDoe := &Director{FirstName: "John", LastName: "Doe"}
	jamesClear := &Director{FirstName: "James", LastName: "Clear"}

	movies = append(movies, Movie{ID: "1", Isbn: "1", Title: "Hello", Director: johnDoe})
	movies = append(movies, Movie{ID: "2", Isbn: "2", Title: "Hi", Director: jamesClear})
	movies = append(movies, Movie{ID: "3", Isbn: "3", Title: "Hey", Director: jamesClear})
}

func listMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	movie.ID = strconv.Itoa(len(movies) + 1)
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			if len(movie.Title) > 0 {
				movies[index].Title = movie.Title
			}

			if len(movie.Isbn) > 0 {
				movies[index].Title = movie.Isbn
			}

			if len(movie.Director.FirstName) > 0 {
				movies[index].Director.FirstName = movie.Director.FirstName
			}

			if len(movie.Director.LastName) > 0 {
				movies[index].Director.LastName = movie.Director.LastName
			}

			movie = movies[index]
			break
		}
	}

	json.NewEncoder(w).Encode(movie)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)

	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[0:index], movies[index+1:]...)
			return
		}
	}
}
