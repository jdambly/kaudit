package fileutils

import (
	"fmt"
	"github.com/spf13/afero"
	"os"
	"regexp"
)

// ListPodFiles returns a list of files in the pod's volume directory
func ListPodFiles(fs afero.Fs, podUID, myRegex string) ([]string, error) {
	basePath := fmt.Sprintf("/var/lib/kubelet/pods/%s/volumes/", podUID)
	var files []string
	r := regexp.MustCompile(myRegex)
	err := afero.Walk(fs, basePath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info == nil {
				return nil
			}
			if info.IsDir() {
				return nil
			}
			if r.MatchString(path) {
				files = append(files, path)
			}
			return nil
		})

	if err != nil {
		return nil, err
	}

	return files, nil
}
