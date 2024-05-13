// Package killprocess provides utility functions for killing running processes.
package killprocess

import "os/exec"

// Kill terminates a process with the specified name.
func Kill(name string) error {
	return exec.Command("killall", "-15", name).Run()
}
