// +build !windows

package userlogin

import (
	"os"

	"github.com/fatih/color"
	"github.com/urfave/cli"
)

// List is only available on windows, this function exits the process
// with a status of 1
func List(context *cli.Context) {
	color.Red("user-login commands are available on windows")
	os.Exit(1)
}

// UpgradeIgnition is only available on windows, this function exits the process
// with a status of 1
func UpgradeIgnition(context *cli.Context) {
	color.Red("user-login commands are available on windows")
	os.Exit(1)
}
