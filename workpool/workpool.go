/*
所有worker从一个全局channel中取job
*/

package workpool

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/net/context"
)

type Job struct {
	JobID int
}

func (job *Job) execute() Result {
	//模拟业务耗时
	time.Sleep(time.Second)
	fmt.Println("good work!")
	return Result{
		JobID: job.JobID,
	}
}

type Result struct {
	JobID int
	err   error
}

type WorkerPool struct {
	workerCount int
	jobs        chan Job
	results     chan Result
	Done        chan struct{}
}

func worker(ctx context.Context, wg sync.WaitGroup, jobs <-chan Job, results chan<- Result) {
	defer wg.Done()

	for {
		select {
		case job, ok := <-jobs:
			if !ok {
				return
			}

			results <- job.execute()

		case <-ctx.Done():
			results <- Result{
				err: fmt.Errorf("worker cancled"),
			}
			return
		}
	}

}

func NewWorkerPool(workerCount int) *WorkerPool {
	return &WorkerPool{
		workerCount: workerCount,
		jobs:        make(chan Job, workerCount),
		results:     make(chan Result, workerCount),
		Done:        make(chan struct{}),
	}
}

func (pool *WorkerPool) Run(ctx context.Context) {
	var wg sync.WaitGroup
	for i := 0; i < pool.workerCount; i++ {
		wg.Add(1)
		go worker(ctx, wg, pool.jobs, pool.results)
	}

	wg.Wait()
	close(pool.results)
	close(pool.Done)
}

func (pool *WorkerPool) Result() <-chan Result {
	return pool.results
}

func (pool *WorkerPool) AddJob(jobs []Job) {
	for _, job := range jobs {
		pool.jobs <- job
	}

	close(pool.jobs)
}
