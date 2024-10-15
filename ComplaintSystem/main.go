package main

import (
	"ComplaintSystem/AdminService/cmd"
	userCmd "ComplaintSystem/UserService/cmd"
	"ComplaintSystem/pkg"
	config "ComplaintSystem/shared"
	"github.com/labstack/echo/v4"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	//todo : what is dev,qa,prod ? explain why we are using them in the lecture
	dbConf := config.GetDBConfig("dev")

	client, err := pkg.GetMongoClient(dbConf.MongoDuration, dbConf.MongoClientURI)
	if err != nil {
		panic(err)
	}

	e := echo.New()

	//cmd.BootUserService(client, e)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Determine which service(s) to start based on command-line argument
	if len(os.Args) < 2 {
		panic("Please provide a service to start: admin,user or both")
	}
	input := os.Args[1]
	switch input {
	case "admin":
		cmd.BootAdminService(client, e)
	case "user":
		userCmd.BootUserService(client, e)
	case "both":
		go cmd.BootAdminService(client, e)
		go userCmd.BootUserService(client, e)
	default:
		panic("Invalid input")
	}
}

//challenge : after you create a func boot order service, manage somehow to run specific project
//description : when you give an input here it should look that input and boot THAT specific project
//if the input says "both" it should

//PS : do not forget to create and call a different column for order service and do not forget to boot order service
//from another port different than customer service

//orderCol, err := pkg.GetMongoCollection(client, "tesodev", "order")
//if err != nil {
//	panic(err)
