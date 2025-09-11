package order

import (
	"net/http"
	"strconv"

	"github.com/danielopara/restaurant-api/request"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderService OrderService
}

func NewOrderHandler(orderService OrderService) *OrderHandler {
	return &OrderHandler{orderService: orderService}
}

func (o *OrderHandler) DeleteOrderById(c *gin.Context){
	idParam := c.Param("id")
	
	if idParam == ""{
		idParam = c.Query("id")
	}

	if idParam == ""{
		c.JSON(http.StatusBadRequest, gin.H{"error": "no id param or query"})
		return
	}

	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := o.orderService.DeleteOrderById(uint(id)); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "order deleted"})
}

// find order by id
func (o *OrderHandler) FindOrderById(c *gin.Context){
	idParam := c.Param("id")
	if idParam == ""{
		idParam = c.Query("id")
	}

	if idParam == ""{
		c.JSON(http.StatusBadRequest, gin.H{"error":"no id in param"})
		return
	}

	id, err := strconv.Atoi(idParam)

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := o.orderService.FindOrderById(uint(id))
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": order})
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