package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/test", handleTest)
	http.ListenAndServe(":8080", nil)
}

func handleTest(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprint(w, name)
}
