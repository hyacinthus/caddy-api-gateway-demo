package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

// HelloServer the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	hostname, _ := os.Hostname()
	log.Println(hostname)
	io.WriteString(w, "Hello, world! I am "+hostname+" .")
}

func main() {
	http.HandleFunc("/", HelloServer)
	log.Fatal(http.ListenAndServe(":12345", nil))
}
