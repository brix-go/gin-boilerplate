package auth

import (
	"GinBoilerplate/config"
	"GinBoilerplate/internal/domain/user"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
	"net/http"
	"strings"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		// c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		token, err := tokenValidation(c.Request)
		var resp map[string]interface{}
		if err != nil {
			resp = map[string]interface{}{"code": "4010", "msg": "You need to be authorized to access this route"}
			c.JSON(http.StatusUnauthorized, resp)
			c.Abort()
			return
		}
		c.Set("JWT", token)
		c.Next()
	}
}

func GetCredential(c *gin.Context) {
	token, _ := c.Get("JWt")
	tokenParse, err := jwt.Parse(token.(string), func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", jwtToken.Header["alg"])
		}
		return []byte(config.Config.AppConfig.Secret), nil
	})
	if err != nil {
		resp := map[string]interface{}{"code": "4010", "msg": "You need to be authorized to access this route"}
		c.JSON(http.StatusUnauthorized, resp)
		c.Abort()
		return
	}
	claims := tokenParse.Claims.(jwt.MapClaims)
	credential := user.User{
		ID:    claims["id"].(string),
		Email: claims["email"].(string),
	}
	c.Set("credentials", credential)
	c.Next()
}

func tokenValidation(r *http.Request) (string, error) {
	bearerToken := r.Header.Get("Authorization")
	if !strings.Contains(bearerToken, "Bearer") {
		return "", errors.New("unauthorized, not bearer token")
	}
	jwtToken := strings.Replace(bearerToken, "Bearer", "", -1)
	return jwtToken, nil
}
