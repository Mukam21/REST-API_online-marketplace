package handlers

import (
	"net/http"
	"regexp"

	"github.com/Mukam21/REST-API_online-marketplace/pkg/jwt"
	"github.com/Mukam21/REST-API_online-marketplace/pkg/models"
	"github.com/Mukam21/REST-API_online-marketplace/pkg/repository"
	"github.com/Mukam21/REST-API_online-marketplace/pkg/service"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthHandler struct {
	repo repository.UserRepository
}

func NewAuthHandler(db *gorm.DB) *AuthHandler {
	return &AuthHandler{
		repo: repository.NewUserRepository(db),
	}
}

type RegisterInput struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (h *AuthHandler) Register(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})

		return
	}

	if len(input.Login) < 4 || len(input.Password) < 6 || !regexp.MustCompile(`^[a-zA-Z0-9_]+$`).MatchString(input.Login) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Логин ≥ 4 символа, пароль ≥ 6, без спецсимволов"})

		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	user := models.User{Login: input.Login, Password: string(hashedPassword)}

	if err := h.repo.Create(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Пользователь уже существует"})

		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})

		return
	}

	user, err := h.repo.GetByLogin(input.Login)
	if err != nil || !service.CheckPasswordHash(input.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный логин или пароль"})

		return
	}

	token, _ := jwt.GenerateToken(user.ID)
	c.JSON(http.StatusOK, gin.H{"token": token})
}
