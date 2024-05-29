package audit

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"os/exec"
)

// CommandRunner is a type for functions that run commands
type CommandRunner func(name string, arg ...string) *exec.Cmd

var DefaultRunner CommandRunner = exec.Command

// CreateAuditRule creates an audit rule for a given path
func CreateAuditRule(runner CommandRunner, path string, podUID string, dryRun bool) (string, error) {
	rule := fmt.Sprintf("-a always,exit -w %s -F perm=wa -k file_deletion:%s", path, podUID)
	cmdRule := fmt.Sprintf("%s %s\n", "auditctl", rule)
	if dryRun {
		log.Debug().Str("rule", cmdRule).Msg("Dry run")
		return cmdRule, nil
	}

	cmd := runner("auditctl", rule)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("error creating audit rule: %v\n%s", err, output)
	}
	log.Info().Str("path", path).Msg("Created audit rule")
	return cmdRule, nil
}
