package main

import (
	"database/sql"
	"fmt"
	"os"
	"quiz-3/controllers"
	"quiz-3/database"
	"quiz-3/process"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	err error
)

func main() {
	// ENVIRONMENT
	err = godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("failed load file env")
	} else {
		fmt.Println("successfully read file env")
	}
	// DB Test
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGDATABASE"))

	DB, err = sql.Open("postgres", psqlInfo)
	err := DB.Ping()
	if err != nil {
		fmt.Println("DB Conn Failed")
	} else {
		fmt.Println("DB Conn Succeeded")
	}

	database.DbMigrate(DB)

	defer DB.Close()

	// Router bangun datar
	r := gin.Default()
	auth := r.Group("/", gin.BasicAuth(gin.Accounts{
		"admin":  "password",
		"editor": "secret",
	}))
	// router bangun datar
	r.GET("/segitiga-sama-sisi", process.SegitigaSamaSisi)
	r.GET("/persegi", process.Persegi)
	r.GET("/persegi-panjang", process.PersegiPanjang)
	r.GET("/lingkaran", process.Lingkaran)
	// router Categories
	r.GET("/categories", controllers.GetAllCategory)
	auth.POST("/categories", controllers.InsertCategory)
	auth.PUT("/categories/:id", controllers.UpdateCategory)
	r.DELETE("/categories/:id", controllers.DeleteCategory)
	// router book
	r.GET("/books", controllers.GetAllBook)
	auth.POST("/books", controllers.InsertBook)
	auth.PUT("/books/:id", controllers.UpdateBook)
	auth.DELETE("/books/:id", controllers.DeleteBook)

	r.Run(":" + os.Getenv("PORT"))

}
