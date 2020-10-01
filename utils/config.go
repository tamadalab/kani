package utils

import (
	"os"
	"path/filepath"
)

/*
FindConfDir returns the path of "~/.config/kani" directory.
ref. https://qiita.com/charon/items/74e49a0fd456e7257dbd,
     https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html
*/
func FindConfDir() string {
	home := findHomeDir()
	kaniConf := filepath.Join(home, ".config", "kani")
	if !ExistDir(kaniConf) {
		Mkdirs(kaniConf)
	}
	return kaniConf
}

func findHomeDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return os.Getenv("HOME")
	}
	return home
}
