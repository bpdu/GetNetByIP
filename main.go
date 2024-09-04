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
	-v, --verbose 	Verbose output including country

`
	HELP_OPTION_SHORT = "-h"
	HELP_OPTION_LONG  = "--help"

	VERBOSE_OPTION_SHORT = "-v"
	VERBOSE_OPTION_LONG  = "--verbose"

	EXIT_STATUS_OK    = 0
	EXIT_STATUS_ERROR = 1

	IP_PATTERN = "^(?:[0-9]{1,3}.){3}[0-9]{1,3}$"

	API_URL = "https://api.incolumitas.com/"
)

func isValidIP(s string) bool {
	probe, _ := regexp.MatchString(IP_PATTERN, s)
	return probe
}

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
		} else if isValidIP(arg) {
			resp, err := http.Get(API_URL + "?q=" + arg)
			if err != nil {
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
			asn_section := result["asn"].(map[string]any)
			fmt.Printf("Network: %s\n", asn_section["route"])
			os.Exit(EXIT_STATUS_OK)
		} else {
			fmt.Printf("Wrong IP address or unknown option, try %s to help\n", "\"go run getNetByIP --help\"")
			os.Exit(EXIT_STATUS_ERROR)
		}
	case 3:
		arg1 := os.Args[1]
		arg2 := os.Args[2]
		if (arg1 == VERBOSE_OPTION_SHORT || arg1 == VERBOSE_OPTION_LONG) && isValidIP(arg2) {
			resp, err := http.Get(API_URL + "?q=" + arg2)
			if err != nil {
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
			asn_section := result["asn"].(map[string]any)
			location_section := result["location"].(map[string]any)
			fmt.Printf("Network: %s\nCountry: %s\n", asn_section["route"], location_section["country"])
			os.Exit(EXIT_STATUS_OK)
		} else {
			fmt.Printf("Wrong IP address or unknown option, try %s to help\n", "\"go run getNetByIP --help\"")
			os.Exit(EXIT_STATUS_ERROR)
		}
	}
}
