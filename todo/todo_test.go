package todo

import (
	"os"
	"testing"
)

func TestAdd(t *testing.T) {
	var l List

	l.Add("Buy groceries")
	l.Add("Walk the dog")

	if len(l) != 2 {
		t.Errorf("expected length 2, got %d", len(l))
	}

	if l[0].Task != "Buy groceries" {
		t.Errorf("expected task 'Buy groceries', got '%s'", l[0].Task)
	}

	if l[1].Task != "Walk the dog" {
		t.Errorf("expected task 'Walk the dog', got '%s'", l[1].Task)
	}
}

func TestComplete(t *testing.T) {
	var l List
	l.Add("Buy groceries")
	l.Add("Walk the dog")

	if err := l.Complete(0); err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if l[0].Task != "Buy groceries" {
		t.Errorf("expected task 'Buy groceries', got '%s'", l[0].Task)
	}
	if !l[0].Done {
		t.Errorf("expected task 'Buy groceries' to be done")
	}

	if err := l.Complete(-1); err == nil {
		t.Error("expected error for index -1, got nil")
	}
	if err := l.Complete(99); err == nil {
		t.Error("expected error for out-of-range index, got nil")
	}
}

func TestDelete(t *testing.T) {
	var l List
	l.Add("Buy groceries")
	l.Add("Walk the dog")
	l.Add("Read a book")

	if err := l.Delete(1); err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if len(l) != 2 {
		t.Errorf("expected length 2 after delete, got %d", len(l))
	}

	if l[0].Task != "Buy groceries" {
		t.Errorf("expected first task to be 'Buy groceries', got '%s'", l[0].Task)
	}
	if l[1].Task != "Read a book" {
		t.Errorf("expected second task to be 'Read a book', got '%s'", l[1].Task)
	}

	if err := l.Delete(-1); err == nil {
		t.Error("expected error for index -1, got nil")
	}
	if err := l.Delete(99); err == nil {
		t.Error("expected error for out-of-range index, got nil")
	}
}

func TestSaveAndGet(t *testing.T) {
	tf, err := os.CreateTemp("", "todo_test_*.json")
	if err != nil {
		t.Fatalf("could not create temp file: %v", err)
	}
	defer os.Remove(tf.Name())
	tf.Close()

	original := List{}
	original.Add("Buy groceries")
	original.Add("Walk the dog")
	original.Complete(0)

	if err := original.Save(tf.Name()); err != nil {
		t.Fatalf("could not save list: %v", err)
	}

	var loaded List
	if err := loaded.Get(tf.Name()); err != nil {
		t.Fatalf("could not get list: %v", err)
	}

	if len(loaded) != len(original) {
		t.Errorf("expected length %d, got %d", len(original), len(loaded))
	}

	for i := range original {
		if loaded[i].Task != original[i].Task {
			t.Errorf("item %d: expected task '%s', got '%s'", i, original[i].Task, loaded[i].Task)
		}
		if loaded[i].Done != original[i].Done {
			t.Errorf("item %d: task '%s': expected Done=%v, got Done=%v", i, original[i].Task, original[i].Done, loaded[i].Done)
		}
	}
}
