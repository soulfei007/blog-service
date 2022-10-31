package main

import (
	"blog-service/global"
	"blog-service/internal/routers"
	"blog-service/pkg/setting"
	"log"
	"net/http"
	"time"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
}

func main() {
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

func setupSetting() error {
	set, err := setting.NewSetting()

	if err != nil {
		return err
	}

	global.ServerSetting = new(setting.ServerSettingS)
	err = set.ReadSection("server", global.ServerSetting)
	if err != nil {
		return err
	}

	global.AppSetting = new(setting.AppSettingS)
	err = set.ReadSection("app", global.AppSetting)

	if err != nil {
		return err
	}

	global.DatabaseSetting = new(setting.DatabaseSettingS)
	err = set.ReadSection("database", global.DatabaseSetting)

	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	return nil

}
