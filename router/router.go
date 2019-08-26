package router

import (
	"github.com/futuretea/yapidoc/conf"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	GinEngine *gin.Engine
)

func init() {
	yapiURL := conf.GetYapiURL()
	GinEngine = gin.New()
	GinEngine.LoadHTMLGlob("templates/*")
	GinEngine.GET("/projects/:token", func(c *gin.Context) {
		token := c.Param("token")
		tag := c.DefaultQuery("tag", "")
		result := render(c, yapiURL, token, tag)
		c.Header("Content-Type", "application/text/plain")
		c.Header("Content-Disposition", "attachment; filename=api.md")
		c.HTML(http.StatusOK, "api.md", result)
	})
}
