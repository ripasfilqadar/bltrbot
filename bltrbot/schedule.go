package main

import (
	"github.com/ripasfilqadar/bltrbot/bltrbot/model"

	"fmt"
	"strconv"

	"github.com/ripasfilqadar/bltrbot/bltrbot/db"

	"time"

	"github.com/jasonlvhit/gocron"

	"encoding/json"
	"io/ioutil"
	"net/http"
)

func RunSchedule() {
	gocron.Every(1).Day().At("20:00").Do(reminderUser)
	gocron.Every(1).Day().At("06:00").Do(updateRemaining)
	Cities := []model.City{}
	db.MysqlDB().Find(&Cities)
	for _, city := range Cities {
		gocron.Every(1).Day().At("01:00").Do(getPrayerTime, city.Name)
	}
	<-gocron.Start()

}

func getPrayerTime(cityName string) {
	scraping_result := model.ScrapingResult{}
	resp, err := http.Get("https://time.siswadi.com/pray/" + cityName)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &scraping_result)
	timeNamePrayer := []string{"Fajr", "Dhuhr", "Asr", "Maghrib", "Isha"}
	for _, name := range timeNamePrayer {
		prayerTime := model.PrayerTime{}
		db.MysqlDB().Where("name = ? and city_name = ? ", name, cityName).First(&prayerTime)
		if prayerTime == (model.PrayerTime{}) {
			prayerTime = model.PrayerTime{
				Name:     name,
				Time:     scraping_result.Data[name],
				CityName: cityName,
			}
			db.MysqlDB().Create(&prayerTime)
		} else {
			db.MysqlDB().Model(&prayerTime).Update("time", scraping_result.Data[name])
		}
	}
	fmt.Println(scraping_result.Data)
	fmt.Println(scraping_result.Data["Sunset"])
}

//Task
func reminderUser() {
	template := "Yang belum laporan \n"
	fmt.Println("Debug")
	groups := []model.Group{}
	db.MysqlDB().Find(&groups)
	for _, group := range groups {
		users := []model.User{}
		db.MysqlDB().Where("group_id = ? and report_today = ? and state = ?", group.GroupId, false, "active").Find(&users)
		var username_users string
		for idx, user := range users {
			fmt.Println(user)
			username_users += strconv.Itoa(idx+1) + ") " + Emoji["not_confirm"] + user.FullName + "(@" + user.UserName + ") (" + strconv.Itoa(user.Target) + ")\n"
			fmt.Println(username_users)
			go Bot.SendToUser("Jangan lupa laporan di group "+group.Name, user.ChatId)
		}
		if len(users) == 0 {
			Bot.SendToGroup(group.GroupId, "Semua User sudah laporan")
		} else {
			Bot.SendToGroup(group.GroupId, template+username_users)
		}
	}
}

func updateRemaining() {
	iqob_date := time.Now().AddDate(0, 0, -1)
	groups := []model.Group{}
	db.MysqlDB().Find(&groups)
	for _, group := range groups {
		template := "Rekap " + DateFormat(iqob_date.Date()) + "\n"
		users := []model.User{}
		db.MysqlDB().Where("group_id = ?", group.GroupId).Find(&users)
		var username_users string
		template += ListMemberToday(users)
		for idx, user := range users {
			fmt.Println(user)
			if !user.ReportToday {
				fmt.Println("active bro")
				if user.State != "active" {
					continue
				}
				Bot.SendToUser("Karena kamu belum laporan di group "+group.Name+" , jangan lupa bayar iqob ya", user.ChatId)
				iqob := model.Iqob{
					UserId:   user.ID,
					State:    "not_paid",
					IqobDate: iqob_date,
					PaidAt:   iqob_date,
				}
				db.MysqlDB().Create(&iqob)
				username_users += strconv.Itoa(idx+1) + " ). " + StateEmoji(user) + " " + user.FullName + "(" + strconv.Itoa(user.Target) + " )\n"
			}
		}
		template += "\nList Iqob " + DateFormat(iqob_date.Date()) + "\n" + username_users

		Bot.SendToGroup(group.GroupId, template)
		db.MysqlDB().Model(&users).UpdateColumn("report_today", false)
	}
}
