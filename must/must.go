package must

import "fmt"

type Task string

func (t Task) Do(err error) {
	Do(string(t), err)
}

func (t Task) Do2(_ interface{}, err error) {
	Do(string(t), err)
}

func (t Task) Do3(_ interface{}, _ interface{}, err error) {
	Do(string(t), err)
}

func Do(task string, err error) {
	if err == nil {
		return
	}
	panic(fmt.Errorf("failed to %s: %w", task, err))
}
