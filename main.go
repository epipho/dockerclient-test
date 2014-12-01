package main

import (
	"github.com/samalba/dockerclient"
	"log"
)

func main() {
	// Init the client
	docker, _ := dockerclient.NewDockerClient("tcp://172.17.8.101:2375", nil)

	// Pull
	log.Println("Pulling")
	err := docker.PullImage("quay.io/epipho/docker:tag37")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Pulling Complete")
}
