package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/peetwerapat/learnhub-go-api/internal/domain"
	"github.com/peetwerapat/learnhub-go-api/internal/interface/controller/dto"
	"github.com/peetwerapat/learnhub-go-api/internal/usecase"
	"github.com/peetwerapat/learnhub-go-api/pkg/middleware"
	"github.com/peetwerapat/learnhub-go-api/pkg/oembed"
	"github.com/peetwerapat/learnhub-go-api/pkg/response"
)

type ContentController struct {
	contentUC *usecase.ContentUsecase
}

func NewContentController(r *gin.Engine, uc *usecase.ContentUsecase) {
	ctrl := &ContentController{contentUC: uc}

	r.POST("/contents", middleware.AuthMiddleware(), ctrl.CreateContent)
}

// @Summary Create Content
// @Description Create a new content
// @Tags Content
// @Accept json
// @Produce json
// @Param content body dto.CreateContentRequest true "Content data"
// @Success 200         {object} response.BaseHttpResponse
// @Failure 400,500 {object} response.BaseHttpResponse
// @Security     ApiKeyAuth
// @Router /contents [post]
func (ctrl *ContentController) CreateContent(c *gin.Context) {
	var req dto.CreateContentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.BaseHttpResponse{
			StatusCode: http.StatusBadRequest,
			Message: response.Message{
				En: "Invalid input",
				Th: "ข้อมูลไม่ถูกต้อง",
			},
		})
		return
	}

	userIdVal, ok := c.Get("userId")
	if !ok {
		c.JSON(http.StatusUnauthorized, response.BaseHttpResponse{
			StatusCode: http.StatusUnauthorized,
			Message: response.Message{
				En: "Unauthorized",
				Th: "ไม่มีสิทธิ์",
			},
		})
		return
	}

	userId := userIdVal.(int)

	oembedData, err := oembed.GetOEmbedInfo(req.VideoUrl)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.BaseHttpResponse{
			StatusCode: http.StatusBadRequest,
			Message: response.Message{
				Th: "ลิงก์วิดีโอไม่ถูกต้อง",
				En: "Invalid video URL",
			},
		})
		return
	}

	content := &domain.Content{
		UserID:       userId,
		VideoUrl:     req.VideoUrl,
		Comment:      req.Comment,
		Rating:       req.Rating,
		VideoTitle:   oembedData.Title,
		ThumbnailUrl: oembedData.ThumbnailURL,
		CreatorName:  oembedData.AuthorName,
	}

	if err := ctrl.contentUC.CreateContent(content); err != nil {
		c.JSON(http.StatusInternalServerError, response.BaseHttpResponse{
			StatusCode: http.StatusInternalServerError,
			Message: response.Message{
				En: "Failed to create content",
				Th: "สร้างคอนเทนต์ไม่สำเร็จ",
			},
		})
		return
	}

	c.JSON(http.StatusOK, response.BaseHttpResponse{
		StatusCode: http.StatusOK,
		Message: response.Message{
			En: "Content created successfully",
			Th: "สร้างคอนเทนต์สำเร็จ",
		},
	})
}
