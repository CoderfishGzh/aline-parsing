package main

import (
	"os"
	"path/filepath"
	"strings"
)

func getDirInfo(path string) []string {
	fileInfo, err := os.Stat(path)
	if err != nil {
		panic(err)
	}
	if !fileInfo.IsDir() {
		panic(err)
	}
	var files []string
	filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !strings.Contains(p, ".git") {
			files = append(files, p)
		}
		return nil
	})
	return files
}

func containsStringPath(arr []string, target string) (string, bool) {
	for _, item := range arr {
		if strings.Contains(item, target) {
			return item, true
		}
	}
	return "", false
}

func containsString(arr []string, target string) bool {
	for _, item := range arr {
		if strings.Contains(item, target) {
			return true
		}
	}
	return false
}
