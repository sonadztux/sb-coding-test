# OMDB Movie Search
Golang Backend HTTP REST API to search movie from OMDB API.

## Getting Started

### Requirements

- [Go](https://golang.org/doc/install) >= 1.17
- [Gin](https://github.com/gin-gonic/gin) >= 1.7.7
- [SQLite3](https://github.com/mattn/go-sqlite3) >= 1.14.10

### How to Run
- Copy or rename file `.env.example` to `.env`.
- Edit the `.env` to your desire configurations.
- Finally run the backend service:
    ```sh
    go run .
    ```
- By default you can access the application on [http://localhost:8080/](http://localhost:8080/)


## API Endpoint Documentation
<details>
<summary><b>Get Movie List</b></summary>
Return json data about all OMDB movies.

- **URL**

  `/movie/searchword={movie_name}&pagination={integer}`

- **Method**

  `GET`

- **Sample Success Response**

    Sample URL: [http://localhost:8080/movie?searchword=endgame&pagination=2](http://localhost:8080/movie?searchword=endgame&pagination=2)

  **Code**: 200 OK

```json
{
  "data": {
    "Search": [
      {
        "ImdbID": "tt11136174",
        "Title": "The Breakroom Lost Episodes: Chapter II - Endgame",
        "Year": "2019",
        "Type": "movie",
        "Poster": "N/A"
      },
      {
        "ImdbID": "tt12861690",
        "Title": "Britten's Endgame",
        "Year": "2013",
        "Type": "movie",
        "Poster": "https://m.media-amazon.com/images/M/MV5BYzk3YTVhYjItNzQ1Ni00MWVmLTljYmEtNGU4YmMxYTRiYWExXkEyXkFqcGdeQXVyMTk0MjQ3Nzk@._V1_SX300.jpg"
      },
      {
        "ImdbID": "tt13180560",
        "Title": "Endgame",
        "Year": "2020",
        "Type": "movie",
        "Poster": "N/A"
      },
      ...
    ]
  },
  "message": "get all movies success",
  "ok": true
}
```

- **Data Type Attributes**

```json
{
  "data": {
    "Search": [
      {
        "ImdbID": string,
        "Title": string,
        "Year": string,
        "Type": string,
        "Poster": string
      }
    ]
  },
  "message": string,
  "ok": boolean
}
```
</details>


<details>
<summary><b>Get Movie Detail by ID</b></summary>
return json data containing movie detail information by id.

- **URL**

  `/movie/:id`

- **Method**

  `GET`

- **Sample Success Response**

    Sample URL: [http://localhost:8080/movie/tt10872600](http://localhost:8080/movie/tt10872600)

  **Code**: 200 OK

```json
{
  "data": {
    "imdbID": "tt10872600",
    "Title": "Spider-Man: No Way Home",
    "Year": "2021",
    "Rated": "PG-13",
    "Released": "17 Dec 2021",
    "Runtime": "148 min",
    "Genre": "Action, Adventure, Fantasy",
    "Director": "Jon Watts",
    "Writer": "Chris McKenna, Erik Sommers, Stan Lee",
    "Actors": "Tom Holland, Zendaya, Benedict Cumberbatch",
    "Plot": "With Spider-Man's identity now revealed, Peter asks Doctor Strange for help. When a spell goes wrong, dangerous foes from other worlds start to appear, forcing Peter to discover what it truly means to be Spider-Man.",
    "Language": "English, Tagalog",
    "Country": "United States",
    "Awards": "N/A",
    "Poster": "https://m.media-amazon.com/images/M/MV5BZWMyYzFjYTYtNTRjYi00OGExLWE2YzgtOGRmYjAxZTU3NzBiXkEyXkFqcGdeQXVyMzQ0MzA0NTM@._V1_SX300.jpg",
    "Ratings": [
      {
        "Source": "Internet Movie Database",
        "Value": "8.9/10"
      },
      {
        "Source": "Rotten Tomatoes",
        "Value": "93%"
      },
      {
        "Source": "Metacritic",
        "Value": "71/100"
      }
    ],
    "Metascore": "71",
    "imdbRating": "8.9",
    "imdbVotes": "257,688",
    "Type": "movie",
    "DVD": "N/A",
    "BoxOffice": "$385,831,855",
    "Production": "N/A",
    "Website": "N/A"
  },
  "message": "get movie detail success",
  "ok": true
}
```

- **Data Type Attributes**

```json
{
  "data": {
    "imdbID": string,
    "Title": string,
    "Year": string,
    "Rated": string,
    "Released": string,
    "Runtime": string,
    "Genre": string,
    "Director": string,
    "Writer": string,
    "Actors": string,
    "Plot": string,
    "Language": string,
    "Country": string,
    "Awards": string,
    "Poster": string,
    "Ratings": array object [
      {
        "Source": string,
        "Value": string
      }
    ],
    "Metascore": string,
    "imdbRating": string,
    "imdbVotes": string,
    "Type": string,
    "DVD": string,
    "BoxOffice": string,
    "Production": string,
    "Website": string
  },
  "message": string,
  "ok": boolean
}
```
</details>