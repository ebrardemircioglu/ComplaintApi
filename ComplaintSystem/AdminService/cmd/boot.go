package cmd

import (
	config3 "ComplaintSystem/AdminService/config"
	"ComplaintSystem/AdminService/internal"
	"ComplaintSystem/pkg"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

func BootAdminService(client *mongo.Client, e *echo.Echo) {
	config := config3.GetAdminConfig("dev")
	adminCol, err := pkg.GetMongoCollection(client, config.DbConfig.DBName, config.DbConfig.ColName)
	if err != nil {
		panic(err)
	}

	repo := internal.NewRepository(adminCol)
	service := internal.NewService(repo)
	internal.NewHandler(e, service)

	e.Logger.Fatal(e.Start(config.Port))

}
