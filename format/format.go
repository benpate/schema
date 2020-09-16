package format

// FormatFunc is a function that takes an optional parameter and generates a StringFormat function
type FormatFunc func(string) StringFormat

// StringFormat verifies that a string matches the desired format, and returns a non-nil error if it does not.
type StringFormat func(string) error
