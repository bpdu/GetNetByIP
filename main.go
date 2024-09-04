package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	USAGE_MESSAGE = `

Get network ID and metadata by requested IP address.

Usage: go run getNetByIP [OPTION] IP_ADDRESS
					
Options:

`
)

func main() {

	if len(os.Args) == 1 {
		fmt.Print(USAGE_MESSAGE)
	}

	address := os.Args[1]
	url := "https://api.incolumitas.com/"

	resp, err := http.Get(url + "?q=" + address)
	if err != nil {
		fmt.Println("Network error")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Data error")
	}

	var result map[string]any
	json.Unmarshal([]byte(body), &result)
	asn := result["asn"].(map[string]any)
	fmt.Println(asn["route"])

}
