package help

import (
	"fmt"

	"github.com/go-errors/errors"
)

// Stacktrace return void and track the error
func Stacktrace(err error) {
	fmt.Println(err.(*errors.Error).ErrorStack())
}
