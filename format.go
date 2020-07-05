package schema

// StringFormat verifies that a string matches the desired format, and returns a non-nil error if it does not.
type StringFormat func(string) error
