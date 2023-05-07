package types

import (
	"regexp"
)

func ValidateAlias(name string) error {

	ok, _ := regexp.MatchString("^[a-zA-Z0-9_]*$", name)
	if !ok || len(name) < 3 || len(name) > 10 {
		return ErrInvalidAlias
	}

	return nil
}
