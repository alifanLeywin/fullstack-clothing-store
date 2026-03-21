package controllers

import(
	"fmt"
	"math/rand"
	"net/http"
	"store-backend/config"
	"store-backend/models"
	"store-backend/utils"
	"time"

	"github.com/gin-gonic/gin"
)

type RegisterInput struct{
	Name string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func Register(c *gin.Context){
	var input RegisterInput


	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format data salah atau kurang lengkap bro!"})
		return
	}

	var existingUser models.User
	
	if err := config.DB.Where("email = ?", input.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Email ini udah dipake bos!"})
		return
	}

	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal encrypt password"})
		return
	}
	
	rand.Seed(time.Now().UnixNano())
	otp := fmt.Sprintf("%06d", rand.Intn(1000000))	

	newUser := models.User{
		Name: 				input.Name,
		Email: 				input.Email,
		Password: 			hashedPassword,
		Role: 				"user",
		IsVerified: 		false,
		VerificationCode: 	otp,
	}

	if err := config.DB.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal nyimpen data ke database"})
		return
	}

	fmt.Printf("\n📩 [EMAIL OTP] Kode buat %s : %s\n\n", newUser.Email, otp)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Berhasil daftar! Silakan cek email buat kode OTP.",
	})
	

}