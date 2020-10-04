package middleware

// import "fmt"

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/meloalright/guora/internal/model"
)

func Administrator() gin.HandlerFunc {
	return func(c *gin.Context) {
		ID, exist := c.Get("uid")
		if !exist {
			c.Redirect(http.StatusTemporaryRedirect, "/error?message=Not exist your uid")
			c.Abort()
			return
		}
		value, ok := ID.(int)
		if !ok {
			c.Redirect(http.StatusTemporaryRedirect, "/error?message=Your uid is not int ")
			c.Abort()
			return
		}

		// check user if Type is 100 or not
		var u model.User
		u.ID = value

		user, err := u.Get()

		if err != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/error?message=User get error")
			c.Abort()
			return
		}
		if user.Type != 100 {
			c.Redirect(http.StatusTemporaryRedirect, "/error?message=You are not Administrator")
			c.Abort()
			return
		}

		// before request

		c.Next()
	}
}
