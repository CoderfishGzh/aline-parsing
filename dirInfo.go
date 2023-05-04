package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
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

func getProjectFromGithub(url string) (string, error) {
	// 创建随机文件夹
	tempDir, err := ioutil.TempDir("", "")
	if err != nil {
		panic(err)
	}
	//defer os.RemoveAll(tempDir)

	fmt.Println("Temporary directory created:", tempDir)
	cmd := exec.Command("git", "clone", url, tempDir)
	err = cmd.Run()
	if err != nil {
		fmt.Println("Failed to execute command:", err)
		return "", err
	}
	return tempDir, nil
}
