package main

import (
	"encoding/json"
	"fmt"
	"log"
	// "maps"
	"strconv"

	// "math/rand"
	"net/http"
	"github.com/gorilla/mux"
)

type Movie struct {
     ID string `json:"id"`
	 Isbn string `json:"isbn"`
	 Title string `json:"title"`
	 Director *Director `json:"director"`
}

type Director struct {
     FullName string `json:"fullName"`
	 Email string `json:"email"`
}

var movies []Movie


func getMovies(w http.ResponseWriter , r *http.Request){
   w.Header().Set("Content-Type", "application/json")
   json.NewEncoder(w).Encode(movies)
}


func deleteMovie(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for idx , item := range movies {
		if item.ID == params["id"]{
			movies = append(movies[:idx] , movies[idx+1:]...)
			json.NewEncoder(w).Encode(map[string]string{"message": "Movie deleted successfully"})
			// json.NewDecoder(w).Encode()
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Movie not found"})
}


func getMovie(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type" , "application/json")
	params := mux.Vars(r)
	for _ , item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error" : "Movie not found"})
}

func generateIsbn(id string) string {
	return id + id + id + id + id
}
func createMovie(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type" , "application/json")
	var movie Movie 
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message" : "invalid request body"})
		return
	}

	maxId := 0
	for _ , m := range movies {
		id , err := strconv.Atoi(m.ID)
		if err != nil {
			continue
		}
		if id > maxId {
			maxId = id
		}
	}

	newId := strconv.Itoa(maxId + 1)
	movie.ID = newId
	movie.Isbn = generateIsbn(newId)
	movies = append(movies , movie)

	json.NewEncoder(w).Encode(movie)
}


func main(){

	//! appends movies
		movies = append(movies, Movie{
		ID:   "1",
		Isbn: "11111",
		Title: "Inception",
		Director: &Director{
			FullName: "Christopher Nolan",
			Email:    "nolan@example.com",
		},
	})

	movies = append(movies, Movie{
		ID:   "2",
		Isbn: "22222",
		Title: "Interstellar",
		Director: &Director{
			FullName: "Jonathan Nolan",
			Email:    "jnolan@example.com",
		},
	})

	movies = append(movies, Movie{
		ID:   "3",
		Isbn: "33333",
		Title: "The Dark Knight",
		Director: &Director{
			FullName: "Christopher Nolan",
			Email:    "batman@example.com",
		},
	})


	r := mux.NewRouter()

	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	// r.HandleFunc("/movies/{id}", updateMovie).Methods("PATCH")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Server is running at port: 3333\n")
	log.Fatal(http.ListenAndServe(":3333" , r) )
}