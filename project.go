package main

import (
	"fmt"
	"log"
	"os"
)

type Project struct {
	ProjectType ProjectType
	GithubAddr  string
	LocalAddr   string
	DirInfo     []string
	Detail      any
}

type ContractDetail struct {
	Ecosystem ContractEcology
	Framework ContractFrameWork
}

type FrontDetail struct {
	FrontFrameWork FrontFrameWork
}

func (cd ContractDetail) print() {
	fmt.Println("ecosystem: ", cd.Ecosystem.string())
	fmt.Println("framework: ", cd.Framework.string())
}

func NewProject(githubAddr string) (Project, error) {
	tmpDir, err := getProjectFromGithub(githubAddr)
	defer os.RemoveAll(tmpDir)
	if err != nil {
		return Project{}, err
	}

	dirInfo := getDirInfo(tmpDir)
	if dirInfo == nil {
		log.Print("get dirInfo error")
		return Project{}, fmt.Errorf("get dirInfo error")
	}
	var project Project
	project.DirInfo = dirInfo
	project.GithubAddr = githubAddr
	// check the project type
	project.projectType()
	// check the ecology
	if project.ProjectType == CONTRACT {
		err := project.getContractDetail()
		if err != nil {
			return Project{}, nil
		}
	}
	return project, nil
}

func (p *Project) projectType() {
	ok := containsString(p.DirInfo, "package.json")
	if ok {
		ffw := getVueType(p.DirInfo)
		if ffw != -1 {
			p.ProjectType = FRONTEND
			p.Detail = FrontDetail{FrontFrameWork: ffw}
			return
		}
	}
	p.ProjectType = CONTRACT
}

func (p *Project) getContractDetail() error {
	ecology, err := judgingEcology(p.DirInfo)
	if err != nil {
		return err
	}
	var contractDetail ContractDetail
	contractDetail.Ecosystem = ecology
	if ecology == EVM {
		frameWork := getEvmFramework(p.DirInfo)
		if frameWork == -1 {
			return fmt.Errorf("parsing EVM framework error")
		}
		contractDetail.Framework = frameWork
	} else if ecology == APTOS {
		contractDetail.Framework = Move
	} else if ecology == SUI {
		contractDetail.Framework = Move
	} else if ecology == STARKWARE {
		contractDetail.Framework = Cairo
	}
	p.Detail = contractDetail
	return nil
}
