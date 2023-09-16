package main

import (
	"connect_mysql_test/service"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

const (
	DB_HOST = "tcp(127.0.0.1:3306)"
	DB_NAME = "go_mysql"
	DB_USER = "root"
	DB_PASS = ""
)

func main() {
	dsn := DB_USER + ":" + DB_PASS + "@" + DB_HOST + "/" + DB_NAME + "?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("Cannot connect to MySQL:", err)
	}

	log.Println("Connected to MySQL:", db)

	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.POST("/items", service.CreateItem(db))
		v1.GET("/items", service.GetListOfItems(db))        // list items
		v1.GET("/items/:id", service.ReadItemById(db))      // get an item by ID
		v1.PUT("/items/:id", service.EditItemById(db))      // edit an item by ID
		v1.DELETE("/items/:id", service.DeleteItemById(db)) // delete an item by ID
	}

	router.Run()
}
