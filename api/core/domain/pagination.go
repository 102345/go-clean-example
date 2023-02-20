package domain

// Pagination is representation of Fetch methods returns products
type Pagination struct {
	Items interface{} `json:"products"`
	Page  interface{} `json:"page"`
}

// Pagination is representation of Fetch methods returns users
type PaginationUsers struct {
	Users interface{} `json:"users"`
	Page  interface{} `json:"page"`
}
