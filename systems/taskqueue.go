package systems

type TaskImplementation func(...interface{})

type Task struct {
	implementation TaskImplementation
	args []interface{}
}

func (t Task) Call(){
	t.implementation(t.args...)
}

func NewTask(impl TaskImplementation, args ...interface{}) Task{
	return Task{
		implementation: impl,
		args: args,
	}
}

type TaskQueue struct {
	tasks []Task
}

func (tq *TaskQueue) AddTask(task Task){
	tq.tasks = append(tq.tasks, task)
}

func (tq *TaskQueue) RunNext(){
	if tq.tasks == nil || len(tq.tasks) == 0 {
		return
	}

	current := tq.tasks[0]
	tq.tasks = tq.tasks[1:]
	current.Call()
}