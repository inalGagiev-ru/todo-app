package main

import (
	"log"
	"os"

	"github.com/inalGagiev-ru/todo-app"
	"github.com/inalGagiev-ru/todo-app/pkg/handler"
	"github.com/inalGagiev-ru/todo-app/pkg/models"
	"github.com/inalGagiev-ru/todo-app/pkg/repository"
	"github.com/inalGagiev-ru/todo-app/pkg/service"

	"fmt"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func main() {
	currentDir, _ := os.Getwd()
	fmt.Printf("Current directory: %s\n", currentDir)

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Trying to load .env from project root...")
		err = godotenv.Load("../.env")
		if err != nil {
			fmt.Println("Trying to load .env from current directory...")
			err = godotenv.Load(".env")
		}
	}

	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	InitDB(db)

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while runninng http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("../configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func InitDB(db *gorm.DB) {
	err := db.AutoMigrate(&models.User{}, &models.Task{}, &models.Category{}, &models.Tag{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
}
