package util

import (
	"encoding/json"
	"os/exec"
)

func ExecAndParse(obj interface{}, cmd string, args ...string) error {
	args = append([]string{"--libxo", "json"}, args...)

	out, err := exec.Command(cmd, args...).Output()
	if err != nil {
		return err
	}

	if err := json.Unmarshal(out, obj); err != nil {
		return err
	}

	return nil
}
