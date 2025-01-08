package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	type ExitCoder interface {
		error
		ExitCode() int
	}
	if err := start(os.Args[1:]); err != nil {
		fmt.Println(err)

		var exitErr ExitCoder
		if errors.As(err, &exitErr) {
			os.Exit(exitErr.ExitCode())
		}
		os.Exit(1)
	}
}

func start(args []string) error {
	if len(args) == 0 {
		binPath, err := os.Executable()
		if err != nil {
			return fmt.Errorf("getting executable name: %w", err)
		}

		stderr := bytes.NewBufferString("")

		cmd := exec.Cmd{
			Path:   binPath,
			Args:   []string{binPath, "run"},
			Stderr: stderr,
			Stdout: os.Stdout,
		}

		err = cmd.Run()
		if err != nil {
			fmt.Println("Look! A Panic!")
			fmt.Println(stderr.String())
			return err
		}
		return nil
	}

	subcommand := args[0]
	switch subcommand {
	case "run":
		return run(args[1:])
	}
	return nil
}

func run(_ []string) error {
	go func() {
		panic("kaboom")
	}()
	time.Sleep(1 * time.Second)
	return nil
}
