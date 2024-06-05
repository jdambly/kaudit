package settings

import (
	"context"
	"github.com/sethvargo/go-envconfig"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Settings is the struct that holds the settings for the application
type Settings struct {
	Verbose       int
	Force         bool   `env:"FORCE, default=false"`                 // Force flag from environment variable
	PodName       string `env:"POD_NAME"`                             // Pod name from environment variable
	Namespace     string `env:"NAMESPACE"`                            // Namespace from environment variable
	LabelSelector string `env:"LABEL_SELECTOR, default=empty_string"` // Label selector from environment variable
	DryRun        bool   `env:"DRY_RUN, default=false"`               // Dry run flag from environment variable
	NodeName      string `env:"NODE_NAME"`                            // Node name from environment variable
	Regex         string `env:"REGEX"`                                // Regex a patter used for matching
}

// GetEnv get the environment variables
func (s *Settings) GetEnv() {
	ctx := context.Background()
	if err := envconfig.Process(ctx, s); err != nil {
		panic(err)
	}
}

// PersistentFlags bolts on the flags to override the env vars
func (s *Settings) PersistentFlags(cmd *cobra.Command) {
	// get the env vars
	s.GetEnv()

	// add the flags
	cmd.PersistentFlags().IntVarP(&s.Verbose, "verbose", "v", s.Verbose,
		"Set the logging level 1 to 8")
	cmd.PersistentFlags().BoolVarP(&s.Force, "force", "f", s.Force,
		"force the creation of the rules")
	cmd.PersistentFlags().StringVarP(&s.PodName, "pod-name", "p", s.PodName,
		"Set the node name")
	cmd.PersistentFlags().StringVarP(&s.Namespace, "namespace", "n", s.Namespace,
		"Namespace from environment variable")
	cmd.PersistentFlags().StringVarP(&s.LabelSelector, "label-selector", "l", s.LabelSelector,
		"Set the label selector, with a comma separated values list")
	cmd.PersistentFlags().StringVarP(&s.Regex, "include", "i", s.Regex,
		"Include flag for filtering file list, with a comma separated values list")
}

// ListOpts returns a metav1.ListOptions struct
func (s *Settings) ListOpts() metav1.ListOptions {
	return metav1.ListOptions{
		LabelSelector: s.LabelSelector,
		FieldSelector: "spec.nodeName=" + s.NodeName,
	}
}
