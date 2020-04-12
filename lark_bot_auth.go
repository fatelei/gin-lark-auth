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

		if v, ok := data["token"]; ok {
			if v != token {
				c.JSON(200, gin.H{})
				c.Abort()
				return
			} else {
				if msgType, ok := data["type"]; ok {
					if msgType == "url_verification" {
						if challenge, ok := data["challenge"]; ok {
							c.JSON(http.StatusOK, gin.H{"challenge": challenge})
							c.Abort()
							return
						}
					} else {
						c.Next()
					}
				}
			}
		} else {
			c.JSON(http.StatusOK, gin.H{})
			c.Abort()
			return
		}
	}
}
