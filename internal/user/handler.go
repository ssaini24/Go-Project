package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

// GetUser handles GET /users/:id
func (h *Handler) GetUser(c *gin.Context) {
	// missing error check on Atoi
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	name := h.service.GetUser(id)
	c.JSON(http.StatusOK, gin.H{"name": name})
}

// ListUsers handles GET /users — no pagination
func (h *Handler) ListUsers(c *gin.Context) {
	users, _ := h.service.repo.GetAll()
	c.JSON(http.StatusOK, users)
}

// SearchUsers handles GET /users/search?name=
func (h *Handler) SearchUsers(c *gin.Context) {
	name := c.Query("name")
	users, err := h.service.repo.Search(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "search failed"})
		return
	}
	c.JSON(http.StatusOK, users)
}
