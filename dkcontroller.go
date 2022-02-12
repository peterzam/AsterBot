package main //docker controller

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"strconv"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func ContainerStatus(id string) bool {
	status := false
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		status = false
		//panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		status = false
		//panic(err)
	}

	for _, container := range containers {
		if container.ID[:len(id)] == id {
			status = true
		}
	}
	defer cli.Close()
	return status
}

func ContainerStart(id string) bool {
	status := true
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		status = false
		//panic(err)
	}

	err = cli.ContainerStart(context.Background(), id, types.ContainerStartOptions{})
	if err != nil {
		status = false
		//panic(err)
	}
	defer cli.Close()
	return status
}

func ContainerStop(id string) bool {
	status := true
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		status = false
		//panic(err)
	}

	err = cli.ContainerStop(context.Background(), id, nil)
	if err != nil {
		status = false
		//panic(err)
	}
	defer cli.Close()
	return status
}

func ContainerRestart(id string) bool {
	status := true
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		status = false
		//panic(err)
	}

	err = cli.ContainerRestart(context.Background(), id, nil)
	if err != nil {
		status = false
		//panic(err)
	}
	defer cli.Close()
	return status
}

func ContainerLog(id string, line int) (bool, string) {
	status := true
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		status = false
		//panic(err)
	}
	reader, err := cli.ContainerLogs(context.Background(), id, types.ContainerLogsOptions{
		ShowStdout: true,
		Follow:     true,
		Tail:       strconv.Itoa(line),
	})
	if err != nil {
		status = false
		//panic(err)
	}

	go func() {
		time.Sleep(time.Second * 2)
		reader.Close()
	}()

	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)

	return status, string(buf.String())
}

func ContainerExec(id string, command string) bool {
	status := true
	cmd := exec.Command("/bin/bash", "-c", fmt.Sprintf(`docker exec  minecraft bash -c 'echo -e "%s" > /tmp/mc-input'`, command))
	err := cmd.Run()
	if err != nil {
		status = false
		//panic(err)
	}
	return status
}
