package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/peetwerapat/learnhub-go-api/internal/domain"
	"github.com/peetwerapat/learnhub-go-api/internal/interface/controller/dto"
	"github.com/peetwerapat/learnhub-go-api/internal/usecase"
	"github.com/peetwerapat/learnhub-go-api/pkg/oembed"
	"github.com/peetwerapat/learnhub-go-api/pkg/response"
)

type ContentController struct {
	contentUC *usecase.ContentUsecase
}

func NewContentController(uc *usecase.ContentUsecase) *ContentController {
	return &ContentController{contentUC: uc}
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

// @Summary Get Contents
// @Description Get contents
// @Tags Content
// @Accept json
// @Produce json
// @Success 200     {object} []dto.ContentResponse
// @Failure 400,500 {object} response.BaseHttpResponse
// @Router /contents [get]
func (ctrl *ContentController) GetContents(c *gin.Context) {
	contents, err := ctrl.contentUC.GetContents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.BaseHttpResponse{
			StatusCode: http.StatusInternalServerError,
			Message: response.Message{
				En: "Failed to fetch content",
				Th: "ดึงข้อมูลคอนเทนต์ไม่สำเร็จ",
			},
		})
		return

	}

	var contentResponses []dto.ContentResponse

	for _, content := range contents {
		user := dto.UserResponse{
			ID:    content.User.ID,
			Email: content.User.Email,
		}

		contentResponses = append(contentResponses, dto.ContentResponse{
			ID:           content.ID,
			VideoTitle:   content.VideoTitle,
			VideoUrl:     content.VideoUrl,
			Comment:      content.Comment,
			Rating:       content.Rating,
			ThumbnailUrl: content.ThumbnailUrl,
			CreatorName:  content.CreatorName,
			User:         user,
		})
	}

	c.JSON(http.StatusOK, response.HttpResponse[[]dto.ContentResponse]{
		BaseHttpResponse: response.BaseHttpResponse{
			StatusCode: http.StatusOK,
			Message: response.Message{
				En: "Fetched contents successfully",
				Th: "ดึงข้อมูลคอนเทนต์สำเร็จ",
			},
		},
		Data: contentResponses,
	})
}

// @Summary Get Content By Id
// @Description Get content by id
// @Tags Content
// @Accept json
// @Produce json
// @Param   id      path      int  true  "Content ID"
// @Success 200     {object} dto.ContentResponse
// @Failure 400,500 {object} response.BaseHttpResponse
// @Router /contents/{id} [get]
func (ctrl *ContentController) GetContenById(c *gin.Context) {
	id := c.Param("id")
	content, err := ctrl.contentUC.GetContentById(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, response.BaseHttpResponse{
			StatusCode: http.StatusInternalServerError,
			Message: response.Message{
				En: "Failed to get content",
				Th: "เกิดข้อผิดพลาดในการดึงคอนเทนต์",
			},
		})
		return
	}
	if content == nil {
		c.JSON(http.StatusNotFound, response.BaseHttpResponse{
			StatusCode: http.StatusNotFound,
			Message: response.Message{
				En: "Content not found",
				Th: "ไม่พบข้อมูลคอนเทนต์",
			},
		})
		return
	}

	contentResponse := dto.ContentResponse{
		ID:           content.ID,
		VideoTitle:   content.VideoTitle,
		VideoUrl:     content.VideoUrl,
		Comment:      content.Comment,
		Rating:       content.Rating,
		ThumbnailUrl: content.ThumbnailUrl,
		CreatorName:  content.CreatorName,
		User: dto.UserResponse{
			ID:    content.User.ID,
			Email: content.User.Email,
		},
	}

	c.JSON(http.StatusOK, response.HttpResponse[dto.ContentResponse]{
		BaseHttpResponse: response.BaseHttpResponse{
			StatusCode: http.StatusOK,
			Message: response.Message{
				En: "Fetched contents successfully",
				Th: "ดึงข้อมูลคอนเทนต์สำเร็จ",
			},
		},
		Data: contentResponse,
	})

}
