package main

import (
	"log"
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	"strconv"
    "database/sql"
	"time"
	"fmt"
    _ "github.com/lib/pq"
)

func dbFunc(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        if _, err := db.Exec("CREATE TABLE IF NOT EXISTS ticks (tick timestamp)"); err != nil {
            c.String(http.StatusInternalServerError,
                fmt.Sprintf("Error creating database table: %q", err))
            return
        }

        if _, err := db.Exec("INSERT INTO ticks VALUES (now())"); err != nil {
            c.String(http.StatusInternalServerError,
                fmt.Sprintf("Error incrementing tick: %q", err))
            return
        }

        rows, err := db.Query("SELECT tick FROM ticks")
        if err != nil {
            c.String(http.StatusInternalServerError,
                fmt.Sprintf("Error reading ticks: %q", err))
            return
        }

        defer rows.Close()
        for rows.Next() {
            var tick time.Time
            if err := rows.Scan(&tick); err != nil {
                c.String(http.StatusInternalServerError,
                    fmt.Sprintf("Error scanning ticks: %q", err))
                return
            }
            c.String(http.StatusOK, fmt.Sprintf("Read from DB: %s\n", tick.String()))
        }
    }
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
    if err != nil {
        log.Fatalf("Error opening database: %q", err)
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Nothing to see here.")
	})

	tfsaMap := map[string]int{
		"2009": 5000,
		"2010": 5000,
		"2011": 5000,
		"2012": 5000,
		"2013": 5500,
		"2014": 5500,
		"2015": 10000,
		"2016": 5500,
		"2017": 5500,
		"2018": 5500,
		"2019": 6000,
		"2020": 6000,
		"2021": 6000,
	}

	router.GET("/TFSA/get", func(c *gin.Context) {
		c.JSON(http.StatusOK, tfsaMap)
	})

	router.GET("/TFSA/get/:year", func(c *gin.Context) {
		year := c.Param("year")
		if val, ok := tfsaMap[year]; ok {
			c.String(http.StatusOK, strconv.Itoa(val))
		} else {
			c.String(http.StatusOK, "Year not found.")
		}
	})

	router.GET("/TFSA/refresh", func(c *gin.Context) {

		c.String(http.StatusOK, "Refreshed")
	})

	router.GET("/db", dbFunc(db))

	router.Run(":" + port)
}
