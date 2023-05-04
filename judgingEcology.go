package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"strings"
)

// 检查合约生态
func judgingEcology(path []string) (ContractEcology, error) {
	_, starkWareOk := containsStringPath(path, "cairo")
	if starkWareOk {
		return STARKWARE, nil
	}
	filePath, moveOk := containsStringPath(path, "Move.toml")
	if moveOk {
		return parsingToml(filePath)
	}
	return EVM, nil
}

// 检查 move.toml文件，判断是sui还是aptos生态
func parsingToml(path string) (ContractEcology, error) {
	var data map[string]any
	if _, err := toml.DecodeFile(path, &data); err != nil {
		return -1, err
	}
	dependenciesData, ok := data["dependencies"].(map[string]interface{})
	if !ok {
		panic("missing dependencies config")
	}
	for key, _ := range dependenciesData {
		if strings.Contains(key, "aptos") || strings.Contains(key, "Aptos") {
			return APTOS, nil
		}
		if strings.Contains(key, "Sui") || strings.Contains(key, "sui") {
			return SUI, nil
		}
	}
	return -1, fmt.Errorf("parsing toml error")
}
