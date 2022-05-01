package routers

import (
	"net/http"

	"malpvapps/web-services-gin/data"

	"github.com/gin-gonic/gin"
)

// getAlbums responds with the list of all albums as JSON.
func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, data.Albums)
}
