package main

import (
	"fmt"
	"ginSessionTest/models"
	"net/http"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	//"github.com/thinkerou/favicon"
)

func init() {
	_ = models.Setup()

}

//CORSMiddleware ...
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

func main() {
	r := gin.Default()

	store := sessions.NewCookieStore([]byte("secret"))
	r.Use(sessions.Sessions("realiza-session", store))

	r.Use(CORSMiddleware())
	r.Use(gzip.Gzip(gzip.DefaultCompression))

	//db.Init()

	//r.POST("/signin", login)
	//r.GET("/logout", logout)
	//r.GET("/home", home)

	r.LoadHTMLGlob("./views/*")

	r.Static("/assets", "./assets")
	//r.Use(favicon.New("./favicon.ico"))

	r.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.tpl", gin.H{
			"oi": "AEEEEEE",
		})
	})

	r.NoRoute(func(c *gin.Context) {
		c.HTML(404, "404.html", gin.H{})
	})

	r.Run(":9000")
}
