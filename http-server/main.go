// USECASE
// $ go run main.go
// $ curl localhost:8080
// $ curl localhost:8080/hello-world
// This can be also opened in a browser

package main

import (
	"net/http"
)

func main() {
	// Without providing handlers, it returns a 404 error page as it serves nothing
	http.HandleFunc("/hello-world", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello-world\n"))
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("homepage\n"))
	})

	http.ListenAndServe(":8080", nil)
}
