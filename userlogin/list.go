// +build windows

package userlogin

import (
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/octoblu/go-meshblu-connector-service/manage"
	"github.com/urfave/cli"
)

// List lists out all services installed on this system
// with service type user-login
func List(context *cli.Context) {
	localAppData := context.String("local-app-data")
	if localAppData == "" {
		cli.ShowCommandHelp(context, "user-login")

		fmt.Println()
		color.Red("Missing required argument --local-app-data, -l, or LOCALAPPDATA")
		os.Exit(1)
	}

	uuids, err := manage.ListUserLogin(localAppData)
	if err != nil {
		log.Fatalln(err.Error())
	}
	for _, uuid := range uuids {
		fmt.Println(uuid)
	}
	os.Exit(0)
}
