// USECASE
// A TCP server protocol
// $ go run main.go
// $ curl localhost:8080
// $ curl localhost:8080/hello-world
// This can be also opened in a browser / done via POSTMAN

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleRoot)
	mux.HandleFunc("POST /users", createUsers)
	// {id} is the wildcard used to identify the path
	mux.HandleFunc("GET /users/{id}", getUser)
	mux.HandleFunc("DELETE /users/{id}", deleteUser)

	fmt.Println("Server listening to 8080...")
	http.ListenAndServe(":8080", mux)
}

// declaring struct data members with json tags.
type User struct {
	Name string `json:"name"`
}

var (
	cacheUser  = make(map[int]User)
	cacheMutex sync.RWMutex
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "helloWorld\n")
}

func createUsers(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if user.Name == "" {
		http.Error(w, "name is required", http.StatusBadRequest)
		return
	}

	cacheMutex.Lock()
	cacheUser[len(cacheUser)+1] = user
	cacheMutex.Unlock()

	// TODO: enhance this by returning the user [COMPLETED]
	// NOTE: WriterHeader should be called at last, in this example calling it earlier than Header().Set("Content-Type", "application/json") doesn't set the header content type
	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	// PathValue is available since go 1.22
	userId, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cacheMutex.RLock()
	user, ok := cacheUser[userId]
	cacheMutex.RUnlock()
	if !ok {
		http.Error(w, "user doesn't exist", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if _, ok := cacheUser[userId]; !ok {
		http.Error(w, "user doesn't exist", http.StatusNotFound)
		return
	}

	cacheMutex.Lock()
	delete(cacheUser, userId)
	cacheMutex.Unlock()

	w.WriteHeader(http.StatusNoContent)
}
