package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthorizationMiddleware(c *gin.Context) {
	s := c.Request.Header.Get("Authorization")

	token := strings.TrimPrefix(s, "Bearer ")

	if err := validateToken(token); err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}

func validateToken(token string) error {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte("secret"), nil
	})

	return err
}

func LoginHandler(c *gin.Context) {
	// implement login logic here
	// user := c.PostForm("user")
	// pass := c.PostForm("pass")

	// // Throws Unauthorized error
	// if user != "john" || pass != "lark" {
	// 	return c.AbortWithStatus(http.StatusUnauthorized)
	// }

	// Create the Claims
	// claims := jwt.MapClaims{
	// 	"name":  "John Lark",
	// 	"admin": true,
	// 	"exp":   time.Now().Add(time.Hour * 72).Unix(),
	// }

	// Create token
	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
	})

	ss, err := token.SignedString([]byte("secret"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"token": ss,
	})
}
