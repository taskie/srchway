package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"github.com/taskie/srchway/lib"
)

type Env struct {
	args         []string
	verbose      bool
	aurFlag      bool
	officialFlag bool
}

func (env Env) repos() (repos []srchway.Repo) {
	repos = make([]srchway.Repo, 0)
	fmt.Println(env)
	if env.officialFlag {
		repos = append(repos, srchway.OfficialRepo {})
	}
	if env.aurFlag {
		repos = append(repos, srchway.UserRepo {})
	}
	return
}

func version(env Env) (exitCode int) {
	fmt.Println("0.0")
	exitCode = 0
	return
}

func search(env Env) (exitCode int) {
	query := strings.Join(env.args, " ")
	repos := env.repos()
	for _, repo := range repos {
		err := repo.PrintSearchResponse(query, srchway.NormalMode)
		if err != nil {
			fmt.Println(err)
		}
	}
	exitCode = 0
	return
}

func info(env Env) (exitCode int) {
	fmt.Println("unimplemented")
	exitCode = 1
	return
}

func get(env Env) (exitCode int) {
	query := strings.Join(env.args, " ")
	repos := env.repos()
	for _, repo := range repos {
		switch repo := repo.(type) {
		case srchway.OfficialRepo:
			repo.Get(query)
		case srchway.UserRepo:
			repo.Get(query)
		}
	}
	exitCode = 0
	return
}

func main() {
	pVerbose := flag.Bool("v", false, "verbose mode")
	pAurFlag := flag.Bool("a", false, "AUR")
	pAurOnlyFlag := flag.Bool("A", false, "AUR only")
	flag.Parse()
	if flag.NArg() == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}
	subcommand := flag.Arg(0)
	env := Env {
		args: flag.Args()[1:], verbose: *pVerbose,
		aurFlag: *pAurFlag || *pAurOnlyFlag, officialFlag: !*pAurOnlyFlag,
	}

	exitCode := 0
	switch subcommand {
	case "version":
		exitCode = version(env)
	case "search":
		exitCode = search(env)
	case "info":
		exitCode = info(env)
	case "get":
		exitCode = get(env)
	default:
		flag.PrintDefaults()
		exitCode = 1
	}
	os.Exit(exitCode)
}
