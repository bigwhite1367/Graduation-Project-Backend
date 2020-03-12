package controllers

import(
	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	models "ncu_14_project/backend/Models"
	_ "ncu_14_project/backend/Database"
)


func GetUser(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	user, _ := c.Get(identityKey)

	
	c.JSON(200, gin.H{
		"userID":   claims[identityKey],
		"userName": user.(*models.User).UserName,
		"privilege": user.(*models.User).Privilege,
	})
}