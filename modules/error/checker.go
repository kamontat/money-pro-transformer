package error

import (
	"os"

	logger "moneypro.kamontat.net/utils-logger"
)

// Wrapper enable more function to manage error
type Wrapper struct {
	err error
}

// Exist return true if error exist
func (e *Wrapper) Exist() bool {
	return e.err != nil
}

// Empty return true if error is empty
func (e *Wrapper) Empty() bool {
	return e.err == nil
}

// OnError will run input function will error is exist
func (e *Wrapper) OnError(fn func(error)) *Wrapper {
	if e.Exist() {
		fn(e.err)
	}
	return e
}

// OnCompleted will run input function will error is empty
func (e *Wrapper) OnCompleted(fn func()) *Wrapper {
	if e.Empty() {
		fn()
	}
	return e
}

// Print will print error to console if exist
func (e *Wrapper) Print(output *logger.Logger, key int) *Wrapper {
	return e.OnError(func(err error) {
		output.Error(key, "%s", err)
	})
}

// Panic will print panic message to console
func (e *Wrapper) Panic() *Wrapper {
	return e.OnError(func(err error) {
		panic(err)
	})
}

// Exit will exit if error exist
func (e *Wrapper) Exit(code int) {
	e.OnError(func(err error) {
		os.Exit(code)
	})
}

// When will generate error wrapper for you to do something
func When(err error) *Wrapper {
	return &Wrapper{
		err: err,
	}
}
