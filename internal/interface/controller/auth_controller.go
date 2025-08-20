package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/peetwerapat/learnhub-go-api/internal/domain"
	"github.com/peetwerapat/learnhub-go-api/internal/interface/controller/dto"
	"github.com/peetwerapat/learnhub-go-api/internal/usecase"
	"github.com/peetwerapat/learnhub-go-api/pkg/response"
	"github.com/peetwerapat/learnhub-go-api/pkg/utils"
)

type AuthController struct {
	userUC *usecase.UserUsecase
}

func NewAuthController(r *gin.Engine, uc *usecase.UserUsecase) {
	ctrl := &AuthController{userUC: uc}

	r.POST("/register", ctrl.CreateUser)
}

// @Summary Create user
// @Description Create a new user
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body dto.CreateUserRequest true "User data"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /register [post]
func (ctrl *AuthController) CreateUser(c *gin.Context) {
	var req dto.CreateUserRequest

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

	if req.Email == "" || req.Password == "" || req.FirstName == "" || req.LastName == "" {
		c.JSON(http.StatusBadRequest, response.BaseHttpResponse{
			StatusCode: http.StatusBadRequest,
			Message: response.Message{
				En: "All fields are required.",
				Th: "กรุณากรอกข้อมูลให้ครบทุกช่อง",
			},
		})
		return
	}

	if !utils.IsValidEmail(req.Email) {
		c.JSON(http.StatusBadRequest, response.BaseHttpResponse{
			StatusCode: http.StatusBadRequest,
			Message: response.Message{
				En: "Invalid email format.",
				Th: "รูปแบบอีเมลไม่ถูกต้อง",
			},
		})
		return
	}

	user := &domain.User{
		Email:     req.Email,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Password:  req.Password,
	}

	if err := ctrl.userUC.CreateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, response.BaseHttpResponse{
			StatusCode: http.StatusInternalServerError,
			Message: response.Message{
				En: "Failed to create user.",
				Th: "สร้างผู้ใช้ไม่สำเร็จ",
			},
		})
		return
	}

	c.JSON(http.StatusOK, response.BaseHttpResponse{
		StatusCode: http.StatusOK,
		Message: response.Message{
			En: "User registered successfully.",
			Th: "สมัครสมาชิกสำเร็จ",
		},
	})
}
