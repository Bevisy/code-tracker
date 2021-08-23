package zz

import (
	"errors"
	"fmt"
)

// The function under test.
//
// Accepts a string and an integer and returns the
// result of sticking them together separated by a dash as a string.
func joinParamsWithDash(str string, num int) (string, error) {
	if str == "" {
		return "", errors.New("string cannot be blank")
	}

	if num <= 0 {
		return "", errors.New("number must be positive")
	}

	return fmt.Sprintf("%s-%d", str, num), nil
}
