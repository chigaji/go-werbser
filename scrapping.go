package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	response, err := http.Get("https://www.devdungeon.com/archive")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	output, err := os.Create("scrapping.html")
	if err != nil {
		log.Fatal(err)
	}
	defer output.Close()

	_, err = io.Copy(output, response.Body)
	if err != nil {
		log.Fatal(err)
	}
}
