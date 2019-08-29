package router

import (
	"fmt"
	"github.com/futuretea/yapidoc/conf"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	GinEngine *gin.Engine
	suffixMap = map[string]string{
		"markdown": "md",
		"json": "json",
	}
)

func init() {
	yapiURL := conf.GetYapiURL()
	GinEngine = gin.New()
	GinEngine.LoadHTMLGlob("templates/*")
	GinEngine.GET("/projects/:token", func(c *gin.Context) {
		token := c.Param("token")
		tag := c.DefaultQuery("tag", "")
		result := render(c, yapiURL, token, tag)
		format := c.DefaultQuery("format", "markdown")
		suffix, formatIsSupport := suffixMap[format]
		if formatIsSupport == false {
			c.AbortWithStatus(404)
		}
		fileName := fmt.Sprintf("%s.api.%s", tag, suffix)
		c.Header("Content-Type", "application/text/plain")
		c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
		if format == "markdown" {
			c.HTML(http.StatusOK, "api.md", result)
		} else if format == "json" {
			c.JSON(http.StatusOK, result)
		} else {
			c.AbortWithStatus(404)
		}
	})
}
