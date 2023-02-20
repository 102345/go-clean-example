package domain

type Page struct {
	Total      int32 `json:"total"`
	PageNumber int32 `json:"pageNumber"`
	Quantity   int32 `json:"quantity"`
}
