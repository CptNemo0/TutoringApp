package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/CptNemo0/TutoringApp/config/db/config"
)

func main() {
    router := gin.Default()
    router.LoadHTMLGlob("./../frontend/htmls/*.html")

    router.GET("/ping", getHello)

    router.GET("/ping/pong", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
          "message": "bong",
        })
    })
    router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func getHello(c *gin.Context) {
  c.HTML(http.StatusOK, "hello.html", nil)
}