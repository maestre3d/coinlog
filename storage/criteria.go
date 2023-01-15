package storage

// Criteria set of parameters used to indicate conditions (e.g. filters, offset, cursor, limit) to a
// Reader -or Repository- when fetching data.
type Criteria struct {
	Limit     int       `json:"limit"`
	PageToken PageToken `json:"page_token"`
}
