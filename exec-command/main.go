package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("ls", "-lah", "non-existent-file")
	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	_ = cmd.Run()
	fmt.Printf("Stdout: %s\nStderr: %s\n", outb.String(), errb.String())
}
