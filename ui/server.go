package ui

import (
	"deedee/groupie-tracker/dal"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const port string = ":8080"

var ArtistDTMs []dal.ArtistDTM
var RelationDTMs []dal.RelationDTM
var ArtistIDToViewModelMap map[int]ArtistViewModel

func init() {
	ArtistDTMs = dal.GetArtistDTMs()
	RelationDTMs = dal.GetRelationDTMs(ArtistDTMs)
	InitIDToViewModelMap(ArtistDTMs)
}

// defaultHandler is responsible for handling all HTTP requests
// received at all endpoints in our application
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	// Verify HTTP method - if method is not 'GET', log an error
	if r.Method != http.MethodGet {
		// Log an error if not a GET method
		log.Printf("%s - %s - %d %s\n", r.Method, r.URL.Path, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		// Writes the 'Bad Request' response header
		w.WriteHeader(http.StatusBadRequest)
		// Write a simple message to the client saying why it was a bad request
		w.Write([]byte("This site does not support non-GET HTTP requests.\n"))
		return
	}
	// Check if the request was made to the home endpoint
	// If it was, then load the home page
	if r.URL.Path == "/" {
		// Initialise a view model for the home page
		var p HomeViewModel = HomeViewModel{Artists: ArtistViewModels}
		// Log the request info to the console
		log.Printf("%s - %s - %d %s\n", r.Method, r.URL.Path, http.StatusOK, http.StatusText(http.StatusOK))
		// load the html file with go code adapted to html format.
		t, err := template.ParseFiles("./wwwroot/MainLayout.html")
		//
		if err != nil {
			// WriteHeader is a method of the http.ResponseWriter interface
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Oops! An error occurred! Try refreshing the page."))
			return
		}
		// Execute the template with the data passed as a struct
		// that contains as fields the relevant information
		err = t.Execute(w, p)
		if err != nil {
			//if for some reason there was a problem executing the temlate
			//informs the client that it was the server's mistake
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Oops! An error occurred! Try refreshing the page."))
		}
		return
	}
	// Strip the forward slash from the path
	var path string = strings.TrimPrefix(r.URL.Path, "/")
	var id int
	var err error
	// Try and convert the string into a number
	id, err = strconv.Atoi(path)
	if err != nil {
		log.Printf("%s - %s - %d %s\n", r.Method, r.URL.Path, http.StatusNotFound, http.StatusText(http.StatusNotFound))
		http.NotFound(w, r)
		return
	}
	var found bool
	var model ArtistViewModel
	// Get the ArtistViewModel value from the entry in the map with key == id
	model, found = ArtistIDToViewModelMap[id]
	// if the model is not found, 404 error
	if !found {
		// Print a 404 not found if page does not exist
		log.Printf("%s - %s - %d %s\n", r.Method, r.URL.Path, http.StatusNotFound, http.StatusText(http.StatusNotFound))
		http.NotFound(w, r)
		return
	}
	// If we reach here, we have the ArtistViewModel
	t, err := template.ParseFiles("./wwwroot/artists.html")
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Oops! An error occurred! Try refreshing the page."))
		return
	}
	err = t.Execute(w, model)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Oops! An error occurred! Try refreshing the page."))
		return
	}
	log.Printf("%s - %s - %d %s\n", r.Method, r.URL.Path, http.StatusOK, http.StatusText(http.StatusOK))
}

func RunServer() {
	// Create a new server mutex to deal with concurrent client requests
	mux := http.NewServeMux()
	// Set up a file server where all of our static files (HTML, CSS) are stored (wwwroot)
	fs := http.FileServer(http.Dir("wwwroot"))
	// Set up the mutex to use the file server
	mux.Handle("/wwwroot/", http.StripPrefix("/wwwroot/", fs))
	// Keep it simple, using one handler to manage requests to any endpoint
	mux.HandleFunc("/", defaultHandler)
	// Print out a log message to show our server is running
	log.Printf("Server listening on port %s...\n", port)
	// Start the server and listen for incoming requests, and on which port
	err := http.ListenAndServe(port, mux)
	// If there is an error with running the server, log the message and exit
	if err != nil {
		log.Fatal(err.Error())
	}
}
