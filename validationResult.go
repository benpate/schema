package schema

type ValidationResult struct {
	Errors []ValidationError
}

type ValidationError struct {
	Path    string
	Message string
}

func NewValidationResult() *ValidationResult {
	return &ValidationResult{
		Errors: make([]ValidationError, 0),
	}
}

func (vr *ValidationResult) Valid() bool {
	return len(vr.Errors) == 0
}

func (vr *ValidationResult) Add(path string, message string) {

	err := ValidationError{
		Path:    path,
		Message: message,
	}

	vr.Errors = append(vr.Errors, err)
}

func (vr *ValidationResult) Combine(other ValidationResult) {

	for _, err := range other.Errors {
		vr.Errors = append(vr.Errors, err)
	}
}
