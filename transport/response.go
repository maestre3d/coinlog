package transport

// BasicResponse generic system response template containing plain messages.
type BasicResponse struct {
	Message string `json:"message"`
}

// DataResponse generic system response template for data retrieval.
type DataResponse struct {
	Data any `json:"data"`
}

// ListDataResponse generic system response template for collections of data retrieval.
type ListDataResponse struct {
	Data          any    `json:"data"`
	Count         int    `json:"count"`
	NextPageToken string `json:"next_page_token"`
}

// ErrorResponse system error template based on RFC 7807 specification (ref: https://www.rfc-editor.org/rfc/rfc7807)
// and Google Cloud's API Design whitepaper (ref: https://cloud.google.com/apis/design/errors).
type ErrorResponse struct {
	// Code HTTP status code.
	Code int `json:"code"`
	// ErrorStatus name of error type cataloged by the system (e.g. ResourceNotFound, MissingParameter).
	//
	// Might be used by external systems to change behaviour at runtime.
	ErrorStatus string `json:"status"`
	// Message description or title of error type.
	Message string `json:"message"`
}

// ErrorsResponse collection of ErrorResponse(s).
type ErrorsResponse struct {
	Errors []ErrorResponse `json:"errors"`
	// Code top-level error code from inner Errors (e.g. 400,404,502 -> 502).
	Code int `json:"code"`
}
