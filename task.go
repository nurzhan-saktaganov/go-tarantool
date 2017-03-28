package tarantool

type Task struct {
	id uint64
	status string
	data interface{}
}

var DEFAULT_TASK = Task{}

func (t Task) GetId() uint64 {
	return t.id
}

func (t Task) GetData() interface{} {
	return t.data
}

func (t Task) GetStatus() string {
	return t.status
}

func (t Task) IsReady() bool {
	return t.status == READY
}

func (t Task) IsTaken() bool {
	return t.status == TAKEN
}

func (t Task) IsDone() bool {
	return t.status == DONE
}

func (t Task) IsBuried() bool {
	return t.status == BURIED
}

func (t Task) IsDelayed() bool {
	return t.status == DELAYED
}