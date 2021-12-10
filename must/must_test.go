package must

import (
	"errors"
	"fmt"
	"testing"
)

func funcPanickedWithError(t *testing.T, f func()) (didpanic bool, err error) {
	t.Helper()
	defer func() {
		var ok bool
		recval := recover()
		if recval != nil {
			didpanic = true
			if err, ok = recval.(error); !ok {
				panic(fmt.Errorf("panic recovered, but recover type was %T, not error", recval))
			}
		}
	}()
	f()
	return
}

func TestDoDoesPanic(t *testing.T) {
	if p, _ := funcPanickedWithError(t, func() { Do("task", errors.New("error")) }); !p {
		t.Fatal("did not panic")
	}
}

func TestDoDoesntPanic(t *testing.T) {
	if p, err := funcPanickedWithError(t, func() { Do("task", nil) }); p {
		t.Fatalf("did panic with error %v", err)
	}
}

func TestTaskDoDoesPanic(t *testing.T) {
	if p, _ := funcPanickedWithError(t, func() { Task("task").Do(errors.New("error")) }); !p {
		t.Fatal("did not panic")
	}
}

func TestTaskDoDoesntPanic(t *testing.T) {
	if p, err := funcPanickedWithError(t, func() { Task("task").Do(nil) }); p {
		t.Fatalf("did panic with error %v", err)
	}
}

func TestTaskDo2DoesPanic(t *testing.T) {
	if p, _ := funcPanickedWithError(t, func() { Task("task").Do2(0, errors.New("error")) }); !p {
		t.Fatal("did not panic")
	}
}

func TestTaskDo2DoesntPanic(t *testing.T) {
	if p, err := funcPanickedWithError(t, func() { Task("task").Do2(0, nil) }); p {
		t.Fatalf("did panic with error %v", err)
	}
}

func TestTaskDo3DoesPanic(t *testing.T) {
	if p, _ := funcPanickedWithError(t, func() { Task("task").Do3(0, 1, errors.New("error")) }); !p {
		t.Fatal("did not panic")
	}
}

func TestTaskDo3DoesntPanic(t *testing.T) {
	if p, err := funcPanickedWithError(t, func() { Task("task").Do3(0, 1, nil) }); p {
		t.Fatalf("did panic with error %v", err)
	}
}
