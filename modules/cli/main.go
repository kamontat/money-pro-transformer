package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"

	connection "moneypro.kamontat.net/connection-common"
	datasource "moneypro.kamontat.net/datasource"
	csv "moneypro.kamontat.net/utils-csv"
	error "moneypro.kamontat.net/utils-error"
	logger "moneypro.kamontat.net/utils-logger"
	measure "moneypro.kamontat.net/utils-measure"
)

// VERSION is commandline version
const VERSION = "v1.0.1"

var logcode = 1000

func _version(name string, version string) string {
	return fmt.Sprintf("%-12s version: %s", name, version)
}

func version(output *logger.Logger) {
	output.Info(0, _version("Core", VERSION))
	if output.IsDebug() {
		output.Info(0, _version("Datasource", datasource.VERSION))
		output.Info(0, _version("Connection", connection.VERSION))
		output.Info(0, _version("CSV Writer", csv.VERSION))

		output.Info(0, _version("Logger", logger.VERSION))
		output.Info(0, _version("Error", error.VERSION))
		output.Info(0, _version("Measure", measure.VERSION))
	}
}

func main() {
	timing := measure.NewTiming()

	stepname := "Step: Load application data"
	exe, err := os.Executable()
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}
	curdir := path.Dir(exe)
	timing.Save(stepname)

	stepname = "Step: Parse script parameters"
	rootdir := flag.String("rootDir", curdir, "base directory for find data")
	inputDir := flag.String("inputDir", "", "directory of input data")
	inputFile := flag.String("inputFile", "moneypro.csv", "input file name (only)")
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

	stepname = "Step: Setup logging"
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

	// Print version information
	version(output)

	i := 10
	for key, duration := range timing.Release() {
		if key != "All" {
			output.Time(logcode+i, key, duration.String())
			i++
		}
	}

	// // open input connection
	// inputConnection, err := connection.NewInputFile(path.Join(*rootDir, *inputDir, *inputFile))
	// error.When(err).Exit(2)
	// inputConnection.AutoClose()

	// // Transform input to struct
	// reader := csv.Reader(inputConnection)
	// profile, err := reader.Transform(output, logcode)
	// error.When(err).Exit(4)
	// profile.Info(output, logcode)
	// profile.Debug(output, logcode)

	// // create output connection
	// outputConnection, err := connection.NewOutputFile(path.Join(*rootDir, *outputDir, *outputFile))
	// error.When(err).Exit(4)
	// outputConnection.AutoClose()

	// // Create csv writer
	// writer := csv.Writer(outputConnection)

	// // Write data to output
	// writer.Write(profile)

	stepname = "Step: Load csv file and transform"
	profile, err := datasource.LoaderV1(output, path.Join(*rootdir, *inputDir, *inputFile))
	error.When(err).Exit(2)
	timing.LogSnapshot(stepname, output, logcode+13).Save(stepname)
	profile.Info(output, logcode)
	profile.Debug(output, logcode)

	stepname = "Step: Create output file"
	outputConnection, err := connection.NewOutputFile(path.Join(*rootdir, *outputDir, *outputFile))
	error.When(err).Exit(3)
	outputConnection.Info(output, logcode)

	timing.LogSnapshot(stepname, output, logcode+14).Save(stepname)

	stepname = "Step: Write data to output file"
	writer := csv.NewWriter(outputConnection, profile)

	size, err := writer.Start(output)
	error.When(err).Exit(4)
	timing.LogSnapshot(stepname, output, logcode+15).Save(stepname)

	output.Info(logcode, "Writing total %d bytes", size)

	output.Info(logcode, "%s", timing.StartTime)
	timing.LogAll("Step: Summary total usage", output, logcode+100)
}
