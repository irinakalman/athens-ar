package main

import (
	"database/sql"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Figure struct {
	ID     int     `json:"id"`
	Lat    float64 `json:"lat"`
	Long   float64 `json:"long"`
	Marker string  `json:"marker"`
	Figure string  `json:"figure"`
}

func main() {
	var err error

	router := gin.Default()

	var POSTGRES_DB = os.Getenv("POSTGRES_DB")
	var POSTGRES_USER = os.Getenv("POSTGRES_USER")
	var POSTGRES_PASSWORD = os.Getenv("POSTGRES_PASSWORD")

	var db *sql.DB
	db, err = sql.Open("postgres", "host=db_postgis dbname="+POSTGRES_DB+" user="+POSTGRES_USER+" password="+POSTGRES_PASSWORD+" sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	router.POST("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data":     "pong",
			"metadata": gin.H{},
		})
	})

	router.POST("/figures/get", func(c *gin.Context) {
		var err error

		var params struct {
			Lat      float64 `json:"lat"`
			Long     float64 `json:"long"`
			Radius_m int     `json:"radius_m"`
		}

		err = c.ShouldBindJSON(&params)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"err": "Invalid JSON format"})
			return
		}

		// Hmiparimitono
		query := `
			SELECT id, lat, long, marker, figure
			FROM figures
			WHERE 
				(6371 * acos(cos(radians($1)) * cos(radians(lat)) * cos(radians(long) - radians($2)) + sin(radians($1)) * sin(radians(lat)))) <= $3
		`
		rows, err := db.Query(query, params.Lat, params.Long, params.Radius_m/1000.0)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		var figures []Figure
		for rows.Next() {
			var figure Figure
			if err := rows.Scan(&figure.ID, &figure.Lat, &figure.Long, &figure.Marker, &figure.Figure); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			figures = append(figures, figure)
		}

		c.JSON(http.StatusOK, gin.H{
			"data": figures,
			"metadata": gin.H{
				"params": params,
			},
		})
	})

	router.POST("/figures/set", func(c *gin.Context) {
		var figure Figure
		if err := c.ShouldBindJSON(&figure); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
			return
		}

		_, err := db.Exec("INSERT INTO figures (lat, long, marker, figure) VALUES ($1, $2, $3, $4)",
			figure.Lat, figure.Long, figure.Marker, figure.Figure)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert figure into database"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Figure added successfully"})
	})

	router.Run(":8081")
}
