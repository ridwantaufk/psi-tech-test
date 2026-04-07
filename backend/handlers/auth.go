package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/ridwantaufk/psi-tech-test/config"
	"github.com/ridwantaufk/psi-tech-test/models"
	"github.com/ridwantaufk/psi-tech-test/utils"
	"golang.org/x/crypto/bcrypt"
)

type RegisterReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	var req RegisterReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// cek usn (unique usn requirement)
	var existing models.AuthUser

	if result := config.DB.Where("username = ?", req.Username).First(&existing); result.Error == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "username sudah dipakai"})
		return
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal hash password"})
		return
	}

	user := models.AuthUser{
		ID:        uuid.New().String(),
		Username:  req.Username,
		Password:  string(hashed),
		CreatedAt: time.Now(),
	}

	config.DB.Create(&user)
	c.JSON(http.StatusCreated, gin.H{"message": "register berhasil", "id": user.ID})
}

func Login(c *gin.Context) {
	var req LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.AuthUser
	if err := config.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user tidak ditemukan"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "password salah"})
		return
	}

	token, err := utils.GenerateToken(user.ID, user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal membuat token"})
		return
	}

	//  cookie token
	c.SetCookie("access_token", token, 3600*24, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message":  "login berhasil",
		"token":    token,
		"id":       user.ID,
		"username": user.Username,
	})
}
func Logout(c *gin.Context) {
	c.SetCookie("access_token", "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "logout berhasil"})
}
