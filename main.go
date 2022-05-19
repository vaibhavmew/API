package main

import (
	"API/models"
	"fmt"
	"github.com/gin-gonic/gin"
    	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql" 
)

var db *gorm.DB
var err error

func main() {
	db, err = gorm.Open("mysql", "<username>:<password>@tcp(127.0.0.1:3306)/<database_name>?charset=utf8&parseTime=True&loc=Local")
	
	if err != nil {
		fmt.Println(err)
	}
	
	defer db.Close()
	db.AutoMigrate(&models.User{})

	r := gin.Default()
	r.GET("/users", GetUsers)
	r.GET("/user/:id", GetUser)
	r.POST("/create", CreateUser)
	r.PATCH("/user/:id", UpdateUser)
	r.DELETE("/user/:id", DeleteUser)

	r.Run(":8080")
}

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatusJSON(400,  gin.H{"error": "Invalid Input"})
	} else {
		if user.Name == "" || user.DOB == "" || user.Address == "" || user.City == "" || user.State == "" || user.Pincode == "" {
			c.AbortWithStatusJSON(400,  gin.H{"error": "Enter all fields"})
		} else {
			db.Create(&user)
			c.JSON(200, user)
		}
	}
}

func GetUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(404,  gin.H{"error": "User Not Found"})
	} else {
		c.JSON(200, user)
	}
}

func GetUsers(c *gin.Context) {
	var users []models.User
	if err := db.Find(&users).Error; err != nil {
		c.AbortWithStatusJSON(404,  gin.H{"error": "User Database is empty"})
	} else {
		c.JSON(200, users)
	}
}

func UpdateUser(c *gin.Context) {
	var user models.User
	id := c.Params.ByName("id")
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(404,  gin.H{"error": "User Not Found"})
	} else {
		if err := c.BindJSON(&user); err != nil {
			c.AbortWithStatusJSON(404,  gin.H{"error": "Invalid Input"})
		} else {
			db.Save(&user)
			c.JSON(200, user)
		}
	}
}

func DeleteUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(404, gin.H{"error": "User Not Found"})
	} else {
		if err := db.Where("id = ?", id).Delete(&user).Error; err != nil {
			c.AbortWithStatusJSON(404, gin.H{"error": "User deletion was not successful"})
		} else {
			c.JSON(200, gin.H{"success": "Deleted the user with id => " + id})
		}
	}
}
