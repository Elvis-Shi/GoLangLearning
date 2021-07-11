package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"ssdd.com/vms/models"
)

// TODO: find safe place to store the key and read it from there.
var SECRETKEY = "abcdefghijk"

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := strings.Split(c.GetHeader("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			models.Log(&models.TraceLog {
				ServiceId: c.GetString("ServiceId"),
				Level: models.Warning,
				Content: fmt.Sprintf("Invalid header 'Authorization': %v", c.GetHeader("Authorization")),
			})

			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized request"})
			c.Abort()
			return
		}

		jwtToken := authHeader[1]
		token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				models.Log(&models.TraceLog {
					ServiceId: c.GetString("ServiceId"),
					Level: models.Warning,
					Content: fmt.Sprintf("Unexpected signing method: %v", token.Header["alg"]),
				})
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(SECRETKEY), nil
		})

		if err != nil {

			models.Log(&models.TraceLog {
				ServiceId: c.GetString("ServiceId"),
				Level: models.Warning,
				Content: fmt.Sprintf("Invalid authorization: %v. %v", jwtToken, err.Error()),
			})
			
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		if !token.Valid {
			models.Log(&models.TraceLog {
				ServiceId: c.GetString("ServiceId"),
				Level: models.Warning,
				Content: fmt.Sprintf("Invalid token: %v", jwtToken),
			})

			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization."})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid authorization."})
			c.Abort()
			return
		} 
		
		// put clains in context in case any API need it.
		c.Set("Claims", claims)
	}
}
