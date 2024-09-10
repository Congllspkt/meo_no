package main

import (
	"database/sql"
	"fmt"
	utildata "meo_no/utilData"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-contrib/cors"
)

var dsn = "root:345FSDF$#@tcp(localhost:3306)/meono"
var db, _ = sql.Open("mysql", dsn)

func main() {

	r := gin.Default()

	r.Use(cors.Default())

	r.GET("/register", register)
	r.GET("/getStatusGame", getStatusGame)

	r.GET("/updateWaitStatus", updateWaitStatus)
	r.GET("/updatePlayStatus", updatePlayStatus)
	r.GET("/updatePlayer", updatePlayer)

	r.GET("/getAllUser", getAllUser)
	r.GET("/updateArr", updateArr)

	r.Run()
}

func register(c *gin.Context) {
	var minID int
	var name = c.Query("name")
	db.QueryRow("SELECT MIN(id) FROM user_tb WHERE username = ''").Scan(&minID)
	db.Exec("UPDATE user_tb SET username = ? WHERE id = ?", name, minID)

	c.JSON(http.StatusOK, gin.H{
		"username": name,
	})

}

func updatePlayer(c *gin.Context) {
	db.Exec("UPDATE game_tb SET playuser = $1", c.Query("p"))
}

func updateWaitStatus(c *gin.Context) {
	db.Exec("UPDATE game_tb SET status = $1", utildata.WAITING)
}

func updatePlayStatus(c *gin.Context) {
	db.Exec("UPDATE game_tb SET status = $1", utildata.PLAYING)
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
		messageStatus = "dang trong tran";
		statusGame = "p"

		status = ""
		db.QueryRow("SELECT arr, status FROM user_tb where id = ?;", c.Query("id")).Scan(&arr, &status)
	
		fmt.Print(status)


		if(status == "p") {
			messageStatusUser = "van dang choi"
			statusUser = "p"
		} else if (status == "d") {
			statusUser = "d"
			messageStatusUser = "ban thua roi"
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"statusGame": statusGame,
		"messageStatus": messageStatus,
		"arr": arr,
		"statusUser" : statusUser,
		"messageStatusUser" : messageStatusUser,
	})
}

func updateArr(c *gin.Context) {
	db.Exec("UPDATE user_tb SET arr = $1 where id = $2", c.Query("arr"), c.Query("id"))
	getAllUser(c)
}

func getAllUser(c *gin.Context) {
	rows, _ := db.Query("SELECT id, username, status  FROM user_tb;")
	var users []map[string]interface{}
	for rows.Next() {
		var id, username, status, arr string
		rows.Scan(&id, &username, &arr, &status)
		user := map[string]interface{}{
			"id":       id,
			"username": username,
			"arr":      arr,
			"status":   status,
		}
		users = append(users, user)

	}

	c.JSON(http.StatusOK, gin.H{
		"allUser": users,
	})
}
