package dto

type CreateContentRequest struct {
	VideoUrl string `json:"videoUrl" binding:"required"`
	Comment  string `json:"comment"`
	Rating   int    `json:"rating" binding:"required"`
}
