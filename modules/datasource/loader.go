package datasource

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	models "kamontat.net/money-pro-models"
)

// Loader will load csv and convert to transaction struct
func Loader(filename string) (*models.Accounts, error) {
	csvFile, err := os.Open(filename)
	defer csvFile.Close()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var nameMapper []string
	accounts := models.NewAccounts()

	for index, line := range csvLines {
		if index == 0 {
			nameMapper = line
		} else {
			mapper := make(map[string]string)

			mapper["Index"] = string(index)
			for colIndex, col := range line {
				keyname := nameMapper[colIndex]
				value := strings.TrimSpace(col)

				if keyname != "" {
					mapper[keyname] = value
				}
			}

			transaction, err := models.NewTransaction(mapper)
			if err == nil {
				accounts.AddTransaction(transaction)
			} else {
				fmt.Println(err)
			}
		}
	}

	return accounts, nil
}
