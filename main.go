package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	url := os.Args[1]
	response, err := http.Get(url)
	if err != nil {
		log.Panic(err)
	}
	defer response.Body.Close()
	data, _ := ioutil.ReadAll(response.Body)
	fmt.Printf("%s\n", data)
}
