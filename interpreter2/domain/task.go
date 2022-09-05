package domain

type Task struct {
	num     int
	done    int
	allDone bool
}

func (t *Task) Add(num int) {
	t.num = num
}

func (t *Task) Done() {
	t.done++
	if t.done == t.num {
		t.allDone = true
	}
}

func (t *Task) HasTask() bool {
	return !t.allDone
}
