package domain

import "net/http"

// StockProductService is a contract of http adapter layer
type StockProductService interface {
	SendMessageStockProduct(response http.ResponseWriter, request *http.Request)
}
