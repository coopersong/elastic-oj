package common

// Error ...
type Error struct {
	statusCode int
	reason     string
}

// NewError ...
func NewError(statusCode int, reason string) *Error {
	return &Error{
		statusCode: statusCode,
		reason:     reason,
	}
}

// Error ...
func (e *Error) Error() string {
	return e.reason
}

// StatusCode ...
func (e *Error) StatusCode() int {
	return e.statusCode
}
