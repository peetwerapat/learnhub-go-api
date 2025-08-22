package controller

import (
	"errors"
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
	r.POST("/login", ctrl.Login)
}

// @Summary Create user
// @Description Create a new user
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body dto.RegisterRequest true "User data"
// @Success 200         {object} response.BaseHttpResponse
// @Failure 400,500 {object} response.BaseHttpResponse
// @Router /register [post]
func (ctrl *AuthController) CreateUser(c *gin.Context) {
	var req dto.RegisterRequest

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

	if len(req.Password) < 6 {
		c.JSON(http.StatusBadRequest, response.BaseHttpResponse{
			StatusCode: http.StatusBadRequest,
			Message: response.Message{
				En: "Password must be at least 6 characters",
				Th: "รหัสผ่านต้องมีอย่างน้อย 6 หลัก",
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

	err := ctrl.userUC.CreateUser(user)
	switch {
	case errors.Is(err, usecase.ErrEmailAlreadyExist):
		c.JSON(http.StatusConflict, response.BaseHttpResponse{
			StatusCode: http.StatusConflict,
			Message: response.Message{
				En: "Email already exists.",
				Th: "อีเมลซ้ำ",
			},
		})
	case err != nil:
		c.JSON(http.StatusInternalServerError, response.BaseHttpResponse{
			StatusCode: http.StatusInternalServerError,
			Message: response.Message{
				En: "Internal server error.",
				Th: "เกิดข้อผิดพลาด",
			},
		})
	default:
		c.JSON(http.StatusOK, response.BaseHttpResponse{
			StatusCode: http.StatusOK,
			Message: response.Message{
				En: "User registered successfully.",
				Th: "สมัครสมาชิกสำเร็จ",
			},
		})
	}
}

// @Summary User login
// @Description User login with username and password
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body dto.LoginRequest true "Username and password"
// @Success 200         {object} response.BaseHttpResponse
// @Failure 400,401,500 {object} response.BaseHttpResponse
// @Router /login [post]
func (ctrl *AuthController) Login(c *gin.Context) {
	var req dto.LoginRequest

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

	token, err := ctrl.userUC.Login(req.Email, req.Password)
	switch {
	case errors.Is(err, usecase.ErrInvalidEmailFormat):
		c.JSON(http.StatusBadRequest, response.BaseHttpResponse{
			StatusCode: http.StatusBadRequest,
			Message: response.Message{
				En: "Invalid email format",
				Th: "รูปแบบอีเมลไม่ถูกต้อง",
			},
		})
	case errors.Is(err, usecase.ErrInvalidEmail):
		c.JSON(http.StatusUnauthorized, response.BaseHttpResponse{
			StatusCode: http.StatusUnauthorized,
			Message: response.Message{
				En: "Invalid email",
				Th: "อีเมลผิด",
			},
		})
	case errors.Is(err, usecase.ErrInvalidPassword):
		c.JSON(http.StatusUnauthorized, response.BaseHttpResponse{
			StatusCode: http.StatusUnauthorized,
			Message: response.Message{
				En: "Invalid password",
				Th: "รหัสผ่านผิด",
			},
		})
	case errors.Is(err, usecase.ErrTokenCreation):
		c.JSON(http.StatusInternalServerError, response.BaseHttpResponse{
			StatusCode: http.StatusInternalServerError,
			Message: response.Message{
				En: "Could not create token",
				Th: "ไม่สามารถออกโทเคนได้",
			},
		})
	case err != nil:
		c.JSON(http.StatusInternalServerError, response.BaseHttpResponse{
			StatusCode: http.StatusInternalServerError,
			Message: response.Message{
				En: "Unexpected error",
				Th: "เกิดข้อผิดพลาด",
			},
		})
	default:
		c.JSON(http.StatusOK, gin.H{
			"statusCode": http.StatusOK,
			"message": response.Message{
				En: "Login successful",
				Th: "เข้าสู่ระบบสำเร็จ",
			},
			"accessToken": token,
		})
	}
}
