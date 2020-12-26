package logger

import "fmt"

// Logger is logging object
type Logger struct {
	level loggerLevel
}

// SetLevel will update current logging level
func (l *Logger) SetLevel(level loggerLevel) {
	l.level = level
}

// IsDebug will return true if on debug mode
func (l *Logger) IsDebug() bool {
	return l.CheckLevel(DEBUG)
}

// CheckLevel is input level should be log or not
func (l *Logger) CheckLevel(level loggerLevel) bool {
	if level.Code <= l.level.Code {
		return true
	}

	return false
}

func (l *Logger) private(level loggerLevel, key int, format string, params ...interface{}) {
	if l.CheckLevel(level) {
		fullFormat := "%s[%05d] " + format + "\n"

		newParams := make([]interface{}, 0)
		newParams = append(newParams, level.Short)
		newParams = append(newParams, key)
		for _, p := range params {
			newParams = append(newParams, p)
		}

		level.Color.Printf(fullFormat, newParams...)
	}
}

// Debug is logging as debug message
func (l *Logger) Debug(key int, format string, params ...interface{}) {
	l.private(DEBUG, key, format, params...)
}

// Time is logging as debug message with time key
func (l *Logger) Time(key int, name string, duration string) {
	l.private(TIME, key, "%-35s: %s", name, duration)
}

// Info is logging as info message
func (l *Logger) Info(key int, format string, params ...interface{}) {
	l.private(INFO, key, format, params...)
}

// Warn is logging as warn message
func (l *Logger) Warn(key int, format string, params ...interface{}) {
	l.private(WARN, key, format, params...)
}

// Error is logging as error message
func (l *Logger) Error(key int, format string, params ...interface{}) {
	l.private(ERROR, key, format, params...)
}

// Log will logging data without any formatting
func (l *Logger) Log(msg string) {
	fmt.Println(msg)
}

var logger *Logger = &Logger{
	level: INFO,
}

// Get will return singleton logger object
func Get() *Logger {
	return logger
}
