package main

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)


func main() {

	

	r := mux.NewRouter()
	// r.HandleFunc(
	// 	"/healthcheck",
	// 	func(w http.ResponseWriter, r *http.Request) {
	// 		h := health{
	// 			Status:   "OK",
	// 			Messages: []string{},
	// 		}

	// 		b, _ := json.Marshal(h)

	// 		w.WriteHeader(http.StatusOK)
	// 		w.Write(b)
	// })
	// r.HandleFunc(
	// 	"/deck/{id}",
	// 	func(w http.ResponseWriter, r *http.Request) {
	// 		fmt.Println("getting env")

	// 		fmt.Println( os.Getenv("DATABASE_URL"))
	// 		vars := mux.Vars(r)
	// 		id := vars["id"]
	// 		h := Error{
	// 			Code:   "001",
	// 			Message: fmt.Sprintf("No Deck with ID %s", id),
	// 		}

	// 		b, _ := json.Marshal(h)

	// 		w.WriteHeader(http.StatusNotFound)
	// 		w.Write(b)
	// })

	s := http.Server{
		Addr:              ":9999",
		Handler:           r,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}
	s.ListenAndServe()
}