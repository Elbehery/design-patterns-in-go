package main

import (
	"fmt"
	"sync"
)

var (
	once      sync.Once
	singleton *WorkerPool
)

type Job struct {
	ID int
}

type Worker struct {
	ID         int
	Quit       chan bool
	JobChannel chan Job
	WorkerPool chan chan Job
}

func NewWorker(workerPool chan chan Job, id int) *Worker {
	return &Worker{
		ID:         id,
		Quit:       make(chan bool),
		JobChannel: make(chan Job),
		WorkerPool: workerPool,
	}
}

func (w *Worker) Start(wg *sync.WaitGroup) {
	go func() {
		for {
			w.WorkerPool <- w.JobChannel
			select {
			case job := <-w.JobChannel:
				fmt.Printf("Wokrer '%v' has been assgined Job '%v'\n", w.ID, job)
				wg.Done()
			case <-w.Quit:
				return
			}
		}
	}()
}

func (w *Worker) Stop() {
	go func() {
		w.Quit <- true
	}()
}

type WorkerPool struct {
	workers    []*Worker
	jobQueue   chan Job
	workerPool chan chan Job
	wg         sync.WaitGroup
}

func (wp *WorkerPool) AddJob(job Job) {
	wp.wg.Add(1)
	wp.jobQueue <- job
}

func (wp *WorkerPool) Wait() {
	wp.wg.Wait()
}

func (wp *WorkerPool) dispatch() {
	for job := range wp.jobQueue {
		workerJobQueue := <-wp.workerPool
		workerJobQueue <- job
	}
}

func GetWorkerPool(maxWorkers int) *WorkerPool {
	once.Do(func() {

		workers := make([]*Worker, maxWorkers)
		jobQueue := make(chan Job, maxWorkers)
		workerPool := make(chan chan Job, maxWorkers)

		pool := &WorkerPool{
			workers:    workers,
			jobQueue:   jobQueue,
			workerPool: workerPool,
			wg:         sync.WaitGroup{},
		}

		for i := 0; i < maxWorkers; i++ {
			w := NewWorker(workerPool, i)
			w.Start(&pool.wg)
			workers[i] = w
		}

		go pool.dispatch()

		singleton = pool

	})

	return singleton
}

func main() {
	wp := GetWorkerPool(5)

	for i := 1; i <= 20; i++ {
		job := Job{ID: i}
		wp.AddJob(job)
	}

	wp.Wait()
}
