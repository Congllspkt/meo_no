package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"math/rand"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var dsn = "root:345FSDF$#@tcp(localhost:3306)/meono"
var db, _ = sql.Open("mysql", dsn)

func main() {

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},                      // Allow specific origin
		AllowMethods:     []string{"GET", "POST", "OPTIONS"}, // Allow specific methods
		AllowHeaders:     []string{"Origin", "Content-Type"}, // Allow specific headers
		ExposeHeaders:    []string{"Content-Length"},         // Expose headers
		AllowCredentials: true,
		MaxAge:           12 * 3600, // Cache duration for preflight requests
	}))

	r.GET("/register", register)
	r.GET("/getStatusGame", getStatusGame)
	r.GET("/startGame", startGame)
	r.GET("/sortBai", sortBai)


	r.GET("/skip", skip)
	r.GET("/reverse", reverse)
	r.GET("/rutbai", rutbai)
	r.GET("/see3", see3)

	r.Run()
}

func see3(c *gin.Context) {

	if !checkID(c.Query("id")) {
		c.Abort()
		return
	}

	// check user have see3
	var arr string
	db.QueryRow("SELECT arr FROM user_tb where id = ?;", c.Query("id")).Scan(&arr)
	bobaiuser := convertStringtoArray(arr)
	exists := false

    for _, num := range bobaiuser {
        if num == 3 {
            exists = true
            break
        }
    }

    if !exists {
        c.Abort()
		return
    }

	// rut see3
	bobaiusernew := removeOne(bobaiuser, 3)
	db.Exec("UPDATE user_tb set arr = ? where id = ?", joinIntSlice(bobaiusernew), c.Query("id"))


	// xem see3
	var bobai string
	db.QueryRow("SELECT bobai FROM game_tb;").Scan(&bobai)
	bobaigame := convertStringtoArray(bobai)
	see3 := bobaigame[:3]

	c.JSON(http.StatusOK, gin.H{
		"see3": see3,
	})
}

func rutbai(c *gin.Context) {
	if !checkID(c.Query("id")) {
		c.Abort()
		return
	}

	var bobai string
	db.QueryRow("SELECT bobai FROM game_tb;").Scan(&bobai)
	bobaigame := convertStringtoArray(bobai)
	bairut := bobaigame[0]
	fmt.Println(bobai)
	fmt.Println(bobaigame)
	fmt.Println(bairut)
	bobaigame = bobaigame[1:]
	fmt.Println(bobaigame)

	db.Exec("UPDATE game_tb set bobai = ?", joinIntSlice(bobaigame))
	var arr string
	db.QueryRow("SELECT arr FROM user_tb where id = ?;", c.Query("id")).Scan(&arr)
	arrNew := arr + "," + strconv.Itoa(bairut)
	db.Exec("UPDATE user_tb set arr = ? where id = ?", arrNew, c.Query("id"))

	// if rut trung mao no {}
	updateSkip(c)
}
func reverse(c *gin.Context) {
	if !checkID(c.Query("id")) {
		c.Abort()
		return
	}
	db.Exec("UPDATE game_tb SET rote = -rote")
	skipBai(c, 5)
}

func sortBai(c *gin.Context) {
	var arr string
	db.QueryRow("SELECT arr FROM user_tb where id = ?;", c.Query("id")).Scan(&arr)
	bobaiuser := convertStringtoArray(arr)
	sort.Ints(bobaiuser)
	db.Exec("UPDATE user_tb set arr = ? where id = ?", joinIntSlice(bobaiuser), c.Query("id"))
}

func skip(c *gin.Context) {
		if !checkID(c.Query("id")) {
		c.Abort()
		return
	}
	skipBai(c, 7)
}

func skipBai(c *gin.Context, bb int) {
	updateSkip(c)
	var arr string
	db.QueryRow("SELECT arr FROM user_tb where id = ?;", c.Query("id")).Scan(&arr)
	fmt.Print(arr)
	bobaiuser := convertStringtoArray(arr)
	fmt.Print(bobaiuser)
	arrNew := removeOne(bobaiuser, bb)
	fmt.Print(arrNew)
	db.Exec("UPDATE user_tb set arr = ? where id = ?", joinIntSlice(arrNew), c.Query("id"))
}

