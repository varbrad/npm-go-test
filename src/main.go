package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

type inputFile struct {
	filepath  string
	seperator string
	pretty    bool
}

func getFileData() (inputFile, error) {
	if len(os.Args) < 2 {
		return inputFile{}, errors.New("a filepath argument is required")
	}

	seperator := flag.String("seperator", "comma", "Column seperator")
	pretty := flag.Bool("pretty", false, "Generate pretty JSON")

	flag.Parse()

	fileLocation := flag.Arg(0)

	if !(*seperator == "comma" || *seperator == "semicolon") {
		return inputFile{}, errors.New("invalid seperator - only 'comma' or 'semicolon' are supported")
	}

	return inputFile{fileLocation, *seperator, *pretty}, nil
}

func main() {
	fileData, err := getFileData()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("File Data", fileData)

}
