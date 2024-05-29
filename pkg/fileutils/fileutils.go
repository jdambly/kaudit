package fileutils

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/spf13/afero"
	"os"
	"path/filepath"
)

// ListPodFiles returns a list of files in the pod's volume directory
func ListPodFiles(fs afero.Fs, podUID string) ([]string, error) {
	basePath := fmt.Sprintf("/var/lib/kubelet/pods/%s/volumes/", podUID)
	var files []string

	err := afero.Walk(fs, basePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// filter out directories
		if !info.IsDir() {
			files = append(files, path)
		}
		// filter out paths that have pg_data in them
		if info.IsDir() && info.Name() == "pgdata" {
			return filepath.SkipDir
		}
		// filter out paths that have pg_wal in them
		if info.IsDir() && info.Name() == "pg_wal" {
			return filepath.SkipDir
		}
		log.Debug().Str("path", path).Msg("found file")
		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}
