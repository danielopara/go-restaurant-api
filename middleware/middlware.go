package middleware

import (
	"net/http"
	"strings"

	"github.com/danielopara/restaurant-api/claims"
	"github.com/danielopara/restaurant-api/models"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc{
	return func(c *gin.Context){
		authHeader := c.GetHeader("Authorization")
		if authHeader == ""{
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization required"})
			c.Abort()
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer"{
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization must start with Bearer"})
			c.Abort()
			return 
		}  

		token := tokenParts[1]
		claims, err := claims.ParseToken(token)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or Expired token"})
			c.Abort()
			return 
		}

		c.Set("userID", claims.ID)
        c.Set("email", claims.Email)
        c.Set("role", claims.Role)

        c.Next()
	}
}

func RoleMiddleWare(allowedRoles ...models.Role) gin.HandlerFunc{
	return func (c *gin.Context)  {
		role, exists := c.Get("role")

		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "role not found"})
			c.Abort()
			return
		}

		userRole, ok := role.(models.Role)
		if !ok{
			c.JSON(http.StatusUnauthorized, gin.H{"error" : "invalid role type"})
			c.Abort()
			return 
		}
		for _, r :=  range allowedRoles{
			if userRole == r{
				c.Next()
				return 
			}
		}
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user does not have access"})
		c.Abort()
	}
}