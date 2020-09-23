package schema

import (
	"github.com/benpate/derp"
	"github.com/benpate/list"
)

// ValidationErrorCode represents HTTP Status Code: 422 "Unproccessable Entity"
const ValidationErrorCode = 422

// ValidationError represents an input validation error, and includes fields necessary to
// report problems back to the end user.
type ValidationError struct {
	Path    string `json:"path"`    // Identifies the PATH (or variable name) that has invalid input
	Message string `json:"message"` // Human-readable message that explains the problem with the input value.
}

// Invalid returns a fully populated ValidationError to the caller
func Invalid(path string, message string) *ValidationError {
	return &ValidationError{
		Path:    path,
		Message: message,
	}
}

// Error returns a string representation of this ValidationError, and implements
// the builtin errors.error interface.
func (v *ValidationError) Error() string {
	return v.Message
}

// ErrorCode returns CodeValidationError for this ValidationError
// It implements the ErrorCodeGetter interface.
func (v *ValidationError) ErrorCode() int {
	return ValidationErrorCode
}

// Rollup bundles a child error into a parent
func Rollup(parent *derp.MultiError, child *derp.MultiError, path string) *derp.MultiError {

	if child != nil {

		for _, err := range *child {

			// If the child error is nil for some reason, then skip this record.
			if err == nil {
				continue
			}

			if validationError, ok := err.(*ValidationError); ok {
				validationError.Path = list.PushHead(validationError.Path, path, ".")
			}

			parent = derp.Append(parent, err)
		}

	}

	return parent
}
