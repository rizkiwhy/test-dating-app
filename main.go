package main

import (
	"fmt"
	"log"
	"rizkiwhy-dating-app/api/handler"
	"rizkiwhy-dating-app/api/router"
	"rizkiwhy-dating-app/config"
	"rizkiwhy-dating-app/pkg/order"
	relationshiptype "rizkiwhy-dating-app/pkg/relationship_type"
	swipehistory "rizkiwhy-dating-app/pkg/swipe_history"
	"rizkiwhy-dating-app/pkg/user"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	app := fiber.New()

	db, err := databaseConnection()
	if err != nil {
		log.Fatal(err)
	}

	swipeHistoryRepository := swipehistory.NewRepository(db)
	relationshipTypeRepository := relationshiptype.NewRepository(db)
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository.(*user.UserRepository), relationshipTypeRepository.(*relationshiptype.RelationshipTypeRepository), swipeHistoryRepository.(*swipehistory.SwipeHistoryRepository))
	userHandler := handler.NewUserHandler(userService.(*user.UserService))
	orderRepository := order.NewRepository(db)
	orderService := order.NewService(orderRepository.(*order.OrderRepository), userRepository.(*user.UserRepository))
	orderHandler := handler.NewOrderHandler(orderService.(*order.OrderService))
	api := app.Group("/api")
	router.UserRouter(api, userHandler.(*handler.UserHandler))
	router.OrderRouter(api, orderHandler.(*handler.OrderHandler))

	app.Listen(":3000")
}

func databaseConnection() (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("%s:@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Config("DB_USERNAME"), config.Config("DB_HOST"), config.Config("DB_PORT"), config.Config("DB_NAME"))

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}

	log.Println("database connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	return
}
