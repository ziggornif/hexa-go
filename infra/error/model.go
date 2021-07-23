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
	if len(err.Kind) == 0 {
		err.Kind = "Error"
	}
	return fmt.Sprintf("%s - %s", err.Kind, err.Error.Error())
}
