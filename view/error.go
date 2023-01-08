package view

// ErrorMessage system error template based on RFC 7807 specification (ref: https://www.rfc-editor.org/rfc/rfc7807).
type ErrorMessage struct {
	// Code HTTP status code.
	Code int `json:"code"`
	// ErrorType name of error type cataloged by the system (e.g. ResourceNotFound, MissingParameter).
	//
	// Might be used by external systems to change behaviour at runtime.
	ErrorType string `json:"error_type"`
	// Message description or title of error type.
	Message string `json:"message"`
}

// ErrorsMessage collection of ErrorMessage(s).
type ErrorsMessage struct {
	Errors []ErrorMessage `json:"errors"`
	// Code top-level error code from inner Errors (e.g. 400,404,502 -> 502).
	Code int `json:"code"`
}
