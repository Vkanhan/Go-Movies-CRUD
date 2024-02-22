package main

import(
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Movie struct {
	ID       string    `json:"id"`
	ISBN     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["ID"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)

}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range movies {
		if item.ID == params["ID"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_, json.NewDecode(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	r.Header().Set("Content-Type", "application/json") //set json content type
	params := mux.Vars(r)
	//json content type
	//params
	//for loop over movies
	//delete movies with the id you send
	//add movies you send with postman

	for index, item := range movies {
		if item.ID == params["ID"] {
			movies = append(movies[:index], movie[index+1:])
			
		}
	}
	json.NewEncode(w).Encode(movie)
}

func main() {

	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", ISBN: "43423", Title: "Movie1", Director: &Director{FirstName: "Martin", LastName: "Scorcese"}})
	movies = append(movies, Movie{ID: "2", ISBN: "43424", Title: "Movie2", Director: &Director{FirstName: "Christofer", LastName: "Nolan"}})

	r.HandleFunc("/movies", getMovies).Method("GET")
	r.HandleFunc("/movies/{id}", getMovie).Method("GET")
	r.HandleFunc("/movies", createMovie).Method("POST")
	r.HandleFunc("/movies", updateMovie).Method("PUT")
	r.HandleFunc("/movie", deleteMovie).Method("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatal(err)
	}

}
