package main

import (
	"fmt"
	"net/http"
	"os"
	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	loadEnv()

	connectionString := getDbConnectionString()

	fmt.Print(connectionString)

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic("Error connecting to db")
	}

	db.DB()

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"hello": "hello",
		})
	})

	router.Run()
}

func loadEnv(){
	if err := godotenv.Load(); err != nil {
        panic("Erorr loading env variables")
    }
}

func getDbConnectionString() string {

	dbName := os.Getenv("POSTGRES_DATABASE")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbPort := os.Getenv("POSTGRES_PORT")
	dbHost := os.Getenv("POSTGRES_HOST")

	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)

}
