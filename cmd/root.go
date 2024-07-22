package cmd

import (
	"github.com/jdambly/kaudit/cmd/add"
	"github.com/jdambly/kaudit/cmd/del"
	"github.com/spf13/cobra"
	"os"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(info VersionInfo) {
	rootCmd := newRootCmd(info)
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

// newRootCmd adds all the sub commands
func newRootCmd(info VersionInfo) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "kaudit",
		Short: "app to used to create auditd rules for k8s objects",
	}
	cmd.AddCommand(
		newVersionCmd(info),
		add.Cmd(), // add the add command
		del.Cmd(), // add the delete command
	)
	return cmd
}
