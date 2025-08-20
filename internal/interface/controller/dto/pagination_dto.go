package dto

type PaginationQuery struct {
	Page   int    `form:"page"`
	Limit  int    `form:"limit"`
	Search string `form:"search"`
	Sort   string `form:"sort"`
	Order  string `form:"order"`
}
