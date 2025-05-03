package docker

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/docker/docker/client"
)

// PIDToContainerName resolves a PID to the container name using cgroup info
func PIDToContainerName(pid string) (string, error) {
	// Step 1: Open the cgroup file
	f, err := os.Open(fmt.Sprintf("/proc/%s/cgroup", pid))
	if err != nil {
		return "", fmt.Errorf("failed to open cgroup file: %w", err)
	}
	defer f.Close()

	var containerID string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if len(parts) != 3 {
			continue
		}
		cgroupPath := parts[2]

		tokens := strings.Split(cgroupPath, "/")
		for _, token := range tokens {
			if len(token) == 64 { // Docker-style container ID
				containerID = token
				break
			}
		}
		if containerID != "" {
			break
		}
	}
	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading cgroup file: %w", err)
	}
	if containerID == "" {
		return "unknown", nil
	}

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return "", fmt.Errorf("failed to create Docker client: %w", err)
	}
	defer cli.Close()

	containerJSON, err := cli.ContainerInspect(context.Background(), containerID)
	if err != nil {
		return "unknown", nil // Could be exited or zombie container
	}

	return strings.TrimPrefix(containerJSON.Name, "/"), nil
}