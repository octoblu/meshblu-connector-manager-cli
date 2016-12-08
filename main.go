package main

import (
	"fmt"
	"log"
	"os"

	"github.com/coreos/go-semver/semver"
	"github.com/octoblu/meshblu-connector-manager-cli/userlogin"
	"github.com/urfave/cli"
	De "github.com/visionmedia/go-debug"
)

var debug = De.Debug("meshblu-connector-manager-cli:main")

func main() {
	app := cli.NewApp()
	app.Name = "meshblu-connector-manager-cli"
	app.Version = version()
	app.Flags = []cli.Flag{}
	app.Commands = []cli.Command{
		{
			Name:    "user-login",
			Aliases: []string{"ul"},
			Usage:   "manage connectors installed as a UserLogin service (windows only)",
			Subcommands: []cli.Command{
				{
					Name:    "list",
					Aliases: []string{"ls"},
					Usage:   "list out currently installed services",
					Action:  userlogin.List,
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:   "local-app-data, l",
							Usage:  "Local AppData directory of the user.",
							EnvVar: "LOCALAPPDATA",
						},
					},
				},
				{
					Name:      "upgrade-ignition",
					Aliases:   []string{"ui"},
					Usage:     "upgrades the ignition script for the installed service",
					ArgsUsage: "<uuid> [<uuid>...]",
					Action:    userlogin.UpgradeIgnition,
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:   "local-app-data, l",
							Usage:  "Local AppData directory of the user.",
							EnvVar: "LOCALAPPDATA",
						},
						cli.BoolFlag{
							Name:   "all, a",
							Usage:  "Upgrade all installed services. Cannot be used in combination with --uuid",
							EnvVar: "MESHBLU_CONNECTOR_MANAGER_ALL",
						},
						cli.StringFlag{
							Name:   "version-tag, t",
							Usage:  "Version to upgrade the ignition script to. (ex: v8.2.0)",
							EnvVar: "MESHBLU_CONNECTOR_MANAGER_UUID",
						},
					},
				},
			},
		},
	}
	app.Run(os.Args)
}

func version() string {
	version, err := semver.NewVersion(VERSION)
	if err != nil {
		errorMessage := fmt.Sprintf("Error with version number: %v", VERSION)
		log.Panicln(errorMessage, err.Error())
	}
	return version.String()
}
