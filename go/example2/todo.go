package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
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

	t := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}
	*l = append(*l, t)
}

func (l *List) Complete(i int) error {
	ls := *l

	if i <= 0 || i > len(ls) {
		return fmt.Errorf("Item %d does not exist", i)
	}

	ls[i-1].Done = true
	ls[i-1].CompletedAt = time.Now()

	return nil
}

func (l *List) Delete(i int) error {
	ls := *l

	if i <= 0 || i > len(ls) {
		return fmt.Errorf("Item %d does not exist", i)
	}
	*l = append(ls[:i-1], ls[i:]...)

	return nil
}

// Save method encodes the List as JSON and saves it​
// ​// using the provided file name​
func (l *List) Save(filename string) error {
	js, err := json.Marshal(l)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, js, 0644)
}

func (l *List) Get(filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	if len(file) == 0 {
		return nil
	}

	return json.Unmarshal(file, l)
}

func (l *List) PrintAll() {
	if len(*l) == 0 {
		fmt.Println("No tasks in the list.")
		return
	}

	for i, t := range *l {
		status := " "
		if t.Done {
			status = "✓"
		}

		fmt.Printf("%d. [%s] %s (created: %s",
			i+1,
			status,
			t.Task,
			t.CreatedAt.Format("2006-01-02 15:04"),
		)

		if t.Done {
			fmt.Printf(", completed: %s",
				t.CompletedAt.Format("2006-01-02 15:04"),
			)
		}

		fmt.Println(")")
	}
}

func (l *List) String() string {
	formatted := " "

	for k, t := range *l {
		prefix := "  "
		if t.Done {
			prefix = "X "
		}

		formatted += fmt.Sprintf("%s%d: %s\n", prefix, k+1, t.Task)
	}
	return formatted
}
