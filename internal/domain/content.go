package domain

import "time"

type Content struct {
	ID           uint       `gorm:"column:id;primaryKey" json:"id"`
	VideoTitle   string     `gorm:"column:video_title" json:"videoTitle"`
	VideoUrl     string     `gorm:"column:video_url" json:"videoUrl"`
	Comment      string     `gorm:"column:comment" json:"comment"`
	Rating       int        `gorm:"column:rating" json:"rating"`
	ThumbnailUrl string     `gorm:"column:thumbnail_url" json:"thumbnailUrl"`
	CreatorName  string     `gorm:"column:creator_name" json:"creatorName"`
	UserID       int        `gorm:"column:user_id"`
	User         User       `gorm:"foreignKey:UserID" json:"user"`
	CreatedAt    time.Time  `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt    time.Time  `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt    *time.Time `gorm:"column:deleted_at" json:"deletedAt"`
}

func (Content) TableName() string {
	return "TB_CONTENT"
}
