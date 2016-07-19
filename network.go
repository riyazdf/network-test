package main

import (
	"fmt"
	"net/http"
	"os"
)


const NotaryUrl = "https://notary.docker.io"

func main() {
	// first get back a challenge to find the auth URL
	req, err := http.NewRequest("GET", NotaryUrl, nil)
	if err != nil {
		fatalf(err.Error())
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fatalf(err.Error())
	}
	resp.Body.Close()
	fmt.Println(resp)
	fmt.Println("SUCCESS!")
}

func fatalf(format string, args ...interface{}) {
	fmt.Printf("* fatal: "+format+"\n", args...)
	os.Exit(1)
}
