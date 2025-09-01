package dto

type CreateContentRequest struct {
	VideoUrl string `json:"videoUrl" binding:"required"`
	Comment  string `json:"comment"`
	Rating   int    `json:"rating" binding:"required"`
}

type ContentResponse struct {
	ID           uint         `json:"id"`
	VideoTitle   string       `json:"videoTitle"`
	VideoUrl     string       `json:"videoUrl"`
	Comment      string       `json:"comment"`
	Rating       int          `json:"rating"`
	ThumbnailUrl string       `json:"thumbnailUrl"`
	CreatorName  string       `json:"creatorName"`
	User         UserResponse `json:"user"`
}
