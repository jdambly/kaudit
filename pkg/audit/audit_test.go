package audit

import (
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Mock CommandRunner
func mockRunner(output string, err error) CommandRunner {
	return func(name string, arg ...string) *exec.Cmd {
		cmd := exec.Command("echo", output)
		if err != nil {
			cmd = exec.Command("false")
		}
		return cmd
	}
}

func TestCreateAuditRule(t *testing.T) {
	t.Run("dry run", func(t *testing.T) {
		expected := "auditctl -a always,exit -w /tmp/testpath -F perm=wa -k file_deletion:pod-UID\n"
		result, err := CreateAuditRule(DefaultRunner, "/tmp/testpath", "pod-UID", true)
		require.NoError(t, err)
		assert.Equal(t, expected, result)
	})

	t.Run("real run", func(t *testing.T) {
		mockOutput := "mocked command"
		runner := mockRunner(mockOutput, nil)

		expected := "auditctl -a always,exit -w /tmp/testpath -F perm=wa -k file_deletion:pod-UID\n"
		result, err := CreateAuditRule(runner, "/tmp/testpath", "pod-UID", false)
		require.NoError(t, err)
		assert.Equal(t, expected, result)
	})
}
