package commands

import "github.com/spf13/cobra"

func CommandApp() *cobra.Command {
	return &cobra.Command{
		Long: "icecast.xml generator",
	}
}

func NewCommandApp() *cobra.Command {
	app := CommandApp()
	app.AddCommand(
		CommandCreate(),
	)
	return app
}
