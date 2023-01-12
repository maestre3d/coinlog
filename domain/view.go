package domain

type ViewList struct {
	Items         any    `json:"items"`
	Count         int    `json:"count"`
	NextPageToken string `json:"next_page_token"`
}
