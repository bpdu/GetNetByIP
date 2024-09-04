package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
)

const (
	USAGE_INFO = `
Get network ID and country information by IP address.

Usage: go run getNetByIP [OPTION] IP_ADDRESS
					
Options:

	-h, --help 	Usage information
	-c, --country 	Include country information

`
	HELP_OPTION_SHORT = "-h"
	HELP_OPTION_LONG  = "--help"

	COUNTRY_OPTION_SHORT = "-c"
	COUNTRY_OPTION_LONG  = "--country"

	EXIT_STATUS_OK    = 0
	EXIT_STATUS_ERROR = 1

	IP_PATTERN = "^(?:[0-9]{1,3}.){3}[0-9]{1,3}$"

	API_URL = "https://api.incolumitas.com/"
)

func main() {

	switch len(os.Args) {
	case 1:
		fmt.Print(USAGE_INFO)
		os.Exit(EXIT_STATUS_OK)
	case 2:
		arg := os.Args[1]
		if arg == HELP_OPTION_SHORT || arg == HELP_OPTION_LONG {
			fmt.Print(USAGE_INFO)
			os.Exit(EXIT_STATUS_OK)
		} else if isValidIP, _ := regexp.MatchString(IP_PATTERN, arg); isValidIP {
			resp, errr := http.Get(API_URL + "?q=" + arg)
			if errr != nil {
				fmt.Println("Network error")
				os.Exit(EXIT_STATUS_ERROR)
			}
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("Data error")
				os.Exit(EXIT_STATUS_ERROR)
			}

			var result map[string]any
			json.Unmarshal([]byte(body), &result)
			asn := result["asn"].(map[string]any)
			fmt.Println(asn["route"])
			os.Exit(EXIT_STATUS_OK)
		} else {
			fmt.Printf("Wrong IP address or unknown option, try %s\n", "\"go run getNetByIP --help\"")
			os.Exit(EXIT_STATUS_ERROR)
		}
	}
}
