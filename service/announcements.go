package service

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

var (
	date string
	text string
)

func Announcements() string {
	request, errSql := db(viper.GetString("DB_CONNECT")).Query("SELECT * FROM announcements ORDER BY date LIMIT 3")
	if errSql != nil {
		log.Fatalln(errSql)
	}
	for request.Next() {
		errScan := request.Scan(&id, &date, &text)
		if errScan != nil {
			log.Fatalln(errScan)
		}
		clientMessageSuccessfully = fmt.Sprintf("🔔 %v", text)
		return clientMessageSuccessfully
	}
	return clientMessageSuccessfully
}
