package main

import (
	"fmt"

	"bufio"

	"os"

	"strings"

	"github.com/ripasfilqadar/bltrbot/bltrbot/db"
	"github.com/ripasfilqadar/bltrbot/bltrbot/model"
)

var Emoji = map[string]string{
	"not_confirm": "👹",
	"smile":       "😇",
	"iqob":        "💀",
	"leave":       "✈",
}

func main() {
	fmt.Println("start")
	initEnv()
	InitRoute()
	InitDB()
	InitTelegram()
	//	reminderUser()
	//	updateRemaining()
	StartTelegram()
}

func InitDB() {
	db.MysqlDB().AutoMigrate(&model.User{}, &model.Report{}, &model.Iqob{}, &model.Group{})
}

func initEnv() {
	file, err := os.Open("../.env")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		envTemp := strings.Split(scanner.Text(), "=")
		os.Setenv(envTemp[0], envTemp[1])
	}
}
