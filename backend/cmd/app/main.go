package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"backend/config/db_config"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

type RequestData struct {
  Name     string `json:"name"`
  Email    string `json:"email"`
  Password string `json:"password"`
}

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
    //qR, err := dbHandle.Query(context.Background() ,"select count(*) from users")
    //if err != nil {
    //  log.Fatalf("Query failed: %v", err)
    //}
    //var a[1]int
    //qR.Scan(a)
    //log.Print(a[0])

    router := gin.Default()
    router.LoadHTMLGlob("./../frontend/app/build/*.html")
    router.Static("/static", "./../frontend/app/build/static")
    
    router.GET("/ping", GetPing)
    router.POST("/ping", PostPing)

    router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func GetPing(c *gin.Context) {
  c.HTML(http.StatusOK, "index.html", nil)
}

func PostPing(c *gin.Context) {

  data, err := io.ReadAll(c.Request.Body)
  if err != nil {
    log.Fatal("Rading PostPing request failed")
  }

  var requestData RequestData
  err = json.Unmarshal(data, &requestData)
	if err != nil {
		log.Fatal(err)
	}

  log.Print(requestData.Name)
  log.Print(requestData.Email)
  log.Print(requestData.Password)
}