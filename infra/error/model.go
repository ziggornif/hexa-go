package error

import (
	"fmt"
)

// HexagoError - error model
type HexagoError struct {
	Kind  string
	Error error
}

func (err HexagoError) String() string {
	return fmt.Sprintf("%s Error - %s", err.Kind, err.Error.Error())
}
