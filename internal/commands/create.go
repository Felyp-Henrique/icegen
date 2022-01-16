package commands

import "github.com/spf13/cobra"

func CommandCreate() *cobra.Command {
	return &cobra.Command{
		Use:   "new",
		Long:  "create a new icecast.xml configuration",
		Short: "new icecast.xml",
		Run: func(self *cobra.Command, args []string) {

		},
	}
}
