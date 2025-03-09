package todocli

import (
	"encoding/json"
	"errors"
	"os"
	"time"
)

type item struct {
	Task         string
	Done         bool
	CompeletedAt time.Time
	CreatedAt    time.Time
}

type Todos []item

func (t *Todos) Load(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}

		return err
	}

	if len(file) == 0 {
		return err
	}

	err = json.Unmarshal(file, t)
	if err != nil {
		return err
	}

	return nil
}

func (t *Todos) Store(filename string) error {
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0o644)
}

func (t *Todos) Add(task string) {
	todo := item{
		Task:         task,
		Done:         false,
		CompeletedAt: time.Time{},
		CreatedAt:    time.Now(),
	}

	*t = append(*t, todo)
}

func (t *Todos) Completed(index int) error {
	ls := *t

	if index <= 0 || index > len(ls) {
		return errors.New("invalid index")
	}

	ls[index-1].CompeletedAt = time.Now()
	ls[index-1].Done = true

	return nil
}

func (t *Todos) Delete(index int) error {
	ls := *t

	if index <= 0 || index > len(ls) {
		return errors.New("invalid index")
	}

	*t = append(ls[:index-1], ls[index:]...)

	return nil
}
