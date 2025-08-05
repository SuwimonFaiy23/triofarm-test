package main

import (
	"log"

	"github.com/SuwimonFaiy23/triofarm-test/internal/db"
	"github.com/SuwimonFaiy23/triofarm-test/internal/handler"
	"github.com/SuwimonFaiy23/triofarm-test/internal/repository"
	"github.com/SuwimonFaiy23/triofarm-test/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	conn := db.Connect()
	sqlDB, err := conn.DB()
	if err != nil {
		log.Fatalf("Failed to get sql.DB from gorm.DB: %v", err)
	}
	defer sqlDB.Close()

	menuRepo := repository.NewMenuRepository(conn)
	menuService := service.NewMenuService(menuRepo)
	menuHandler := handler.NewMenuHandler(menuService)

	itemRepo := repository.NewItemRepository(conn)
	itemService := service.NewItemService(itemRepo, menuRepo)
	itemHandler := handler.NewItemHandler(itemService)

	r.POST("/api/v1/menus", menuHandler.CreateMenu)
	r.PUT("/api/v1/menus", menuHandler.UpdateMenu)
	r.DELETE("/api/v1/menus/:id", menuHandler.DeleteMenu)
	r.GET("/api/v1/menus", menuHandler.GetMenuList)

	r.POST("/api/v1/items", itemHandler.CreateItem)
	r.PUT("/api/v1/items", itemHandler.UpdateItem)
	r.PUT("/api/v1/items/index", itemHandler.UpdateIndexMenu)
	r.DELETE("/api/v1/items/:id", itemHandler.DeleteItem)
	r.GET("/api/v1/items/:id", itemHandler.GetItemList)

	r.Run(":8080")
}
