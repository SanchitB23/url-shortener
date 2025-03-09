package middlewares

import (
	"github.com/gin-gonic/gin"
)

var userIdStr = "userId"

func AuthMiddleware(context *gin.Context) {
	// token := context.GetHeader("Authorization")
	// if token == "" {
	// 	context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing Authorization header"})
	// 	return
	// }
	// userId, err := utils.VerifyJWT(token)
	// if err != nil {
	// 	context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	// 	return
	// }
	//fixme hardcoding the user id for now
	context.Set(userIdStr, int64(999))
	context.Next()
}

func GetCurrentUser(context *gin.Context) int64 {
	return context.GetInt64(userIdStr)
}
