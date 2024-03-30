package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data":     "pong",
			"metadata": gin.H{},
		})
	})

	router.POST("/figures", func(c *gin.Context) {
		var err error

		var figures map[string]string

		var params struct {
			Lat     float64 `json:"lat"`
			Long    float64 `json:"long"`
			Range_m int     `json:"range_m"`
		}

		err = c.ShouldBindJSON(&params)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"err": "Invalid JSON format"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": figures,
			"metadata": gin.H{
				"paramsParsed": params,
			},
		})
	})

	router.Run(":8080")
}
