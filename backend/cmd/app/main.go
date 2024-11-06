package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"backend/config/db_config"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
    log.Print("Starting program")

    log.Print("Loading Config")
    dbConfig := db_config.LoadConfig()

    log.Print(dbConfig.DBHost)
    log.Print(dbConfig.DBName)
    log.Print(dbConfig.DBUser)
    log.Print(dbConfig.DBPassword)
    log.Print(dbConfig.DBPort)

    log.Print("Connecting database")
    dbConnectionStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", 
                        dbConfig.DBUser, 
                        dbConfig.DBPassword,
                        dbConfig.DBHost, 
                        dbConfig.DBPort, 
                        dbConfig.DBName)
                        

    log.Print(dbConnectionStr)
    dbHandle, err := pgxpool.Connect(context.Background(), dbConnectionStr)
    if err != nil {
        log.Fatalf("unable to connect to database: %v", err)
    }
    defer dbHandle.Close()
    qR, err := dbHandle.Query(context.Background() ,"select count(*) from users")
    if err != nil {
      log.Fatalf("Query failed: %v", err)
    }
    var a[1]int
    qR.Scan(a)
    log.Print(a[0])

    //router := gin.Default()
    //router.LoadHTMLGlob("./../frontend/htmls/*.html")
//
    //router.GET("/ping", getHello)
//
    //router.GET("/ping/pong", func(c *gin.Context) {
    //    c.JSON(http.StatusOK, gin.H{
    //      "message": "bong",
    //    })
    //})
    //router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func getHello(c *gin.Context) {
  c.HTML(http.StatusOK, "hello.html", nil)
}