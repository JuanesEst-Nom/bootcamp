package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

var (
	binName  = "todo"
	fileName = ".todo.json"
)

func TestMain(m *testing.M) {
	fmt.Println("Building tool...")

	build := exec.Command("go", "build", "-o", binName)
	if out, err := build.CombinedOutput(); err != nil {
		fmt.Fprintf(os.Stderr, "Cannot build %s: %s\n%s", binName, err, string(out))
		os.Exit(1)
	}

	result := m.Run()

	os.Remove(binName)
	os.Remove(fileName)

	os.Exit(result)
}

func TestAddNewTask(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	cmdPath := filepath.Join(dir, binName)
	task := "buy groceries"

	cmd := exec.Command(cmdPath, "-task", task)
	if err := cmd.Run(); err != nil {
		t.Fatalf("expected no error, got %s", err)
	}
}

func TestListTasks(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	cmdPath := filepath.Join(dir, binName)

	cmd := exec.Command(cmdPath, "-list")
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("expected no error, got %s", err)
	}

	expected := "buy groceries"
	if !strings.Contains(string(out), expected) {
		t.Errorf("expected %q in output, got %q", expected, string(out))
	}
}

func TestCompleteTask(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	cmdPath := filepath.Join(dir, binName)

	cmd := exec.Command(cmdPath, "-complete", "1")
	if err := cmd.Run(); err != nil {
		t.Fatalf("expected no error, got %s", err)
	}

	listCmd := exec.Command(cmdPath, "-list")
	out, err := listCmd.CombinedOutput()
	if err != nil {
		t.Fatalf("expected no error, got %s", err)
	}

	if strings.Contains(string(out), "buy groceries") {
		t.Error("expected completed task to not appear in incomplete task list")
	}
}

func TestDeleteTask(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	cmdPath := filepath.Join(dir, binName)

	addCmd := exec.Command(cmdPath, "-task", "task to delete")
	if err := addCmd.Run(); err != nil {
		t.Fatalf("expected no error adding task, got %s", err)
	}

	// After TestCompleteTask: [buy groceries (done=1)]
	// After adding "task to delete": [buy groceries (done=1), task to delete (done=2)]
	delCmd := exec.Command(cmdPath, "-delete", "2")
	if err := delCmd.Run(); err != nil {
		t.Fatalf("expected no error deleting task, got %s", err)
	}

	listCmd := exec.Command(cmdPath, "-list")
	out, err := listCmd.CombinedOutput()
	if err != nil {
		t.Fatalf("expected no error listing tasks, got %s", err)
	}

	if strings.Contains(string(out), "task to delete") {
		t.Error("expected deleted task to not appear in list")
	}
}
