package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func ExistDir(path string) bool {
	status, err := os.Stat(path)
	return err == nil && status.IsDir()
}

func ExistFile(path string) bool {
	status, err := os.Stat(path)
	return err == nil && !status.Mode().IsRegular()
}

func Mkdirs(dirName string) error {
	if ExistDir(dirName) {
		return nil
	}
	parentDir := filepath.Dir(dirName)
	if !ExistDir(parentDir) {
		if err := Mkdirs(parentDir); err != nil {
			return err
		}
	}
	return os.Mkdir(dirName, 0755)
}

func FindProjectDir(basePath string) (string, error) {
	path := filepath.Clean(basePath)
	for path != "/" {
		gitDir := filepath.Join(path, ".git")
		if ExistDir(gitDir) {
			return filepath.Clean(path), nil
		}
		path = filepath.Dir(path)
	}
	return "", fmt.Errorf("%s: project directory not found", basePath)
}

func KaniHome() (string, error) {
	homePaths := []string{
		"/usr/local/Cellar/kani",
		".",
	}
	for _, path := range homePaths {
		if ExistDir(path) {
			return path, nil
		}
	}
	return "", fmt.Errorf("kani is not installed")
}
