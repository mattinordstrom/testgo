package httphandling

import (
	"fmt"
	"net/http"
)

func httpStupidFunction(someNumber int) int {
	return someNumber + 3
}

//CreateHTTPListener creates http listener
func CreateHTTPListener() {
	http.HandleFunc("/", httpHandler)
	http.HandleFunc("/test", httpHandler2)

	http.ListenAndServe(":8080", nil)
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func httpHandler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world 2")
}
