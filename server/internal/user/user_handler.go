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

//paths information

// @summary Create a new user
// @description Create a new user
// @tags user
// @accept json
// @produce json
// @param userRequest body CreateUserRequest true "Create User Request"
// @success 200 {object} CreateUserResponse
// @router /signup [post]
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
	////check if email is already in use
	//notUsedEmail, _ := util.IsEmailExist(h.db, userRequest.Email)
	//if !notUsedEmail {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": "email already in use"})
	//	return
	//}
	////check if username is already in use
	//notUsedUsername, _ := util.IsUsernameExist(h.db, userRequest.Username)
	//if !notUsedUsername {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": "username already in use"})
	//	return
	//}

	userResponse, err := h.Service.CreateUser(c.Request.Context(), &userRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, userResponse)
}

// @summary Login user
// @description Login user
// @tags user
// @accept json
// @produce json
// @param userRequest body LoginUserRequest true "Login User Request"
// @success 200 {object} LoginUserResponse
// @router /login [post]
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

// @summary Logout user
// @description Logout user
// @tags user
// @accept json
// @produce json
// @success 200 {object} string
// @router /logout [get]
func (h *Handler) Logout(c *gin.Context) {
	c.SetCookie("access_token", "",
		-1,
		"/",
		"localhost",
		false,
		true)

	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}
