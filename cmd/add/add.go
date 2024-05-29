package add

import (
	"github.com/jdambly/kaudit/cmd/settings"
	"github.com/jdambly/kaudit/pkg/audit"
	"github.com/jdambly/kaudit/pkg/fileutils"
	"github.com/jdambly/kaudit/pkg/helpers"
	"github.com/jdambly/kaudit/pkg/kubeapi"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

// Cmd initialize the add command
func Cmd() *cobra.Command {
	s := &settings.Settings{}
	cmd := &cobra.Command{
		Use:   "add",
		Short: "add auditd rules for k8s objects",
		RunE: func(cmd *cobra.Command, args []string) error {
			level, _ := cmd.Flags().GetInt("verbose")
			helpers.SetLogLevel(level)
			err := runAdd(s)
			return err
		}}
	s.PersistentFlags(cmd)
	return cmd
}

// runAdd do all the work for the cmd
func runAdd(s *settings.Settings) error {
	// creat the k8s client
	k8sClient, err := kubeapi.NewKubeClient()
	if err != nil {
		return err
	}
	// get the list of pod UIDs based on the label selector and namespace
	UIDList, err := kubeapi.GetPodUIDList(k8sClient, s.Namespace, s.LabelSelector, s.NodeName)
	if err != nil {
		return err
	}
	var rules []string
	fs := afero.NewOsFs()
	for _, podUID := range UIDList {
		files, err := fileutils.ListPodFiles(fs, podUID)
		if err != nil {
			return err
		}
		for _, file := range files {
			rule, err := audit.CreateAuditRule(audit.DefaultRunner, file, podUID, s.DryRun)
			if err != nil {
				return err
			}
			rules = append(rules, rule)
		}
	}
	return nil
}
