package errors

import (
	"fmt"
	"os"
)

// ErrorStat general struct error type.
// this struct can be used the code that is not
// a specific http/api/protocol error.
// A more generic error type
type ErrorStat struct {
	Message string // the actual error message
}

// struct ErrorStat implements Stringer interface
// Now we can use this with the fmt package.
func (e ErrorStat) String() string {
	return fmt.Sprintf("[ > ]\tError : %s\n", e.Message)
}

// Print func it's just a util function that we can use
// independently as it is.
func (e ErrorStat) Print() {
	fmt.Fprintf(os.Stderr, "[ > ]\t %s\n", e.Message)
}

// struct ErrorStat implements Error interface
// Now we can use this with return error types.
func (e ErrorStat) Error() string {
	return fmt.Sprintf("[ > ]\tError : %s\n", e.Message)
}
