package module

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var context *gin.Context

// Router gin
func Router() *gin.Engine {
	r := gin.New()

	// Global middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	return r
}

// SetContext format
func SetContext(c *gin.Context) {
	context = c
}

// Error format
func Error(name string, msg string) {
	context.JSON(http.StatusBadRequest, gin.H{
		"panic": false,
		"name":  name,
		"msg":   msg,
	})
}

// Panic format
func Panic(name string, msg string) {
	context.JSON(http.StatusInternalServerError, gin.H{
		"panic": true,
		"name":  name,
		"msg":   msg,
	})
}

// Output format
func Output(data interface{}) {
	context.JSON(http.StatusOK, data)
}
