package auth

import (
	"github.com/gin-gonic/gin"
)

func Basic(cred string) func(*gin.Context) {
	return func(c *gin.Context) {
		// basic authen get username and password from header
		username, password, ok := c.Request.BasicAuth()
		if !ok {
			c.AbortWithStatus(401)
			return
		}

		// check if username and password is correct
		if username != "e2eagent" || password != "weakp!$$" {
			c.AbortWithStatus(401)
			return
		}

		c.Next()
	}
}
