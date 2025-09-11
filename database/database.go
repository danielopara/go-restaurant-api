package database

import (
	"log"
	"os"

	"github.com/danielopara/restaurant-api/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() *gorm.DB{
	err := godotenv.Load()

	if err != nil{
		log.Fatal("error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")

	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err !=nil{
		log.Fatalf("failed to connect to database %v", err)
	}

	err = DB.AutoMigrate(&models.User{}, &models.Menu{}, &models.Order{}, &models.OrderItem{})

	if err != nil{
		log.Fatal("failed to migrate database")
	}

	log.Println("Database connection established")

	return DB

}