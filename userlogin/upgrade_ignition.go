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

// UpgradeIgnition upgrades the ignition scripts in services that have been installed as user-login
func UpgradeIgnition(context *cli.Context) {
	var err error
	localAppData, versionTag, all, uuids := upgradeIgnitionGetOpts(context)

	if all {
		uuids, err = manage.ListUserLogin(localAppData)
		if err != nil {
			log.Fatalln(err.Error())
		}
	}

	for _, uuid := range uuids {
		err = manage.UpgradeIgnitionUserLogin(localAppData, versionTag, uuid)
		if err != nil {
			log.Fatalln("Error updating", uuid, err.Error())
		}

	}
	fmt.Println("UpdateIgnition", localAppData, versionTag, all, uuids)
}

func upgradeIgnitionGetOpts(context *cli.Context) (string, string, bool, []string) {
	localAppData := context.String("local-app-data")
	versionTag := context.String("version-tag")
	all := context.Bool("all")

	uuids := make([]string, len(context.Args()))
	for i, arg := range context.Args() {
		uuids[i] = arg
	}

	if localAppData == "" || versionTag == "" || (all && len(uuids) > 0) || (!all && len(uuids) == 0) {
		cli.ShowCommandHelp(context, "user-login")

		fmt.Println()
		if localAppData == "" {
			color.Red("Missing required argument --local-app-data, -l, or LOCALAPPDATA")
		}
		if versionTag == "" {
			color.Red("Missing required argument --version-tag, -t, MESHBLU_CONNECTOR_MANAGER_VERSION_TAG")
		}
		if all && len(uuids) > 0 {
			color.Red("Received both --all and at least one <uuid>. Only one at a time may be used.")
		}
		if !all && len(uuids) == 0 {
			color.Red("Must be called with either --all or at least one <uuid>")
		}

		os.Exit(1)
	}

	return localAppData, versionTag, all, uuids
}
