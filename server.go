package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"math/rand"
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
		ExposeHeaders:    []string{"Content-Length"},             // Expose headers
		AllowCredentials: true,
		MaxAge:           12 * 3600, // Cache duration for preflight requests
	}))
	r.GET("/register", register)
	r.GET("/getStatusGame", getStatusGame)
	r.GET("/startGame", startGame)

	r.GET("/updatePlayer", updatePlayer)

	r.Run()
}

func startGame(c *gin.Context) {
	// get all id nguoi choi
	db.Exec("UPDATE user_tb set status = ''")
	db.Exec("UPDATE user_tb set status = 'p' where username != ''")
	
	var ids string
	db.QueryRow("SELECT GROUP_CONCAT(id) as ids FROM user_tb where username != ''").Scan(&ids)
	parts := strings.Split(ids, ",")
	var numbers []int
	for _, part := range parts {
		num, _ := strconv.Atoi(part)
		numbers = append(numbers, num)
	}
	fmt.Println("ids: ",  numbers)

	numberPlayers := len(numbers)
	
	// meodefause2 := 1
	// meobom1 := numberPlayers - 1


	meosee3 := 4
	meogive4 := 4
	meoreverse5 := 4
	meosuffle6 := 4
	meoskip7 := 4

	arrBobai := []int{}
	arrBobai = appendBobai(arrBobai, meosee3, 3)
	arrBobai = appendBobai(arrBobai, meogive4, 4)
	arrBobai = appendBobai(arrBobai, meoreverse5, 5)
	arrBobai = appendBobai(arrBobai, meosuffle6, 6)
	arrBobai = appendBobai(arrBobai, meoskip7, 7)
	fmt.Println("Tao Bo Bai : ",  arrBobai)

	// trom bai
	shuffleSlice(arrBobai)
	shuffleSlice(arrBobai)
	shuffleSlice(arrBobai)
	fmt.Println("Hoan Vi Bo Bai : ",  arrBobai)

	// chi bai
	for i := 0; i < numberPlayers; i++ {
		baiUser := arrBobai[:4]
		arrBobai = arrBobai[4:]

		fmt.Println("Bo Bai : ", i, arrBobai)


		var bai = []int{2}
		bai = append(bai, baiUser...)
		fmt.Println("bai user", i, bai)

		// save arr
		joinbai := joinIntSlice(bai)
		db.Exec("UPDATE user_tb SET arr = $1 where id = $2", joinbai, numbers[i])
	}

	fmt.Println("Bo Bai : ", arrBobai)
	db.Exec("UPDATE game_tb SET arr = $1", joinIntSlice(arrBobai))












	






	
}

func joinIntSlice(numbers []int) string {
	var stringNumbers []string
	for _, num := range numbers {
		stringNumbers = append(stringNumbers, strconv.Itoa(num))
	}
	return strings.Join(stringNumbers, ",")
}

func shuffleSlice(arr []int) {
	rand.Seed(time.Now().UnixNano()) 
	for i := len(arr) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func appendBobai(arrBobai []int, n int, index int) [] int{
	for i := 0; i < n; i++ {
		arrBobai = append(arrBobai, index)
	}
	return arrBobai
}

func register(c *gin.Context) {
	var minID int
	var name = c.Query("name")
	db.QueryRow("SELECT MIN(id) FROM user_tb WHERE username = ''").Scan(&minID)
	db.Exec("UPDATE user_tb SET username = ?, status = 'w' WHERE id = ?", name, minID)

	c.JSON(http.StatusOK, gin.H{
		"username": name,
		"iduser":   minID,
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
