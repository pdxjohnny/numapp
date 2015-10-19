package commands

import (
	"github.com/spf13/cobra"

	"github.com/pdxjohnny/numapp/get"
	"github.com/pdxjohnny/numapp/put"
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
	&cobra.Command{
		Use:   "put",
		Short: "Put a doc in the database",
		Run: func(cmd *cobra.Command, args []string) {
			ConfigBindFlags(cmd)
			put.Run()
		},
	},
}

func init() {
	ConfigDefaults(Commands...)
}
