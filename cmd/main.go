package main

import (
	"github.com/AndrewFilkin/todo-app"
	"github.com/AndrewFilkin/todo-app/pkg/handler"
	"github.com/AndrewFilkin/todo-app/pkg/repository"
	"github.com/AndrewFilkin/todo-app/pkg/service"
	"log"
	"github.com/spf13/viper"
)

func main() {

	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	err := srv.Run(viper.GetString("8000"), handlers.InitRoutes())
	if err != nil {
		log.Fatalf("error occured while running http server %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}