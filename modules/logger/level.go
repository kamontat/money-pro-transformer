package logger

import (
	"strings"

	"github.com/fatih/color"
)

type loggerLevel struct {
	Code  uint8
	Name  string
	Short string
	Color *color.Color
}

func (l loggerLevel) String() string {
	return strings.ToLower(l.Name)
}

var (
	// DEBUG is lowest level in logging
	DEBUG loggerLevel = loggerLevel{
		Code:  uint8(4),
		Name:  "Debug",
		Short: "DBG",
		Color: color.New(color.Italic),
	}

	// INFO is common message
	INFO loggerLevel = loggerLevel{
		Code:  uint8(3),
		Name:  "Info",
		Short: "INF",
		Color: color.New(),
	}

	// WARN is a calm error message
	WARN loggerLevel = loggerLevel{
		Code:  uint8(2),
		Name:  "Warn",
		Short: "WRN",
		Color: color.New(color.Bold).Add(color.FgHiYellow),
	}

	// ERROR is a critical error message
	ERROR loggerLevel = loggerLevel{
		Code:  uint8(1),
		Name:  "Error",
		Short: "ERR",
		Color: color.New(color.Bold).Add(color.FgRed),
	}

	// SILENT is a special level for mute all logging
	SILENT loggerLevel = loggerLevel{
		Code:  uint8(0),
		Name:  "Silent",
		Short: "",
		Color: color.New(),
	}
)
