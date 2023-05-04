package main

import "fmt"

func main() {
	//getProjectFromGithub("git@github.com:mohaijiang/my-sui-nfttest4.git")
	project, err := NewProject("git@github.com:mohaijiang/my-sui-nfttest4.git")
	// truffle: /home/gzh/work/checkCon/my-contract-Erc919/
	//project, err := NewProject("/home/gzh/work/checkCon/my-contract-Erc919/", "")

	// foundry: /home/gzh/work/checkCon/hello_foundry/
	//project, err := NewProject("/home/gzh/work/checkCon/hello_foundry/", "")

	// stark: /home/gzh/work/checkCon/stark-wa/
	//project, err := NewProject("/home/gzh/work/checkCon/stark-wa/", "")

	// sui: /home/gzh/work/checkCon/my-sui-nfttest2/
	//project, err := NewProject("/home/gzh/work/checkCon/my-sui-nfttest2/", "")

	// aptos: /home/gzh/work/checkCon/my-aptos-todo/
	//project, err := NewProject("/home/gzh/work/checkCon/my-aptos-todo/", "")

	// vue3: /home/gzh/work/checkFront/T_FrontEnd_Container_Vue/
	//project, err := NewProject("/home/gzh/work/checkFront/T_FrontEnd_Container_Vue/", "")

	// react: /home/gzh/work/checkFront/react/
	//project, err := NewProject("/home/gzh/work/checkFront/react/", "")

	if err != nil {
		panic(err)
	}
	fmt.Println("project type: ", project.ProjectType.string())
	if project.ProjectType == CONTRACT {
		project.Detail.(ContractDetail).print()
	}
	if project.ProjectType == FRONTEND {
		fmt.Println("front framework: ", project.Detail.(FrontDetail).FrontFrameWork.string())
	}
}
