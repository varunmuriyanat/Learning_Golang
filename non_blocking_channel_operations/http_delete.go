package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	req, err := http.NewRequest(http.MethodDelete, "http://localhost:8080/v1/organisation/accounts/bd27e265-9605-4b4b-a0e5-3003ea9cc4dc?version=0", nil)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	// read response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// convert the body to string
	sb := string(body)
	log.Printf(sb)
}
