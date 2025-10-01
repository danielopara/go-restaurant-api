package router

import (
	"net/http"

	"github.com/danielopara/restaurant-api/cache"
	"github.com/danielopara/restaurant-api/internal/auth"
	"github.com/danielopara/restaurant-api/internal/menu"
	"github.com/danielopara/restaurant-api/internal/order"
	"github.com/danielopara/restaurant-api/internal/user"
	"github.com/danielopara/restaurant-api/middleware"
	"github.com/danielopara/restaurant-api/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)




func Router(db *gorm.DB) *gin.Engine{
	r := gin.Default()

	newUserCache := cache.NewCache()


	userRepo := user.NewUserRepository(db)

	userCache := cache.NewUserCache(userRepo, newUserCache)
	userService := user.NewUserService(userCache)
	userHandler := user.NewUserHandler(userService)
	
	userGroup := r.Group("api/v1/user")

	{

/** 
		register user
		get all users
		get logged in user

*/

		userGroup.POST("/register", userHandler.Register)
		userGroup.GET("/users", middleware.AuthMiddleware(), middleware.RoleMiddleWare(models.RoleAdmin, models.RoleManager), userHandler.GetAllUsers)
		userGroup.GET("/", middleware.AuthMiddleware(), userHandler.GetLoggedInUser)
		
	}

	authService := auth.NewAuthService(userRepo)
	authHandler := auth.NewAuthHandler(authService)

	authGroup := r.Group("api/v1/auth")

	{
/** 
		login

*/
		authGroup.POST("/login", authHandler.Login)
	}

	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"running": true} )
	
	})

	menuRepo := menu.NewMenuRepository(db)
	cacheStore := cache.NewCache()
	cacheRepo := cache.NewCacheRepo(menuRepo, cacheStore)
	menuService := menu.NewMenuService(cacheRepo)
	menuHandlers := menu.NewMenuHandlers(menuService)

	menuGroup := r.Group("/api/v1/menu")
	{
/** 
		create menu item
		fetch a menu item
		get menu
		update menu item
		delete menu item

*/

		menuGroup.POST("/create-menu-item", 
		middleware.AuthMiddleware(), 
		middleware.RoleMiddleWare(models.RoleAdmin, models.RoleManager, models.RoleChef),
		menuHandlers.CreateMenuItem )

		menuGroup.GET("/item/:name", middleware.AuthMiddleware(), menuHandlers.FetchMenuItem)
		menuGroup.GET("/item", middleware.AuthMiddleware(), menuHandlers.FetchMenuItem)

		menuGroup.GET("/", middleware.AuthMiddleware(), menuHandlers.FetchMenu)

		menuGroup.PATCH("/", middleware.AuthMiddleware(),
		middleware.RoleMiddleWare(models.RoleAdmin, models.RoleManager, models.RoleChef),
		menuHandlers.PatchMenuItem)

		menuGroup.PATCH("/:id", middleware.AuthMiddleware(),
		middleware.RoleMiddleWare(models.RoleAdmin, models.RoleManager, models.RoleChef),
		menuHandlers.PatchMenuItem)
		
		menuGroup.DELETE("/", middleware.AuthMiddleware(),
		middleware.RoleMiddleWare(models.RoleAdmin, models.RoleManager, models.RoleChef),
		menuHandlers.DeleteMenuItem)

		menuGroup.DELETE("/:id", middleware.AuthMiddleware(),
		middleware.RoleMiddleWare(models.RoleAdmin, models.RoleManager, models.RoleChef),
		menuHandlers.DeleteMenuItem)
		
	}

	orderRepo := order.NewOrderRepository(db)
	orderService := order.NewOrderService(orderRepo, menuRepo, userRepo)
	orderHandler := order.NewOrderHandler(orderService)

	orderGroup := r.Group("api/v1/order")

	{
		orderGroup.GET("/orders", middleware.AuthMiddleware(), orderHandler.FindAllOrders)
		orderGroup.POST("/create-order", middleware.AuthMiddleware(), middleware.RoleMiddleWare(models.RoleWaiter),orderHandler.CreateOrder)
		orderGroup.GET("/id/:id", middleware.AuthMiddleware(), orderHandler.FindOrderById) 
		orderGroup.DELETE("/id/:id", middleware.AuthMiddleware(), middleware.RoleMiddleWare(models.RoleManager),  orderHandler.DeleteOrderById) 
		orderGroup.PUT("/id/:id", middleware.AuthMiddleware(), middleware.RoleMiddleWare(models.RoleAdmin, models.RoleCashier, models.RoleChef, models.RoleWaiter, models.RoleManager), orderHandler.UpdateOrderStatusById)
	}

	return r
}