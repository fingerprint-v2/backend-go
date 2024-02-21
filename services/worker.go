package services

type Job struct {
	Function func() error
}

type WorkerService interface {
	Work(j *Job) error
}

type WorkerServiceImpl struct{}

func NewWokerService() WorkerService {
	return &WorkerServiceImpl{}
}

func (p *WorkerServiceImpl) Work(j *Job) error {
	return j.Function()
}
