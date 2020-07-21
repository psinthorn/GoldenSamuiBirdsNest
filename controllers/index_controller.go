package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type index struct{}

var Index index

func (i *index) Welcome(c *gin.Context) {
	c.HTML(http.StatusOK, "welcome.html", nil)
}
