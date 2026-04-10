package todo

import (
	"encoding/json"
	"errors"
	"os"
	"time"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type List []item

func (l *List) Add(task string) {
	nuevoItem := item{
		Task:      task,
		Done:      false,
		CreatedAt: time.Now(),
	}

	*l = append(*l, nuevoItem)
}

func (l *List) Complete(i int) error {
	ls := *l

	if i < 0 || i >= len(ls) {

		return errors.New("invalid index")
	}

	ls[i].Done = true
	ls[i].CompletedAt = time.Now()
	return nil
}

func (l *List) Delete(i int) error {
	ls := *l

	if i < 0 || i >= len(ls) {
		return errors.New("invalid index")
	}

	*l = append(ls[:i], ls[i+1:]...)
	return nil
}

func (l *List) Save(filename string) error {

	data, err := json.Marshal(l)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}

func (l *List) Get(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	if len(data) == 0 {
		return nil
	}

	return json.Unmarshal(data, l)
}
