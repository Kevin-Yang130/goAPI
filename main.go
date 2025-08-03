package main

import (
	"fmt"
	"net/http"
)

// *http.Request: The client request with the parameters and metadata
// ResponseWriter: What the backend server uses to write back a response to the user
func apiHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func main() {
	http.HandleFunc("/api", apiHandler)
	http.HandleFunc("/api/user", apiHandler)
	
	fmt.Println("Starting server...")
		
	//opens the port 8080 for all interfaces/IP addresses.
	http.ListenAndServe(":8080", nil)
}