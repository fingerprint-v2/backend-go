package services

import (
	"context"
	"fmt"
	"sync"
)

type DispatcherService interface {
	Start(ctx context.Context)
	Wait()
	Add(job *Job)
	Stop()
	Loop(ctx context.Context)
}
type DispatcherServiceImpl struct {
	sem       chan struct{}
	jobBuffer chan *Job
	errCh     chan error
	worker    WorkerService
	wg        sync.WaitGroup
}

func NewDispatcherService(worker WorkerService) DispatcherService {
	return &DispatcherServiceImpl{
		sem:       make(chan struct{}, 10), // 10 = max workers
		jobBuffer: make(chan *Job, 100),    // 100 = max jobs
		errCh:     make(chan error),
		worker:    worker,
	}
}

func (d *DispatcherServiceImpl) Start(ctx context.Context) {
	d.wg.Add(1)
	go d.Loop(ctx)
}

func (d *DispatcherServiceImpl) Wait() {
	d.wg.Wait()
}

func (d *DispatcherServiceImpl) Add(job *Job) {
	d.jobBuffer <- job
}

func (d *DispatcherServiceImpl) Stop() {
	d.wg.Done()
}

func (d *DispatcherServiceImpl) Loop(ctx context.Context) {
	var wg sync.WaitGroup
Loop:
	for {
		select {
		case <-ctx.Done():
			// Block until all the jobs finish
			wg.Wait()
			break Loop

		case job := <-d.jobBuffer:
			// Increment the waitgroup
			wg.Add(1)
			// Decrement a semaphore count
			d.sem <- struct{}{}
			go func(job *Job) {
				defer wg.Done()
				// After the job finished, increment a semaphore count
				defer func() { <-d.sem }()
				if err := d.worker.Work(job); err != nil {
					d.errCh <- err
				}
			}(job)

		case err := <-d.errCh:
			fmt.Printf("Error: %s\n", err.Error())
		}
	}

	d.Stop()
}
