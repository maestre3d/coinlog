package view

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
	Items         any    `json:"items"`
	Count         int    `json:"count"`
	NextPageToken string `json:"next_page_token"`
}
