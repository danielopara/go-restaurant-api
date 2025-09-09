package user

import (
	"net/http"

	"github.com/danielopara/restaurant-api/models"
	"github.com/danielopara/restaurant-api/request"
	"github.com/danielopara/restaurant-api/response"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService UserService
}

func NewUserHandler(userService UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) GetLoggedInUser( c *gin.Context){
	userEmail, exists := c.Get("email")
	if !exists{
		c.JSON(http.StatusUnauthorized, gin.H{"error": "no logged in user found"})
		return
	}

	user, err := h.userService.GetUserByEmail(userEmail.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req := response.UserResponse{
		FirstName: user.FirstName,
		LastName: user.LastName,
		Email: user.Email,
		Role: user.Role,
	}
	c.JSON(http.StatusOK, gin.H{"data": req})
}

func (h *UserHandler) GetAllUsers(c *gin.Context){
	users,  err := h.userService.GetAllUsers()
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

func (h *UserHandler) Register(c *gin.Context) {
	var req request.Register

	if err := c.ShouldBindJSON(&req); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := &models.User{
		FirstName: req.FirstName,
		LastName: req.LastName,
		Role: models.Role(req.Role),
		Email: req.Email,
		Password: req.Password,
	}
	res,  err := h.userService.RegisterUser(user)

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


	c.JSON(http.StatusCreated, gin.H{"data": res})
}