package csv

import (
	"encoding/csv"
	"strconv"
	"strings"

	connection "moneypro.kamontat.net/connection-common"
	logger "moneypro.kamontat.net/utils-logger"
)

// Reader is
type Reader struct {
	log       *logger.Logger
	connector connection.Interface
}

func (r *Reader) Read() ([]map[string]string, error) {
	csvLines, err := csv.NewReader(r.connector.GetReader()).ReadAll()
	if err != nil {
		return nil, err
	}

	r.log.Debug(logcode, "Loading csv size: %d line", len(csvLines)-1)

	var nameMapper []string
	var result []map[string]string
	for index, line := range csvLines {
		if index == 0 {
			nameMapper = line
		} else {
			mapper := make(map[string]string)

			mapper["Index"] = strconv.Itoa(index)
			for colIndex, col := range line {
				keyname := nameMapper[colIndex]
				value := strings.TrimSpace(col)

				if keyname != "" {
					mapper[keyname] = value
				}
			}

			result = append(result, mapper)
		}
	}

	return result, nil
}

// NewReader will use input connector as csv format and mapping to map string
func NewReader(connector connection.Interface, output *logger.Logger) *Reader {
	return &Reader{
		log:       output,
		connector: connector,
	}
}
