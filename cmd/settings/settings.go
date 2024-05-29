package settings

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/sethvargo/go-envconfig"
	"github.com/spf13/cobra"
	"strings"
)

// Settings is the struct that holds the settings for the application
type Settings struct {
	Verbose       int
	Force         bool   `env:"FORCE, default=false"`
	PodName       string `env:"POD_NAME"`
	Namespace     string `env:"NAMESPACE"`
	LabelSelector string `env:"LABEL_SELECTOR"`
	DryRun        bool   `env:"DRY_RUN, default=false"`
	NodeName      string `env:"NODE_NAME"`
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
	cmd.PersistentFlags().StringVarP(&s.LabelSelector, "label-selector", "l", s.LabelSelector,
		"Set the label selector")
	cmd.PersistentFlags().BoolVarP(&s.DryRun, "dry-run", "d", s.DryRun,
		"Run the command in dry-run mode")
	cmd.PersistentFlags().StringVarP(&s.NodeName, "node-name", "n", s.NodeName,
		"Set the node name")

}

// ParseLabelSelector parses the label selector string into a map
func (s *Settings) ParseLabelSelector(selector string) (map[string]string, error) {
	if selector == "" {
		return map[string]string{}, fmt.Errorf("label selector is empty")
	}
	labels := map[string]string{}
	pairs := strings.Split(selector, ",")
	for _, pair := range pairs {
		kv := strings.Split(pair, "=")
		if len(kv) != 2 {
			log.Warn().Msgf("invalid label selector: %s", pair)
			continue
		}
		labels[kv[0]] = kv[1]
	}
	return labels, nil
}
