package handlers

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/Mukam21/REST-API_online-marketplace/pkg/models"
	"github.com/Mukam21/REST-API_online-marketplace/pkg/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrderHandler struct {
	repo repository.OrderRepository
}

func NewOrderHandler(db *gorm.DB) *OrderHandler {
	return &OrderHandler{
		repo: repository.NewOrderRepository(db),
	}
}

type CreateOrderInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
	Price       uint   `json:"price"`
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	userID := c.GetUint("user_id")

	var input CreateOrderInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат"})
		return
	}

	if len(input.Title) > 100 || len(input.Description) > 500 || input.Price == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Превышен размер полей или некорректная цена"})
		return
	}

	if _, err := url.ParseRequestURI(input.ImageURL); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректная ссылка"})
		return
	}

	order := models.Order{
		Title:       input.Title,
		Description: input.Description,
		ImageURL:    input.ImageURL,
		Price:       input.Price,
		UserID:      userID,
	}

	if err := h.repo.Create(&order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании"})
		return
	}

	c.JSON(http.StatusOK, order)
}

func getUserIDFromContext(c *gin.Context) uint {
	userIDInterface, exists := c.Get("userID")
	if !exists {
		return 0
	}
	userID, ok := userIDInterface.(uint)
	if !ok {
		return 0
	}
	return userID
}

func (h *OrderHandler) GetOrders(c *gin.Context) {

	minPriceStr := c.DefaultQuery("min_price", "0")
	maxPriceStr := c.DefaultQuery("max_price", "0")
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("page_size", "10")
	sortBy := c.DefaultQuery("sort_by", "created_at")
	order := c.DefaultQuery("order", "desc")

	minPrice, _ := strconv.Atoi(minPriceStr)
	maxPrice, _ := strconv.Atoi(maxPriceStr)
	page, _ := strconv.Atoi(pageStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)

	orders, err := h.repo.GetFiltered(minPrice, maxPrice, page, pageSize, sortBy, order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
		return
	}

	userID := getUserIDFromContext(c)

	var result []gin.H
	for _, order := range orders {
		isMine := false
		if userID != 0 && order.UserID == userID {
			isMine = true
		}
		result = append(result, gin.H{
			"id":          order.ID,
			"title":       order.Title,
			"description": order.Description,
			"price":       order.Price,
			"image_url":   order.ImageURL,
			"user_id":     order.UserID,
			"is_mine":     isMine,
			"created_at":  order.CreatedAt,
		})
	}

	c.JSON(http.StatusOK, result)
}
