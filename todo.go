package todocli

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/alexeyco/simpletable"
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

	return os.WriteFile(filename, []byte(data), 0o644)
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

func (t *Todos) Print() {
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Task"},
			{Align: simpletable.AlignCenter, Text: "Done"},
			{Align: simpletable.AlignCenter, Text: "Completed At"},
			{Align: simpletable.AlignCenter, Text: "Created At"},
		},
	}

	var cells [][]*simpletable.Cell

	for idx, item := range *t {
		idx++

		task := blue(item.Task)

		if item.Done {
			task = green(fmt.Sprintf("\u2705 %s", item.Task))
		}

		cells = append(cells, *&[]*simpletable.Cell{
			{Text: fmt.Sprintf("%d", idx)},
			{Text: task},
			{Text: fmt.Sprintf("%t", item.Done)},
			{Text: item.CompeletedAt.Format(time.RFC822)},
			{Text: item.CreatedAt.Format(time.RFC822)},
		})
	}

	table.Body = &simpletable.Body{Cells: cells}

	table.Footer = &simpletable.Footer{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Span: 5, Text: "You todos here"},
		},
	}

	table.SetStyle(simpletable.StyleUnicode)

	table.Println()
}
