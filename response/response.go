package response

import "github.com/gin-gonic/gin"

func ErrorResponse(res *gin.Context, statusCode int, message string) {
	res.JSON(statusCode, gin.H{"error": message})
}

func SuccessResponse(res *gin.Context, status int, result any) {
	res.JSON(status, gin.H{"result": result})
}
