package main

import (
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		fmt.Println(string(body))
		fmt.Fprintf(w, "Request Path: %q", html.EscapeString(r.URL.Path))
	})

	err := http.ListenAndServe(":7879", nil)
	if err != nil {
		fmt.Println(err)
	}
}
