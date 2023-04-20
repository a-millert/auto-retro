package fail

import (
	"fmt"
	"os"
)

type EnrichedError struct {
	err error
	msg string
}

func New(err error, msg string) *EnrichedError {
	return &EnrichedError{err, msg}
}

func (ee *EnrichedError) Abort() error {
	fmt.Fprintf(os.Stderr, "[ERROR] %s: %s\n", ee.msg, ee.err)
	return ee.err
}
