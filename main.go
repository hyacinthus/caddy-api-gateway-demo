package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// HelloServer the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	hostname, _ := os.Hostname()
	fmt.Print("Hello, world! I am " + hostname + " :)\n")
	io.WriteString(w, "Hello, world! I am "+hostname+" :)\n")
}

func main() {
	http.HandleFunc("/", HelloServer)
	log.Fatal(http.ListenAndServe(":12345", nil))
}
