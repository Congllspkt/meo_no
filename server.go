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
	r.GET("/datmeono", datmeono)
	r.GET("/xaobai", xaobai)
	r.GET("/stealbai", stealbai)
	r.GET("/givesource", givesource)

	r.Run()
}

func givesource(c *gin.Context) {
	db.Exec("UPDATE game_tb set gd = ?", c.Query("idd"))
}

func stealbai(c *gin.Context) {
	db.Exec("UPDATE game_tb set gs = ?", c.Query("id"))
	rows, _ := db.Query("SELECT id, username, arr FROM user_tb where status = 'p' and id != ?;", c.Query("id"))
	var users []map[string]interface{}
	for rows.Next() {
		var id, username, arr string
		rows.Scan(&id, &username, &arr)
		nums := convertStringtoArray(arr)

		if len(nums) < 1 {
			continue
		}

		user := map[string]interface{}{
			"id":       id,
			"username": username,
		}
		users = append(users, user)
	}

	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func xaobai(c *gin.Context) {
	if !checkID(c.Query("id")) {
		c.Abort()
		return
	}

	var username string
	db.QueryRow("SELECT username FROM user_tb WHERE id = ?", c.Query("id")).Scan(&username)
	db.Exec("insert into log_tb (mm) values (?)", username+": da xao bai")
	db.Exec("UPDATE game_tb set bai = ?", 6)

	// check user have xaobai
	var arr string
	db.QueryRow("SELECT arr FROM user_tb where id = ?;", c.Query("id")).Scan(&arr)
	bobaiuser := convertStringtoArray(arr)
	exists := false

	for _, num := range bobaiuser {
		if num == 6 {
			exists = true
			break
		}
	}

	if !exists {
		c.Abort()
		return
	}

	// rut saobai
	bobaiusernew := removeOne(bobaiuser, 6)
	bobaiusernew1 := removeOne(bobaiusernew, 0)
	db.Exec("UPDATE user_tb set arr = ? where id = ?", joinIntSlice(bobaiusernew1), c.Query("id"))

	// xaobai
	var bobai string
	db.QueryRow("SELECT bobai FROM game_tb;").Scan(&bobai)
	bobaigameee := convertStringtoArray(bobai)
	shuffleSlice(bobaigameee)
	bobaigameee = removeOne(bobaigameee, 0)
	db.Exec("UPDATE game_tb set bobai = ?", joinIntSlice(bobaigameee))
}

func datmeono(c *gin.Context) {
	if !checkID(c.Query("id")) {
		c.Abort()
		return
	}

	var username string
	db.QueryRow("SELECT username FROM user_tb WHERE id = ?", c.Query("id")).Scan(&username)
	db.Exec("insert into log_tb (mm) values (?)", username+": rut meo no")
	db.Exec("insert into log_tb (mm) values (?)", username+": bi mat 1 la defuse")
	db.Exec("insert into log_tb (mm) values (?)", username+": da nhet 1 la meo no vao bo bai")
	db.Exec("UPDATE game_tb set bai = ?", 1)
	db.Exec("UPDATE user_tb set bom = '1' where id = ?", c.Query("id"))

	var bobai string
	db.QueryRow("SELECT bobai FROM game_tb;").Scan(&bobai)
	bobaigame := convertStringtoArray(bobai)
	rand.Seed(time.Now().UnixNano())

	typeG := c.Query("type")
	var idg int
	if typeG == "1" {
		idg = 0
	} else if typeG == "2" {
		idg = 1
	} else if typeG == "n" {
		idg = len(bobaigame)
	} else {
		idg = rand.Intn(len(bobaigame))
	}
	bobainew := insertAtPosition(bobaigame, 1, idg)
	bobainew = removeOne(bobainew, 0)

	db.Exec("UPDATE game_tb set bobai = ?", joinIntSlice(bobainew))
	updateSkip(c)
}

func insertAtPosition(arr []int, num int, position int) []int {
	if position < 0 || position > len(arr) {
		fmt.Println("Invalid position")
		return arr
	}
	arr = append(arr[:position], append([]int{num}, arr[position:]...)...)
	return arr
}

func see3(c *gin.Context) {
	if !checkID(c.Query("id")) {
		c.Abort()
		return
	}

	var username string
	db.QueryRow("SELECT username FROM user_tb WHERE id = ?", c.Query("id")).Scan(&username)
	db.Exec("insert into log_tb (mm) values (?)", username+": danh bai see a furture")
	db.Exec("UPDATE game_tb set bai = ?", 3)

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
	bobaiusernew1 := removeOne(bobaiuser, 3)
	bobaiusernew2 := removeOne(bobaiusernew1, 0)
	db.Exec("UPDATE user_tb set arr = ? where id = ?", joinIntSlice(bobaiusernew2), c.Query("id"))

	// xem see3
	var bobai string
	db.QueryRow("SELECT bobai FROM game_tb;").Scan(&bobai)
	bobaigame := convertStringtoArray(bobai)
	see3 := bobaigame[:3]

	see3 = removeOne(see3, 0)
	see3 = removeOne(see3, 0)
	see3 = removeOne(see3, 0)

	c.JSON(http.StatusOK, gin.H{
		"see3": see3,
	})
}

func rutbai(c *gin.Context) {
	if !checkID(c.Query("id")) {
		c.Abort()
		return
	}

	var username string
	db.QueryRow("SELECT username FROM user_tb WHERE id = ?", c.Query("id")).Scan(&username)
	db.Exec("insert into log_tb (mm) values (?)", username+": rut bai")

	var bobai string
	db.QueryRow("SELECT bobai FROM game_tb;").Scan(&bobai)
	bobaigame := convertStringtoArray(bobai)
	bairut := bobaigame[0]
	bobaigame = bobaigame[1:]
	bobaigame = removeOne(bobaigame, 0)

	db.Exec("UPDATE game_tb set bobai = ?", joinIntSlice(bobaigame))

	//trung meo no
	if bairut == 1 {
		var arr string
		db.QueryRow("SELECT arr FROM user_tb where id = ?;", c.Query("id")).Scan(&arr)
		bobaiuser := convertStringtoArray(arr)
		exists := false

		for _, num := range bobaiuser {
			if num == 2 {
				exists = true
				break
			}
		}
		if exists {
			bobaiusernew1 := removeOne(bobaiuser, 2)
			bobaiusernew2 := removeOne(bobaiusernew1, 0)
			db.Exec("UPDATE user_tb set arr = ? where id = ?", joinIntSlice(bobaiusernew2), c.Query("id"))
			c.JSON(http.StatusOK, gin.H{
				"datmeono": "datmeono",
			})
		} else {
			updateSkip(c)
			db.Exec("UPDATE user_tb set status = 'd' where id = ?", c.Query("id"))
			var username string
			db.QueryRow("SELECT username FROM user_tb WHERE id = ?", c.Query("id")).Scan(&username)
			db.Exec("insert into log_tb (mm) values (?)", username+": thua")
			db.Exec("UPDATE game_tb set bai = ?", 1)
		}
		return
	}

	var arr string
	db.QueryRow("SELECT arr FROM user_tb where id = ?;", c.Query("id")).Scan(&arr)
	arrNew := arr + "," + strconv.Itoa(bairut)
	db.Exec("UPDATE user_tb set arr = ? where id = ?", arrNew, c.Query("id"))
	updateSkip(c)

}
func reverse(c *gin.Context) {

	if !checkID(c.Query("id")) {
		fmt.Println("??")
		c.Abort()
		return
	}

	var username string
	db.QueryRow("SELECT username FROM user_tb WHERE id = ?", c.Query("id")).Scan(&username)
	fmt.Println(username)
	db.Exec("insert into log_tb (mm) values (?)", username+": danh bai reverse")
	db.Exec("UPDATE game_tb set bai = ?", 5)

	db.Exec("UPDATE game_tb SET rote = -rote")
	skipBai(c, 5)
}

func sortBai(c *gin.Context) {
	var arr string
	db.QueryRow("SELECT arr FROM user_tb where id = ?;", c.Query("id")).Scan(&arr)
	bobaiuser := convertStringtoArray(arr)
	sort.Ints(bobaiuser)
	bobaiuser = removeOne(bobaiuser, 0)
	db.Exec("UPDATE user_tb set arr = ? where id = ?", joinIntSlice(bobaiuser), c.Query("id"))
}

func skip(c *gin.Context) {
	if !checkID(c.Query("id")) {
		c.Abort()
		return
	}

	var username string
	db.QueryRow("SELECT username FROM user_tb WHERE id = ?", c.Query("id")).Scan(&username)
	db.Exec("insert into log_tb (mm) values (?)", username+": danh bai skip")
	db.Exec("UPDATE game_tb set bai = ?", 7)

	skipBai(c, 7)
}

func skipBai(c *gin.Context, bb int) {
	updateSkip(c)
	var arr string
	db.QueryRow("SELECT arr FROM user_tb where id = ?;", c.Query("id")).Scan(&arr)
	fmt.Print(arr)
	bobaiuser := convertStringtoArray(arr)
	fmt.Print(bobaiuser)
	arrNew1 := removeOne(bobaiuser, bb)
	arrNew2 := removeOne(arrNew1, 0)
	db.Exec("UPDATE user_tb set arr = ? where id = ?", joinIntSlice(arrNew2), c.Query("id"))
}

func updateSkip(c *gin.Context) {
	var ids string
	var next int
	db.QueryRow("SELECT GROUP_CONCAT(id) as ids FROM user_tb where username != '' and status = 'p'").Scan(&ids)

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
	fmt.Println(playuser, id)
	return playuser == id
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

	db.Exec("DELETE from log_tb")
	db.Exec("insert into log_tb (mm) values ('bat dau game')")

	// get all id nguoi choi
	db.Exec("UPDATE user_tb set status = ''")
	db.Exec("UPDATE user_tb set status = 'p', bom = '0' where username != ''")

	var ids string
	db.QueryRow("SELECT GROUP_CONCAT(id) as ids FROM user_tb where username != ''").Scan(&ids)
	var numbers = convertStringtoArray(ids)
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

	// trom bai
	shuffleSlice(arrBobai)
	shuffleSlice(arrBobai)
	shuffleSlice(arrBobai)

	// chi bai
	for i := 0; i < numberPlayers; i++ {
		baiUser := arrBobai[:4]
		arrBobai = arrBobai[4:]
		var bai = []int{2}
		bai = append(bai, baiUser...)
		// save arr
		bai = removeOne(bai, 0)
		joinbai := joinIntSlice(bai)
		db.Exec("UPDATE user_tb SET arr = ? where id = ?", joinbai, numbers[i])
	}
	arrBobai = append(arrBobai, 2)
	arrBobai = appendBobai(arrBobai, meobom1, 1)
	shuffleSlice(arrBobai)
	shuffleSlice(arrBobai)
	shuffleSlice(arrBobai)

	nguoididau := getRandomElement(numbers)
	arrBobai = removeOne(arrBobai, 0)
	db.Exec("UPDATE game_tb SET bobai = ?, playuser = ?, rote = 1, bai = 0", joinIntSlice(arrBobai), nguoididau)

	var username string
	db.QueryRow("SELECT username FROM user_tb WHERE id = ?", nguoididau).Scan(&username)
	db.Exec("insert into log_tb (mm) values (?)", username+" di dau tien")

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
	var status, arr, statusUser, messageStatusUser, playuser, rote, bobai, bai, bom string
	db.QueryRow("SELECT status, playuser, rote, bobai, bai FROM game_tb;").Scan(&status, &playuser, &rote, &bobai, &bai)

	var messageStatus string
	var statusGame string
	var bobaiarr = convertStringtoArray(bobai)

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

	var arrs string
	rows, _ := db.Query("SELECT id, username, status, bom, arr as arrs  FROM user_tb where username != '';")
	var users []map[string]interface{}
	for rows.Next() {
		var id, username, status string
		rows.Scan(&id, &username, &status, &bom, &arrs)
		arrUsers := convertStringtoArray(arrs)
		arrUsers = removeOne(arrUsers, 0)
		user := map[string]interface{}{
			"id":       id,
			"username": username,
			"status":   status,
			"bom":      bom,
			"arrs":     len(arrUsers),
		}
		users = append(users, user)
	}

	rowms, _ := db.Query("SELECT mm FROM log_tb ORDER BY id desc LIMIT 20;")
	var logs []string
	for rowms.Next() {
		var mm string
		rowms.Scan(&mm)
		logs = append(logs, mm)

	}

	c.JSON(http.StatusOK, gin.H{
		"statusGame":        statusGame,
		"messageStatus":     messageStatus,
		"arr":               arr,
		"statusUser":        statusUser,
		"messageStatusUser": messageStatusUser,
		"allUser":           users,
		"playUser":          playuser,
		"rote":              rote,
		"sobaiconlai":       len(bobaiarr),
		"bai":               bai,
		"log":               logs,
	})
}
