package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	sourceId, err := getOptions()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	apiKey, err := readSecrets()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(sourceId)
	fmt.Println(apiKey)
}

func getOptions() (string, error) {
	sourceId := flag.String("id", "", "The ID of the data source to update")
	flag.Parse()

	if *sourceId == "" {
		flag.Usage()
		return "", errors.New("No data source ID given")
	}

	return *sourceId, nil
}

func readSecrets() (string, error) {
	const path string = "secrets.json"
	apiKeyJson, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	var apiKey struct { Value string `json:"api-key"` }
	err = json.Unmarshal(apiKeyJson, &apiKey)
	if err != nil {
		return "", err
	}

	return apiKey.Value, nil
}
