package response

type Message struct {
	Th string `json:"th"`
	En string `json:"en"`
}

type Pagination struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	TotalItems int `json:"totalItems"`
	TotalPages int `json:"totalPages"`
}

type BaseHttpResponse struct {
	StatusCode int     `json:"statusCode"`
	Message    Message `json:"message"`
}

type HttpResponse[T any] struct {
	BaseHttpResponse
	Data T `json:"data,omitempty"`
}

type HttpResponseWithPagination[T any] struct {
	BaseHttpResponse
	Data       T           `json:"data,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
}
