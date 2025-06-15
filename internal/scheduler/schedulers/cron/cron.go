package schedulders

import (
	"errors"

	c "github.com/robfig/cron/v3"
)

type CronScheduler struct {
	cron *c.Cron
	jobs map[int]c.EntryID
	next int
}

func NewCronScheduler(
	cron *c.Cron,
	jobs map[int]c.EntryID,
) *CronScheduler {
	return &CronScheduler{
		cron: cron,
		jobs: jobs,
		next: 1,
	}
}

func (cs *CronScheduler) AddJob(
	spec string,
	job func(),
) (int, error) {
	id, err := cs.cron.AddFunc(spec, job)
	if err != nil {
		return 0, err
	}
	cs.next++
	cs.jobs[cs.next] = id
	return cs.next, nil
}

func (cs *CronScheduler) RemoveJob(id int) error {
	entryID, exists := cs.jobs[id]
	if !exists {
		return errors.New("задача с таким ID не найдена")
	}
	delete(cs.jobs, id)
	cs.cron.Remove(entryID)
	return nil
}

func (cs *CronScheduler) Start() {
	cs.cron.Start()
}

func (cs *CronScheduler) Stop() {
	cs.cron.Stop()
}
