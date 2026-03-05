package main

import (
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("ls", "-lah")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
