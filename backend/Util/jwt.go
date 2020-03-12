package util

import (
	"time"
	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	controller "ncu_14_project/backend/Controllers"
)
var identityKey = "id"

var AuthMiddleware, err = jwt.New(&jwt.GinJWTMiddleware{
	Realm:       "test zone",
	Key:         []byte("secret key"),
	Timeout:     time.Hour,
	MaxRefresh:  time.Hour,
	IdentityKey: identityKey,
	PayloadFunc: controller.PayloadFunc,
	IdentityHandler: controller.IdentityHandler,
	Authenticator: controller.Authenticator,
	Authorizator: controller.Authorizator,
	Unauthorized: func(c *gin.Context, code int, message string) {
		c.JSON(code, gin.H{
			"code":    code,
			"message": message,
		})
	},
	// TokenLookup is a string in the form of "<source>:<name>" that is used
	// to extract token from the request.
	// Optional. Default value "header:Authorization".
	// Possible values:
	// - "header:<name>"
	// - "query:<name>"
	// - "cookie:<name>"
	// - "param:<name>"
	TokenLookup: "header: Authorization, query: token, cookie: jwt",
	// TokenLookup: "query:token",
	// TokenLookup: "cookie:token",

	// TokenHeadName is a string in the header. Default value is "Bearer"
	TokenHeadName: "Bearer",

	// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
	TimeFunc: time.Now,
})
