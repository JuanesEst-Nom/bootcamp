package main

import (
	"flag"
	"fmt"
	"os"
	"todo"
)

const todoFile = ".todo.json"

func main() {
	list := flag.Bool("list", false, "List all incomplete tasks")
	task := flag.String("task", "", "Add a new task")
	complete := flag.Int("complete", 0, "Mark task as complete by number")
	del := flag.Int("delete", 0, "Delete a task by number")

	flag.Parse()

	l := &todo.List{}
	if err := l.Get(todoFile); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case *list:
		for _, t := range *l {
			if !t.Done {
				fmt.Printf("Title: %s, Done: %t, CreatedAt: %s, CompletedAt: %s\n",
					t.Task, t.Done, t.CreatedAt, t.CompletedAt)
			}
		}
	case *complete > 0:
		if err := l.Complete(*complete - 1); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if err := l.Save(todoFile); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *del > 0:
		if err := l.Delete(*del - 1); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if err := l.Save(todoFile); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *task != "":
		l.Add(*task)
		if err := l.Save(todoFile); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		fmt.Fprintln(os.Stderr, "no command provided")
		os.Exit(1)
	}
}
