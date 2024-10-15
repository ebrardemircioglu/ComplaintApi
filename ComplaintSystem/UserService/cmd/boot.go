package cmd

import (
	config2 "ComplaintSystem/UserService/config"
	"ComplaintSystem/UserService/internal"
	"ComplaintSystem/pkg"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

func BootUserService(client *mongo.Client, e *echo.Echo) {
	config := config2.GetUserConfig("dev")
	userCol, err := pkg.GetMongoCollection(client, config.DbConfig.DBName, config.DbConfig.ColName)
	if err != nil {
		panic(err)
	}

	repo := internal.NewRepository(userCol)
	service := internal.NewService(repo)
	internal.NewHandler(e, service)

	e.Logger.Fatal(e.Start(config.Port))

}
