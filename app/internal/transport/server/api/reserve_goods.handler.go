package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/laterius/service_architecture_hw3/app/internal/service"
	"net/http"
)

// ReserveGoodsHandler handles HTTP request to reserve goods
func ReserveGoodsHandler(service service.Service) func(c *gin.Context) {
	type Body struct {
		OrderID uuid.UUID   `json:"order_id"`
		GoodIds []uuid.UUID `json:"goods_ids"`
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

		err := service.Reserve(body.OrderID, body.GoodIds)
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
			"data": gin.H{
				"order_id": body.OrderID,
				"good_ids": body.GoodIds,
			},
		})
	}
}
