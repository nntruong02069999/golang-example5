package main

import (
	"log"
	"github.com/urfave/cli/v2"
	"golang/example5/db"
	pb "golang/example5/proto"
	"os"
	"context"
)

var database = new(db.Db)

func main() {
	ctx , _ := context.WithCancel(context.Background())
	err := createCliGolang()
	if err != nil {
		panic("Stop program !!")
	}

	userPartner := &pb.UserPartnerRequest{
		UserId : "123",
		Phone : "0375860699",
		Limit : 10,
	}
	dsUserPartner, err := database.GetUserPartner(ctx, userPartner)
	if err != nil {
		log.Println(err)
	} else if len(dsUserPartner) != 0 {
		log.Println(dsUserPartner)
	} else {
		log.Println("Khong tim thay user partner")
	}

}

func createDb(c *cli.Context) error {
	err := database.ConnectDb()
	if err != nil {
		panic(err)
	}
	err = database.InitDatabase()
	if err != nil {
		panic("Create database faild")
	}
	return nil
}

func startApp(c *cli.Context) error {
	err := database.ConnectDb()
	if err != nil {
		panic(err)
	}
	return nil
}

func createCliGolang() error {
	app := cli.NewApp()
	app.Name = "cli-golang"
	app.Version = "0.0.1"
	app.Usage = "Using cli in golang to run app"
	addCommandCli(app)
	err := app.Run(os.Args)
	if err != nil {
		return err
	}
	return nil
}

func addCommandCli(app *cli.App) {
	app.Commands = []*cli.Command{
		{
			Name:   "createDb",
			Usage:  "Run command to create database",
			Action: createDb,
		},
		{
			Name:   "startApp",
			Usage:  "Run command to running app",
			Action: startApp,
		},
	}
}

