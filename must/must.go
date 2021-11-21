package must

import "fmt"

func Do(task string, err error) {
	if err == nil {
		return
	}
	panic(fmt.Errorf("failed to %s: %w", task, err))
}
