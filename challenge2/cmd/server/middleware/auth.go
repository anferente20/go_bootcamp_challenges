package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// NewAuthenticator returns a new instance of Authenticator.
func NewAuthenticator(token string) *Authenticator {
	return &Authenticator{
		token: token,
	}
}

// Authenticator is an interface that represents an authenticator.
type Authenticator struct {
	// token is the token used to authenticate the user.
	token string
}

func (a *Authenticator) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		// before
		token := c.GetHeader("Authorization")
		if token != a.token {
			c.JSON(http.StatusUnauthorized, map[string]any{"message": "unauthorized"})
			c.Abort()
			return
		}

		// now
		c.Next()

		// after
		// ...
	}
}
