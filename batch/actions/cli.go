package actions

import (
	"example/models"
	"log"
	"os"

	"github.com/urfave/cli"
)

func UpdateTest(c *cli.Context) error {
	log.Println("start UpdateTest")

	err := models.UpdateTest()

	log.Println("finish UpdateTest")
    
	return err
}

func Run() {
	app := cli.NewApp()
	app.Commands = []cli.Command{
		{
			Name:   "UpdateTest",
			Action: UpdateTest,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
