package scheduler

import (
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	"log"
	"strconv"
	"time"
)

const (
	StatusCompleted = "done"
	StatusPending   = "backlog"
)

/**
 * On initlization module checks for pending tasks
 * Each pending task is resumed
 */
func init() {
	resumePendingTasks()
}

// Filter the tasks on the basis of time
func handleTask(task Task) error {

	printDetails(task)

	data, err := getBSON(task)
	if err != nil {
		log.Println("ERROR:", err)
		return err
	}

	err = addNewTasksEntity(data)
	if err != nil {
		log.Println("ERROR:", err)
		return err
	}

	if time.Since(task.ScheduleAt).Seconds() < 0 {
		time.AfterFunc(task.ScheduleAt.Sub(time.Now()), func() { runTask(task) })
	} else {
		go runTask(task)
	}

	return nil
}

// Runs all the processes of a task parallelly, Marks the task as completed when all the child processes are done
func runTask(task Task) error {

	log.Println("Running Major Task")
	done := make(chan bool)
	log.Println("Total processes:", len(task.Processes))
	for i := range task.Processes {
		go callProcess(task.Processes[i], done)
	}

	for i := 0; i < len(task.Processes); i++ {
		<-done
	}

	log.Println("Major task completed!")

	updateTaskStatusEntity(task.Name, StatusCompleted)
	return nil
}

// Runs a process on a different thread while mainting status with a channel
func callProcess(process Process, done chan<- bool) {
	log.Println(process.Name, "Started")
	time.Sleep(time.Second * time.Duration(process.WorkDone))
	log.Println(process.Name, "Completed")
	updateProcessStatusEntity(process.Name, StatusCompleted)
	done <- true
}

// Prints details of task
func printDetails(task Task) {

	details := "\n\nTask Name: " + task.Name
	details += "\nNo. of Child Processes: " + strconv.Itoa(len(task.Processes))
	if time.Since(task.ScheduleAt).Seconds() < 0 {
		details += "\nScheduled at " + task.ScheduleAt.String()
	} else {
		details += "\nRunning it now"
	}
	log.Println(details, "\n\n")
}

// Resumes all pending tasks
func resumePendingTasks() error {

	tasks, err := readTaskDetailsEntity()
	if err != nil {
		log.Println("ERROR:", err)
		return err
	}

	for i := range tasks {
		go handleTask(tasks[i])
	}
	return nil
}
func getBSON(task Task) (bson.M, error) {

	jsonData, err := json.Marshal(task)
	if err != nil {
		log.Println("ERROR:", err)
		return bson.M{}, err
	}

	var data bson.M
	err = bson.Unmarshal(jsonData, &data)
	if err != nil {
		log.Println("ERROR:", err)
	}
	return data, err
}
