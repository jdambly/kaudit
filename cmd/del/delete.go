package del

import (
	"github.com/jdambly/kaudit/cmd/settings"
	"github.com/jdambly/kaudit/pkg/audit"
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
			all, _ := cmd.Flags().GetBool("all")
			helpers.SetLogLevel(level)
			err := runDelete(all)
			return err
		}}
	s.PersistentFlags(cmd)
	cmd.Flags().Bool("all", true, "Delete all rules")
	return cmd
}

// runDelete do all the work for the cmd
func runDelete(all bool) error {
	// delete all rules
	if all {
		err := audit.CleanAuditRules(audit.DefaultRunner)
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}
