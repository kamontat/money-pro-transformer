# Money Pro Tranformer

This is a transformer CSV data from Money Pro application to support Tablau datasource

- [User](#user)
  - [Usage](#usage)
- [Contributors](#contributors)
  - [Modules](#modules)
    - [Core](#core)
    - [Repository](#repository)
    - [Model](#model)
    - [Utility](#utility)

## User

I provide 2 choice for user to run this application

1. Run via compiled application, you will find the list of support machine in [execution](resources/execution) path.
2. Run via go command, you might install golang in your local machine first in order to run the application. I have a script to run in [scripts](scripts/run.sh) directory.

### Usage

You can run with this syntax (assume you using compiled option)

```bash
money-pro-transformer \
  [-rootDir=<ROOT_DIR>] \       # Root dir is root directory of all path to be resolve in application (default = $pwd)
  [-inputDir=<INPUT_DIR>] \     # input dir is optional dir in root that input file exist (default = <empty>)
  [-inputFile=<INPUT_FILE>] \   # input file is datasource file name (default = 'test.csv')
  [-outputDir=<OUTPUT_DIR>] \   # output dir is optional dir in root that result csv will write to (default = <input_dir>)
  [-outputFile=<OUTPUT_FILE>] \ # output file is output file name (default = new-<input_file>)
  [-debug] \                    # print debug message and timing measurement
  [-silent] \                   # print nothing
```

## Contributors

Since Golang support formatter, lint, and many more by it self so we no need to install anything.
In order to run application you can use [scripts/run.sh](scripts/run.sh) script for run terminal or use [.vscode/launch.json](.vscode/launch.json) run VSCode with debug supported.

### Modules

All module can be separate 4 types. [core, repository, model, utility]

#### Core

Currently, This project has only 1 core application which is [cli](modules/cli). This contains main commandline interface for run application via terminal

#### Repository

Repository is a module that relate to files or databases. This contains 2 major module and 1 minor module.

1. [Datasource](modules/datasource) - File loader and use model transformer to transform raw data format to golang struct
2. [Writer](modules/writer) - File writer enable support logging, error, and timing from utilities modules
3. [CSV](modules/csv) - The csv utilities for writer in csv format

#### Model

Model is a golang struct of input data for datasource to transform to. This contains 1 core module and 2 minor.

1. [Models](modules/models) - This is core model with only a few of api and logic
2. [Transaction](modules/transaction) - A transaction models and transformer
3. [Currency](modules/currency) - A Currency models and transformer

#### Utility

Utility function using on each modules. This contains 1 common module and 3 minor.

1. [Common](modules/utils) - This is a common utilities, this should NOT depend on anything
2. [Logger](modules/logger) - Logging module for log inside application
3. [Error](modules/error) - Error management module for handle any kind of error
4. [Measure](modules/measure) - Measurement module for timing application speed
