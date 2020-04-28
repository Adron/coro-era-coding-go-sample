package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", RootHandler)
	http.HandleFunc("/response", ResponseExampleHandler)
	http.HandleFunc("/errexample", ErrorResponseHandler)
	http.ListenAndServe(":8080", nil)
}

// A good example of handling the "request" body.
func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is my content.")
	fmt.Fprintln(w, r.Header)

	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintln(w, err)
	}
	fmt.Fprintln(w, string(body))
}

func ResponseExampleHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	fmt.Fprintln(w, "Testing status code. Manually added a 200 Status OK.")
	fmt.Fprintln(w, "Another line.")
}

func ErrorResponseHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(500)
	fmt.Fprintln(w, "Server error.")
}