func updateSkip(c *gin.Context) {
	var ids string
	var next int
	db.QueryRow("SELECT GROUP_CONCAT(id) as ids FROM user_tb where username != ''").Scan(&ids)

	var rote int
	db.QueryRow("SELECT rote FROM game_tb").Scan(&rote)
	var numbers = convertStringtoArray(ids)

	for i, value := range numbers {
		if strconv.Itoa(value) == c.Query("id") {

			if rote == 1 && i == len(numbers)-1 {
				next = numbers[0]
			} else if rote == -1 && i == 0 {
				next = numbers[len(numbers)-1]
			} else {
				next = numbers[i+rote]
			}
			break
		}
	}
	db.Exec("UPDATE game_tb SET playuser = ?", next)
}

func checkID(id string) bool {
	var playuser string
	db.QueryRow("SELECT playuser FROM game_tb").Scan(&playuser)

	if playuser != id {
		return false
	}
	return true
}

func removeOne(arr []int, num int) []int {
	for i, value := range arr {
		if value == num {
			return append(arr[:i], arr[i+1:]...)
		}
	}
	return arr
}

func convertStringtoArray(ids string) []int {
	parts := strings.Split(ids, ",")
	var numbers []int
	for _, part := range parts {
		num, _ := strconv.Atoi(part)
		numbers = append(numbers, num)
	}
	return numbers
}

func startGame(c *gin.Context) {
	// get all id nguoi choi
	db.Exec("UPDATE user_tb set status = ''")
	db.Exec("UPDATE user_tb set status = 'p' where username != ''")

	var ids string
	db.QueryRow("SELECT GROUP_CONCAT(id) as ids FROM user_tb where username != ''").Scan(&ids)
	var numbers = convertStringtoArray(ids)
	fmt.Println("ids: ", numbers)

	numberPlayers := len(numbers)

	// meodefause2 := 1
	meobom1 := numberPlayers - 1

	meosee3 := 6
	meogive4 := 6
	meoreverse5 := 6
	meosuffle6 := 6
	meoskip7 := 6

	arrBobai := []int{}
	arrBobai = appendBobai(arrBobai, meosee3, 3)
	arrBobai = appendBobai(arrBobai, meogive4, 4)
	arrBobai = appendBobai(arrBobai, meoreverse5, 5)
	arrBobai = appendBobai(arrBobai, meosuffle6, 6)
	arrBobai = appendBobai(arrBobai, meoskip7, 7)
	fmt.Println("Tao Bo Bai : ", arrBobai)

	// trom bai
	shuffleSlice(arrBobai)
	shuffleSlice(arrBobai)
	shuffleSlice(arrBobai)
	fmt.Println("Hoan Vi Bo Bai : ", arrBobai)

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
		db.Exec("UPDATE user_tb SET arr = ? where id = ?", joinbai, numbers[i])
	}

	fmt.Println("Bo Bai : ", arrBobai)

	arrBobai = append(arrBobai, 2)
	arrBobai = appendBobai(arrBobai, meobom1, 1)
	shuffleSlice(arrBobai)
	shuffleSlice(arrBobai)
	shuffleSlice(arrBobai)

	db.Exec("UPDATE game_tb SET bobai = ?, playuser = ?", joinIntSlice(arrBobai), getRandomElement(numbers))

}

func getRandomElement(arr []int) int {
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(arr))
	return arr[randomIndex]
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

func appendBobai(arrBobai []int, n int, index int) []int {
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

func getStatusGame(c *gin.Context) {
	var status, arr, statusUser, messageStatusUser, playuser string
	db.QueryRow("SELECT status, playuser FROM game_tb;").Scan(&status, &playuser)

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

	rows, _ := db.Query("SELECT id, username, status  FROM user_tb where username != '';")
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
		"playUser":          playuser,
	})
}
