package commands

import (
	"github.com/spf13/cobra"

	"github.com/pdxjohnny/numapp/get"
)

// Commands
var Commands = []*cobra.Command{
	&cobra.Command{
		Use:   "get",
		Short: "Get number",
		Run: func(cmd *cobra.Command, args []string) {
			ConfigBindFlags(cmd)
			get.Run()
		},
	},
}

func init() {
	ConfigDefaults(Commands...)
}
