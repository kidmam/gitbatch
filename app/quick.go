package app

import (
	"fmt"
	"sync"
	"time"

	"github.com/isacikgoz/gitbatch/core/command"
	"github.com/isacikgoz/gitbatch/core/git"
)

func quick(directories []string, mode string) {
	var wg sync.WaitGroup
	start := time.Now()
	for _, dir := range directories {
		wg.Add(1)
		go func(d string, mode string) {
			defer wg.Done()
			err := operate(d, mode)
			if err != nil {
				fmt.Printf("%s: %s\n", d, err.Error())
			} else {
				fmt.Printf("%s: successful\n", d)
			}
		}(dir, mode)
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("%d repositories finished in: %s\n", len(directories), elapsed)
}

func operate(directory, mode string) error {
	r, err := git.FastInitializeRepo(directory)
	if err != nil {
		return err
	}
	switch mode {
	case "fetch":
		return command.Fetch(r, &command.FetchOptions{
			RemoteName: "origin",
			Progress:   true,
		})
	case "pull":
		return command.Pull(r, &command.PullOptions{
			RemoteName: "origin",
			Progress:   true,
		})
	}
	return nil
}
