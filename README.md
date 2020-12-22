# Money Pro Tranformer

This is a transformer CSV data from Money Pro application to support Tablau datasource

## User

I provide 2 choice for user to run this application

1. Run via compiled application, you will find the list of support machine in [execution](resources/execution) path.
2. Run via go command, you might install golang in your local machine first in order to run the application. I have a script to run in [scripts](scripts/run.sh) directory.

### Usage

You can run with this syntax (assume you using compiled option)

```bash
money-pro-transformer \
  [-rootDir=<ROOT_DIR>] \     # Root dir is root directory of all path to be resolve in application (default = $pwd)
  [-inputDir=<INPUT_DIR>] \   # input dir is optional dir in root that input file exist (default = <empty>)
  [-inputFile=<INPUT_FILE>] \ # input file is datasource file name (default = 'test.csv')
  [-outputDir=<OUTPUT_DIR>] \ # output dir is optional dir in root that result csv will write to (default = <input_dir>)
  [-outputFile=<OUTPUT_FILE>] # output file is output file name (default = new-<input_file>)
```

## Contributors

Since Golang support formatter, lint, and many more by it self so we no need to install anything.
In order to run application you can use [scripts/run.sh](scripts/run.sh) script for run terminal or use [.vscode/launch.json](.vscode/launch.json) run VSCode with debug supported.

### Modules

1. [utils](modules/utils) - Utilities module full with helper function
2. [models](modules/models) - Global models for several modules
3. [datasource](modules/datasource) - Input data loader and transformer to models via mapping
4. [writer](modules/writer) - Output writer for write data to file
5. [cli](modules/cli) - Main commandline interface for run application
