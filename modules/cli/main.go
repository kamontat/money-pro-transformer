package main

import (
	"flag"
	"log"
	"os"
	"path"

	datasource "moneypro.kamontat.net/datasource"
	csv "moneypro.kamontat.net/utils-csv"
	error "moneypro.kamontat.net/utils-error"
	logger "moneypro.kamontat.net/utils-logger"
	writer "moneypro.kamontat.net/writer"
)

var logcode = 1000

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

	onDebugMode := flag.Bool("debug", false, "set log level to debug mode")
	onErrorMode := flag.Bool("error", false, "set log level to error only mode")
	onSilentMode := flag.Bool("silent", false, "set log level to silent mode")

	flag.Parse()

	if *outputDir == "" {
		outputDir = inputDir
	}
	if *outputFile == "" {
		newString := "new-" + *inputFile
		outputFile = &newString
	}

	output := logger.Get()

	// Set to debug level
	if *onDebugMode {
		output.SetLevel(logger.DEBUG)
	} else if *onErrorMode {
		output.SetLevel(logger.ERROR)
	} else if *onSilentMode {
		output.SetLevel(logger.SILENT)
	}

	// Loading file and convert to Account struct
	application, err := datasource.Loader(output, path.Join(*rootdir, *inputDir, *inputFile))
	error.When(err).Print(output, logcode).Panic().Exit(2)

	creator, err := writer.NewFileCreator(path.Join(*rootdir, *outputDir, *outputFile))
	error.When(err).Print(output, logcode).Panic().Exit(2)

	writer := csv.NewWriter(creator, application)
	size, err := writer.Start(output)
	error.When(err).Panic().Exit(3)

	output.Info(logcode, "Writing total %d bytes", size)

	// // Create output file via File Creator
	// creator, err := writer.NewFileCreator(path.Join(*rootdir, *outputDir, *outputFile))
	// if err != nil {
	// 	log.Fatal(err)
	// 	os.Exit(2)
	// }

	// // Create output format and write data
	// writer := writer.NewCsvWriter(creator, accounts)
	// byteSize, err := writer.Writing()
	// if err != nil {
	// 	log.Fatal(err)
	// 	os.Exit(2)
	// }

	// // Print the result
	// fmt.Printf("Write data totally %d bytes\n", byteSize)
}
