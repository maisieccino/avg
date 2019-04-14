package server

import (
	"fmt"
	"net/http"
)

// Start starts the server
func Start(host, port string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Average"))
	})
	addr := fmt.Sprintf("%s:%s", host, port)
	http.ListenAndServe(addr, nil)
}
