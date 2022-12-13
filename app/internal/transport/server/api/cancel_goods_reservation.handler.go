package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/laterius/service_architecture_hw3/app/internal/service"
	"net/http"
)

// CancelGoodsReservationHandler handles HTTP request to cancel goods
func CancelGoodsReservationHandler(service service.Service) func(c *gin.Context) {
	type Body struct {
		OrderId uuid.UUID `json:"order_id"`
	}

	return func(c *gin.Context) {
		body := Body{}
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": err.Error(),
				"data":    gin.H{},
			})

			return
		}

		err := service.CancelReservation(body.OrderId)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"success": false,
				"message": err.Error(),
				"data":    gin.H{},
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "",
			"data":    gin.H{},
		})
	}
}
