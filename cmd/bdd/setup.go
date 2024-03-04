package bdd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

func Setup() error {
	if err := executeDockerCompose(); err != nil {
		return fmt.Errorf("error starting docker-compose: %w", err)
	}
	return nil
}

func executeDockerCompose() error {
	var cmd *exec.Cmd
	ChdirIfInDirectory("../", "cmd")
	cmd = exec.Command("docker", "compose", "up", "-d")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to execute docker compose up: %w", err)
	}

	for {
		cmd = exec.Command("docker-compose", "ps", "-q")

		output, err := cmd.Output()
		if err != nil {
			log.Fatalf("failed to check docker-compose status: %s", err)
		}

		if len(output) > 0 {
			break
		}

		log.Println("Waiting for docker-compose to start...")
		time.Sleep(5 * time.Second)
	}

	return nil
}

func ChdirIfInDirectory(newDir string, currentDir string) {
	if wd, _ := os.Getwd(); strings.HasSuffix(wd, currentDir) {
		_ = os.Chdir(newDir)
		wd, _ := os.Getwd()
		fmt.Println("Changed to directory: ", wd)
	}
}
