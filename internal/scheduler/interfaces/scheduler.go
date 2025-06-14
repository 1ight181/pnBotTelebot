package interfaces

type Scheduler interface {
	AddJob(spec string, job func()) (int, error)
	RemoveJob(id int) error
	Start()
	Stop()
}
