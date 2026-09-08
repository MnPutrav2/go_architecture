package httperror

import (
	"errors"
	"fmt"
)

var (
	ErrIsUnauthorization = errors.New("Unauthorization")
)

func Custom(message string) error {
	return fmt.Errorf("%s", message)
}
