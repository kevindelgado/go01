package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go01",
	Short: "go01 lets you score darts in the terminal",
	Long:  "go01 lets you score darts in the terminal",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("game type", args[0], "num teams", args[1])
	},
}

func Execute() error {
	return rootCmd.Execute()
}
