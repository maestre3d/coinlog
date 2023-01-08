package view

// BasicResponse generic system response template containing plain messages.
type BasicResponse struct {
	Message string `json:"message"`
}

// DataResponse generic system response template for data retrieval.
type DataResponse struct {
	Data any `json:"data"`
}
