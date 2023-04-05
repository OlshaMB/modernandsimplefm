package server

import (
	// "bufio"
	// "bytes"
	"net/http"
	"path"
	// "text/scanner"
	"time"

	"github.com/OlshaMB/modernandsimplefm/pkg/api"
	"github.com/OlshaMB/modernandsimplefm/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/noirbizarre/gonja"
)
var listTemplate = gonja.Must(gonja.FromFile("pkg/view/list.j2"))
func Server() {
  r := gin.Default()
  r.GET("/list/*path", func(c *gin.Context) {
	path := c.Param("path");
	entries, err := api.ListDir(path)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}
    c.JSON(http.StatusCreated, entries)
  })
  r.GET("/view/*path", func(c *gin.Context) {
	pathToFile := c.Param("path");
	isFile, err := api.IsFile(pathToFile)
	println(isFile)
	if isFile {
		file, err := api.ReadFile(pathToFile);
		if err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		// file_content := make([]byte, 10)
		// scanner = bufio.NewScanner(file);
		// scanner.
		// fileContentBuffer := bytes.NewReader(file_content);
		http.ServeContent(c.Writer, c.Request, "text.go", time.Now(), file)
		// c.Status(http.StatusAccepted)
		return
	}
	if err!=nil {
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
		"entries": entries,
		"formatTime": utils.FormatTime,
		"formatPath": utils.StripPathSeparator,
		"path": path.Dir(pathToFile),
	})
	c.Writer.WriteString(ren)
	c.Writer.Header().Add("Content-Type", "text/html")
	c.Status(200)
  })
  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}