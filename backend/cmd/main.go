package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/ridwantaufk/psi-tech-test/config"
	"github.com/ridwantaufk/psi-tech-test/models"
	"github.com/ridwantaufk/psi-tech-test/routes"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("tidak ada .env")
	}

	config.ConnectDB()

	// migrate tabel
	config.DB.AutoMigrate(
		&models.User{},
		&models.Company{},
		&models.Voucher{},
		&models.AuthUser{},
	)
	seedData()

	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type,Authorization")
		c.Header("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	routes.Setup(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}

func seedData() {
	var count int64
	config.DB.Model(&models.User{}).Count(&count)
	if count > 0 {
		return
	}

	users := []models.User{
		{ID: "12qwer", Nama: "Imron", Email: "", Telp: "081234567890"},
		{ID: "321rewq", Nama: "Juli", Email: "sammy@mail.com", Telp: "087654321"},
	}
	config.DB.Create(&users)

	companies := []models.Company{
		{ID: "trew098", UserID: "12qwer", CompanyCode: "SPI", CompanyName: ""},
		{ID: "poiuyt1234", UserID: "321rewq", CompanyCode: "PIC", CompanyName: "Samudera"},
	}
	config.DB.Create(&companies)

	voucher := models.Voucher{
		ID:       "v001",
		Code:     "DISKON50",
		Discount: 50,
		IsActive: true,
	}
	config.DB.Create(&voucher)
}
