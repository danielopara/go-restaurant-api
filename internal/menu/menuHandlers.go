package menu

import (
	"net/http"
	"strconv"

	"github.com/danielopara/restaurant-api/models"
	"github.com/danielopara/restaurant-api/request"
	"github.com/danielopara/restaurant-api/response"
	"github.com/gin-gonic/gin"
)

type MenuHandlers struct {
	menuService MenuService
}

func NewMenuHandlers(menuService MenuService) *MenuHandlers {
	return &MenuHandlers{menuService: menuService}
}

// delete menu item
func(m *MenuHandlers) DeleteMenuItem(c *gin.Context){
	idParam := c.Param("id")
	if idParam == ""{
		idParam = c.Query("id")
	}

	if idParam == ""{
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is invalid"})
		return
	}

	if err= m.menuService.DeleteMenuItem(uint(id)); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "item deleted"})
}

func(m *MenuHandlers) PatchMenuItem(c *gin.Context){

	idParam := c.Param("id")
	if idParam == ""{
		idParam = c.Query("id")
	}

	if idParam == ""{
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is invalid"})
		return
	}

	var updates map[string]interface{}
	if err := c.ShouldBindJSON(&updates); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err:=m.menuService.UpdateMenuItem(uint(id), updates); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


	c.JSON(http.StatusOK, gin.H{"message": "updated successfully"})

}

func(m *MenuHandlers) FetchMenu(c *gin.Context){

	menu, err := m.menuService.GetMenu()

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": menu})
}


func (m *MenuHandlers) FetchMenuItem(c *gin.Context){
	nameParam := c.Param("name")

	if nameParam == ""{
		nameParam = c.Query("name")
	}

	if nameParam == ""{
		c.JSON(http.StatusBadRequest, gin.H{"error": "name is required"})
		return
	}

	menu, err := m.menuService.GetFoodByName(nameParam)

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res:=response.MenuResponse{
		Name: menu.Name,
		Price: menu.Price,
		Category: menu.Category,
	}

	c.JSON(http.StatusOK, gin.H{"data": res})
}


//create menuItem
func (m *MenuHandlers) CreateMenuItem(c *gin.Context){
	var req request.CreateMenuItem
	userEmail, exists := c.Get("email")

	if !exists{
		c.JSON(http.StatusUnauthorized, gin.H{"error": "no logged in user found"})
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	menuItems := &models.Menu{
		Name: req.Name,
		Category: models.Category(req.Category),
		Description: req.Description,
		Price: req.Price,
		Available: req.Available,
		CreatedBy: userEmail.(string),
	}

	res, err := m.menuService.CreateFood(menuItems)

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
	}

	menuRes := response.MenuResponse{
		Name: res.Name,
		Price: res.Price,
		Category: res.Category,
	}

	c.JSON(http.StatusCreated, gin.H{"data": menuRes})
}