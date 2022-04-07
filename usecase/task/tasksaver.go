package task

import "toporet/hop/goclean/entity"

type TaskSaver interface {
	SaveNewTask(t entity.Task) (*entity.TaskId, error)
	SaveTask(t entity.Task) error
}
