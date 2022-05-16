package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	go getData("https://catfact.ninja/fact")

}

type Data struct {
	Fact   string
	Length uint
}

func getData(url string) {

	data := Data{}
	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal([]byte(string(body)), &data)

	fmt.Println(data.Fact)
}
