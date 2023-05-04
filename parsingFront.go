package main

import (
	"fmt"
	"github.com/tidwall/gjson"
	"os"
	"strings"
)

func getVueType(files []string) FrontFrameWork {
	packageJsonPath := getPackageJson(files)
	if len(packageJsonPath) == 0 {
		panic(fmt.Errorf("file: package.json not exit"))
	}
	json := getPackageJson(files)
	v := gjson.Get(json, "dependencies").Get("vue").String()
	if v != "" {
		if strings.Contains(v, "2.") {
			return Vue2
		}
		if strings.Contains(v, "3.") {
			return Vue3
		}
	}
	r := gjson.Get(json, "dependencies").Get("react").String()
	if r != "" {
		return React
	}
	return -1
}

func getPackageJson(files []string) string {
	var packageJsonPath string
	for _, item := range files {
		if strings.Contains(item, "package.json") {
			packageJsonPath = item
			break
		}
	}
	if len(packageJsonPath) == 0 {
		panic(fmt.Errorf("file: package.json not exit"))
	}
	buf, err := os.ReadFile(packageJsonPath)
	if err != nil {
		panic(err)
	}
	json := string(buf)
	return json
}
