package audit

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"syscall"

	"github.com/rs/zerolog/log"
)

// CommandRunner is a type for functions that run commands
type CommandRunner func(name string, arg ...string) *exec.Cmd

var DefaultRunner CommandRunner = exec.Command

// Validate checks to make sure auditd is running, returns an error if it's not running
func Validate() error {
	socketPath := "/var/run/auditd.sock"
	info, err := os.Stat(socketPath)
	if err != nil {
		return fmt.Errorf("error checking auditd socket: %v", err)
	}
	if info.Mode()&os.ModeSocket == 0 {
		return fmt.Errorf("auditd socket is not a socket file")
	}

	u, err := user.Current()
	if err != nil {
		return fmt.Errorf("error getting current user: %v", err)
	}

	uid, err := strconv.ParseUint(u.Uid, 10, 32)
	if err != nil {
		return fmt.Errorf("error converting UID to uint32: %v", err)
	}

	stat, ok := info.Sys().(*syscall.Stat_t)
	if !ok {
		return fmt.Errorf("failed to get file info")
	}

	if info.Mode().Perm()&0200 == 0 && stat.Uid != uint32(uid) {
		return fmt.Errorf("auditd socket is not writable by the current user")
	}

	return nil
}

// CreateAuditRule creates an audit rule for a given path
func CreateAuditRule(runner CommandRunner, path string, podUID string, dryRun bool) (string, error) {
	rule := fmt.Sprintf("-a always,exit -w %s -F perm=wa -k file_deletion:%s", path, podUID)
	cmdRule := fmt.Sprintf("%s %s\n", "auditctl", rule)
	if dryRun {
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
