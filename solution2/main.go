package main

import (
	"net/http"
	"os"
	"solution2/omdb-movie-search/controllers"
	"solution2/omdb-movie-search/databases"

	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"
)

func init() {
	// load .env file
	gotenv.Load()

	// Initialize database
	databases.InitDB()
}

func main() {
	// Initialize HTTP routes
	r := gin.Default()

	r.GET("ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "pong")
	})

	movies := r.Group("/movie")
	{
		movies.GET("", controllers.GetMovieList)
		movies.GET("/:id", controllers.GetMovieDetail)
	}

	host := os.Getenv("APP_HOST")
	port := os.Getenv("APP_PORT")

	_ = r.Run(host + ":" + port)
}
