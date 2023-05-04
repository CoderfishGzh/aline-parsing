package main

func getEvmFramework(files []string) ContractFrameWork {
	okTruffle := isTruffle(files)
	if okTruffle {
		return Truffle
	}
	okFoundry := isFoundry(files)
	if okFoundry {
		return Foundry
	}
	okHardhat := isHardhat(files)
	if okHardhat {
		return Hardhat
	}
	return -1
}

func isTruffle(files []string) bool {
	if containsString(files, "truffle-config.js") && containsString(files, "contracts") && containsString(files, "migrations") {
		return true
	}
	return false
}

func isFoundry(files []string) bool {
	if containsString(files, "foundry.toml") {
		return true
	}
	return false
}

func isHardhat(files []string) bool {
	if containsString(files, "hardhat.config.js") {
		return true
	}
	return false
}
