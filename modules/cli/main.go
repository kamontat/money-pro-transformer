package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"

	datasource "kamontat.net/money-pro-datasource"
	writer "kamontat.net/money-pro-writer"
)

func main() {
	exe, err := os.Executable()
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}
	curdir := path.Dir(exe)

	rootdir := flag.String("rootDir", curdir, "base directory for find data")
	inputDir := flag.String("inputDir", "", "directory of input data")
	inputFile := flag.String("inputFile", "test.csv", "input file name (only)")
	outputDir := flag.String("outputDir", "", "directory of output data")
	outputFile := flag.String("outputFile", "", "output file name (only)")

	flag.Parse()

	// Loading file and convert to Account struct
	accounts, err := datasource.Loader(path.Join(*rootdir, *inputDir, *inputFile))
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}

	if *outputDir == "" {
		outputDir = inputDir
	}
	if *outputFile == "" {
		newString := "new-" + *inputFile
		outputFile = &newString
	}

	// Create output file via File Creator
	creator, err := writer.NewFileCreator(path.Join(*rootdir, *outputDir, *outputFile))
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}

	// Create output format and write data
	writer := writer.NewCsvWriter(creator, accounts)
	byteSize, err := writer.Writing()
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}

	// Print the result
	fmt.Printf("Write data totally %d bytes\n", byteSize)
}
