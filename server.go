package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-contrib/cors"
)

var dsn = "root:345FSDF$#@tcp(localhost:3306)/meono"
var db, _ = sql.Open("mysql", dsn)

func main() {

	r := gin.Default()

	r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://172.25.219.197:5500"}, // Allow specific origin
        AllowMethods:     []string{"GET", "POST", "OPTIONS"},     // Allow specific methods
        AllowHeaders:     []string{"Origin", "Content-Type"},     // Allow specific headers
        ExposeHeaders:    []string{"Content-Length"},              // Expose headers
        AllowCredentials: true,
        MaxAge:           12 * 3600, // Cache duration for preflight requests
    }))
	r.GET("/register", register)
	r.GET("/getStatusGame", getStatusGame)

	r.GET("/updatePlayer", updatePlayer)

	r.GET("/updateArr", updateArr)

	r.Run()
}

func register(c *gin.Context) {
	var minID int
	var name = c.Query("name")
	db.QueryRow("SELECT MIN(id) FROM user_tb WHERE username = ''").Scan(&minID)
	db.Exec("UPDATE user_tb SET username = ?, status = 'w' WHERE id = ?", name, minID)

	c.JSON(http.StatusOK, gin.H{
		"username": name,
	})

}

func updatePlayer(c *gin.Context) {
	db.Exec("UPDATE game_tb SET playuser = $1", c.Query("p"))
}

func getStatusGame(c *gin.Context) {
	var status, arr, statusUser, messageStatusUser string
	db.QueryRow("SELECT status FROM game_tb;").Scan(&status)

	var messageStatus string
	var statusGame string

	if status == "w" {
		messageStatus = "Cho game 1 xiu"
		statusGame = "w"
	} else {
		messageStatus = "dang trong tran"
		statusGame = "p"

		status = ""
		db.QueryRow("SELECT arr, status FROM user_tb where id = ?;", c.Query("id")).Scan(&arr, &status)

		fmt.Print(status)

		if status == "p" {
			messageStatusUser = "van dang choi"
			statusUser = "p"
		} else if status == "d" {
			statusUser = "d"
			messageStatusUser = "ban thua roi"
		}
	}

	rows, _ := db.Query("SELECT id, username, status  FROM user_tb where username != '' and id != ?;", c.Query("id"))
	var users []map[string]interface{}
	for rows.Next() {
		var id, username, status string
		rows.Scan(&id, &username, &status)
		user := map[string]interface{}{
			"id":       id,
			"username": username,
			"status":   status,
		}
		users = append(users, user)

	}

	c.JSON(http.StatusOK, gin.H{
		"statusGame":        statusGame,
		"messageStatus":     messageStatus,
		"arr":               arr,
		"statusUser":        statusUser,
		"messageStatusUser": messageStatusUser,
		"allUser":           users,
	})
}

func updateArr(c *gin.Context) {
	db.Exec("UPDATE user_tb SET arr = $1 where id = $2", c.Query("arr"), c.Query("id"))
}
