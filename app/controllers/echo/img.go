package echo

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func ImgBase(c *gin.Context) {

	ff, _ := os.Open("Vector-5.png")
	defer ff.Close()

	sourcebuffer := make([]byte, 500000)
	n, _ := ff.Read(sourcebuffer)
	sourcestring := base64.StdEncoding.EncodeToString(sourcebuffer[:n])

	c.JSON(http.StatusOK, gin.H{
		"data": "data:image;base64," + sourcestring,
	})
}
