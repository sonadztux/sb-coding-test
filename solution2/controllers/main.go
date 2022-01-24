package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"solution2/omdb-movie-search/databases"
	"solution2/omdb-movie-search/models"
	"time"

	"github.com/gin-gonic/gin"
)

func GetMovieList(ctx *gin.Context) {
	var client = &http.Client{}
	var dataResponse models.MovieList
	db := databases.GetDB()

	searchQuery, isAnySearchQuery := ctx.GetQuery("searchword")
	pageQuery, _ := ctx.GetQuery("pagination")

	if !isAnySearchQuery || searchQuery == "" {
		newLogCall := models.LogCall{
			CallType:   "GetMovieList",
			ParamQuery: fmt.Sprintf("searchword: %s, pagination: %s", searchQuery, pageQuery),
			Status:     http.StatusBadRequest,
			CreatedAt:  time.Now(),
		}
		databases.AddLogCall(db, newLogCall)

		ctx.JSON(http.StatusBadRequest, gin.H{
			"ok":      false,
			"message": "argument 'searchword' is required and can't be empty!",
		})
		return
	}

	apiUrl := os.Getenv("API_BASE_URL") + "?apikey=" + os.Getenv("API_KEY") + "&s=" + searchQuery + "&page=" + pageQuery

	request, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		newLogCall := models.LogCall{
			CallType:   "GetMovieList",
			ParamQuery: fmt.Sprintf("searchword: %s, pagination: %s", searchQuery, pageQuery),
			Status:     http.StatusInternalServerError,
			CreatedAt:  time.Now(),
		}
		databases.AddLogCall(db, newLogCall)

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"ok":      false,
			"message": err.Error(),
		})
		return
	}

	response, err := client.Do(request)
	if err != nil {
		newLogCall := models.LogCall{
			CallType:   "GetMovieList",
			ParamQuery: fmt.Sprintf("searchword: %s, pagination: %s", searchQuery, pageQuery),
			Status:     http.StatusInternalServerError,
			CreatedAt:  time.Now(),
		}
		databases.AddLogCall(db, newLogCall)

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"ok":      false,
			"message": err.Error(),
		})
		return
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&dataResponse)
	if err != nil {
		log.Fatal(err)
	}

	newLogCall := models.LogCall{
		CallType:   "GetMovieList",
		ParamQuery: fmt.Sprintf("searchword: %s, pagination: %s", searchQuery, pageQuery),
		Status:     http.StatusOK,
		CreatedAt:  time.Now(),
	}
	databases.AddLogCall(db, newLogCall)

	ctx.JSON(http.StatusOK, gin.H{
		"ok":      true,
		"message": "get all movies success",
		"data":    &dataResponse,
	})
}

func GetMovieDetail(ctx *gin.Context) {
	var client = &http.Client{}
	var dataResponse models.MovieDetail
	db := databases.GetDB()

	movieId := ctx.Param("id")
	fmt.Println(movieId)

	apiUrl := os.Getenv("API_BASE_URL") + "?apikey=" + os.Getenv("API_KEY") + "&i=" + movieId

	request, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		newLogCall := models.LogCall{
			CallType:   "GetMovieDetail",
			ParamQuery: fmt.Sprintf("id: %s", movieId),
			Status:     http.StatusInternalServerError,
			CreatedAt:  time.Now(),
		}
		databases.AddLogCall(db, newLogCall)

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"ok":      false,
			"message": err.Error(),
		})
		return
	}

	response, err := client.Do(request)
	if err != nil {
		newLogCall := models.LogCall{
			CallType:   "GetMovieDetail",
			ParamQuery: fmt.Sprintf("id: %s", movieId),
			Status:     http.StatusInternalServerError,
			CreatedAt:  time.Now(),
		}
		databases.AddLogCall(db, newLogCall)

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"ok":      false,
			"message": err.Error(),
		})
		return
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&dataResponse)
	if err != nil {
		log.Fatal(err)
	}

	newLogCall := models.LogCall{
		CallType:   "GetMovieDetail",
		ParamQuery: fmt.Sprintf("id: %s", movieId),
		Status:     http.StatusOK,
		CreatedAt:  time.Now(),
	}
	databases.AddLogCall(db, newLogCall)

	ctx.JSON(http.StatusOK, gin.H{
		"ok":      true,
		"message": "get movie detail success",
		"data":    &dataResponse,
	})
}
