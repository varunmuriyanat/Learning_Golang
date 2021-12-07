package main

import (
    "io/ioutil"
    "log"
    "net/http"
)

func main() { 
    resp, err := http.Get("http://localhost:8080/v1/organisation/accounts/ad27e265-9605-4b4b-a0e5-3003ea9cc4dc")
    if err != nil {
        log.Fatalln(err)
    }

    // read response
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatalln(err)
    }

    // convert the body to string
    sb := string(body)
    log.Printf(sb)
}

