package main

import( "fmt" 
        "log" 
		"encoding/json" 
		"math/rand" 
		"net/http" 
		"strconv" 
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

func main(){
	r := mux.NewRouter()

	r.HandleFunc("/movies/", createMovie).Methods("POST")
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PATCH")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

}