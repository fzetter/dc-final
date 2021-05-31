package api

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/fzetter/dc-final/api/src/routes"
)

var router *gin.Engine

func Start() {
    app := gin.Default()
    app.MaxMultipartMemory = 8 << 20 // 8 MiB

    // ******
    // PUBLIC
    // ******
    app.LoadHTMLGlob("api/public/*")

    app.GET("/", func(c *gin.Context) {
      c.HTML(
          http.StatusOK,
          "index.html",
          gin.H{
              "title": "Home Page",
          },
      )
    })

    // ******
    // ROUTES
    // ******
    routes.Init(app)

    app.Run(":8080")
}
