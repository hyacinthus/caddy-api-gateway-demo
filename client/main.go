package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	// dist count
	res := make(map[string]int)
	// start concurrent request
	fin := make(chan string)
	for i := 0; i < 30; i++ {
		go req(fin)
	}
	for i := 0; i < 30; i++ {
		resp := <-fin
		_, ok := res[resp]
		if ok {
			res[resp]++
		} else {
			res[resp] = 1
		}
	}
	log.Printf("statistics:\n%+v", res)
	if len(res) <= 1 {
		os.Exit(1)
	}
}

func req(fin chan string) {
	// docker0 interface is 172.10.0.1
	res, err := http.Get("http://172.17.0.1:2015")
	if err != nil {
		log.Fatal(err)
	}
	content, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s", content)
	fin <- string(content)
}
