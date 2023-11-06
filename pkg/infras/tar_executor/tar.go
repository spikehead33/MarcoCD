package tar_executor

import (
	"bytes"
	"errors"
	"os/exec"
)

type TarExecutor struct {
}

func (executor *TarExecutor) Tar(tarName string, files []string) error {
	args := []string{"-cf", tarName}
	args = append(args, files...)

	var cmdErr bytes.Buffer

	cmd := exec.Command("tar", args...)
	cmd.Stderr = &cmdErr
	if err := cmd.Run(); err != nil {
		return errors.New(cmdErr.String())
	}

	return nil
}

func (executor *TarExecutor) UnTar(file string) error {
	args := []string{"-xf", file}

	var cmdErr bytes.Buffer

	cmd := exec.Command("tar", args...)
	cmd.Stderr = &cmdErr
	if err := cmd.Run(); err != nil {
		return errors.New(cmdErr.String())
	}

	return nil
}
