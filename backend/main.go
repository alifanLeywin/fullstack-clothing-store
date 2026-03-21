package main

import (
	"fmt"
	"store-backend/config"
	"store-backend/models"

	"github.com/gin-gonic/gin"
	
)

func main(){
		config.ConnectDatabase()

		fmt.Println("Database berhasil terhubung")
		config.DB.AutoMigrate(
			&models.User{},

		)

		r := gin.Default()

		r.GET("/ping", func(c *gin.Context){
			c.JSON(200, gin.H{
				"status": "success",
				"message": "Server V1",
			})
		})

		r.Run()
		
}