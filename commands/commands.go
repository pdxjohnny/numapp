package commands

import (
	"github.com/spf13/cobra"

	"github.com/pdxjohnny/numapp/db"
	"github.com/pdxjohnny/numapp/db/get"
	"github.com/pdxjohnny/numapp/db/put"
	"github.com/pdxjohnny/numapp/http"
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
	&cobra.Command{
		Use:   "http",
		Short: "Start the http server",
		Run: func(cmd *cobra.Command, args []string) {
			ConfigBindFlags(cmd)
			http.Run()
		},
	},
	&cobra.Command{
		Use:   "db",
		Short: "Start the db service",
		Run: func(cmd *cobra.Command, args []string) {
			ConfigBindFlags(cmd)
			db.Run()
		},
	},
}

func init() {
	ConfigDefaults(Commands...)
}
