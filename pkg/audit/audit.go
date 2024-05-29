package audit

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"os/exec"
	"strings"
)

// CommandRunner is a type for functions that run commands
type CommandRunner func(name string, arg ...string) *exec.Cmd

var DefaultRunner CommandRunner = exec.Command

// CreateAuditRule creates an audit rule for a given path
func CreateAuditRule(runner CommandRunner, path string, podUID string, dryRun bool) (string, error) {
	cmdSlice := []string{"-a", "always,exit", "-F", fmt.Sprintf("dir=%s", path), "-F", "perm=wa", "-k", fmt.Sprintf("file_deletion:%s", podUID)}
	if dryRun {
		log.Debug().Str("rule", strings.Join(cmdSlice, " ")).Msg("Dry run")
		return strings.Join(cmdSlice, " "), nil
	}
	// note to self cmd needs to take a slice as each element is a separate argument
	cmd := runner("auditctl", cmdSlice...)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("error creating audit rule: %v\n%s", err, output)
	}
	log.Info().Str("path", path).Msg("Created audit rule")
	return strings.Join(cmdSlice, " "), nil
}
