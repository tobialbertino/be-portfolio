package exception

import "fmt"

// Custom wrap error
type AuthorizationError struct {
	Code    int
	Message string
}

func (w *AuthorizationError) Error() string {
	return fmt.Sprintf(`%v: %v `, w.Code, w.Message)
}

func NewAuthorizationError(message string, code ...int) *AuthorizationError {
	codeStatus := 403
	if len(code) != 0 {
		codeStatus = code[0]
	}
	return &AuthorizationError{
		Code:    codeStatus,
		Message: message,
	}
}
