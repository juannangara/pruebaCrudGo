package handler

import (
	"net/http"
	service "purebaCrud/API/Service"

	"github.com/gin-gonic/gin"
)

type AuthHandler interface {
	Login(c *gin.Context)
}

type authHandler struct {
}

func CreateAuthHandler() AuthHandler {
	return &authHandler{}
}

func (h authHandler) Login(c *gin.Context) {

	token, err := service.CreateToken("user")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, token)

}

func authToken(c *gin.Context) error {
	token := c.Request.Header.Values("Authorization")

	return service.AuthToken(token[0])

}
