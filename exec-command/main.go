package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

type CapturingPassThroughWriter struct {
	buf bytes.Buffer
	w   io.Writer
}

func (w *CapturingPassThroughWriter) Write(d []byte) (int, error) {
	w.buf.Write(d)
	return w.w.Write(d)
}

func main() {
	fmt.Println("Starting command with progress...")

	cmd := exec.Command("ls", "-lah")

	stdoutWriter := &CapturingPassThroughWriter{w: os.Stdout}
	cmd.Stdout = stdoutWriter
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n--- Final Summary (from buffer) ---\nSize: %d bytes\n", stdoutWriter.buf.Len())
}
