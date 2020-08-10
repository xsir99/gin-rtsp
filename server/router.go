package server

import (
	"ginrtsp/api"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"net/http"
)

// NewRouter Gin 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())

	// 路由
	r.GET("/ping", api.Ping)
	r.LoadHTMLFiles("./html/view-stream.html")
	r.Static("/h", "./html")
	r.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "view-stream.html", nil)
	})
	route := r.Group("/stream")
	{
		route.POST("/play", api.PlayRTSP)
		route.POST("/upload/:channel", api.Mpeg1Video)
		route.GET("/live/:channel", api.Wsplay)
	}


	return r
}
