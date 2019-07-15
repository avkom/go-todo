package main

type TaskService interface {
	GetList() (tasks []TaskData, err error)
	GetById(id string) (task TaskData, err error)
	Create(task TaskData) error
	Update(task TaskData) error
	Delete(id string) error
}

type TaskServiceImpl struct {
}

func NewTaskService() TaskService {
	return &TaskServiceImpl{}
}

func (s *TaskServiceImpl) GetList() (tasks []TaskData, err error) {
	return
}

func (s *TaskServiceImpl) GetById(id string) (task TaskData, err error) {
	return
}

func (s *TaskServiceImpl) Create(task TaskData) error {
	return nil
}

func (s *TaskServiceImpl) Update(task TaskData) error {
	return nil
}

func (s *TaskServiceImpl) Delete(id string) error {
	return nil
}
