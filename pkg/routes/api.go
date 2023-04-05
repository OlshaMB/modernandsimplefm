package routes

import (
	"net/http"

	"github.com/OlshaMB/modernandsimplefm/pkg/api"
	"github.com/gin-gonic/gin"
)

func ListDir(c *gin.Context) {
	path := c.Param("path")
	entries, err := api.ListDir(path)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}
	c.JSON(http.StatusCreated, entries)
}