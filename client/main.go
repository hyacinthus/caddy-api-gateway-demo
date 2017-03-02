package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	fin := make(chan bool)
	for i := 0; i < 30; i++ {
		go req(fin)
	}
	for i := 0; i < 30; i++ {
		<-fin
	}
}

func req(fin chan bool) {
	res, err := http.Get("http://127.0.0.1:2015")
	if err != nil {
		log.Fatal(err)
	}
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s", robots)
	fin <- true
}
