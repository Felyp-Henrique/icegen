package commands

import (
	"github.com/Felyp-Henrique/icegen/internal/models"
	"github.com/Felyp-Henrique/icegen/internal/templates"
	"github.com/Felyp-Henrique/icegen/internal/usecases"
	"github.com/spf13/cobra"
)

func CommandCreate() *cobra.Command {
	command := &cobra.Command{
		Use:   "new",
		Long:  "create a new icecast.xml configuration",
		Short: "new icecast.xml",
		Run: func(cmd *cobra.Command, args []string) {
			icecast := models.NewIcecastModel()
			icecast.SetNumClients(cmd.Flag("numclients").Value.String())
			icecast.SetNumSources(cmd.Flag("numsources").Value.String())
			icecast.SetQueue(cmd.Flag("queue").Value.String())
			icecast.SetCliTimeout(cmd.Flag("clitimeout").Value.String())
			icecast.SetHdrTimeout(cmd.Flag("hdrtimeout").Value.String())
			icecast.SetSrcTimeout(cmd.Flag("srctimeout").Value.String())
			icecast.SetBurst(cmd.Flag("burst").Value.String())
			icecast.SetSrcPass(cmd.Flag("srcpass").Value.String())
			icecast.SetAdmin(cmd.Flag("admin").Value.String())
			icecast.SetAdminPass(cmd.Flag("adminpass").Value.String())
			icecast.SetHost(cmd.Flag("host").Value.String())
			icecast.SetPort(cmd.Flag("port").Value.String())

			engine := templates.NewEngineFile()
			engine.SetContext(icecast.ToContext())

			manager, err := templates.NewManagerFile(engine, cmd.Flag("dirtemplates").Value.String(), "icecast.xml")

			if err != nil {
				panic(err)
			}

			params := usecases.NewCreateParams()
			params.SetManager(manager)

			usecases.UseCaseCreate(params.ToParams())
		},
	}

	command.Flags().String("dirtemplates", "./templates", "Path of dir templates")
	command.Flags().String("numclients", "1000", "Set the limit of connection in your server")
	command.Flags().String("numsources", "1", "Set the limit of source mount in your server")
	command.Flags().String("queue", "524288", "Set queue size")
	command.Flags().String("clitimeout", "30", "Timeout of client")
	command.Flags().String("hdrtimeout", "15", "Timeout of header")
	command.Flags().String("srctimeout", "10", "Timeout of source")
	command.Flags().String("burst", "65535", "Set burst size")
	command.Flags().String("srcpass", "hackme", "Set the passoword of user source")
	command.Flags().String("admin", "admin", "Set the name of administrator")
	command.Flags().String("adminpass", "hackme", "Set the password of administrator")
	command.Flags().String("host", "localhost", "Set the hostname of server")
	command.Flags().String("port", "8000", "Set the port of server")

	return command
}
