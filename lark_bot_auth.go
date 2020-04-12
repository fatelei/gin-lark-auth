package gin_lark_auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func VerifyLarkBot(token string) gin.HandlerFunc {
	return func(c *gin.Context) {
		data := make(map[string]interface{})
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusOK, gin.H{})
			c.Abort()
			return
		}

		if msgType, ok := data["type"]; ok {
			if msgType == "url_verification" {
				if receivedToken, ok := data["token"]; ok {
					if receivedToken != token {
						c.JSON(200, gin.H{})
						c.Abort()
						return
					} else {
						if challenge, ok := data["challenge"]; ok {
							c.JSON(http.StatusOK, gin.H{"challenge": challenge})
							c.Abort()
							return
						}
					}
				}
			}
		}
		c.Next()
	}
}
