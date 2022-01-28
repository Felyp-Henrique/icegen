package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func CommandVersion() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "version of icegen",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("icegen v0.1.0 beta - go v1.17.5")
		},
	}
}
