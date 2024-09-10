package main

import (
	"database/sql"
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

	r.GET("/getGame", getGame)
	r.GET("/updateWaitStatus", updateWaitStatus)
	r.GET("/updatePlayStatus", updatePlayStatus)
	r.GET("/updatePlayer", updatePlayer)

	r.GET("/getAllUser", getAllUser)
	r.GET("/updateArr", updateArr)

	r.Run()
}

func register(c *gin.Context) {
	var minID int

	var name = c.Query("name");
	db.QueryRow("SELECT MIN(id) FROM user_tb WHERE username = ''").Scan(&minID)
	db.Exec("UPDATE user_tb SET username = ? WHERE id = ?", name, minID)


}

func updatePlayer(c *gin.Context) {
	db.Exec("UPDATE game_tb SET playuser = $1", c.Query("p"))
	getGame(c)
}

func updateWaitStatus(c *gin.Context) {
	db.Exec("UPDATE game_tb SET status = $1", utildata.WAITING)
	getGame(c)
}

func updatePlayStatus(c *gin.Context) {
	db.Exec("UPDATE game_tb SET status = $1", utildata.PLAYING)
	getGame(c)
}

func getGame(c *gin.Context) {
	rows, _ := db.Query("SELECT * FROM game_tb;")
	var games []map[string]interface{}
	for rows.Next() {
		var idgame int
		var status, playuser string
		rows.Scan(&idgame, &status, &playuser)
		game := map[string]interface{}{
			"idgame":   idgame,
			"status":   status,
			"playuser": playuser,
		}
		games = append(games, game)

	}

	g := games[0]
	c.JSON(http.StatusOK, gin.H{
		"idgame":   g["idgame"],
		"status":   g["status"],
		"playuser": g["playuser"],
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
			"id":   id,
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
