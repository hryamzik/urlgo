package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
)

func checkErrFatal(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	url := os.Args[1]
	matched, err := regexp.MatchString("https?://", url)
	checkErrFatal(err)
	if !matched {
		url = "http://" + url
	}
	response, err := http.Get(url)
	checkErrFatal(err)
	defer response.Body.Close()
	data, _ := ioutil.ReadAll(response.Body)
	fmt.Printf("%s", data)
}
