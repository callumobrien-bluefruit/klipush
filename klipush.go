package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

func main() {
	sourceId, err := getOptions()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(sourceId)
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
