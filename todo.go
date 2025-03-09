package todocli

import "time"

type item struct {
	Task        string
	Done        bool
	ComplatedAt time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

type Todos []item

func (t *Todos) Add(task string) {
	todo := item{
		Task:        task,
		Done:        false,
		ComplatedAt: time.Time{},
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   time.Time{},
	}

	*t = append(*t, todo)
}
