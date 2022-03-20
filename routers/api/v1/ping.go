package v1

import (
	"github.com/gin-gonic/gin"
)

type Pong struct {
	Code    int    `json:"code" example:"200"`
	Message string `json:"message" example:"Pong"`
}
// HTTPError example
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

// Ping godoc
// @Summary Ping
// @Description response to ping
// @Tags api/v1
// @Accept  json
// @Produce  json
// @Success 200 {object} Pong
// @Failure 400 {object} HTTPError
// @Failure 404 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /ping [get]
func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong! (DoctorX API Server is up running.)",
	})
}