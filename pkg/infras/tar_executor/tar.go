package tar_executor

import "os/exec"

type TarExecutor struct {
}

func (executor *TarExecutor) Tar(tarName string, files []string) error {
	args := []string{"-cf", tarName}
	args = append(args, files...)

	cmd := exec.Command("tar", args...)
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func (executor *TarExecutor) UnTar(file string) error {
	args := []string{"-xf", file}

	cmd := exec.Command("tar", args...)
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
