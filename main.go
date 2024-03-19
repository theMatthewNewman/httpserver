package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

// represents data about album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Returning all albums")
	json.NewEncoder(w).Encode(albums)
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Got / request \n")
	io.WriteString(w, "This is my Website! \n")
}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}

func methodHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Got /method request \n")
	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "GET request received \n")
	case "POST":
		fmt.Fprintf(w, "POST request received \n")
	default:
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)
	http.HandleFunc("/method", methodHandler)
	http.HandleFunc("/albums", getAlbums)
	err := http.ListenAndServe(":3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("Server Closed \n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
