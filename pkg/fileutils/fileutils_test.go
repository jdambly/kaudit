package fileutils

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListPodFiles_WithRegex(t *testing.T) {
	fs := afero.NewMemMapFs()

	podUID := "2029920a-5b67-4021-bb19-87d8b6ee7b86"
	volumesDir := fmt.Sprintf("/var/lib/kubelet/pods/%s/volumes", podUID)

	filesToCreate := []string{
		filepath.Join(volumesDir, "file1.txt"),
		filepath.Join(volumesDir, "file2.txt"),
		filepath.Join(volumesDir, "fakeDir/file3.txt"),
		filepath.Join(volumesDir, "fakeDir/file4.txt"),
		filepath.Join(volumesDir, "pg_data/file5.yaml"),
		filepath.Join(volumesDir, "pg_wal/file6.tgz"),
	}

	for _, file := range filesToCreate {
		_, err := fs.Create(file)
		require.NoError(t, err, "Failed to create file")
	}

	files, err := ListPodFiles(fs, podUID, ".*txt")
	require.NoError(t, err, "listPodFiles returned an error")

	expectedFiles := []string{
		filepath.Join(volumesDir, "file1.txt"),
		filepath.Join(volumesDir, "file2.txt"),
		filepath.Join(volumesDir, "fakeDir/file3.txt"),
		filepath.Join(volumesDir, "fakeDir/file4.txt"),
	}

	assert.ElementsMatch(t, expectedFiles, files, "The list of files does not match the expected list")
}

func TestListPodFiles_EmptyString(t *testing.T) {
	fs := afero.NewMemMapFs()

	podUID := "2029920a-5b67-4021-bb19-87d8b6ee7b86"
	volumesDir := fmt.Sprintf("/var/lib/kubelet/pods/%s/volumes", podUID)

	filesToCreate := []string{
		filepath.Join(volumesDir, "file1.txt"),
		filepath.Join(volumesDir, "file2.txt"),
		filepath.Join(volumesDir, "fakeDir/file3.txt"),
		filepath.Join(volumesDir, "fakeDir/file4.txt"),
		filepath.Join(volumesDir, "pg_data/file5.yaml"),
		filepath.Join(volumesDir, "pg_wal/file6.tgz"),
	}

	for _, file := range filesToCreate {
		_, err := fs.Create(file)
		require.NoError(t, err, "Failed to create file")
	}

	files, err := ListPodFiles(fs, podUID, "")
	require.NoError(t, err, "listPodFiles returned an error")

	expectedFiles := []string{
		filepath.Join(volumesDir, "file1.txt"),
		filepath.Join(volumesDir, "file2.txt"),
		filepath.Join(volumesDir, "fakeDir/file3.txt"),
		filepath.Join(volumesDir, "fakeDir/file4.txt"),
		filepath.Join(volumesDir, "pg_data/file5.yaml"),
		filepath.Join(volumesDir, "pg_wal/file6.tgz"),
	}

	assert.ElementsMatch(t, expectedFiles, files, "The list of files does not match the expected list")
}

func TestListPodFiles_WithOR(t *testing.T) {
	fs := afero.NewMemMapFs()

	podUID := "2029920a-5b67-4021-bb19-87d8b6ee7b86"
	volumesDir := fmt.Sprintf("/var/lib/kubelet/pods/%s/volumes", podUID)

	filesToCreate := []string{
		filepath.Join(volumesDir, "file1.txt"),
		filepath.Join(volumesDir, "file2.txt"),
		filepath.Join(volumesDir, "fakeDir/file3.txt"),
		filepath.Join(volumesDir, "fakeDir/file4.txt"),
		filepath.Join(volumesDir, "pg_data/file5.yaml"),
		filepath.Join(volumesDir, "pg_wal/file6.tgz"),
	}

	for _, file := range filesToCreate {
		_, err := fs.Create(file)
		require.NoError(t, err, "Failed to create file")
	}

	files, err := ListPodFiles(fs, podUID, "pg_data|pg_wal")
	require.NoError(t, err, "listPodFiles returned an error")

	expectedFiles := []string{
		filepath.Join(volumesDir, "fakeDir/file4.txt"),
		filepath.Join(volumesDir, "pg_data/file5.yaml"),
		filepath.Join(volumesDir, "pg_wal/file6.tgz"),
	}

	assert.ElementsMatch(t, expectedFiles, files, "The list of files does not match the expected list")
}
