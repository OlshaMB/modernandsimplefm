package server

import (
	"github.com/OlshaMB/modernandsimplefm/pkg/routes"
	"github.com/gin-gonic/gin"
	"github.com/noirbizarre/gonja"
)
var listTemplate = gonja.Must(gonja.FromFile("pkg/view/list.j2"))
func Server() {
  r := gin.Default()
  r.GET("/list/*path", routes.ListDir)
  r.GET("/view/*path", routes.View)
  r.Run()
}