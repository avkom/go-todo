package main

type TaskService interface {
	GetList() (tasks []TaskData, err error)
	GetByID(id string) (task TaskData, err error)
	Create(task TaskData) error
	Update(task TaskData) error
	Delete(id string) error
}

type TaskRepository interface {
	GetList() (tasks []TaskData, err error)
	GetByID(id string) (task TaskData, err error)
	Create(task TaskData) error
	Update(task TaskData) error
	Delete(id string) error
}

type TaskServiceImpl struct {
	repository TaskRepository
}

func NewTaskService(repository TaskRepository) TaskService {
	return &TaskServiceImpl{repository}
}

func (s *TaskServiceImpl) GetList() (tasks []TaskData, err error) {
	return s.repository.GetList()
}

func (s *TaskServiceImpl) GetByID(id string) (task TaskData, err error) {
	return s.repository.GetByID(id)
}

func (s *TaskServiceImpl) Create(task TaskData) error {
	return s.repository.Create(task)
}

func (s *TaskServiceImpl) Update(task TaskData) error {
	return s.repository.Update(task)
}

func (s *TaskServiceImpl) Delete(id string) error {
	return s.repository.Delete(id)
}
