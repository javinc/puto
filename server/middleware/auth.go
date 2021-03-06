package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/javinc/mango/errors"
	"github.com/javinc/mango/server/auth"
)

// Auth middleware checks if has valid token
func Auth(checkPayload func(*gin.Context, map[string]interface{}) error) gin.HandlerFunc {
	return func(c *gin.Context) {
		// validates token
		p, err := auth.CheckToken(
			getToken(c.Request.Header.Get("authorization")))
		if err != nil {
			c.JSON(http.StatusUnauthorized, errors.NewError("AUTH_REQUIRED", err))
			c.Abort()

			return
		}

		// validates payload
		err = checkPayload(c, p)
		if err != nil {
			// NOTE you can check use 404
			c.JSON(http.StatusForbidden, errors.NewError("AUTH_INVALID", err))
			c.Abort()

			return
		}

		c.Next()
	}
}

func getToken(authHeader string) string {
	s := strings.Split(strings.TrimSpace(authHeader), " ")
	if len(s) != 2 {
		return ""
	}

	return s[1]
}
