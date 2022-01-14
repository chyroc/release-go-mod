package internal

import (
	"bytes"
	"os"
	"os/exec"
)

func execCommand(command string, args ...string) (string, error) {
	out := new(bytes.Buffer)
	er := new(bytes.Buffer)

	cmd := exec.Command(command, args...)
	cmd.Env = os.Environ()
	cmd.Stdin = os.Stdin
	cmd.Stdout = out
	cmd.Stderr = er
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return out.String(), nil
}

func execCommandStd(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Env = os.Environ()
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
