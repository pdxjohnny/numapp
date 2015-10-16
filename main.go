package main

import (
	"runtime"

	"github.com/spf13/cobra"

	"github.com/pdxjohnny/numapp/commands"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var rootCmd = &cobra.Command{Use: "numapp"}
	rootCmd.AddCommand(commands.Commands...)
	rootCmd.Execute()
}
