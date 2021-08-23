package zz

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJoinParamsWithDash(t *testing.T) {
	assert := assert.New(t)

	// Type used to hold function parameters and expected results.
	type testData struct {
		param1         string
		param2         int
		expectedResult string
		expectError    bool
	}

	// List of tests to run including the expected results
	data := []testData{
		// Failure scenarios
		{"", -1, "", true},
		{"", 0, "", true},
		{"", 1, "", true},
		{"foo", 0, "", true},
		{"foo", -1, "", true},

		// Success scenarios
		{"foo", 1, "foo-1", false},
		{"bar", 42, "bar-42", false},
	}

	// Run the tests
	for i, d := range data {
		// Create a test-specific string that is added to each assert
		// call. It will be displayed if any assert test fails.
		msg := fmt.Sprintf("test[%d]: %+v", i, d)

		// Call the function under test
		result, err := joinParamsWithDash(d.param1, d.param2)

		// update the message for more information on failure
		msg = fmt.Sprintf("%s, result: %q, err: %v", msg, result, err)

		if d.expectError {
			assert.Error(err, msg)

			// If an error is expected, there is no point
			// performing additional checks.
			continue
		}

		assert.NoError(err, msg)
		assert.Equal(d.expectedResult, result, msg)
	}
}
