package routes

import (
	"net/http"
	"path"

	"github.com/OlshaMB/modernandsimplefm/pkg/api"
	"github.com/OlshaMB/modernandsimplefm/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/noirbizarre/gonja"
)
var listTemplate = gonja.Must(gonja.FromFile("pkg/view/list.j2"))

func View(c *gin.Context) {
	pathToFile := c.Param("path")
	isFile, err := api.IsFile(pathToFile)
	println(isFile)
	if isFile {
		file, err := api.ReadFile(pathToFile)
		if err != nil {
			c.Status(http.StatusNotFound)
			return
		}

		http.ServeContent(c.Writer, c.Request, "text.go", time.Now(), file)
		return
	}
	if err != nil {
		print(err.Error())
		c.Status(http.StatusNotFound)
		return
	}
	entries, err := api.ListDir(pathToFile)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	ren, _ := listTemplate.Execute(gonja.Context{
		"entries":    entries,
		"formatTime": utils.FormatTime,
		"formatPath": utils.StripPathSeparator,
		"path":       path.Dir(pathToFile),
	})
	c.Writer.WriteString(ren)
	c.Writer.Header().Add("Content-Type", "text/html")
	c.Status(200)
}