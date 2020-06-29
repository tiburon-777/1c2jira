package main

import (
	"github.com/go-martini/martini"
	"github.com/natefinch/lumberjack"
	"github.com/tiburon-777/1c2jira/src/params"
	"log"
)

func main() {
	log.Println("Запускаем сервис...")
	log.Println("Запускаем сервис...")
	pms := params.GetInstance()
	m := martini.Classic()
	// Запускаем логи в файл и воркер, только на ПРОДЕ!
	if pms.Development {
		f := &lumberjack.Logger{
			Filename:   pms.LogFile,
			MaxSize:    5, // megabytes
			MaxBackups: 20,
			MaxAge:     30,   //days
			Compress:   true, // disabled by default
		}
		log.SetOutput(f)
		m.Map(log.New(f, "[martini]", log.LstdFlags))
	}
	//
	// Запускаем сам сервис
	m.RunOnAddr("localhost:8080")
}
