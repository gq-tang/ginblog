package routers

import (
	"os"
	"path/filepath"
)

func getTempalteFiles(layoutDir, ext string) []string {
	fileNames := make([]string, 0, 10)
	filepath.Walk(layoutDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if info.IsDir() {
			var pattern string
			if path[len(path)-1] == '/' {
				pattern = path + "*." + ext
			} else {
				pattern = path + "/*." + ext
			}
			files, err := filepath.Glob(pattern)
			if err != nil {
				return err
			}
			fileNames = append(fileNames, files...)
		}
		return nil
	})
	return fileNames
}
