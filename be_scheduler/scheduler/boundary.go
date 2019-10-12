package scheduler

import "time"

// Master task consisting of different processes
type Task struct {
	Name       string    `json:"name"`
	ScheduleAt time.Time `json:"schedule_at"`
	Processes  []Process `json:"processes"`
}

// Child process
type Process struct {
	Name     string `json:"name"`
	WorkDone int    `json:"work_done"`
	Status   string `json:"status"`
}

/* Schedule task takes a master task and breaks it into multiple child processes
 * Master task can either be scheduled or invoked ad-hoc
 */
func HandleTask(task Task) error {
	return handleTask(task)
}
