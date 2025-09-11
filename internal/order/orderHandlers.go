package order

import (
	"net/http"

	"github.com/danielopara/restaurant-api/request"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderService OrderService
}

func NewOrderHandler(orderService OrderService) *OrderHandler {
	return &OrderHandler{orderService: orderService}
}

func (o *OrderHandler) CreateOrder(c *gin.Context){
	var req request.CreateOrder

	if err:= c.ShouldBindJSON(&req); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := o.orderService.MakeOrder(req.WaiterID, int(req.TableNo), req.Items)

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}



	c.JSON(http.StatusOK, gin.H{"success": true, "data": order})
}