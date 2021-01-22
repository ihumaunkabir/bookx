package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	books "github.com/oasiscse/bookx/api/books"
)

//GetRoutes gets all api routes
func GetRoutes() {

	r := mux.NewRouter()

	r.HandleFunc("/", homepage).Methods("GET")
	r.HandleFunc("/api/books/info/all", books.GetBooks).Methods("GET")
	r.HandleFunc("/api/books/info/{bookid}", books.GetByID).Methods("GET")
	r.HandleFunc("/api/books/remove/{bookid}", books.DeleteByID).Methods("DELETE")
	r.HandleFunc("/api/books/info/all/{year}", books.GetByYear).Methods("GET")
	r.HandleFunc("/api/books/create", books.Create).Methods("POST")

	log.Fatal(http.ListenAndServe(":3000", r))
}

func homepage(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"bookx" : "A clever monkey is cooking for you, wait..."}`))
}
