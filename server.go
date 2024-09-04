package main

import (
	"database/sql"
	utildata "meo_no/utilData"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var connStr = "user=root password=345FSDF$# dbname=meono sslmode=disable"
var db, _ = sql.Open("postgres", connStr)

func main() {
	r := gin.Default()

	r.GET("/getGame", getGame)
	r.GET("/UpdateWaitStatus", UpdateWaitStatus)
	r.GET("/UpdatePlayStatus", UpdatePlayStatus)
	r.GET("/UpdatePlayer", UpdatePlayer)

	
	r.GET("/getAllUser", getAllUser)
	r.GET("/UpdateArr", UpdateArr)

	r.Run()
}

func UpdatePlayer(c *gin.Context) {
	db.Exec("UPDATE game_tb SET playuser = $1", c.Query("p"))
	getGame(c)
}

func UpdateWaitStatus(c *gin.Context) {
	db.Exec("UPDATE game_tb SET status = $1", utildata.WAITING)
	getGame(c)
}

func UpdatePlayStatus(c *gin.Context) {
	db.Exec("UPDATE game_tb SET status = $1", utildata.PLAYING)
	getGame(c)
}

func getGame(c *gin.Context) {
	rows, _ := db.Query("SELECT * FROM game_tb;")
	var games []map[string]interface{}
	for rows.Next() {
		var id int
		var status, playuser string
		rows.Scan(&id, &status, &playuser)
		game := map[string]interface{}{
			"id":       id,
			"status":   status,
			"playuser": playuser,
		}
		games = append(games, game)

	}

	g := games[0]
	c.JSON(http.StatusOK, gin.H{
		"id":       g["id"],
		"status":   g["status"],
		"playuser": g["playuser"],
	})
}

func UpdateArr(c *gin.Context) {
	db.Exec("UPDATE user_tb SET arr = $1 where id = $2", c.Query("arr"), c.Query("id"))
	getAllUser(c)
}

func getAllUser(c *gin.Context) {
	rows, _ := db.Query("SELECT * FROM user_tb where play is not null;")
	var users []map[string]interface{}
	for rows.Next() {
		var id int
		var username, arr, play string
		rows.Scan(&id, &username, &arr, &play)
		user := map[string]interface{}{
			"id":       id,
			"username": username,
			"arr":      arr,
			"play":     play,
		}
		users = append(users, user)

	}

	c.JSON(http.StatusOK, gin.H{
		"allUser": users,
	})
}
