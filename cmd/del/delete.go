package del

import (
	"github.com/jdambly/kaudit/cmd/settings"
	"github.com/jdambly/kaudit/pkg/helpers"
	"github.com/spf13/cobra"
)

// Cmd initialize the del command
func Cmd() *cobra.Command {
	s := &settings.Settings{}
	cmd := &cobra.Command{
		Use:   "del",
		Short: "del auditd rules for k8s objects",
		RunE: func(cmd *cobra.Command, args []string) error {
			level, _ := cmd.Flags().GetInt("verbose")
			helpers.SetLogLevel(level)
			err := runDelete(s)
			return err
		}}
	s.PersistentFlags(cmd)
	return cmd
}

// runDelete do all the work for the cmd
func runDelete(s *settings.Settings) error {
	return nil
}
