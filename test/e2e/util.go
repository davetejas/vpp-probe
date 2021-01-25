package e2e

import (
	"log"
	"os"
	"os/exec"
	"testing"
	"time"
)

const (
	waitCreateCluster = time.Minute
)

func kubectl(t *testing.T, context string, args ...string) {
	args = append([]string{
		"--context", context,
	}, args...)
	runCmd(t, "kubectl", args...)
}

func createCluster(name string) {
	runCmd(nil, "kind", "create", "cluster", "--name", name, "--wait", waitCreateCluster.String())
}

func deleteCluster(name string) {
	runCmd(nil, "kind", "delete", "cluster", "--name", name)
}

func runCmd(t *testing.T, cmd string, args ...string) {
	c := exec.Command(cmd, args...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	if t != nil {
		t.Logf("[EXEC] %v", c)
	} else {
		log.Printf("[EXEC] %v", c)
	}
	if err := c.Run(); err != nil {
		if t == nil {
			log.Fatalf("ERROR: %v", err)
		} else {
			t.Fatalf("ERROR: %v", err)
		}
	}
}
