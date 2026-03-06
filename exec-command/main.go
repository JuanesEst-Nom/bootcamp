package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("ls", "-lah")
	var b bytes.Buffer
	cmd.Stdout = &b
	cmd.Stderr = &b
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Manual buffer output:\n%s", b.Bytes())
}
