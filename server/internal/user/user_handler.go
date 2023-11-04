package user

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"server/util"
)

type Handler struct {
	Service
	db *sql.DB
}

func NewHandler(s Service) *Handler {
	return &Handler{Service: s}
}

func (h *Handler) CreateUser(c *gin.Context) {
	var userRequest CreateUserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//validate email
	validEmail, err := util.IsEmailValid(userRequest.Email)
	if !validEmail {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//validate username
	validUsername, err := util.IsUsernameValid(userRequest.Username)
	if !validUsername {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userResponse, err := h.Service.CreateUser(c.Request.Context(), &userRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, userResponse)
}

func (h *Handler) Login(c *gin.Context) {
	var user LoginUserRequest

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userResponse, err := h.Service.LoginUser(c.Request.Context(), &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie(
		"access_token",
		userResponse.accessToken,
		3600,
		"/",
		"localhost",
		false,
		true)

	res := &LoginUserResponse{
		Username: userResponse.Username,
		ID:       userResponse.ID,
	}
	c.JSON(http.StatusOK, res)
}

func (h *Handler) Logout(c *gin.Context) {
	c.SetCookie("access_token", "",
		-1,
		"/",
		"localhost",
		false,
		true)

	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}
