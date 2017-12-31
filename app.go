package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Args Args `json:"args"`
}

type Args struct {
	Test string `json:"test"`
}

func main() {
	fmt.Println("Running Example...")

	response, err := http.Get("http://httpbin.org/get?test=1")

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	fmt.Println("Args: " + responseObject.Args.Test)
}
