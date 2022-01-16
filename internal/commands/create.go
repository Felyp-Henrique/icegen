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
			icecast.SetRelayHost(cmd.Flag("relayhost").Value.String())
			icecast.SetRelayUpdateInterval(cmd.Flag("relayupdateinterval").Value.String())
			icecast.SetRelayUser(cmd.Flag("relayuser").Value.String())
			icecast.SetRelayPassword(cmd.Flag("relaypassword").Value.String())
			icecast.SetRelayDemand(cmd.Flag("relaydemand").Value.String())
			icecast.SetRelayPort(cmd.Flag("relayport").Value.String())
			icecast.SetRelayMount(cmd.Flag("relaymount").Value.String())
			icecast.SetRelayMountLocal(cmd.Flag("relaymountlocal").Value.String())
			icecast.SetRelayOnDemand(cmd.Flag("relayondemand").Value.String())
			icecast.SetRelayShoutcast(cmd.Flag("relayshoutcast").Value.String())

			engine := templates.NewEngineFile()
			engine.SetContext(icecast.Context)

			manager, _ := templates.NewManagerFile(engine, "", "")

			params := usecases.NewCreateParams()
			params.SetManager(manager)

			usecases.UseCaseCreate(params.Params)
		},
	}

	command.Flags().String("numclients", "", "")
	command.MarkFlagRequired("numclients")

	command.Flags().String("numsources", "", "")
	command.MarkFlagRequired("numsources")

	command.Flags().String("queue", "", "")
	command.MarkFlagRequired("queue")

	command.Flags().String("clitimeout", "", "")
	command.MarkFlagRequired("clitimeout")

	command.Flags().String("hdrtimeout", "", "")
	command.MarkFlagRequired("hdrtimeout")

	command.Flags().String("srctimeout", "", "")
	command.MarkFlagRequired("srctimeout")

	command.Flags().String("burst", "", "")
	command.MarkFlagRequired("burst")

	command.Flags().String("srcpass", "", "")
	command.MarkFlagRequired("srcpass")

	command.Flags().String("admin", "", "")
	command.MarkFlagRequired("admin")

	command.Flags().String("adminpass", "", "")
	command.MarkFlagRequired("adminpass")

	command.Flags().String("host", "", "")
	command.MarkFlagRequired("host")

	command.Flags().String("port", "", "")
	command.MarkFlagRequired("port")

	command.Flags().String("relayhost", "", "")
	command.MarkFlagRequired("relayhost")

	command.Flags().String("relayupdateinterval", "", "")
	command.MarkFlagRequired("relayupdateinterval")

	command.Flags().String("relayuser", "", "")
	command.MarkFlagRequired("relayuser")

	command.Flags().String("relaypassword", "", "")
	command.MarkFlagRequired("relaypassword")

	command.Flags().String("relaydemand", "", "")
	command.MarkFlagRequired("relaydemand")

	command.Flags().String("relayport", "", "")
	command.MarkFlagRequired("relayport")

	command.Flags().String("relaymount", "", "")
	command.MarkFlagRequired("relaymount")

	command.Flags().String("relaymountlocal", "", "")
	command.MarkFlagRequired("relaymountlocal")

	command.Flags().String("relayondemand", "", "")
	command.MarkFlagRequired("relayondemand")

	command.Flags().String("relayshoutcast", "", "")
	command.MarkFlagRequired("relayshoutcast")

	return command
}
