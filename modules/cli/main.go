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
	measure "moneypro.kamontat.net/utils-measure"
	writer "moneypro.kamontat.net/writer"
)

var logcode = 1000

func main() {
	timing := measure.NewTiming()

	stepname := "Load current data"
	exe, err := os.Executable()
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}
	curdir := path.Dir(exe)
	timing.Save(stepname)

	stepname = "Parse script parameters"
	rootdir := flag.String("rootDir", curdir, "base directory for find data")
	inputDir := flag.String("inputDir", "", "directory of input data")
	inputFile := flag.String("inputFile", "test.csv", "input file name (only)")
	outputDir := flag.String("outputDir", "", "directory of output data")
	outputFile := flag.String("outputFile", "", "output file name (only)")

	onDebugMode := flag.Bool("debug", false, "set log level to debug mode")
	onErrorMode := flag.Bool("error", false, "set log level to error only mode")
	onSilentMode := flag.Bool("silent", false, "set log level to silent mode")

	flag.Parse()
	timing.Save(stepname)

	if *outputDir == "" {
		outputDir = inputDir
	}
	if *outputFile == "" {
		newString := "new-" + *inputFile
		outputFile = &newString
	}

	stepname = "Setup logging"
	output := logger.Get()
	// Set to debug level
	if *onDebugMode {
		output.SetLevel(logger.DEBUG)
	} else if *onErrorMode {
		output.SetLevel(logger.ERROR)
	} else if *onSilentMode {
		output.SetLevel(logger.SILENT)
	}
	timing.Save(stepname)

	i := 10
	for key, duration := range timing.Release() {
		if key != "All" {
			output.Time(logcode+i, key, duration.String())
			i++
		}
	}

	stepname = "Load csv file and transform"
	profile, err := datasource.Loader(output, path.Join(*rootdir, *inputDir, *inputFile))
	error.When(err).Exit(2)
	timing.LogSnapshot(stepname, output, logcode+13).Save(stepname)
	profile.Info(output, logcode)
	profile.Debug(output, logcode)

	stepname = "Create output file"
	creator, err := writer.NewFileCreator(path.Join(*rootdir, *outputDir, *outputFile))
	error.When(err).Exit(3)
	timing.LogSnapshot(stepname, output, logcode+14).Save(stepname)

	stepname = "Write data to output file"
	writer := csv.NewWriter(creator, profile)
	writer.Info(output)

	size, err := writer.Start(output)
	error.When(err).Exit(4)
	timing.LogSnapshot(stepname, output, logcode+15).Save(stepname)

	output.Info(logcode, "Writing total %d bytes", size)

	timing.LogAll("All", output, logcode+100)
}
