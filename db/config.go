package db

import (
	"fmt"
	"log"
	"os"

	"github.com/itsharshitk/2_Blog/model"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() error {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env file")
	}

	user := os.Getenv("username")
	pass := os.Getenv("password")
	host := os.Getenv("hostname")
	port := os.Getenv("port")
	dbName := os.Getenv("dbName")

	dsn := user + ":" + pass + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=utf8&parseTime=True&loc=Local"

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error in database connection: ", err)
		return err
	}

	mysql, err := DB.DB()
	if err != nil {
		log.Fatal("Failed to connect mysql: ", err)
	}

	err = mysql.Ping()
	if err != nil {
		log.Fatal("Failed to ping the database: ", err)
	}

	fmt.Println("Database Connected Successfully.")

	allModels := []any{
		&model.User{},
		&model.Post{},
		&model.Comment{},
		&model.Like{},
	}

	if err := DB.AutoMigrate(allModels...); err != nil {
		log.Fatalf("Failed to Automigrate the databases: %v", err)
		return err
	}

	return nil
}
