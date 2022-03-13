package auth

import (
	"fmt"
	"log"
	"site/api/config/env"
	"site/api/internal/logger"
	"time"

	ctrl "site/api/pkg/controllers/user"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var identityKey = "id"

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func AuthHandler(logger logger.Logger, db *gorm.DB) *jwt.GinJWTMiddleware {
	userCtrl, err := ctrl.New(ctrl.WithGorm(db))
	if err != nil {
		log.Fatal(fmt.Scanf("User Controller error in AuthHandler: %w", err.Error()))
		return nil
	}
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       env.GetDefault("SECRET_REALM", "test zone"),
		Key:         []byte(env.GetDefault("SECRET_KEY", "secret key")),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*User); ok {
				return jwt.MapClaims{
					identityKey: v.UserName,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &User{
				UserName: claims[identityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			userID := loginVals.Username
			password := loginVals.Password

			user, err := userCtrl.Login(userID, password)
			if err != nil {
				return nil, fmt.Errorf("%w: %s", jwt.ErrFailedAuthentication, err)
			}
			return User{
				UserName: user.Username,
				FullName: user.Fullname,
				Role:     Roles(user.Role),
			}, nil
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*User); ok && v.Role == ADMIN {
				return true
			}

			return false
		},
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

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
		return nil
	}
	return authMiddleware
}
