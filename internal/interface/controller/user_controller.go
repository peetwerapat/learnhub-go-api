package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/peetwerapat/learnhub-go-api/internal/usecase"
)

type UserController struct {
	userUC *usecase.UserUsecase
}

func NewUserController(r *gin.Engine, uc *usecase.UserUsecase) {
	ctrl := &UserController{userUC: uc}

	r.GET("/users/:id", ctrl.GetUserById)
}

// @Summary Get user by ID
// @Description Get user by ID from path
// @Tags User
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} domain.User
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /users/{id} [get]
func (ctrl *UserController) GetUserById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	user, err := ctrl.userUC.GetUserById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}
