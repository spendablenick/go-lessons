// // https://api.example.com/user/123?type=object

// example.com/123436

// example.com/r/123436

// spendable.uk/1234567

// spendable.org.uk

// example.com/api/spendcash

// nickshop.com

// api.nickshop.com

// test.nicksshop.com

package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	t := vars["resource"]
	fmt.Fprintf(w, "Resource type = %s!\n", t)
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

func createResource(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]bool{"Create": true})
}
func updateResource(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]bool{"Update": true})
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/user", createResource).Methods("POST")
	r.HandleFunc("/user", updateResource).Methods("PUT")

	r.HandleFunc("/health", jsonHandler)
	r.HandleFunc("/{resource}", handler)
	http.ListenAndServe(":8080", r)
	// log.Fatal(http.ListenAndServe(":8080", nil))
}



/// url shortner 

shorturl = map[string]string

