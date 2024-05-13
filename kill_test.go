package killprocess

import (
	"os/exec"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func processExist(name string) bool {
	err := exec.Command("pgrep", name).Run()
	return err == nil
}

func TestKillProcess(t *testing.T) {
	if !processExist("sleep") {
		err := exec.Command("sleep", "3600").Start()
		require.NoError(t, err)
		require.True(t, processExist("sleep"))
	}

	err := Kill("sleep")
	require.NoError(t, err)

	endTime := time.Now().Add(10 * time.Second)
	for time.Now().Before(endTime) {
		if !processExist("sleep") {
			return
		}
		time.Sleep(1 * time.Second)
	}
	require.Fail(t, "the process should be killed after 10 seconds")
}
