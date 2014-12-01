package main

import (
	"github.com/samalba/dockerclient"
	"log"
)

func main() {
	// Init the client
	docker, _ := dockerclient.NewDockerClient("tcp://172.17.8.101:2375", nil)

	// Get only running containers
	containers, err := docker.ListContainers(false, true, "")
	if err != nil {
		log.Fatal(err)
	}
	for _, c := range containers {
		log.Println(c.Id, c.Names)
	}

	// Inspect the first container returned
	if len(containers) > 0 {
		id := containers[0].Id
		info, _ := docker.InspectContainer(id)
		log.Println(info)
	}

	// Pull
	err = docker.PullImage("quay.io/epipho/docker:tag37")
	if err != nil {
		log.Fatal(err)
	}

	// Create a container
	containerConfig := &dockerclient.ContainerConfig{Image: "ubuntu:12.04", Cmd: []string{"bash"}}
	containerId, err := docker.CreateContainer(containerConfig, "")
	if err != nil {
		log.Fatal(err)
	}

	// Start the container
	err = docker.StartContainer(containerId, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Stop the container (with 5 seconds timeout)
	docker.StopContainer(containerId, 5)
}
