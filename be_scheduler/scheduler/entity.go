package scheduler

import (
	"../base"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Writes details about new task and all its child components
func addNewTasksEntity(data bson.M) error {

	log.Println("adding new task", data)
	session, err := mgo.Dial(base.MONGO_BASE_URL)
	if err != nil {
		log.Println("ERROR:", err)
	}
	defer session.Close()

	c := session.DB(base.DB_ACTYV).C(base.COL_TASKS)

	err = c.Insert(data)
	if err != nil {
		log.Println("ERROR:", err)
	}

	return err
}

// Updates the details and status of the child processes
func updateProcessStatusEntity(name, status string) error {

	log.Println("Updating process status", name, status)
	session, err := mgo.Dial(base.MONGO_BASE_URL)
	if err != nil {
		log.Println("ERROR:", err)
	}
	defer session.Close()

	c := session.DB(base.DB_ACTYV).C(base.COL_TASKS)

	query := bson.M{
		"processes": bson.M{
			"$elemMatch": bson.M{
				"name": name,
			},
		},
	}
	update := bson.M{
		"$set": bson.M{
			"processes.$.status": status,
		},
	}
	err = c.Update(query, update)
	if err != nil {
		log.Println("ERROR:", err)
	}

	return err
}

// Updates the task status as soon as child processes are completed
func updateTaskStatusEntity(name, status string) error {

	log.Println("updating task status", name, status)

	session, err := mgo.Dial(base.MONGO_BASE_URL)
	if err != nil {
		log.Println("ERROR:", err)
	}
	defer session.Close()

	c := session.DB(base.DB_ACTYV).C(base.COL_TASKS)

	query := bson.M{
		"name": name,
	}
	update := bson.M{
		"$set": bson.M{
			"status": status,
		},
	}
	err = c.Update(query, update)
	if err != nil {
		log.Println("ERROR:", err)
	}

	return err
}

// Read all pending tasks and their corresponding child processes
func readTaskDetailsEntity() ([]Task, error) {

	log.Println("reading task details")

	session, err := mgo.Dial(base.MONGO_BASE_URL)
	if err != nil {
		log.Println("ERROR:", err)
	}
	defer session.Close()

	c := session.DB(base.DB_ACTYV).C(base.COL_TASKS)

	query := bson.M{
		"status": StatusPending,
	}

	var tasks []Task
	err = c.Find(query).All(&tasks)
	if err != nil {
		log.Println("ERROR:", err)
	}
	log.Println("Tasks remaining:", len(tasks))

	return tasks, err
}
