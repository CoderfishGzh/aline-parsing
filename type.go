package main

type ProjectType int
type ContractEcology int
type ContractFrameWork int
type FrontFrameWork int

func (pt ProjectType) string() string {
	switch pt {
	case CONTRACT:
		return "CONTRACT"
	case FRONTEND:
		return "FRONTEND"
	default:
		return ""
	}
}

func (ce ContractEcology) string() string {
	switch ce {
	case EVM:
		return "EVM"
	case APTOS:
		return "APTOS"
	case STARKWARE:
		return "STARKWARE"
	case SUI:
		return "SUI"
	default:
		return ""
	}
}

func (cf ContractFrameWork) string() string {
	switch cf {
	case Truffle:
		return "Truffle"
	case Foundry:
		return "Foundry"
	case Hardhat:
		return "Hardhat"
	case Move:
		return "Move"
	case Cairo:
		return "Cairo"
	default:
		return ""
	}
}

func (ffw FrontFrameWork) string() string {
	switch ffw {
	case Vue2:
		return "Vue2"
	case Vue3:
		return "Vue3"
	case React:
		return "React"
	default:
		return ""
	}
}

// 前端框架
const (
	Vue2 FrontFrameWork = iota + 2
	Vue3
	React
)

// 合约框架
const (
	Truffle ContractFrameWork = iota
	Foundry
	Hardhat
	Move
	Cairo
)

// Project 类型
const (
	CONTRACT ProjectType = iota
	FRONTEND
)

// 合约生态
const (
	EVM ContractEcology = iota
	APTOS
	STARKWARE
	SUI
)
