package service

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"math"
)

var (
	lessonDay                 = current.Weekday()
	clientMessageSuccessfully string
	id                        int
	subject                   string
	weekDay                   int8
	lessonBeginTime           string
	teacher                   string
	teacherSex                string
	room                      string
)

type Scheduler interface {
	Previous() Response
	Current() Response
	Next() Response
}

type Request struct {
	User string
}

type Response struct {
	Status, Result string
}

func getTeacher(teacherName, teacherSex string) (teacher string) {
	if teacherSex == "male" {
		teacher = fmt.Sprintf("👨‍🏫 Учитель: %v.\n", teacherName)
	} else {
		teacher = fmt.Sprintf("👩‍🏫 Учительница: %v.\n", teacherName)
	}
	return teacher
}

func (r Request) Previous() Response {
	sqlPrevious := "SELECT * FROM lesson_table WHERE week_day = $1 AND begin_time = $2"
	request, errSql := db(viper.GetString("DB_CONNECT")).Query(sqlPrevious, int(lessonDay), previousLesson().lessonStartTime)
	if errSql != nil {
		log.Fatalf("Ошибка запроса: %v", errSql)
	}
	for request.Next() {
		errScan := request.Scan(&id, &subject, &weekDay, &lessonBeginTime, &teacher, &teacherSex, &room)
		if errScan != nil {
			log.Fatalf("Ошибка получения данных: %v", errScan)
		}
		clientMessageSuccessfully = fmt.Sprintf(
			"🤖 Привет %v! %v-я пара уже закончил. ⛔\n\n"+
				"%v"+
				"📚 Предмет: %v.\n "+
				"👉 Кабинет: %v.\n",
			r.User, previousLesson().lessonStage,
			getTeacher(teacher, teacherSex),
			subject,
			room,
		)
		return Response{Status: "SUCCESSFULLY", Result: clientMessageSuccessfully}
	}
	return Response{Status: "ERROR", Result: "❌ По расписание занятий предыдущей уроки не найдено."}
}

func (r Request) Current() Response {
	sqlCurrent := "SELECT * FROM lesson_table WHERE week_day = $1 AND begin_time = $2"
	request, errSql := db(viper.GetString("DB_CONNECT")).Query(sqlCurrent, int(lessonDay), currentLesson().lessonStartTime)
	if errSql != nil {
		log.Fatalf("Ошибка запроса: %v", errSql)
	}
	for request.Next() {
		errScan := request.Scan(&id, &subject, &weekDay, &lessonBeginTime, &teacher, &teacherSex, &room)
		if errScan != nil {
			log.Fatalf("Ошибка получения данных: %v", errScan)
		}
		clientMessageSuccessfully = fmt.Sprintf(
			"🤖 Привет %v! Сейчас %v-я пара. ✅\n\n"+
				"%v"+
				"📚 Предмет: %v.\n "+
				"👉 Кабинет: %v.\n"+
				"⏰ Пара заканчивается через: %v минут.",
			r.User, currentLesson().lessonStage,
			getTeacher(teacher, teacherSex),
			subject,
			room,
			math.Round(lessonCountdown),
		)
		return Response{Status: "SUCCESSFULLY", Result: clientMessageSuccessfully}
	}
	return Response{Status: "ERROR", Result: "❌ По расписание занятий уроки не найдено."}
}

func (r Request) Next() Response {
	sqlNext := "SELECT * FROM lesson_table WHERE week_day = $1 AND begin_time = $2"
	request, errSql := db(viper.GetString("DB_CONNECT")).Query(sqlNext, int(lessonDay), nextLesson().lessonStartTime)
	if errSql != nil {
		log.Fatalf("Ошибка запроса: %v", errSql)
	}
	for request.Next() {
		errScan := request.Scan(&id, &subject, &weekDay, &lessonBeginTime, &teacher, &teacherSex, &room)
		if errScan != nil {
			log.Fatalf("Ошибка получения данных: %v", errScan)
		}
		clientMessageSuccessfully = fmt.Sprintf(
			"🤖 Привет %v! Следующий %v-я пара. ⏳\n\n"+
				"%v"+
				"📚 Предмет: %v.\n "+
				"👉 Кабинет: %v.\n"+
				"⏰ До следующего пара осталось: %v минут.",
			r.User, lessonStage,
			getTeacher(teacher, teacherSex),
			subject,
			room,
			math.Round(lessonCountdown),
		)
		return Response{Status: "SUCCESSFULLY", Result: clientMessageSuccessfully}
	}
	return Response{Status: "ERROR", Result: "❌ Это последний пара. После этого уроки не планировали."}
}
