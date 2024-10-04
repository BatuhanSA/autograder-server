package tasks

import (
	"sync"

	"github.com/edulinq/autograder/internal/config"
)

// An object that can be used to control some task functionality.
// This is a way to break the reliance on the task package for things like stopping the tasks for a course.
// The task package will monitor this object and act on it.
// Note that all methods are noops in testing mode.
var Handler *TasksHandler = newTaskHandler()

type TasksHandler struct {
	lock     sync.Mutex
	DoneChan chan error

	StopCourseChan chan string
	ScheduleChan   chan *SchedulePayload
}

type SchedulePayload struct {
	Course any
	Target ScheduledTask
}

func newTaskHandler() *TasksHandler {
	return &TasksHandler{
		DoneChan:       make(chan error),
		StopCourseChan: make(chan string),
		ScheduleChan:   make(chan *SchedulePayload),
	}
}

// See task.StopCourse().
func (this *TasksHandler) StopCourse(courseID string) {
	this.lock.Lock()
	defer this.lock.Unlock()

	// Skip in testing mode.
	if config.TESTING_MODE.Get() {
		return
	}

	this.StopCourseChan <- courseID

	// StopCourse() cannot return errors.
	<-this.DoneChan
}

// See task.Schedule().
// |course| must be a *model.Course, and |target| must be a pointer
// (the same semantics as task.Schedule()).
func (this *TasksHandler) Schedule(course any, target ScheduledTask) error {
	this.lock.Lock()
	defer this.lock.Unlock()

	// Skip in testing mode.
	if config.TESTING_MODE.Get() {
		return nil
	}

	payload := &SchedulePayload{
		Course: course,
		Target: target,
	}

	this.ScheduleChan <- payload

	return <-this.DoneChan
}