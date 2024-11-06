package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RequestData struct {
  Name     string `json:"name"`
  Email    string `json:"email"`
  Password string `json:"password"`
}

func main() {
    log.Print("Starting program")

    /*log.Print("Loading Database Config")
    dbConfig := db_config.LoadConfig()
    
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
    defer dbHandle.Close()*/

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
    router.StaticFile("favicon.ico",  "./../frontend/app/build/favicon.ico")
    router.StaticFile("manifest.json",  "./../frontend/app/build/manifest.json")
    router.StaticFile("logo192.png",  "./../frontend/app/build/logo192.png")

    router.GET("/ping", GetPing)
    router.POST("/ping", PostPing)

    router.Run()
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