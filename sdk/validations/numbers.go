package validations

import (
	"fmt"
)

func ValidateIsPositive(num int64, name string) error {
	if num >= 0 {
		return nil
	}

	return fmt.Errorf("%s must be positive.", name)
}
