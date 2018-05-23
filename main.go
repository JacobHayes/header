// header runs a simple web server that prints all request headers.
//
// The port can be set with the `PORT` envvar and defaults to 8080.
package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"
)

// printHeaders writes request headers to the response
func printHeaders(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, string(dump))
}

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	http.HandleFunc("/", printHeaders)
	err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
	if err != nil {
		fmt.Print(err)
	}
}
