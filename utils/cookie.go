package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetCookie(c *gin.Context, token string) {
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Auth", token, 3600*24*100, "", "", false, true) // false because we're testing in local in prod change it to true

}
