package main


import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}
type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)

}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, movi := range movies {
		if movi.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, movi := range movies {
		if movi.ID == params["id"] {
			json.NewEncoder(w).Encode(movi)
			break
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var m Movie
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		log.Fatal(err)
	}
	m.ID = strconv.Itoa(rand.Intn(10000))
	movies = append(movies, m)
	json.NewEncoder(w).Encode(m)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	params := mux.Vars(r)
	for index, movi := range movies {
		if movi.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var m Movie
			_ = json.NewDecoder(r.Body).Decode(&m)
			m.ID = params["id"]
			movies = append(movies, m)
			json.NewEncoder(w).Encode(movies)
			return
		}
	}

}

var movies []Movie

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "123", Title: "movie1", Director: &Director{FirstName: "john", LastName: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "1234", Title: "movie2", Director: &Director{FirstName: "karan", LastName: "johar"}})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	fmt.Println("starting server at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}


